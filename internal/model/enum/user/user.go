package user

type Role string

const (
	ADMIN     Role = "ADMIN"
	USER      Role = "USER"
	MODERATOR Role = "MODERATOR"
	MOCK      Role = "MOCK"
)

var UserRoles = [2]Role{ADMIN, USER}
