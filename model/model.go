package model

import "fmt"

func GetTablesSchema() string {
	return fmt.Sprintf(`
	%v

	%v

	%v
	`, ProjectSchema, FieldSchema, UserSchema)
}
