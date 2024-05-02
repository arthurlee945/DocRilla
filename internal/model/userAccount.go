package model

import (
	"time"

	userEnum "github.com/arthurlee945/Docrilla/internal/model/enum/user"
)

type Account struct {
	ID                *uint64    `json:"id" db:"id"`
	UserID            *uint64    `json:"userID" db:"user_id"`
	Type              *string    `json:"type" db:"type"`
	Provide           *string    `json:"provider" db:"provider"`
	ProviderAccountID *string    `json:"providerAccountID" db:"provider_account_id"`
	RefreshToken      *string    `json:"refreshToken" db:"refresh_token"`
	AccessToken       *string    `json:"accessToken" db:"access_token"`
	ExpiresAt         *time.Time `json:"expiresAt" db:"expires_at"`
	TokenType         *string    `json:"tokenType" db:"token_type"`
	Scope             *string    `json:"scope" db:"scope"`
	IDToken           *string    `json:"idToken" db:"id_token"`
	SessionState      *string    `json:"sessionState" db:"session_state"`
}

type User struct {
	ID                     *uint64        `json:"id" db:"id"`
	Name                   *string        `json:"name" db:"name"`
	Email                  *string        `json:"email" db:"email"`
	EmailVerified          *bool          `json:"emailVerified" db:"email_verified"`
	EmailVerificationToken *string        `json:"emailVerificationToken" db:"email_verification_token"`
	Password               *string        `json:"password" db:"password"`
	PasswordChangedAt      *time.Time     `json:"passwordChangedAt" db:"password_changed_at"`
	ResetPasswordToken     *string        `json:"resetPasswordToken" db:"reset_password_token"`
	ResetPasswordExpires   *time.Time     `json:"resetPasswordExpires" db:"reset_password_expires"`
	Role                   *userEnum.Role `json:"role" db:"role"`
	Active                 *bool          `json:"active" db:"active"`
	CreatedAt              *time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt              *time.Time     `json:"updatedAt" db:"updated_at"`
	Projects               []Project      `json:"projects"`
	Accounts               []Account      `json:"accounts"`
	Sessions               []Session      `json:"session"`
}

type Session struct {
	ID           *uint64    `json:"id" db:"id"`
	SessionToken *string    `json:"sessionToken" db:"session_token"`
	Expires      *time.Time `json:"expires" db:"expires"`
}

type VerificationToken struct {
	Identifier string    `json:"id" db:"id"`
	Token      string    `json:"token" db:"token"`
	Expires    time.Time `json:"expires" db:"expires"`
}
