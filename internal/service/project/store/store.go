package store

import (
	"context"
	"sync"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	ErrProjectFailedCreate = errors.Error("project_failed_create: project couldn't be created.")
	ErrProjectFailedUpdate = errors.Error("project_failed_update: project couldn't update.")
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
	if !rows.Next() {
		return "", ErrProjectFailedCreate.Wrap(errors.ErrNotFound)
	}
	var uuid string
	if err := rows.Scan(uuid); err != nil {
		return "", errors.ErrUnknown.Wrap(err)
	}
	return uuid, nil
}

func (pr *Store) UpdateProject(ctx context.Context, user *model.User, proj *model.Project) error {
	tx, err := pr.db.Beginx()
	if err != nil {
		return err
	}

	txCtx, cancel := context.WithCancel(ctx)

	errChan := make(chan error)
	waitChan := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(len(*proj.Fields) + 1)
	go func() {
		go func() {

			if _, err := tx.NamedExecContext(txCtx,
				`UPDATE project
			SET
			title = COALESCE(:title, title),
			description = COALESCE(:description, description),
			document_url = COALESCE(:document_url, document_url),
			archived = COALESCE(:archived, archived),
			visitedAt = COALESCE(:visted_at, visted_at)
			WHERE id = :id AND user_id = :user_id
			`, proj); err != nil {
				cancel()
				errChan <- err
			}
			wg.Done()
		}()

		for field := range *proj.Fields {
			go func() {
				if _, err := tx.NamedExecContext(txCtx,
					`UPDATE field
				SET
				x1 = COALESCE(:x1, x1),
				y1 = COALESCE(:y1, y1),
				x2 = COALESCE(:x2, x2),
				y2 = COALESCE(:y2, y2),
				page = COALESCE(:page, page),
				WHERE id = :id AND project_id = :project_id
				`,
					field); err != nil {
					cancel()
					errChan <- err
				}
				wg.Done()
			}()
		}

		wg.Wait()
		close(waitChan)
	}()

	select {
	case err := <-errChan:
		tx.Rollback()
		return ErrProjectFailedUpdate.Wrap(err)
	case <-waitChan:
		if err := tx.Commit(); err != nil {
			return ErrProjectFailedUpdate.Wrap(err)
		}
		return nil
	}
}
