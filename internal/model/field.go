package model

import "github.com/arthurlee945/Docrilla/internal/model/enum/field"

type Field struct {
	ID        *uint64     `json:"-" db:"id"`
	UUID      *string     `json:"uuid" db:"uuid"`
	ProjectID *string     `json:"projectID" db:"project_id"`
	X         *float64    `json:"x" db:"x"`
	Y         *float64    `json:"y" db:"y"`
	Width     *float64    `json:"width" db:"width"`
	Height    *float64    `json:"height" db:"height"`
	Page      *uint32     `json:"page" db:"page"`
	Type      *field.Type `json:"type" db:"type"`
}
