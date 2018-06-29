ChatopsBOT

---- 設定
環境変数を2個使っている
.env: 主にSlack Key
.credentials: AWS設定

export PWD="$(echo $('pwd') | sed -e 's/^\/mnt//')"
export BOT_TOKEN=
export BOT_NAME=BOT名
export CHANNEL_ID=チャンネルID

export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=
export AWS_DEFAULT_REGION=
export AWS_DEFAULT_OUTPUT=
export AWS_DEFAULT_PROFILE=


また AWSの環境設定は下記
aws.json

---- Dockerを起動しシェル起動
$ docker build
$ ./docker_run

---- Build
# build.sh

---- Run
# source .env
で環境変数を反映
Buildで作成した実行ファイルを起動すればOK




---- Deploy
まだない




---- コマンド
Slack上で 
BOT名 aws help
をするとHelp表示される