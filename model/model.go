package model

import "fmt"

func GetTablesSchema() string {
	return fmt.Sprintf(`
%v
%v
%v
%v
	`, updateTimestampFunc, userSchema, projectSchema, fieldSchema)
}
