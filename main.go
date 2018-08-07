package main

import (
	"fmt"
	"os"

	"github.com/theblueskies/GoTodoLists/handlers"
	"github.com/theblueskies/GoTodoLists/models"
)

func main() {
	_ = models.GetDBMap()
	r := handlers.GetRouter()
	port := ":" + os.Getenv("PORT")
	fmt.Println(port)
	r.Run(":9000")
}
