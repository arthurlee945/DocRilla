package field

import "github.com/arthurlee945/Docrilla/internal/errors"

const (
	//REPOSITORY
	ErrRepoGet    = errors.Error("field_repo_get: couldn't get the field.")
	ErrRepoCreate = errors.Error("field_repo_create: couldn't create field.")
	ErrRepoUpdate = errors.Error("field_repo_update: coudln't update field.")
	ErrRepoDelete = errors.Error("field_repo_delete: couldn't delete field.")
)
