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
	Start  bool     `json:"start"`
	Users  []*User  `json:"users"`
	Cancel []string `json:"cancel"`
	Size   int
}
