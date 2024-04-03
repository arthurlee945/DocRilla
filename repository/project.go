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

// 018ea1b1-b9ba-79af-81ce-81bae9930afa
func (pr *ProjectRepository) GetOverview(endpoint string) (*model.Project, error) {
	query := `
	SELECT endpoint, title, description, archived, created_at, visited_at 
	FROM project WHERE endpoint = $1
	`
	proj := model.Project{}
	err := pr.db.Get(&proj, query, endpoint)
	if err != nil {
		return nil, err
	}
	return &proj, nil
}

func (pr *ProjectRepository) GetDetail(endpoint string) (*model.Project, error) {
	proj := new(model.Project)
	fields := []model.Field{}
	projErr := pr.db.Get(proj, `SELECT * FROM project WHERE endpoint = $1`, endpoint)
	if projErr != nil {
		return nil, projErr
	}
	fieldErr := pr.db.Select(&fields, `SELECT * FROM field WHERE project_id = $1`, proj.ID)
	if fieldErr != nil {
		return nil, fieldErr
	}
	proj.Fields = fields
	return proj, nil
}

func (pr *ProjectRepository) Create() error {
	return nil
}
