package dao

import "time"

type App struct {
	ID    int64     `json:"id"`
	AppID string    `json:"app_id"`
	Name  string    `json:"name"`
	Port  int16     `json:"port"`
	State int8      `json:"state"`
	Desc  string    `json:"desc"`
	Mid   int64     `json:"mid"`
	Uid   int64     `json:"uid"`
	Ctime time.Time `json:"ctime"`
}
