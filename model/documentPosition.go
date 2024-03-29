package model

type DocumentPosition struct {
	ID      string  `db:"id"`
	Project Project `json:"project"`
}
