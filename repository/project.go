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

func (pr *ProjectRepository) GetOverview(id string) (*model.Project, error) {
	query := `
	SELECT endpoint, title, description, archived, created_at, visited_at 
	FROM project where id = ?
	`
	var project *model.Project
	err := pr.db.Get(project, query, id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (pr *ProjectRepository) GetDetail(id string) (*model.Project, error) {
	var project *model.Project
	var fields []model.Field
	projErr := pr.db.Get(project, `SELECT * FROM project WHERE id = ?`, id)
	if projErr != nil {
		return nil, projErr
	}
	fieldErr := pr.db.Select(&fields, `SELECT * FROM field WHERE project_id = ?`, id)
	if fieldErr != nil {
		return nil, fieldErr
	}
	return project, nil
}

func (pr *ProjectRepository) Create() error {
	return nil
}
