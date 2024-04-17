package null

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type String struct {
	sql.NullString
}

func (s *String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

type Int64 struct {
	sql.NullInt64
}

func (i64 *Int64) MarshalJSON() ([]byte, error) {
	if !i64.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i64.Int64)
}

type Bool struct {
	sql.NullBool
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Bool)
}

type Time struct {
	sql.NullTime
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", t.Time.Format(time.RFC3339))
	return json.Marshal(val)
}
