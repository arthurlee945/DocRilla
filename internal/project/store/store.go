package store

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db,
	}
}

// 018ea1b1-b9ba-79af-81ce-81bae9930afa
func (pr *Store) GetProjectOverview(ctx context.Context, user *model.User, endpoint string) (*model.Project, error) {
	query := `
	SELECT endpoint, title, description, archived, created_at, visited_at 
	FROM project WHERE endpoint = $1 AND user_id = $2
	`
	proj := new(model.Project)
	err := pr.db.GetContext(ctx, proj, query, endpoint, user.ID)
	if err != nil {
		return nil, err
	}
	return proj, nil
}

func (pr *Store) GetProjectDetail(ctx context.Context, user *model.User, endpoint string) (*model.Project, error) {
	proj := new(model.Project)
	fields := new([]model.Field)
	projErr := pr.db.GetContext(ctx, proj, `SELECT * FROM project WHERE endpoint = $1 AND user_id = $2`, endpoint, user.ID)
	if projErr != nil {
		return nil, projErr
	}
	fieldErr := pr.db.SelectContext(ctx, fields, `SELECT * FROM field WHERE project_id = $1`, proj.ID)
	if fieldErr != nil {
		return nil, fieldErr
	}
	proj.Fields = fields
	return proj, nil
}

func (pr *Store) CreateProject(ctx context.Context, user *model.User, proj *model.Project) (string, error) {
	query := `INSERT INTO project (endpoint, user_id, title, description, documentUrl) VALUES (:endpoint, :user_id, :title, :description, :documentUrl) RETURNING endpoint`
	rows, err := pr.db.NamedQueryContext(ctx, query, proj)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var endpoint string
	for rows.Next() {
		err = rows.Scan(endpoint)
		break
	}
	if err != nil {
		return "", err
	}
	return endpoint, nil
}

func (pr *Store) UpdateProject(ctx context.Context, user *model.User, proj *model.Project) error {
	return nil
}
