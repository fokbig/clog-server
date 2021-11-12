package dao

import "time"

type Machine struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	IP       string    `json:"ip"`
	State    int8      `json:"state"`
	Location string    `json:"location"`
	Ctime    time.Time `json:"ctime"`
}
