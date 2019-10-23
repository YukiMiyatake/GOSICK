#ChatopsBOT

サーバレスに生まれ変わらせる！
コンテナとの併用は行えればいいが考えない


#設定ファイル
##slack設定
sample/slack.json.orgをもとにルートディレクトリにslack.jsonを設置する


環境変数を2個使っている
.env: 主にSlack Key
.credentials: AWS設定

export PWD="$(echo $('pwd') | sed -e 's/^\/mnt//')"

export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=
export AWS_DEFAULT_REGION=
export AWS_DEFAULT_OUTPUT=
export AWS_DEFAULT_PROFILE=


また AWSの環境設定は下記
aws.json

---------------------------------------- 起動方法
---- Dockerを起動しシェル起動
プロジェクトルートディレクトリにて実行
$ docker build -t gosick ./
$ sh ./scripts/docker_run

---- Dockerを起動しシェル起動（docker-composeの場合）
ビルド&起動
$ docker-compose up --build -d

各コンテナログイン
$ sh ./scripts/docker_app_run
$ sh ./scripts/docker_redis_run

停止
$ docker-compose stop

停止およびコンテナ削除したい場合はこちら
$ docker-compose down

---- Build
プロジェクトルートディレクトリにて実行
# sh ./scripts/build.sh

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