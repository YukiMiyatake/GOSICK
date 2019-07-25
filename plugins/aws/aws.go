// AWS関連をまとめたい
package main

import (
	"strings"
	// 	"plugin"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type Route struct {
	route *regexp.Regexp
	fnc   func([]string) string
	help  string
	//	arg   *[]string
}

type SlackEc2 struct {
	session *session.Session
	ec2     *ec2.EC2
	ecr     *ecr.ECR
	ecs     *ecs.ECS
	routes  *[]Route
}

type AwsEnv struct {
	Cluster             string `json:"Cluster"`
	Service             string `json:"Service"`
	RepositoryName      string `json:"RepositoryName"`
	TaskDefinition      string `json:"TaskDefinition"`
	ContainerDefinition string `json:"ContainerDefinition"`
}

var awsEnv map[string]AwsEnv

// TODO: 設定ファイルからリスト読み込み
const DEV_INSTANCE_ID = "your instance id"

var s SlackEc2

func Init() {

	s.session = session.Must(session.NewSession(aws.NewConfig().WithLogLevel(aws.LogDebugWithRequestRetries | aws.LogDebugWithRequestErrors).WithRegion("ap-northeast-1")))
	log.Printf("region=" + *s.session.Config.Region)
	s.ec2 = ec2.New(s.session)
	s.ecr = ecr.New(s.session)
	s.ecs = ecs.New(s.session)

	bytes, err := ioutil.ReadFile("aws.json")
	if err != nil {
		log.Fatal("ReadJson[aws.json] error :" + err.Error())
	}
	if err := json.Unmarshal(bytes, &awsEnv); err != nil {
		log.Fatal("UnmarshalJson[aws.json] error :" + err.Error())
	}

	routes := []Route{}
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^help`), fnc: cmdHelp, help: "'help'   ヘルプ表示"})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^env`), fnc: cmdEnv, help: "'env'   環境一覧表示"})

	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^list\s+(.*)\s+(.*)$`), fnc: cmdList, help: "'list env n'  最新のコンテナイメージ {n}件表示"})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^list\s+(.*)$`), fnc: cmdList, help: "'list env'   最新のコンテナイメージ3件表示"})

	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^deploy\s+(.*)\s+(.*)$`), fnc: cmdDeploy, help: "'deploy env img'  コンテナイメージ {img} をDeploy"})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^deploy\s+(.*)$`), fnc: cmdDeploy, help: "'deploy env'  最新のコンテナイメージ をDeploy"})

	// TODO: 複数インスタンスを登録し指定したい
	// 開発用インスタンス
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^up$`), fnc: cmdUp, help: "'up'  開発用インスタンスのUP"})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^down$`), fnc: cmdDown, help: "'down' 　開発用インスタンスのDown"})

	//debug
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^ListClusters\s+(.*)$`), fnc: cmdListClusters, help: "'ListClusters n'  クラスターリストn件表示"})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^ListClusters$`), fnc: cmdListClusters, help: "'ListClusters' 　クラスターリスト表示"})

	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^DescribeClusters\s+(.*)$`), fnc: cmdDescribeClusters, help: "'DescribeClusters s'  クラスター s の詳細情報"})

	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^DescribeTaskDefinition\s+(.*)$`), fnc: cmdDescribeTaskDefinition, help: "'DescribeTaskDefinition s' s タスク定義詳細"})

	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^DescribeInstances\s+(.*)$`), fnc: cmdDescribeInstances})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^DescribeInstances$`), fnc: cmdDescribeInstances})
	//	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^list\s+(.*)`), fnc: listImage})
	s.routes = &routes
}

func OnMention(msgs []string) string {

	var response = ""
	log.Printf("length=%d", len(msgs))
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

func cmdEnv(msgs []string) string {
	var response string = ""

	for key, _ := range awsEnv {
		response += key
		response += ", "
	}

	return (response)
}

func listImage(env AwsEnv, list_count int64) ([]*ecr.ImageDetail, error) {
	log.Printf("listImage")

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
		return nil, err
	} else {
		// sort by ImagePushedAt desc
		sort.Slice(result.ImageDetails, func(i, j int) bool {
			return result.ImageDetails[i].ImagePushedAt.After(*result.ImageDetails[j].ImagePushedAt)
		})

		return result.ImageDetails[:list_count], nil
	}
}

func cmdList(msgs []string) string {
	log.Printf("cmdList")

	var response = ""
	var list_count int64 = 3

	switch len(msgs) {
	case 1:

	case 2:
		i, err := strconv.ParseInt(msgs[1], 10, 64)
		if err == nil {
			list_count = i
		}

	default:
		return ("コマンド違うよ")
	}

	env, b := awsEnv[msgs[0]]
	if !b {
		return ("その環境は存在しないよ : " + msgs[0])
	}

	imageList, err := listImage(env, list_count)
	if err != nil {
		log.Printf("Error listImage: " + err.Error())

	} else {
		log.Printf("Success: listImage:")
		// Timezone変換？
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)

		if list_count > int64(len(imageList)) {
			list_count = int64(len(imageList))
		}
		for _, value := range imageList {
			response += *value.ImageTags[0] + "    " + value.ImagePushedAt.In(jst).String() + "\n"
		}
	}

	//s.rtm.SendMessage(s.rtm.NewOutgoingMessage(response, s.channelID))

	//	fmt.Println(response)

	return (response)
}

func deploy(env AwsEnv, img string) (err_ error) {

	describeTaskDefinitionInput := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(env.TaskDefinition),
	}

	describeTaskDefinitionOutput, err_ := s.ecs.DescribeTaskDefinition(describeTaskDefinitionInput)
	var container *ecs.ContainerDefinition
	if err_ != nil {
		return
	} else {
		for _, value := range describeTaskDefinitionOutput.TaskDefinition.ContainerDefinitions {
			//			if( *value.Name == "phpfpm"){
			if *value.Name == env.ContainerDefinition {
				container = value
			}
		}
	}

	idx := strings.LastIndex(*container.Image, ":")
	*container.Image = ((*container.Image)[:idx] + ":" + img)

	registerTaskDefinitionInput := &ecs.RegisterTaskDefinitionInput{
		//		ContainerDefinitions:   []*ecs.ContainerDefinition{container},
		ContainerDefinitions: describeTaskDefinitionOutput.TaskDefinition.ContainerDefinitions,
		Cpu:                  describeTaskDefinitionOutput.TaskDefinition.Cpu,
		ExecutionRoleArn:     describeTaskDefinitionOutput.TaskDefinition.ExecutionRoleArn,

		Family:                  describeTaskDefinitionOutput.TaskDefinition.Family,
		Memory:                  describeTaskDefinitionOutput.TaskDefinition.Memory,
		NetworkMode:             describeTaskDefinitionOutput.TaskDefinition.NetworkMode,
		PlacementConstraints:    describeTaskDefinitionOutput.TaskDefinition.PlacementConstraints,
		RequiresCompatibilities: describeTaskDefinitionOutput.TaskDefinition.RequiresCompatibilities,
		TaskRoleArn:             describeTaskDefinitionOutput.TaskDefinition.TaskRoleArn,
		Volumes:                 describeTaskDefinitionOutput.TaskDefinition.Volumes,
	}

	registerTaskDefinitionOutput, err_ := s.ecs.RegisterTaskDefinition(registerTaskDefinitionInput)

	if err_ != nil {
		return
	}

	taskDefinicationRevision := env.TaskDefinition + ":" + strconv.FormatInt(*registerTaskDefinitionOutput.TaskDefinition.Revision, 10)

	_, err_ = s.ecs.UpdateService(&ecs.UpdateServiceInput{
		Cluster:        aws.String(env.Cluster),
		TaskDefinition: &taskDefinicationRevision,
		Service:        aws.String(env.Service),
	})

	if err_ != nil {
		return
	}

	log.Printf(taskDefinicationRevision)

	return nil
}

func cmdDeploy(msgs []string) string {

	response := ""
	var img string
	var err error

	if len(msgs) < 1 {
		return ("コマンド違うよ")
	}

	env, b := awsEnv[msgs[1]]
	if !b {
		return ("その環境は存在しないよ")
	}

	switch len(msgs) {
	case 1:
		hoge, err := listImage(env, 1)
		if err == nil {
			img = *hoge[0].ImageTags[0]
		}

	case 2:
		img = msgs[1]

	default:
		return ("コマンド違うよ")
	}

	if err == nil {
		response = img + ": を Deployするね！\n"
		err = deploy(env, img)
	}

	if err != nil {
		return ("Deployエラー発生:" + err.Error())
	} else {
		response += "Deployしたよ。migrateは自分でしてね！"
		return (response)
	}
}

func cmdUp(msgs []string) string {
	response := ""

	// TODO: IP取得できるのでは？
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{aws.String(DEV_INSTANCE_ID)},
	}
	r, err := s.ec2.StartInstances(input)

	if err != nil {
		log.Printf("Error StartInstancesInput: " + err.Error())
	} else {
		log.Printf("%s\n", r)
		response += "おこすね！\n"
	}
	return (response)
}

func cmdDown(msgs []string) string {
	response := ""

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{aws.String(DEV_INSTANCE_ID)},
	}
	r, err := s.ec2.StopInstances(input)

	if err != nil {
		log.Printf("Error StopInstancesInput: " + err.Error())
	} else {
		log.Printf("%s\n", r)
		response += "かわりに殺しておくね～\n"
	}

	return (response)
}

func cmdDescribeTaskDefinition(msgs []string) string {
	var response string = ""

	input := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(msgs[0]),
	}

	output, err := s.ecs.DescribeTaskDefinition(input)

	if err == nil {
		response += "ContainerDefinition \n"
		for _, value := range output.TaskDefinition.ContainerDefinitions {
			response += *value.Name + "    " + *value.Image + "\n"
		}
	}

	return (response)
}

func cmdListClusters(msgs []string) string {

	var response string = ""
	var list_count int64 = 20

	if len(msgs) > 0 {

		//		list_count_, err := strconv.Atoi(msgs[0])
		i, err := strconv.ParseInt(msgs[0], 10, 64)
		if err == nil {
			list_count = i
		}
	}

	input := &ecs.ListClustersInput{
		MaxResults: aws.Int64(list_count),
	}

	output, err := s.ecs.ListClusters(input)

	if err == nil {
		response += "ClusterArns\n"
		for _, value := range output.ClusterArns {
			response += *value + "\n"
		}
	}

	return (response)
}

func cmdDescribeInstances(msgs []string) string {

	var response string = ""
	/*
		var instances  string = nil

		if( len(msgs) > 0){
			instances = msgs[0]
		}
	*/
	result, err := s.ec2.DescribeInstances(nil)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}
	return (response)
}
func cmdDescribeClusters(msg []string) string {

	var response string = ""
	//		var list_count int64 = 20
	msgs := strings.Split(msg[0], " ")

	var cls []*string

	for _, value := range msgs {
		cls = append(cls, aws.String(value))
	}

	input := &ecs.DescribeClustersInput{
		Clusters: cls,
	}

	output, err := s.ecs.DescribeClusters(input)

	if err == nil {
		response += "Clusters\n"
		for _, value := range output.Clusters {
			response += *value.ClusterName + "\n"
		}
	}

	return (response)
}
