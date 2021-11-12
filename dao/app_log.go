package dao

import "time"

type AppLog struct {
	ID        int64     `json:"id"`
	AppID     string    `json:"app_id"`
	LogNo     int64     `json:"log_no"`
	Content   string    `json:"content"`
	Level     int8      `json:"level"`
	LevelDesc string    `json:"level_desc"`
	Ctime     time.Time `json:"ctime"`
}
