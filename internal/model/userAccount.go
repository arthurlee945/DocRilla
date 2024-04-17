package model

import (
	"time"

	"github.com/arthurlee945/Docrilla/internal/model/enum/user"
	"github.com/arthurlee945/Docrilla/internal/model/null"
)

type Account struct {
	ID                uint64      `json:"id" db:"id"`
	UserID            uint64      `json:"userID" db:"user_id"`
	Type              string      `json:"type" db:"type"`
	Provide           string      `json:"provider" db:"provider"`
	ProviderAccountID null.String `json:"providerAccountID" db:"provider_account_id"`
	RefreshToken      null.String `json:"refreshToken" db:"refresh_token"`
	AccessToken       null.String `json:"accessToken" db:"access_token"`
	ExpiresAt         null.Time   `json:"expiresAt" db:"expires_at"`
	TokenType         null.String `json:"tokenType" db:"token_type"`
	Scope             null.String `json:"scope" db:"scope"`
	IDToken           null.String `json:"idToken" db:"id_token"`
	SessionState      null.String `json:"sessionState" db:"session_state"`
	User              *User       `json:"user"`
}

type User struct {
	ID                     uint64      `json:"id" db:"id"`
	Name                   string      `json:"name" db:"name"`
	Email                  string      `json:"email" db:"email"`
	EmailVerified          bool        `json:"emailVerified" db:"email_verified"`
	EmailVerificationToken null.String `json:"emailVerificationToken" db:"email_verification_token"`
	Password               null.String `json:"password" db:"password"`
	PasswordChangedAt      null.Time   `json:"passwordChangedAt" db:"password_changed_at"`
	ResetPasswordToken     null.String `json:"resetPasswordToken" db:"reset_password_token"`
	ResetPasswordExpires   null.Time   `json:"resetPasswordExpires" db:"reset_password_expires"`
	Role                   user.Role   `json:"role" db:"role"`
	Active                 bool        `json:"active" db:"active"`
	CreatedAt              time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt              time.Time   `json:"updatedAt" db:"updated_at"`
	Projects               []Project   `json:"projects"`
	Accounts               []Account   `json:"accounts"`
	Sessions               []Session   `json:"session"`
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
