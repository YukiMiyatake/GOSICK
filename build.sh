#TODO: makefile
cd plugins/echo
go build --buildmode=plugin
cd ../..

cd plugins/aws
go build --buildmode=plugin
cd ../..

cd plugins/sqs
go build --buildmode=plugin
cd ../..

go build

