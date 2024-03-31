package model

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          uint64
	Endpoint    string
	UserID      int `db:"user_id"`
	Title       string
	Description sql.NullString
	DocumentUrl string `db:"document_url"`
	Archived    bool
	VisitedAt   sql.NullTime `db:"visited_at"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	User        User
	Fields      []Field
}
