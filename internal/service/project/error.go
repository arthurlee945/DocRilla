package project

import "github.com/arthurlee945/Docrilla/internal/errors"

const (
	//REPOSITORY
	ErrRepoGet    = errors.Error("project_repo_get: couldn't find the project.")
	ErrRepoCreate = errors.Error("project_repo_create: project couldn't be created.")
	ErrRepoUpdate = errors.Error("project_repo_update: project couldn't update.")
	ErrRepoDelete = errors.Error("project_repo_delete: project couldn't delete.")
	//SERVICE
	ErrServiceGet    = errors.Error("project_service_get: couldn't find the project.")
	ErrInvalidUUID   = errors.Error("project_service_bad_request: uuid provided is not valid format.")
	ErrInvalidReqObj = errors.Error("project_service_bad_request: validation failed for request object")
	ErrServiceCreate = errors.Error("project_service_create: project couldn't be created.")
)
