package model

import "github.com/arthurlee945/Docrilla/model/enum/field"

type Field struct {
	FieldId string `db:"id"`
	X1      float64
	Y1      float64
	X2      float64
	Y2      float64
	Page    uint32
	Type    field.Type
	Value   string
	Project Project
}
