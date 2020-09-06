package model

import "time"

type Collect struct {
	Id int
	UserId int
	Type int
	SongId int
	SongListId int
	CreateTime time.Time
}