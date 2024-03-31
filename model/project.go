package model

import "time"

type Project struct {
	ProjectID   uint64
	Endpoint    string
	UserID      int `db:"user_id"`
	Title       string
	Description string
	DocumentUrl string `db:"document_url"`
	Archived    bool
	VisitedAt   time.Time `db:"visited_at"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	User        User
	Fields      []Field
}
