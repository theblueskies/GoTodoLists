package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJSONMarshalling(t *testing.T) {
	td := Todos{
		ID:        5,
		ListID:    2,
		Name:      "Todo marshals",
		Notes:     "Todo Notes",
		DueDate:   time.Date(2049, 1, 25, 0, 0, 0, 0, time.UTC),
		Completed: true,
	}

	// Marshal the struct
	JSONify, err := json.Marshal(td)
	assert.Equal(t, nil, err)

	// Unmarshal the JSON and populate Todos{}
	var regeneratedTodo Todos
	err = json.Unmarshal(JSONify, &regeneratedTodo)
	assert.Equal(t, nil, err)

	assert.Equal(t, td.ID, regeneratedTodo.ID)
	assert.Equal(t, td.ListID, regeneratedTodo.ListID)
	assert.Equal(t, td.Name, regeneratedTodo.Name)
	assert.Equal(t, td.Notes, regeneratedTodo.Notes)
	assert.Equal(t, td.Completed, regeneratedTodo.Completed)

	// The date time needs to be checked more explicitly
	assert.Equal(t, 2049, regeneratedTodo.DueDate.Year())
	assert.Equal(t, "January", regeneratedTodo.DueDate.Month().String())
	assert.Equal(t, 25, regeneratedTodo.DueDate.Day())
}
