package model

import "time"

type Comment struct {

	Id int
	UserId int
	SongId int
	SongListId int
	Content string
	CreateTime time.Time
	Type int
	Up int
}
