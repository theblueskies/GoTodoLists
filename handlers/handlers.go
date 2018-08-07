package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/theblueskies/GoTodoLists/models"
)

// GetRouter registers the various handlers
func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.POST("/list", CreateList)
	r.POST("/todo", CreateTodo)
	r.DELETE("/todo/:todoID", DeleteTodo)
	r.PUT("/todo/:todoID", UpdateTodo)
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
	q := fmt.Sprintf(`INSERT INTO todos (id, name, notes, list_id, completed, due_date)
					  VALUES (nextval('todos_id_seq'), '%s', '%s', %d, %t, '%s') RETURNING id;`, td.Name, td.Notes, td.ListID, td.Completed, td.DueDate.Format(time.RFC3339))
	err := d.QueryRow(q).Scan(&tdID)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Todo could not be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message":   "Todo Created",
		"id":        tdID,
		"name":      td.Name,
		"notes":     td.Notes,
		"list_id":   td.ListID,
		"completed": td.Completed,
		"due_date":  td.DueDate,
	})
}

// DeleteTodo removes a Todo permanently
func DeleteTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	q := fmt.Sprintf(`DELETE FROM todos WHERE id = %s;`, todoID)
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

// UpdateTodo updates a Todo
func UpdateTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	var td models.Todos
	if err := c.ShouldBindWith(&td, binding.JSON); err != nil {
		c.JSON(500, gin.H{
			"error":   "Error",
			"message": err.Error(),
		})
		return
	}

	db := models.GetDBMap()
	q := fmt.Sprintf(`UPDATE todos SET list_id=%d, name='%s', notes='%s', completed=%v, due_date='%s'
					  WHERE id=%s;`, td.ListID, td.Name, td.Notes, td.Completed, td.DueDate.Format(time.RFC3339), todoID)

	_, err := db.Exec(q)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Todo could not be updated",
			"error":   err.Error(),
		})
	}

	responseJSON, err := json.Marshal(td)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(responseJSON)
}

// GetTodos gets Todos based on the query params: completed and/or name
func GetTodos(c *gin.Context) {
	var td []models.Todos
	completed := c.DefaultQuery("completed", "true")
	tdName := c.Query("name")

	q := fmt.Sprintf(`SELECT * FROM todos WHERE completed=%s`, completed)
	if tdName != "" {
		q += fmt.Sprintf(` AND name ILIKE '%s%%'`, tdName)
	}
	q += ";"

	db := models.GetDBMap()
	_, err := db.Select(&td, q)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Todo could not be fetched",
			"error":   err.Error(),
		})
	}

	responseJSON, err := json.Marshal(td)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(responseJSON)
}
