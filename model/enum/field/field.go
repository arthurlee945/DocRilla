package field

type FieldType string

const (
	TEXT   FieldType = "TEXT"
	IMAGE  FieldType = "IMAGE"
	NUMBER FieldType = "NUMBER"
)

var FieldTypes = [3]FieldType{TEXT, IMAGE, NUMBER}
