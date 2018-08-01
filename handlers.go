package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func createList(c *gin.Context) {
	errBody := 5
	if errBody != 5 {
		// out, _ = json.Marshal(errBody)
		fmt.Println("body present")
	} else {
		fmt.Println("body absent")
	}

	c.JSON(201, gin.H{
		"message": "pong"})

}
