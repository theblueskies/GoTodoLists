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
	r.DELETE("/todo/:todoID", DeleteTodo)
	r.PATCH("/complete/:todoID", UpdateTodo)
	r.GET("/todo", GetTodos)
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
	var td models.Todos
	if err := c.ShouldBindWith(&td, binding.JSON); err != nil {
		c.JSON(500, gin.H{
			"error":   "Error",
			"message": err.Error(),
		})
		return
	}

	d := models.GetDBMap()
	var tdID int
	q := fmt.Sprintf(`INSERT INTO todos (id, name, notes, list_id)
					  VALUES (nextval('todos_id_seq'), '%s', '%s', %d) RETURNING id;`, td.Name, td.Notes, td.ListID)
	err := d.QueryRow(q).Scan(&tdID)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Todo could not be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Todo Created",
		"id":      tdID,
		"name":    td.Name,
		"notes":   td.Notes,
		"list_id": td.ListID,
	})
}

// DeleteTodo removes a Todo permanently
func DeleteTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	q := fmt.Sprintf(`DELETE FROM todos WHERE id = %s`, todoID)
	db := models.GetDBMap()

	_, err := db.Exec(q)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Todo could not be deleted",
			"error":   err.Error(),
		})
	}

	c.JSON(204, gin.H{
		"message": "",
	})
}

func UpdateTodo(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": nil,
	})
}

func GetTodos(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": nil,
	})
}
