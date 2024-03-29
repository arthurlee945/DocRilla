package model

import "time"

type Default struct {
	ID        uint64
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
