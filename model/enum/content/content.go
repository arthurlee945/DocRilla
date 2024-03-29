package content

type ContentType string

const (
	TEXT   ContentType = "TEXT"
	IMAGE  ContentType = "IMAGE"
	NUMBER ContentType = "NUMBER"
)

var ContentTypes = [3]ContentType{TEXT, IMAGE, NUMBER}
