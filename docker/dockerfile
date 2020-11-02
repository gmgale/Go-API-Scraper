FROM golang:latest
ADD . /go/src/github.com/gmgale/result_task
WORKDIR /go/src/github.com/gmgale/result_task
RUN go get github.com/gorilla/mux
RUN go install
ENTRYPOINT ["/go/bin/api-test"]
EXPOSE 8080