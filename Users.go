package main

import (
    "time"
)

type User struct {
    Name        string          `json:"name"`
    // First       string      `json:"first"`
    // Last        string      `json:"last"`
    Created     time.Time       `json:"created"`
}

type Users []User
