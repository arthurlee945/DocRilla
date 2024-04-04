package model

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          uint64         `json:"id" db:"id"`
	Endpoint    string         `json:"endpoint" db:"endpoint"`
	UserID      int            `json:"userID" db:"user_id"`
	Title       string         `json:"title" db:"title"`
	Description sql.NullString `json:"description" db:"description"`
	DocumentUrl string         `json:"documentUrl" db:"document_url"`
	Archived    bool           `json:"archived" db:"archived"`
	VisitedAt   sql.NullTime   `json:"vistedAt" db:"visited_at"`
	CreatedAt   time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time      `json:"updatedAt" db:"updated_at"`
	User        *User          `json:"user"`
	Fields      *[]Field       `json:"fields"`
}
