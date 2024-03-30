package model

import (
	"fmt"

	"github.com/arthurlee945/Docrilla/model/enum/field"
)

var fieldSchema = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS field (
	id SERIAL PRIMARY KEY,
	x1 NUMERIC NOT NULL,
	y1 NUMERIC NOT NULL,
	x2 NUMERIC NOT NULL,
	y2 NUMERIC NOT NULL,
	page INT NOT NULL,
	type TEXT NOT NULL,
	field_id TEXT NOT NULL,
	value TEXT NOT NULL,
	project_id INT,
	CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES project(id),
	CONSTRAINT check_type check (type in ('%v', '%v', '%v'));
);
`, field.TEXT, field.NUMBER, field.IMAGE)

type Field struct {
	ContentID uint64 `db:"id"`
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	Page      uint32
	Type      field.FieldType
	FieldId   string `db:"field_id"`
	Value     string
	Project   Project
}
