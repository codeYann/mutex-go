package users

import (
	"math/rand"
)

// User has fields: name, id and role and mutex for sync
type User struct {
	name    string
	id      int
	role    string
	content string
}

// Create new user
func CreateUser(name, role string) *User {
	return &User{
		name: name,
		id:   rand.Intn(1000),
		role: role,
	}
}

// Get user name
func (u *User) GetName() string {
	return u.name
}

// Get user id
func (u *User) GetId() int {
	return u.id
}

// Get user role
func (u *User) GetRole() string {
	return u.role
}

// Get user content
func (u *User) GetContent() string {
	return u.content
}

// Set user content
func (u *User) SetContent(content string) {
	u.content = content
}
