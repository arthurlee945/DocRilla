package model

import (
	"time"

	"github.com/arthurlee945/Docrilla/model/enum/user"
)

type Account struct {
	AccountID         uint64 `db:"id"`
	UserID            uint64
	Type              string
	Provide           string
	ProviderAccountID string    `db:"provider_account_id"`
	RefreshToken      string    `db:"refresh_token"`
	AccessToken       string    `db:"access_token"`
	ExpiresAt         time.Time `db:"expires_at"`
	TokenType         string    `db:"token_type"`
	Scope             string
	IDToken           string `db:"id_token"`
	SessionState      string `db:"session_state"`
	User
}

type User struct {
	UserID                 uint64 `db:"id"`
	Name                   string
	Email                  string
	EmailVerified          bool   `db:"email_verified"`
	EmailVerificationToken string `db:"email_verification_token"`
	Password               string
	PasswordChangedAt      time.Time `db:"password_changed_at"`
	ResetPasswordToken     time.Time `db:"reset_password_token"`
	ResetPasswordExpires   time.Time `db:"reset_password_expires"`
	Role                   user.Role
	Active                 bool
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
	Projects               []Project
	Accounts               []Account
	Sessions               []Session
}

type Session struct {
	SessionID    uint64 `db:"id"`
	SessionToken string `db:"session_token"`
	Expires      time.Time
	User
}

type VerificationToken struct {
	Identifier string
	Token      string
	Expires    time.Time
}
