package user

import (
	"time"
)

// User defines the information about an individual user.
type User struct {
	Name string `json:"name"`
	// First       string      `json:"first"`
	// Last        string      `json:"last"`
	Created time.Time `json:"created"`
}

// Users is a list of users.
type Users []User
