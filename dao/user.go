package dao

import (
	"log"
	"time"
)

type User struct {
	ID       int64     `json:"id"`
	Nickname string    `json:"nickname"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Ctime    time.Time `json:"ctime"`
}

func QueryAllUser() []User {
	var users []User
	db.Find(&users)
	return users
}

func QueryByUsername(username string) *User {
	var user User
	if err := db.Model(&User{}).Where("username = ?", username).First(user).Error; err != nil {
		log.Println("用户查找失败: ", err)
		return nil
	}
	return &user
}

func (u *User) CreateUser() bool {
	u.Ctime = time.Now()
	if err := db.Create(&u).Error; err != nil {
		log.Println("创建用户失败: ", err)
		return false
	}
	return true
}

func (u *User) DeleteUser() {
	if err := db.Delete(u).Error; err != nil {
		log.Println("删除用户失败: ", err)
	}
}

func (u *User) UpdateUser() {
	if err := db.Updates(u).Error; err != nil {
		log.Println("更新用户失败: ", err)
	}
}
