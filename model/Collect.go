package model

import "time"

type Collect struct {
	Id int
	UserId int
	Type int
	SongId int
	CreateTime time.Time
}