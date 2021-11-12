package dao

import "time"

type User struct {
	ID       int64     `json:"id"`
	Nickname string    `json:"nickname"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Ctime    time.Time `json:"ctime"`
}
