package project

import "github.com/arthurlee945/Docrilla/internal/errors"

const (
	//REPOSITORY
	ErrRepoGet    = errors.Error("project_repo_get: couldn't get the project.")
	ErrRepoCreate = errors.Error("project_repo_create: couldn't create project.")
	ErrRepoUpdate = errors.Error("project_repo_update: couldn't update project.")
	ErrRepoDelete = errors.Error("project_repo_delete: couldn't delete project.")
	//SERVICE
	ErrInvalidUUID   = errors.Error("project_service_bad_request: uuid provided is not valid format.")
	ErrInvalidReqObj = errors.Error("project_service_bad_request: validation failed for request object")
	ErrServiceGet    = errors.Error("project_service_get: couldn't find the project.")
	ErrServiceCreate = errors.Error("project_service_create: project failed to created.")
	ErrServiceUpdate = errors.Error("project_service_update: project failed to updated.")
	ErrServiceDelete = errors.Error("project_service_delete: project failed to updated.")
)
