package model

import (
	"fmt"
	"time"
)

type Default struct {
	ID        uint64
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var updateTimestampFunc = `
CREATE OR REPLACE FUNCTION auto_updated_at() RETURNS TRIGGER
    LANGUAGE plpgsql
    AS $$
BEGIN
	NEW.updated_at := now();
	RETURN NEW;
END;
$$;
`

func addAutoUpdatedAtTrigger(name string) string {
	return fmt.Sprintf(`
CREATE TRIGGER trigger_password_changed
BEFORE UPDATE ON %v
FOR EACH ROW
EXECUTE PROCEDURE auto_updated_at();
	`, name)
}
