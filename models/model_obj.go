package models

// Lists is the model for list items
type Lists struct {
	ID   int64  `db:"id, primarykey, autoincrement"`
	Name string `db:"name" json:"name" binding:"required"`
}

// Todos model
type Todos struct {
	ID     int64  `db:"id, primarykey, autoincrement" json:"id" form:"id"`
	ListID int64  `db:"list_id" json:"list_id" binding:"required" form:"list_id"`
	Name   string `db:"name" json:"name" binding:"required" form:"name"`
	Notes  string `db:"notes" json:"notes" binding:"required"`
	// DueDate   time.Time `db:"due_date" json:"due_date"`
	Completed bool `db:"completed" json:"completed" form:"completed"`
}
