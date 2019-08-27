#TODO: makefile

if [ $TEST = "test" ]; then
    #cd plugins/memo
    #go build ${TEST} --buildmode=plugin
    #cd ../..

    cd plugins/echo
#    golint
    go test -v ./... --buildmode=plugin
    cd ../..

    #cd plugins/cmd
    #go build --buildmode=plugin
    #cd ../..

    #cd plugins/aws
    #go build --buildmode=plugin
    #cd ../..

    #cd plugins/sqs
    #go build --buildmode=plugin
    #cd ../..

#    golint
    go test -v ./...
else
    #cd plugins/memo
    #go build ${TEST} --buildmode=plugin
    #cd ../..

    cd plugins/echo
    go build --buildmode=plugin
    cd ../..

    #cd plugins/cmd
    #go build --buildmode=plugin
    #cd ../..

    #cd plugins/aws
    #go build --buildmode=plugin
    #cd ../..

    #cd plugins/sqs
    #go build --buildmode=plugin
    #cd ../..

    go build

fi

