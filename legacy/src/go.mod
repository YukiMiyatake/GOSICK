module GOSICK

go 1.12

replace util => ./util

replace InputPlugins => ./InputPlugins

replace Slack => ./InputPlugins/Slack

require (
	Slack v0.0.0-00010101000000-000000000000
	github.com/YukiMiyatake/GOSICK v0.0.0-20190724221603-25478903779c
	github.com/aws/aws-sdk-go v1.23.8
	github.com/golangci/golangci-lint v1.17.1 // indirect
	github.com/nlopes/slack v0.6.0
	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472 // indirect
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422 // indirect
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/sys v0.0.0-20190830080133-08d80c9d36de // indirect
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/tools v0.0.0-20190830082254-f340ed3ae274 // indirect
	util v0.0.0-00010101000000-000000000000
)
