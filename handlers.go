package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/theblueskies/GoTodoLists/models"
)

func createList(c *gin.Context) {
	_ = models.InitDb()
	errBody := 5
	// var out []byte
	// Log the error body if present
	if errBody != 5 {
		// out, _ = json.Marshal(errBody)
		fmt.Println("body present")
	} else {
		fmt.Println("body absent")
	}

	c.JSON(201, gin.H{
		"message": "pong"})

}
