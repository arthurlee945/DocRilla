package project

import (
	"encoding/base64"
	"time"

	"github.com/arthurlee945/Docrilla/internal/errors"
)

const (
	testDSN    = "postgresql://public_user:Qwer1234@localhost:5432/docrilla?sslmode=disable"
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

func DecodeCursor(encodedTime string) (time.Time, error) {
	byts, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, errors.ErrDecodeCursor.Wrap(err)
	}
	return time.Parse(timeFormat, string(byts))
}

func EncodeCursor(t time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(t.Format(timeFormat)))
}
