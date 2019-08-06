module GOSICK

go 1.12

replace util => ./util

replace slackUtil => ./util/slack

require (
	github.com/YukiMiyatake/GOSICK v0.0.0-20190724221603-25478903779c
	github.com/nlopes/slack v0.5.0
	util v0.0.0-00010101000000-000000000000
)
