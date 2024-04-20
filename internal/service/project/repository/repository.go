package repository

import (
	"context"
	"database/sql"
	"sync"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	ErrProjectFailedGet    = errors.Error("project_failed_get: couldn't find the project.")
	ErrProjectFailedCreate = errors.Error("project_failed_create: project couldn't be created.")
	ErrProjectFailedUpdate = errors.Error("project_failed_update: project couldn't update.")
	ErrProjectFailedDelete = errors.Error("project_failed_delete: project couldn't delete.")
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) GetProjectOverview(ctx context.Context, uuid string) (*model.Project, error) {
	proj := new(model.Project)
	if err := r.db.GetContext(ctx, proj, `
	SELECT uuid, title, description, archived, created_at, visited_at 
	FROM project WHERE uuid = $1
	`, uuid); err != nil {
		return nil, err
	}
	return proj, nil
}

func (r *Repository) GetProjectDetail(ctx context.Context, uuid string) (*model.Project, error) {
	proj, fields := new(model.Project), []model.Field{}
	if err := r.db.GetContext(ctx, proj, `SELECT * FROM project WHERE uuid = $1`, uuid); err != nil {
		return nil, err
	}
	if err := r.db.SelectContext(ctx, &fields, `SELECT * FROM field WHERE project_id = $1`, proj.ID); err != nil {
		return nil, err
	}
	proj.Fields = fields
	return proj, nil
}

func (r *Repository) CreateProject(ctx context.Context, proj *model.Project) (*model.Project, error) {
	rows, err := r.db.NamedQueryContext(ctx, `
		INSERT INTO project ( user_id, title, description, document_url) 
		VALUES (:user_id, :title, :description, :document_url) RETURNING *
		`, proj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, ErrProjectFailedCreate.Wrap(errors.ErrNotFound)
	}
	newProj := new(model.Project)
	if err := rows.StructScan(newProj); err != nil {
		return nil, errors.ErrUnknown.Wrap(err)
	}
	return newProj, nil
}

func (r *Repository) UpdateProject(ctx context.Context, proj *model.Project) error {
	txCtx, txCancel := context.WithCancel(ctx)
	defer txCancel()
	tx, err := r.db.BeginTxx(txCtx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	errChan, waitChan := make(chan error), make(chan struct{})

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(proj.Fields) + 1)
		go func() {
			defer wg.Done()
			if _, err := tx.NamedExecContext(txCtx,
				`UPDATE project
			SET
			title = COALESCE(:title, title),
			description = COALESCE(:description, description),
			document_url = COALESCE(:document_url, document_url),
			archived = COALESCE(:archived, archived),
			visited_at = COALESCE(:visited_at, visited_at)
			WHERE uuid=:uuid
			`, proj); err != nil {
				errChan <- errors.Error("Project Update Error").Wrap(err)
				txCancel()
			}
		}()

		for _, field := range proj.Fields {
			go func(field *model.Field) {
				defer wg.Done()
				if _, err := tx.NamedExecContext(txCtx,
					`UPDATE field
				SET
				x1 = COALESCE(:x1, x1),
				y1 = COALESCE(:y1, y1),
				x2 = COALESCE(:x2, x2),
				y2 = COALESCE(:y2, y2),
				page = COALESCE(:page, page),
				type = COALESCE(:type, type)
				WHERE uuid=:uuid
				`,
					field); err != nil {
					errChan <- errors.Error("Field Update Error").Wrap(err)
					txCancel()
				}
			}(&field)
		}

		wg.Wait()
		close(waitChan)
	}()

	select {
	case err := <-errChan:
		return ErrProjectFailedUpdate.Wrap(err)
	case <-waitChan:
		if err := tx.Commit(); err != nil {
			return ErrProjectFailedUpdate.Wrap(err)
		}
		return nil
	}
}

func (r *Repository) DeleteProject(ctx context.Context, uuid string) error {
	if _, err := r.db.ExecContext(ctx, `
	DELETE FROM project
	WHERE uuid = $1 
	`, uuid); err != nil {
		return ErrProjectFailedDelete.Wrap(err)
	}
	return nil
}
