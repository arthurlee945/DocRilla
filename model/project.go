package model

import (
	"fmt"
	"time"
)

var projectSchema = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS project (
	id SERIAL PRIMARY KEY,
	endpoint UUID DEFAULT gen_random_uuid() UNIQUE,
	title VARCHAR(128) NOT NULL,
	description TEXT,
	document_url TEXT NOT NULL,
	archived BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	user_id INT,
	CONSTRAINT fk_account FOREIGN KEY(user_id) REFERENCES user_account(id)
);

%v
`, addAutoUpdatedAtTrigger("project"))

type Project struct {
	ProjectID   uint64
	Endpoint    string
	UserID      int `db:"user_id"`
	Title       string
	Description string
	DocumentUrl string `db:"document_url"`
	Archived    bool
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	User        User
}
