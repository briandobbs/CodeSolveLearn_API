package entity

type Article struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Author      string `db:"author" json:"author"`
	Body        string `db:"body" json:"body"`
}
