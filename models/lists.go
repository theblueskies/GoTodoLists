package models

// Lists is the model for list items
type Lists struct {
	ID   int64  `db:"id, primarykey, autoincrement"`
	Name string `db:"name" json:"name" binding:"required"`
}
