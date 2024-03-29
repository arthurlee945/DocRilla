package model

import (
	"fmt"

	"github.com/arthurlee945/Docrilla/model/enum/content"
)

var ContentSchema = fmt.Sprintf(`
	CREATE TYPE IF NOT EXITS ContentType AS ENUM ('%v', '%v', '%v')

	CREATE TABLE IF NOT EXISTS Content (
		id SERIAL PRIMARY KEY,
		x1 NUMERIC NOT NULL,
		y1 NUMERIC NOT NULL,
		x2 NUMERIC NOT NULL,
		y2 NUMERIC NOT NULL,
		type 
		project_id INT NOT NULL,
		CONTRAINT fk_Project FOREIGN KEY(project_id) REFERENCES Project(id) ON DELETE CASCADE
	)
`, content.TEXT, content.NUMBER, content.IMAGE)

type Content struct {
	ContentID uint64 `db:"id"`
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	Type      content.ContentType
	Project   Project
}
