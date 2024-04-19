package model

import "github.com/arthurlee945/Docrilla/internal/model/enum/field"

type Field struct {
	ID        *uint64     `json:"id" db:"id"`
	UUID      *string     `json:"uuid" db:"uuid"`
	ProjectID *uint64     `json:"projectID" db:"project_id"`
	X1        *float64    `json:"x1" db:"x1"`
	Y1        *float64    `json:"y1" db:"y1"`
	X2        *float64    `json:"x2" db:"x2"`
	Y2        *float64    `json:"y2" db:"y2"`
	Page      *uint32     `json:"page" db:"page"`
	Type      *field.Type `json:"type" db:"type"`
}
