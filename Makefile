.PHONY:
	build

LOCAL_REPO_PATH=$(GOPATH)/src/github.com/theblueskies/GoTodoLists
CONTAINER_GOPATH=/go/src/github.com/theblueskies/GoTodoLists

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo-service .
	docker build --no-cache -t todo-service .

pgb:
	docker run -p 5432:5432 --name postgres postgres:10.3
