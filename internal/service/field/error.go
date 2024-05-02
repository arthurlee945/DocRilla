package field

import "github.com/arthurlee945/Docrilla/internal/errors"

const (
	//REPOSITORY
	ErrRepoGet    = errors.Error("field_repo_get: couldn't get the field.")
	ErrRepoCreate = errors.Error("field_repo_create: couldn't create field.")
	ErrRepoUpdate = errors.Error("field_repo_update: coudln't update field.")
	ErrRepoDelete = errors.Error("field_repo_delete: couldn't delete field.")
	//SERVICE
	ErrInvalidUUID   = errors.Error("field_service_bad_request: uuid provided is not valid format.")
	ErrInvalidReqObj = errors.Error("field_service_bad_request: validation failed for request object")
	ErrServiceGet    = errors.Error("field_service_get: couldn't find the field.")
	ErrServiceCreate = errors.Error("field_service_create: field failed to created.")
	ErrServiceUpdate = errors.Error("field_service_update: field failed to updated.")
	ErrServiceDelete = errors.Error("field_service_delete: field failed to updated.")
)
