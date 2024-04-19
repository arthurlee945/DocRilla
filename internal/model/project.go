package model

import (
	"time"
)

type Project struct {
	ID          *uint64    `json:"id" db:"id"`
	UUID        *string    `json:"uuid" db:"uuid"`
	UserID      *uint64    `json:"userID" db:"user_id"`
	Title       *string    `json:"title" db:"title"`
	Description *string    `json:"description" db:"description"`
	DocumentUrl *string    `json:"documentUrl" db:"document_url"`
	Archived    *bool      `json:"archived" db:"archived"`
	VisitedAt   *time.Time `json:"visitedAt" db:"visited_at"`
	CreatedAt   *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
	Endpoint    *Endpoint  `json:"endpoint" db:"endpoint"`
	Fields      []Field    `json:"fields"`
}
