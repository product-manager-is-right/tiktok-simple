package mw

type User struct {
	Id       int64
	Username string
	Password string
}

const IdentityKey = "user"
