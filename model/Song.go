package model

import "time"

type Song struct {

	Id int
	singerId int
	Name string
	Introduction string
	CreateTime time.Time
	UpdateTime time.Time
	Pic string
	Lyric string
	Url string

}
