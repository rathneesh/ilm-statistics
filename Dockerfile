FROM golang:1.6

RUN go get github.com/gorilla/mux && go get gopkg.in/yaml.v2 && go get github.com/robfig/cron

ENV PATH $PATH:/go/bin:/usr/local/go/bin

COPY . /go/src/github.com/ilm-statistics/ilm-statistics

WORKDIR /go/src/github.com/ilm-statistics/ilm-statistics

EXPOSE 8084

RUN go build

CMD ./ilm-statistics
