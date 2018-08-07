package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/theblueskies/GoTodoLists/models"
)

func createList(router *gin.Engine) *httptest.ResponseRecorder {
	newList := models.Lists{Name: "New List"}
	jsonList, _ := json.Marshal(newList)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/list", bytes.NewBuffer(jsonList))
	router.ServeHTTP(w, req)
	return w
}

func createTodo(router *gin.Engine) (models.Lists, models.Todos, int) {
	w := createList(router)
	var ls models.Lists
	_ = json.Unmarshal(w.Body.Bytes(), &ls)

	w = httptest.NewRecorder()
	newTodo := models.Todos{
		Name:      "NewTodo",
		Notes:     "New Notes",
		ListID:    ls.ID,
		Completed: true,
		DueDate:   time.Now(),
	}
	jsonTodo, err := json.Marshal(newTodo)
	if err != nil {
		panic(err)
	}
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonTodo))
	router.ServeHTTP(w, req)
	var td models.Todos
	_ = json.Unmarshal(w.Body.Bytes(), &td)

	return ls, td, w.Code
}

func TestHealth(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	var b map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &b)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "ok", b["status"])

}

func TestCreateListSuccess(t *testing.T) {
	router := GetRouter()

	// Create a List
	w := createList(router)
	var nl models.Lists
	_ = json.Unmarshal(w.Body.Bytes(), &nl)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "New List", nl.Name)
}

func TestCreateTodoSuccess(t *testing.T) {
	router := GetRouter()

	// Create  a Todo attached to the list
	ls, td, code := createTodo(router)

	assert.Equal(t, "NewTodo", td.Name)
	assert.Equal(t, "New Notes", td.Notes)
	assert.Equal(t, ls.ID, td.ListID)
	assert.NotZero(t, td.ID)
	assert.Equal(t, true, td.Completed)
	assert.Equal(t, 201, code)
}

func TestDeleteTodo(t *testing.T) {
	defer teardown()
	router := GetRouter()

	// Create  a List and a Todo attached to it
	_, td, _ := createTodo(router)

	// Delete todo
	uri := fmt.Sprintf("/todo/%d", td.ID)
	req, _ := http.NewRequest("DELETE", uri, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var b map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &b)

	assert.Equal(t, 204, w.Code)
	assert.Equal(t, 0, len(b))
}

func TestUpdateTodo(t *testing.T) {
	defer teardown()
	router := GetRouter()

	// Create  a List and a Todo attached to it
	ls, td, _ := createTodo(router)

	// Update Todo
	td.Name = "Updated Name"
	td.Completed = true
	td.Notes = "Updated Notes"
	td.DueDate = time.Date(2049, 1, 1, 0, 0, 0, 0, time.UTC)
	jsonTodo, err := json.Marshal(td)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	uri := fmt.Sprintf("/todo/%d", td.ID)
	req, _ := http.NewRequest("PUT", uri, bytes.NewBuffer(jsonTodo))
	router.ServeHTTP(w, req)

	var tdResponse models.Todos
	err = json.Unmarshal(w.Body.Bytes(), &tdResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, td.ID, tdResponse.ID)
	assert.Equal(t, ls.ID, tdResponse.ListID)
	assert.Equal(t, "Updated Name", tdResponse.Name)
	assert.Equal(t, "Updated Notes", tdResponse.Notes)
	assert.Equal(t, td.DueDate, tdResponse.DueDate)
	assert.Equal(t, true, tdResponse.Completed)
}

func TestGetTodo(t *testing.T) {
	defer teardown()
	router := GetRouter()

	// Create  a List and a Todo attached to it
	_, _, _ = createTodo(router)

	// Search with "completed" and "name". These fields can be used individually or together
	w := httptest.NewRecorder()
	uri := "/todo?completed=true&name=NewT"
	req, _ := http.NewRequest("GET", uri, nil)
	router.ServeHTTP(w, req)

	var data []models.Todos
	_ = json.Unmarshal(w.Body.Bytes(), &data)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(data))
	assert.Equal(t, true, data[0].Completed)
}

func teardown() {
	db := models.GetDBMap()
	_, err := db.Exec(`DELETE FROM TODOS; DELETE FROM LISTS;`)
	if err != nil {
		panic(err)
	}
}
