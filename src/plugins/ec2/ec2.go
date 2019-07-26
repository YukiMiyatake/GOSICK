package main

import (
	"strings"
	// 	"plugin"
	"log"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Route struct {
	route *regexp.Regexp
	fnc   func([]string) string
	//	arg   *[]string
}

type SlackEc2 struct {
	// TODO: SessionはAWSだいたい共通だが今は Singletonにしない
	// TODO: List作る
	session *session.Session
	service *ec2.EC2
	routes  *[]Route
}

const DEV_INSTANCE_ID = "your instance id";

var s SlackEc2

func Init() {
	s.session = session.Must(session.NewSession(aws.NewConfig().WithLogLevel(aws.LogDebugWithRequestRetries | aws.LogDebugWithRequestErrors).WithRegion("ap-northeast-1")))
	log.Printf("region=" + *s.session.Config.Region)
	s.service = ec2.New(s.session)

	// router設定する
	routes := []Route{}
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^list$`), fnc: cmdList})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^list\s+(.*)`), fnc: cmdList})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^up$`), fnc: cmdUp})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^down$`), fnc: cmdDown})
	//	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^list\s+(.*)`), fnc: listImage})
	s.routes = &routes
}

func OnMention(msgs []string) string {
	// TODO: 正規表現による routes的な機能つけたい

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
	//	fmt.Printf("%s\n", msgs)
	return (strings.Join(msgs, ` `))
}

func cmdUp(msgs []string) string {
	response := ""

	// TODO: IP取得できるのでは？
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{aws.String(DEV_INSTANCE_ID)},
	}
	r, err := s.service.StartInstances(input)

	if err != nil {
		log.Printf("Error StartInstancesInput: " + err.Error())
	} else {
		log.Printf("%s\n",r);
		response += "おこすね！\n"
	}
	return (response)
}

func cmdDown(msgs []string) string {
	response := ""

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{aws.String(DEV_INSTANCE_ID)},
	}
	r, err := s.service.StopInstances(input)

	if err != nil {
		log.Printf("Error StopInstancesInput: " + err.Error())
	} else {
		log.Printf("%s\n",r);
		response += "かわりに殺しておくね～\n"
	}

	return (response)
}

