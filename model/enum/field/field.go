package field

type Type string

const (
	TEXT   Type = "TEXT"
	IMAGE  Type = "IMAGE"
	NUMBER Type = "NUMBER"
)

var FieldTypes = [3]Type{TEXT, IMAGE, NUMBER}
