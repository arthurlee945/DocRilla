package repo

import (
	"github.com/arthurlee945/Docrilla/model"
	"github.com/jmoiron/sqlx"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{
		db,
	}
}

// 3b73b142-13c7-4e9b-b89c-22a4b6d45c9f
func (pr *ProjectRepository) GetOverview(endpoint string) (*model.Project, error) {
	query := `
	SELECT endpoint, title, description, archived, created_at, visited_at 
	FROM project where endpoint = ?
	`
	var proj *model.Project
	err := pr.db.Get(proj, query, endpoint)
	if err != nil {
		return nil, err
	}
	return proj, nil
}

func (pr *ProjectRepository) GetDetail(endpoint string) (*model.Project, error) {
	var proj *model.Project
	var fields []model.Field
	projErr := pr.db.Get(proj, `SELECT * FROM project WHERE endpoint = ?`, endpoint)
	if projErr != nil {
		return nil, projErr
	}
	fieldErr := pr.db.Select(&fields, `SELECT * FROM field WHERE project_id = ?`, proj)
	if fieldErr != nil {
		return nil, fieldErr
	}
	return proj, nil
}

func (pr *ProjectRepository) Create() error {
	return nil
}
