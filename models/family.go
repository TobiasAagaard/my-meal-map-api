package models

type Family struct {
	ID          uint
	Name        string
	Users       []User
	InviteToken string
}
