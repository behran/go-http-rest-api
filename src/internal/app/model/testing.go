package model

import "testing"

// TestUser return user
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user-2@exmaple.org",
		Password: "PasswordHard",
	}
}
