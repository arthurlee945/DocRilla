package model

import (
	"time"

	"github.com/arthurlee945/Docrilla/model/user"
)

type User struct {
	ID                     string    `json:"id"`
	Name                   string    `json:"name"`
	Email                  string    `json:"email"`
	EmailVerified          bool      `json:"-"`
	EmailVerificationToken string    `json:"-"`
	Password               string    `json:"-"`
	PasswordChangedAt      time.Time `json:"-"`
	ResetPasswordToken     time.Time `json:"-"`
	ResetPasswordExpires   time.Time `json:"-"`
	AccessToken            string    `json:"accessToken"`
	TokenExpiresAt         time.Time `json:"_"`
	Role                   user.Role `json:"-"`
	Active                 bool      `json:"-"`
	CreatedAt              time.Time `json:"-"`
	UpdatedAt              time.Time `json:"-"`
	Projects               []Project `json:"projects"`
}
