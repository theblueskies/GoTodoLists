FROM golang:alpine

RUN apk add --no-cache git curl
RUN mkdir -p /go/src/github.com/theblueskies/GoTodoLists
WORKDIR /go/src/github.com/theblueskies/GoTodoLists

COPY . ./
EXPOSE 9000

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    go install && \
    dep ensure

RUN go build -o todo-service .

ENTRYPOINT ["./entrypoint.sh"]
