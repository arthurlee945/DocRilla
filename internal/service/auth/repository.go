package auth

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetById(ctx context.Context, id uint64) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, user *model.User) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.GetContext(ctx, user, `SELECT * FROM usr WHERE email=$1`, email); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetById(ctx context.Context, id uint64) (*model.User, error) {
	user := new(model.User)
	if err := r.db.GetContext(ctx, user, `SELECT * FROM usr WHERE id=$1`, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	rows, err := r.db.NamedQueryContext(ctx, `
	INSERT INTO usr (name, email, password)
	VALUES (:name, :email, :password) RETURNING *
	`, user)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.ErrUnknown
	}
	newUser := new(model.User)
	if err := rows.StructScan(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *repository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	rows, err := r.db.NamedQueryContext(ctx, `
	UPDATE usr
	SET
	name = COALESCE(:name, name),
	email = COALESCE(:email, email),
	email_verified = COALESCE(:email_verified, email_verified),
	email_verification_token = COALESCE(:email_verification_token, email_verification_token),
	password = COALESCE(:password, password),
	role = COALESCE(:role, role),
	password_changed_at = COALESCE(:password_changed_at, password_changed_at),
	reset_password_token = COALESCE(:reset_password_token, reset_password_token),
	reset_password_expires = COALESCE(:reset_password_expires, reset_password_expires),
	active = COALESCE(:active, active)
	WHERE user_id = :user_id RETURNING *
	`, user)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.ErrUnknown
	}
	updatedUser := new(model.User)
	if err := rows.StructScan(updatedUser); err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (r *repository) Delete(ctx context.Context, user *model.User) error {
	if _, err := r.db.NamedExecContext(ctx, `DELETE FROM usr WHERE id=:id`, user); err != nil {
		return err
	}
	return nil
}
