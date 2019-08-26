module GOSICK

go 1.12

replace util => ./util

replace slackUtil => ./util/slack

require (
	github.com/YukiMiyatake/GOSICK v0.0.0-20190724221603-25478903779c
	github.com/aws/aws-sdk-go v1.23.8
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/nlopes/slack v0.5.0
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
	golang.org/x/tools v0.0.0-20190826060629-95c3470cfb70 // indirect
	util v0.0.0-00010101000000-000000000000
)
