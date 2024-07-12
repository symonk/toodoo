package model

// User encompasses a user of the system
type User struct {
	id       int
	username string
	password string
	admin    bool
}
