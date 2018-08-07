package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theblueskies/GoTodoLists/models"
)

func TestCreateListSuccess(t *testing.T) {
	router := GetRouter()

	// Create a List
	newList := models.Lists{Name: "New List"}
	jsonList, _ := json.Marshal(newList)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/list", bytes.NewBuffer(jsonList))
	router.ServeHTTP(w, req)

	var nl models.Lists
	_ = json.Unmarshal(w.Body.Bytes(), &nl)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, newList.Name, nl.Name)
}

func TestCreateTodoSuccess(t *testing.T) {
	router := GetRouter()

	// Create a List first
	newList := models.Lists{Name: "New List"}
	jsonList, _ := json.Marshal(newList)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/list", bytes.NewBuffer(jsonList))
	router.ServeHTTP(w, req)
	var ls models.Lists
	_ = json.Unmarshal(w.Body.Bytes(), &ls)

	// Create  a Todo attached to the list
	v := httptest.NewRecorder()
	newTodo := models.Todos{Name: "NewTodo", Notes: "New Notes", ListID: ls.ID}
	jsonTodo, err := json.Marshal(newTodo)
	if err != nil {
		panic(err)
	}
	req, _ = http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonTodo))
	router.ServeHTTP(v, req)

	var td models.Todos
	_ = json.Unmarshal(v.Body.Bytes(), &td)

	assert.Equal(t, newTodo.Name, td.Name)
	assert.Equal(t, newTodo.Notes, td.Notes)
	assert.Equal(t, ls.ID, td.ListID)
	assert.NotZero(t, td.ID)
	assert.Equal(t, 201, v.Code)
}
