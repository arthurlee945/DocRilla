package field

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetById(ctx context.Context, uuid string) (*model.Field, error)
	Create(ctx context.Context, field *model.Field) (*model.Field, error)
	Update(ctx context.Context, field *model.Field) (*model.Field, error)
	Delete(ctx context.Context, uuid string) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) GetById(ctx context.Context, uuid string) (*model.Field, error) {
	field := &model.Field{}
	if err := r.db.GetContext(ctx, field, `SELECT * FROM field WHERE uuid = $1`, uuid); err != nil {
		return nil, errors.RepoError(err)
	}
	return field, nil
}

func (r *repository) Create(ctx context.Context, field *model.Field) (*model.Field, error) {
	rows, err := r.db.NamedQueryContext(ctx,
		`
	INSERT INTO field (project_id, x, y, width, height, page, type)
	VALUES (:project_id, :x, :y, :width, :height, :page, :type) RETURNING * 	
	`, field)
	if err != nil {
		return nil, errors.RepoError(err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.ErrNotFound
	}
	createdProj := &model.Field{}
	if err := rows.StructScan(createdProj); err != nil {
		return nil, errors.RepoError(err)
	}
	return createdProj, nil
}

func (r *repository) Update(ctx context.Context, field *model.Field) (*model.Field, error) {
	rows, err := r.db.NamedQueryContext(ctx,
		`UPDATE field
	SET
	x = COALESCE(:x, x),
	y = COALESCE(:y, y),
	width = COALESCE(:width, width),
	height = COALESCE(:height, height),
	page = COALESCE(:page, page),
	type = COALESCE(:type, type)
	WHERE uuid=:uuid AND project_id=:project_id RETURNING *
	`, field)
	if err != nil {
		return nil, errors.RepoError(err)
	}
	defer rows.Close()
	updatedField := &model.Field{}
	if !rows.Next() {
		return nil, errors.ErrNotFound
	}
	if err := rows.StructScan(updatedField); err != nil {
		return nil, errors.RepoError(err)
	}
	return updatedField, nil
}

func (r *repository) Delete(ctx context.Context, uuid string) error {
	if _, err := r.db.ExecContext(ctx, `
	DELETE FROM field
	WHERE uuid = $1 
	`, uuid); err != nil {
		return errors.RepoError(err)
	}
	return nil
}
