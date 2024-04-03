package model

import (
	"database/sql"
	"time"

	"github.com/arthurlee945/Docrilla/model/enum/user"
)

type Account struct {
	ID                uint64
	UserID            uint64 `db:"user_id"`
	Type              string
	Provide           string
	ProviderAccountID sql.NullString `db:"provider_account_id"`
	RefreshToken      sql.NullString `db:"refresh_token"`
	AccessToken       sql.NullString `db:"access_token"`
	ExpiresAt         sql.NullTime   `db:"expires_at"`
	TokenType         sql.NullString `db:"token_type"`
	Scope             sql.NullString
	IDToken           sql.NullString `db:"id_token"`
	SessionState      sql.NullString `db:"session_state"`
	*User
}

type User struct {
	ID                     uint64
	Name                   string
	Email                  string
	EmailVerified          bool           `db:"email_verified"`
	EmailVerificationToken sql.NullString `db:"email_verification_token"`
	Password               sql.NullString
	PasswordChangedAt      sql.NullTime   `db:"password_changed_at"`
	ResetPasswordToken     sql.NullString `db:"reset_password_token"`
	ResetPasswordExpires   sql.NullTime   `db:"reset_password_expires"`
	Role                   user.Role
	Active                 bool
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
	Projects               []Project
	Accounts               []Account
	Sessions               []Session
}

type Session struct {
	ID           uint64
	SessionToken string `db:"session_token"`
	Expires      time.Time
	*User
}

type VerificationToken struct {
	Identifier string
	Token      string
	Expires    time.Time
}
