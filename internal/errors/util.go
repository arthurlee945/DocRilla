package errors

import (
	"database/sql"
)

func RepoError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case Is(err, sql.ErrNoRows):
		return ErrNotFound.Wrap(err)
	case Is(err, ErrDecodeCursor):
		return ErrInvalidRequest.Wrap(err)
	default:
		return ErrUnknown.Wrap(err)
	}
}

func AuthError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case Is(err, sql.ErrNoRows):
		fallthrough
	case Is(err, sql.ErrNoRows):
		return ErrUnauthorized
	default:
		return ErrUnknown.Wrap(err)
	}
}
