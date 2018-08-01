package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/theblueskies/GoTodoLists/models"
)

func main() {
	_ = models.GetDBMap()
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
