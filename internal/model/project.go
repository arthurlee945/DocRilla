package model

import (
	"time"

	"github.com/arthurlee945/Docrilla/internal/model/null"
)

type Project struct {
	ID          uint64      `json:"id" db:"id"`
	UUID        string      `json:"uuid" db:"uuid"`
	UserID      uint64      `json:"userID" db:"user_id"`
	Title       string      `json:"title" db:"title"`
	Description null.String `json:"description" db:"description"`
	DocumentUrl string      `json:"documentUrl" db:"document_url"`
	Archived    bool        `json:"archived" db:"archived"`
	VisitedAt   null.Time   `json:"visitedAt" db:"visited_at"`
	CreatedAt   time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time   `json:"updatedAt" db:"updated_at"`
	User        *User       `json:"user"`
	Fields      []Field     `json:"fields"`
}
