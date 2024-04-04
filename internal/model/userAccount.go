package model

import (
	"database/sql"
	"time"

	"github.com/arthurlee945/Docrilla/internal/model/enum/user"
)

type Account struct {
	ID                uint64         `json:"id" db:"id"`
	UserID            uint64         `json:"userID" db:"user_id"`
	Type              string         `json:"type" db:"type"`
	Provide           string         `json:"provider" db:"provider"`
	ProviderAccountID sql.NullString `json:"providerAccountID" db:"provider_account_id"`
	RefreshToken      sql.NullString `json:"refreshToken" db:"refresh_token"`
	AccessToken       sql.NullString `json:"accessToken" db:"access_token"`
	ExpiresAt         sql.NullTime   `json:"expiresAt" db:"expires_at"`
	TokenType         sql.NullString `json:"tokenType" db:"token_type"`
	Scope             sql.NullString `json:"scope" db:"scope"`
	IDToken           sql.NullString `json:"idToken" db:"id_token"`
	SessionState      sql.NullString `json:"sessionState" db:"session_state"`
	User              *User          `json:"user"`
}

type User struct {
	ID                     uint64         `json:"id" db:"id"`
	Name                   string         `json:"name" db:"name"`
	Email                  string         `json:"email" db:"email"`
	EmailVerified          bool           `json:"emailVerified" db:"email_verified"`
	EmailVerificationToken sql.NullString `json:"emailVerificationToken" db:"email_verification_token"`
	Password               sql.NullString `json:"password" db:"password"`
	PasswordChangedAt      sql.NullTime   `json:"passwordChangedAt" db:"password_changed_at"`
	ResetPasswordToken     sql.NullString `json:"resetPasswordToken" db:"reset_password_token"`
	ResetPasswordExpires   sql.NullTime   `json:"resetPasswordExpires" db:"reset_password_expires"`
	Role                   user.Role      `json:"role" db:"role"`
	Active                 bool           `json:"active" db:"active"`
	CreatedAt              time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt              time.Time      `json:"updatedAt" db:"updated_at"`
	Projects               []Project      `json:"projects"`
	Accounts               []Account      `json:"accounts"`
	Sessions               []Session      `json:"session"`
}

type Session struct {
	ID           uint64    `json:"id" db:"id"`
	SessionToken string    `json:"sessionToken" db:"session_token"`
	Expires      time.Time `json:"expires" db:"expires"`
	User         *User     `json:"user"`
}

type VerificationToken struct {
	Identifier string    `json:"id" db:"id"`
	Token      string    `json:"token" db:"token"`
	Expires    time.Time `json:"expires" db:"expires"`
}
