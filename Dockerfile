FROM golang:1.9.2-alpine

ENV GOBIN=$GOPATH/bin
WORKDIR /chaps
RUN mkdir /root/.aws
COPY .credentials /root/.aws/
COPY . .
RUN apk add --no-cache git gcc musl-dev
RUN \
	mkdir -p /aws && \
	apk -Uuv add groff less python py-pip && \
	pip install awscli && \
	apk --purge -v del py-pip && \
	rm /var/cache/apk/*


RUN go get





