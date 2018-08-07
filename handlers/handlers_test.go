package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/theblueskies/GoTodoLists/models"
)

func createlist(router *gin.Engine) *httptest.ResponseRecorder {
	newList := models.Lists{Name: "New List"}
	jsonList, _ := json.Marshal(newList)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/list", bytes.NewBuffer(jsonList))
	router.ServeHTTP(w, req)
	return w
}

func createtodo(router *gin.Engine, ls models.Lists) *httptest.ResponseRecorder {
	v := httptest.NewRecorder()
	newTodo := models.Todos{Name: "NewTodo", Notes: "New Notes", ListID: ls.ID}
	jsonTodo, err := json.Marshal(newTodo)
	if err != nil {
		panic(err)
	}
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonTodo))
	router.ServeHTTP(v, req)
	return v
}

func TestCreateListSuccess(t *testing.T) {
	router := GetRouter()

	// Create a List
	w := createlist(router)
	var nl models.Lists
	_ = json.Unmarshal(w.Body.Bytes(), &nl)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "New List", nl.Name)
}

func TestCreateTodoSuccess(t *testing.T) {
	router := GetRouter()

	// Create a List first
	w := createlist(router)
	var ls models.Lists
	_ = json.Unmarshal(w.Body.Bytes(), &ls)

	// Create  a Todo attached to the list
	v := createtodo(router, ls)
	var td models.Todos
	_ = json.Unmarshal(v.Body.Bytes(), &td)

	assert.Equal(t, "NewTodo", td.Name)
	assert.Equal(t, "New Notes", td.Notes)
	assert.Equal(t, ls.ID, td.ListID)
	assert.NotZero(t, td.ID)
	assert.Equal(t, false, td.Completed)
	assert.Equal(t, 201, v.Code)
}

func TestDeleteTodo(t *testing.T) {
	router := GetRouter()

	// Create a List first
	w := createlist(router)
	var ls models.Lists
	_ = json.Unmarshal(w.Body.Bytes(), &ls)

	// Create  a Todo attached to the list
	w = createtodo(router, ls)
	var td models.Todos
	_ = json.Unmarshal(w.Body.Bytes(), &td)

	// Delete todo
	uri := fmt.Sprintf("/todo/%d", td.ID)
	req, _ := http.NewRequest("DELETE", uri, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var b map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &b)

	assert.Equal(t, 204, w.Code)
	assert.Equal(t, 0, len(b))
}

func TestUpdateTodo(t *testing.T) {
	router := GetRouter()

	// Create a List first
	w := createlist(router)
	var ls models.Lists
	_ = json.Unmarshal(w.Body.Bytes(), &ls)

	// Create  a Todo attached to the list
	w = createtodo(router, ls)
	var td models.Todos
	_ = json.Unmarshal(w.Body.Bytes(), &td)

	// Update Todo
	td.Name = "Updated Name"
	td.Completed = true
	td.Notes = "Updated Notes"
	jsonTodo, err := json.Marshal(td)
	if err != nil {
		panic(err)
	}

	w = httptest.NewRecorder()
	uri := fmt.Sprintf("/todo/%d", td.ID)
	req, _ := http.NewRequest("PUT", uri, bytes.NewBuffer(jsonTodo))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Updated Name", td.Name)
	assert.Equal(t, "Updated Notes", td.Notes)
	assert.Equal(t, ls.ID, td.ListID)
	assert.NotZero(t, td.ID)
	assert.Equal(t, true, td.Completed)
}
