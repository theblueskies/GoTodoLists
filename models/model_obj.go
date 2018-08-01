package models

// Lists is the model for list items
type Lists struct {
	ID   int64  `db:"id, primarykey, autoincrement"`
	Name string `db:"name" json:"name" binding:"required"`
}

// Todo model
type Todo struct {
	ID          int64  `db:"id, primarykey, autoincrement"`
	Name        string `db:"name" json:"name" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
	ListID      int64  `db:"list_id" json:"list_id"`
}
