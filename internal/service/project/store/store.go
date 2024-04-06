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

func (pr *Store) GetProjectOverview(ctx context.Context, user *model.User, uuid string) (*model.Project, error) {
	query := `
	SELECT uuid, title, description, archived, created_at, visited_at 
	FROM project WHERE uuid = $1 AND user_id = $2
	`
	proj := new(model.Project)
	err := pr.db.GetContext(ctx, proj, query, uuid, user.ID)
	if err != nil {
		return nil, err
	}
	return proj, nil
}

func (pr *Store) GetProjectDetail(ctx context.Context, user *model.User, uuid string) (*model.Project, error) {
	proj := new(model.Project)
	fields := new([]model.Field)
	projErr := pr.db.GetContext(ctx, proj, `SELECT * FROM project WHERE uuid = $1 AND user_id = $2`, uuid, user.ID)
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
	query := `INSERT INTO project ( user_id, title, description, documentUrl) VALUES (:user_id, :title, :description, :documentUrl) RETURNING uuid`
	rows, err := pr.db.NamedQueryContext(ctx, query, proj)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var uuid string
	for rows.Next() {
		err = rows.Scan(uuid)
		break
	}
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (pr *Store) UpdateProject(ctx context.Context, user *model.User, proj *model.Project) error {
	return nil
}
