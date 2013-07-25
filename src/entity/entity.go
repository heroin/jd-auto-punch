package entity

import (
	"time"
)

type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Start    bool   `json:"start"`
	Trigger  int64  `json:"trigger"`
	Date     time.Time
}

type Task struct {
	Users []*User `json:"users"`
	Start bool    `json:"start"`
	Size  int
}
