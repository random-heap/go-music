package dao

import (
	"fmt"
	"go-music/model"
	"testing"
)

func testInsertAdmin(t *testing.T) {

	adminDao := new(AdminDao)

	adminDao.Insert(model.Admin{
		Name: "guest",
		Password: "guest",
	})

}

func TestSelectByNameAndPassword(t *testing.T) {

	adminDao := new(AdminDao)

	count := adminDao.SelectByNameAndPassword("admin", "123")
	fmt.Println(count)

}

