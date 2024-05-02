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
		return nil, ErrRepoGet.Wrap(err)
	}
	return field, nil
}

func (r *repository) Create(ctx context.Context, field *model.Field) (*model.Field, error) {
	rows, err := r.db.NamedQueryContext(ctx,
		`
	INSERT INTO field (project_id, x1, y1, x2, y2, page, type)
	VALUES (:project_id, :x1, :y1, :x2, :y2, :page, :type) RETURNING * 	
	`, field)
	if err != nil {
		return nil, ErrRepoCreate.Wrap(err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, ErrRepoCreate.Wrap(errors.ErrNotFound)
	}
	createdProj := &model.Field{}
	if err := rows.StructScan(createdProj); err != nil {
		return nil, ErrRepoCreate.Wrap(err)
	}
	return createdProj, nil
}

func (r *repository) Update(ctx context.Context, field *model.Field) (*model.Field, error) {
	rows, err := r.db.NamedQueryContext(ctx,
		`UPDATE field
	SET
	x1 = COALESCE(:x1, x1),
	y1 = COALESCE(:y1, y1),
	x2 = COALESCE(:x2, x2),
	y2 = COALESCE(:y2, y2),
	page = COALESCE(:page, page),
	type = COALESCE(:type, type)
	WHERE uuid=:uuid RETURNING *
	`, field)
	if err != nil {
		return nil, ErrRepoUpdate.Wrap(err)
	}
	defer rows.Close()
	updatedField := &model.Field{}
	if !rows.Next() {
		return nil, ErrRepoUpdate.Wrap(errors.ErrNotFound)
	}
	if err := rows.StructScan(updatedField); err != nil {
		return nil, ErrRepoUpdate.Wrap(err)
	}
	return updatedField, nil
}

func (r *repository) Delete(ctx context.Context, uuid string) error {
	if _, err := r.db.ExecContext(ctx, `
	DELETE FROM field
	WHERE uuid = $1 
	`, uuid); err != nil {
		return ErrRepoDelete.Wrap(err)
	}
	return nil
}
