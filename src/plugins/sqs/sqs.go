// AWS関連をまとめたい
package main

import (
	"strings"
//	"strconv"
	"log"
	"regexp"
//	"sort"
	"fmt"
//	"time"
	"sync"
//	"encoding/json"
//	"io/ioutil"
	
    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Route struct {
	route 		*regexp.Regexp
	fnc   		func([]string) string
	help  		string
	//	arg   *[]string
}

type SlackEc2 struct {
	session *session.Session
//	ec2 	*ec2.EC2
	sqs		*sqs.SQS
	routes  *[]Route
}

type AwsEnv struct{
	Cluster				string		`json:"Cluster"`
	Service    			string		`json:"Service"`
	RepositoryName  	string		`json:"RepositoryName"`
	TaskDefinition		string		`json:"TaskDefinition"`
	ContainerDefinition	string		`json:"ContainerDefinition"`
}

var awsEnv map[string]AwsEnv

const (
    QUEUE_URL  = "your Queue URL"
)
var s SlackEc2

func Init() {
	
	s.session = session.Must(session.NewSession(aws.NewConfig().WithLogLevel(aws.LogDebugWithRequestRetries | aws.LogDebugWithRequestErrors).WithRegion("ap-northeast-1")))
	log.Printf("region=" + *s.session.Config.Region)
//	s.ec2 = ec2.New(s.session)
    s.sqs = sqs.New(s.session)
/*
	bytes, err := ioutil.ReadFile("aws.json")
    if err != nil {
        log.Fatal("ReadJson[aws.json] error :" + err.Error())
    }
    if err := json.Unmarshal(bytes, &awsEnv); err != nil {
        log.Fatal("UnmarshalJson[aws.json] error :" + err.Error())
    }
*/
	routes := []Route{}
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^help`), fnc: cmdHelp, help: "'help'   ヘルプ表示" })
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^env`), fnc: cmdEnv, help: "'env'   環境一覧表示" })

	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^put\s+(.*)$`), fnc: cmdPut, help: "'put {msg}'   msgをQueueに追加" })
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^get$`), fnc: cmdGet, help: "'get'   Queueを取得" })
	
	s.routes = &routes
}

func OnMention(msgs []string) string {

	var response = ""

	ss := strings.TrimSpace(strings.Join(msgs[:], " "))

	for _, value := range *s.routes {
		match := value.route.FindStringSubmatch(ss)

		if match != nil {
			response = value.fnc(match[1:])
			break
		}
	}

	return response
}


func cmdEnv(msgs []string) string {
	var response string = ""

	for key, _ := range awsEnv {
		response += key
		response += ", "
	}

	return (response)
}

func cmdHelp(msgs []string) string {
	var response string = ""

	for _, value := range *s.routes {
		help := value.help

		if len(help) > 0 {
			response += help
			response += "\n"
		}
	}

	return (response)
}


func cmdPut(msgs []string) string {
	msg := "test message"

	if(len(msgs) > 0){
		msg = msgs[0]
	}

    params := &sqs.SendMessageInput{
        MessageBody:  aws.String(msg),
        QueueUrl:     aws.String(QUEUE_URL),
        DelaySeconds: aws.Int64(1),
    }

    sqsRes, err := s.sqs.SendMessage(params)
    if err != nil {
        return err.Error()
    }

    fmt.Println("SQSMessageID", *sqsRes.MessageId)

	return (*sqsRes.MessageId)
}



func cmdGet(msgs []string) string {
	respmsg := ""

    params := &sqs.ReceiveMessageInput{
        QueueUrl: aws.String(QUEUE_URL),
        MaxNumberOfMessages: aws.Int64(10),
        WaitTimeSeconds: aws.Int64(1),
    }
    resp, err := s.sqs.ReceiveMessage(params)

    if err != nil {
        return err.Error()
    }

    fmt.Printf("messages count: %d\n", len(resp.Messages))

    if len(resp.Messages) == 0 {
        return "empty queue."
    }

    // メッセージの数だけgoroutineを実行してみる。
    var wg sync.WaitGroup
    for _, m := range resp.Messages {
        wg.Add(1)
        go func(msg *sqs.Message) {
            defer wg.Done()
            respmsg += (*msg.Body) + "\n"
            if err := DeleteMessage(msg); err != nil {
            }
        }(m)
    }

    wg.Wait()

	return (respmsg)
}

func DeleteMessage(msg *sqs.Message) error {
    params := &sqs.DeleteMessageInput{
        QueueUrl:      aws.String(QUEUE_URL),
        ReceiptHandle: aws.String(*msg.ReceiptHandle),
    }
    _, err := s.sqs.DeleteMessage(params)

    if err != nil {
        return err
    }
    return nil
}
