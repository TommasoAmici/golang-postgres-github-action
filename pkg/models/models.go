package models

// User is a struct to contain user records from database
type User struct {
	ID       int
	Username string
}

// UserModel is an interface for database bindings of User
type UserModel interface {
	Get(id int) (*User, error)
}
