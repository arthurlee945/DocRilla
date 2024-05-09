package project

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetAll(ctx context.Context, limit uint8, cursor string, userID uint64) (res []model.Project, nextCursor string, err error)
	GetOverviewById(ctx context.Context, uuid string, userID uint64) (*model.Project, error)
	GetDetailById(ctx context.Context, uuid string, userID uint64) (*model.Project, error)
	Create(ctx context.Context, proj *model.Project) (*model.Project, error)
	Update(ctx context.Context, proj *model.Project) (*model.Project, error)
	Delete(ctx context.Context, uuid string, userID uint64) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db,
	}
}

type Order string

const (
	DESC Order = "DESC"
	ASC  Order = "ASC"
)

func (r *repository) GetAll(ctx context.Context, limit uint8, cursor string, userID uint64) ([]model.Project, string, error) {
	decodedCursor, err := DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", errors.RepoError(err)
	}
	projects := []model.Project{}
	if err := r.db.SelectContext(ctx, &projects, `
	SELECT uuid, title, description, archived, token, route, created_at, visited_at 
	FROM project 
	WHERE user_id = $1 AND created_at > $2 ORDER BY created_at LIMIT $3`,
		userID, decodedCursor, limit); err != nil {
		return nil, "", errors.RepoError(err)
	}
	var nextCursor string

	if len(projects) == int(limit) {
		nextCursor = EncodeCursor(*projects[len(projects)-1].CreatedAt)
	}
	return projects, nextCursor, nil
}

func (r *repository) GetOverviewById(ctx context.Context, uuid string, userID uint64) (*model.Project, error) {
	proj := new(model.Project)
	if err := r.db.GetContext(ctx, proj, `
	SELECT uuid, title, description, archived, token, route, created_at, visited_at 
	FROM project WHERE user_id = $1 AND uuid = $2
	`, userID, uuid); err != nil {
		return nil, errors.RepoError(err)
	}
	return proj, nil
}

func (r *repository) GetDetailById(ctx context.Context, uuid string, userID uint64) (*model.Project, error) {
	proj, fields := new(model.Project), []model.Field{}
	if err := r.db.GetContext(ctx, proj, `SELECT * FROM project WHERE user_id = $1 AND uuid = $2`, userID, uuid); err != nil {
		return nil, errors.RepoError(err)
	}
	if err := r.db.SelectContext(ctx, &fields, `SELECT * FROM field WHERE project_id = $1`, proj.UUID); err != nil {
		return nil, errors.RepoError(err)
	}
	proj.Fields = fields
	return proj, nil
}

func (r *repository) Create(ctx context.Context, proj *model.Project) (*model.Project, error) {
	rows, err := r.db.NamedQueryContext(ctx, `
		INSERT INTO project ( user_id, title, description, document_url, token, route) 
		VALUES (:user_id, :title, :description, :document_url, :token, :route) RETURNING *
		`, proj)
	if err != nil {
		return nil, errors.RepoError(err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.ErrUnknown
	}
	newProj := new(model.Project)
	if err := rows.StructScan(newProj); err != nil {
		return nil, errors.RepoError(err)
	}
	return newProj, nil
}

func (r *repository) Update(ctx context.Context, proj *model.Project) (*model.Project, error) {
	rows, err := r.db.NamedQueryContext(ctx,
		`UPDATE project
	SET
	title = COALESCE(:title, title),
	description = COALESCE(:description, description),
	document_url = COALESCE(:document_url, document_url),
	route = COALESCE(:route, route),
	token = COALESCE(:token, token),
	archived = COALESCE(:archived, archived),
	visited_at = COALESCE(:visited_at, visited_at)
	WHERE uuid=:uuid AND user_id=:user_id RETURNING *
	`, proj)
	if err != nil {
		return nil, errors.RepoError(err)
	}
	defer rows.Close()
	updatedProj := &model.Project{}
	if !rows.Next() {
		return nil, errors.ErrNotFound
	}
	if err := rows.StructScan(updatedProj); err != nil {
		return nil, errors.RepoError(err)
	}
	return updatedProj, nil
}

func (r *repository) Delete(ctx context.Context, uuid string, userID uint64) error {
	if _, err := r.db.ExecContext(ctx, `
	DELETE FROM project
	WHERE uuid = $1 AND user_id = $2
	`, uuid, userID); err != nil {
		return errors.RepoError(err)
	}
	return nil
}
