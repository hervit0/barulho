FROM golang:1.13-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/hervit0/barulho
COPY . ./
RUN dep ensure

EXPOSE 5050

RUN go build app/main.go

CMD ["./main"]
