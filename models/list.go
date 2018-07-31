package models

// List is the model for list items
type List struct {
	ID   int64  `db:"list_id"`
	Name string `db:"name"`
}
