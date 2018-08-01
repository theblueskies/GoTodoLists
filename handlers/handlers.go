package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/theblueskies/GoTodoLists/models"
)

// GetRouter registers the various handlers
func GetRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/list", CreateList)
	r.POST("/todo", CreateTodo)
	return r
}

// CreateList creates a list
func CreateList(c *gin.Context) {
	var lb models.Lists
	if err := c.ShouldBindWith(&lb, binding.JSON); err != nil {
		c.JSON(500, gin.H{
			"error":   "Error",
			"message": err.Error(),
		})
		return
	}

	d := models.GetDBMap()
	if err := d.Insert(&lb); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "List could not be created",
			"error":   err.Error(),
		})
	}

	c.JSON(201, gin.H{
		"message": "List Created",
		"id":      lb.ID,
		"name":    lb.Name,
	})
}

// CreateTodo creates the Todo and attaches it to the related List
func CreateTodo(c *gin.Context) {
	var td models.Todo
	if err := c.ShouldBindWith(&td, binding.JSON); err != nil {
		c.JSON(500, gin.H{
			"error":   "Error",
			"message": err.Error(),
		})
		return
	}

	d := models.GetDBMap()
	fmt.Println("ZORRO: about to add to db")
	q := fmt.Sprintf(`INSERT INTO todos (name, description, list_id) VALUES ("%s", "%s", %d);`, td.Name, td.Description, td.ListID)
	r, err := d.Exec(q)
	if err != nil {
		a := err.Error()
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Todo could not be created",
			"error":   a,
		})
		return
	}
	fmt.Println(r)
	c.JSON(201, gin.H{
		"message":     "Todo Created",
		"id":          td.ID,
		"name":        td.Name,
		"description": td.Description,
		"list":        td.ListID,
	})
}
