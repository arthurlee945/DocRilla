package model

import (
	"fmt"
	"time"

	"github.com/arthurlee945/Docrilla/model/enum/user"
)

var userSchema = fmt.Sprintf(`	
CREATE TABLE IF NOT EXISTS user_account (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	email_verified BOOLEAN DEFAULT FALSE,
	email_verification_token TEXT,
	password TEXT NOT NULL,
	password_changed_at TIMESTAMP,
	reset_password_token TEXT NULL,
	reset_password_expires TIMESTAMP NULL,
	access_token TEXT NULL,
	token_expires_at TIMESTAMP,
	role Text DEFAULT '%v',
	active BOOLEAN DEFAULT TRUE,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	CONSTRAINT check_role check (role in ('%v', '%v'))
);

%v
`, user.ADMIN, user.ADMIN, user.USER, addAutoUpdatedAtTrigger("user_account"))

type User struct {
	UserID                 uint64
	Name                   string
	Email                  string
	EmailVerified          bool   `db:"email_verified"`
	EmailVerificationToken string `db:"email_verification_token"`
	Password               string
	PasswordChangedAt      time.Time `db:"password_changed_at"`
	ResetPasswordToken     time.Time `db:"reset_password_token"`
	ResetPasswordExpires   time.Time `db:"reset_password_expires"`
	AccessToken            string    `db:"access_token"`
	TokenExpiresAt         time.Time `db:"token_expires_at"`
	Role                   user.Role
	Active                 bool
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
	Projects               []Project
}
