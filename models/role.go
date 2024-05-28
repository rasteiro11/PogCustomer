package models

type Role struct {
	ID   uint
	Name string
}

type DefaultRole struct {
	ID   uint
	Name string
}

type UserRole struct {
	UserID uint
	RoleID uint
}
