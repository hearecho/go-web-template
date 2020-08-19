FROM golang:latest

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

WORKDIR  /go-web
COPY . .

RUN go build -o app .

EXPOSE 8888
ENTRYPOINT ["./app"]
