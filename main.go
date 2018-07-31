package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := getRouter()
	port := ":" + os.Getenv("PORT")
	fmt.Println(port)
	r.Run(":9000")
}

func getRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/list", createList)
	return r
}
