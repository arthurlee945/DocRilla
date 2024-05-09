package model

import (
	"time"
)

type Project struct {
	ID          *uint64    `json:"-" db:"id"`
	UUID        *string    `json:"uuid" db:"uuid"`
	UserID      *uint64    `json:"userID" db:"user_id"`
	Route       *string    `json:"route" db:"route"`
	Token       *string    `json:"token" db:"token"`
	Title       *string    `json:"title" db:"title"`
	Description *string    `json:"description" db:"description"`
	DocumentUrl *string    `json:"documentUrl" db:"document_url"`
	Archived    *bool      `json:"archived" db:"archived"`
	VisitedAt   *time.Time `json:"visitedAt" db:"visited_at"`
	CreatedAt   *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
	Fields      []Field    `json:"fields"`
}
