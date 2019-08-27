FROM golang:1.12.7


CMD go install github.com/golangci/golangci-lint/cmd/golangci-lint

WORKDIR /app
#COPY src .
#CMD ./build.sh
#CMD ls

#CMD ./GOSICK

