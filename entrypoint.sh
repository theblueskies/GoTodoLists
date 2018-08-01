#!/bin/sh

# wait for Postgresql server to start
sleep 10

# start development server on public ip interface, on port 9000
./todo-service
