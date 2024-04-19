package model

type Endpoint struct {
	ID        *uint64 `json:"id" db:"id"`
	Route     *string `json:"route" db:"route"`
	Token     *string `json:"token" db:"token"`
	ProjectId *uint64 `json:"projectId" db:"project_id"`
}
