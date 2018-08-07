.PHONY:
	build
	removePG
	testfresh

LOCAL_REPO_PATH=$(GOPATH)/src/github.com/theblueskies/GoTodoLists
CONTAINER_GOPATH=/go/src/github.com/theblueskies/GoTodoLists

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo-service .
	docker build --no-cache -t todo-service .

removePG:
	docker stop postgres
	docker rm -f postgres

pgb:
	docker run -d -p 5432:5432 --name postgres postgres:10.3

runtest:
	# Sleep 5s so that PostgreSQL DB can come up.
	sleep 5
	go clean -testcache
	ENV=local go test ./handlers

test: pgb runtest

testfresh: removePG pgb runtest
