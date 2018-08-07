package models

// Lists is the model for list items
type Lists struct {
	ID   int64  `db:"id, primarykey, autoincrement"`
	Name string `db:"name" json:"name" binding:"required"`
}

// Todos model
type Todos struct {
	ID     int64  `db:"id, primarykey, autoincrement"`
	ListID int64  `db:"list_id" json:"list_id" binding:"required"`
	Name   string `db:"name" json:"name" binding:"required"`
	Notes  string `db:"notes" json:"notes" binding:"required"`
}
