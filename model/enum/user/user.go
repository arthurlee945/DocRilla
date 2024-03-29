package user

type Role string

const (
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

var UserRoles = [2]Role{ADMIN, USER}
