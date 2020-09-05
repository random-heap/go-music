package model

import "time"

type Consumer struct {

	Id int
	Username string
	Password string
	Sex int
	PhoneNum string
	Email string
	Birth time.Time
	Introduction string
	Location string
	Avator string
	CreateTime time.Time
	UpdateTime time.Time
}
