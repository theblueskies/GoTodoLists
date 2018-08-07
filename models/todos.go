package models

import (
	"encoding/json"
	"strings"
	"time"
)

// Todos model
type Todos struct {
	ID        int64     `db:"id, primarykey, autoincrement" form:"id"`
	ListID    int64     `db:"list_id" binding:"required" form:"list_id"`
	Name      string    `db:"name" binding:"required" form:"name"`
	Notes     string    `db:"notes" binding:"required"`
	DueDate   time.Time `db:"due_date"`
	Completed bool      `db:"completed" form:"completed"`
}

//UnmarshalJSON implements function to Unmarshal
func (t *Todos) UnmarshalJSON(j []byte) error {
	var rawStrings map[string]interface{}

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "id" {
			t.ID = int64(v.(float64))
		}
		if strings.ToLower(k) == "name" {
			t.Name = v.(string)
		}
		if strings.ToLower(k) == "notes" {
			t.Notes = v.(string)
		}
		if strings.ToLower(k) == "completed" {
			t.Completed = v.(bool)
		}
		if strings.ToLower(k) == "list_id" {
			t.ListID = int64(v.(float64))
		}

		if strings.ToLower(k) == "due_date" {
			date, err := time.Parse(time.RFC3339, v.(string))
			if err != nil {
				return err
			}
			t.DueDate = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
		}
	}

	return nil
}

//MarshalJSON implements function to Marshal
func (t Todos) MarshalJSON() ([]byte, error) {
	basicTodo := struct {
		ID        int64  `json:"id"`
		ListID    int64  `json:"list_id"`
		Name      string `json:"name"`
		Notes     string `json:"notes"`
		DueDate   string `json:"due_date"`
		Completed bool   `json:"completed"`
	}{
		ID:        t.ID,
		ListID:    t.ListID,
		Name:      t.Name,
		Notes:     t.Notes,
		DueDate:   t.DueDate.Format(time.RFC3339),
		Completed: t.Completed,
	}

	return json.Marshal(basicTodo)
}
