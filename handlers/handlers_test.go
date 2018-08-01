package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/theblueskies/GoTodoLists/models"
)

func TestCreateTodo(t *testing.T) {
	router := GetRouter()
	_ = models.GetDBMap()

	newTodo := models.Todo{Name: "NewTodo", Description: "NewDescription", ListID: 1}
	jsonTodo, err := json.Marshal(&newTodo)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonTodo))
	router.ServeHTTP(w, req)

	// expectedBody := map[string]string{
	// 	"error": "No Authorization header found",
	// }
	// expectedBodyJSON, _ := json.Marshal(expectedBody)
	// assert.Equal(t, 403, w.Code)
	// assert.Equal(t, expectedBodyJSON, w.Body.Bytes())

}
