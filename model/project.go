package model

import "time"

var ProjectSchema = `
	CREATE TABLE IF NOT EXISTS Project (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		endpoint UUID DEFAULT gen_random_uuid() UNIQUE,
		title VARCHAR(128) NOT NULL,
		description NVARCHAR(512),
		document_url varchar(256) NOT NULL,
		archived BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
		CONTRAINT fk_User FOREIGN KEY(user_id) REFERENCES User(id) ON DELETE CASCADE
	)
`

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
