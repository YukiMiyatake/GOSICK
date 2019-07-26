package main

import (
	"strings"
	// 	"plugin"
	"strconv"
	"log"
	"regexp"
	"sort"
	"fmt"
	"time"

	
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

type Route struct {
	route *regexp.Regexp
	fnc   func([]string) string
}

type SlackEc2 struct {
	// TODO: SessionはAWSだいたい共通だが今は Singletonにしない
	session *session.Session
	ecr 	*ecr.ECR
	routes  *[]Route
}




var s SlackEc2

func Init() {
	s.session = session.Must(session.NewSession(aws.NewConfig().WithLogLevel(aws.LogDebugWithRequestRetries | aws.LogDebugWithRequestErrors).WithRegion("ap-northeast-1")))
	log.Printf("region=" + *s.session.Config.Region)
	s.ecr = ecr.New(s.session)
	
	// router設定する
	routes := []Route{}
	// bot ecr list
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^list$`), fnc: cmdList})
	// bot ecr list n
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^list\s+(.*)`), fnc: cmdList})

	s.routes = &routes
}

func OnMention(msgs []string) string {

	var response = ""

	ss := strings.Join(msgs[:], " ")

	for _, value := range *s.routes {
		match := value.route.FindStringSubmatch(ss)

		if match != nil {
			response = value.fnc(match[1:])
			break
		}
	}

	return response
}

func cmdList(msgs []string) string {
	log.Printf("listImage")

	var response = ""
	var list_count int64 = 3

	if( len(msgs) > 0){
		log.Printf(msgs[0])
		//		list_count_, err := strconv.Atoi(msgs[0])
		i, err := strconv.ParseInt(msgs[0], 10, 64)
		if( err == nil){
			list_count = i
		}		
	}

	env, b := awsEnv[msgs[0]];
	if(!b){
		return ( "その環境は存在しないよ : " + msgs[0] )		
	}

	input := &ecr.DescribeImagesInput{
		RepositoryName: aws.String(env.RepositoryName),
		MaxResults:     aws.Int64(100),
		Filter: &ecr.DescribeImagesFilter{
			TagStatus: aws.String("TAGGED"),
		},
	}

	result, err := s.ecr.DescribeImages(input)

	if err != nil {
		log.Printf("Error DescribeImages: " + err.Error())
	} else {
		// sort by ImagePushedAt desc
		sort.Slice(result.ImageDetails, func(i, j int) bool {
			return result.ImageDetails[i].ImagePushedAt.After(*result.ImageDetails[j].ImagePushedAt)
		})

		// Timezone変換？
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		
		if( list_count > int64(len(result.ImageDetails)) ){
			list_count = int64(len(result.ImageDetails))
		}

		for _, value := range result.ImageDetails[:list_count] {

			response += *value.ImageTags[0] + "    " + value.ImagePushedAt.In(jst).String() + "\n"
		}

		//s.rtm.SendMessage(s.rtm.NewOutgoingMessage(response, s.channelID))

		fmt.Println(result)
	}	
	
	//	fmt.Printf("%s\n", msgs)
	return (response)
}

