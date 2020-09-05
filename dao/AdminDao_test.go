package dao

import (
	"fmt"
	"go-music/model"
	"testing"
)

func testInsertAdmin(t *testing.T) {

	adminDao := new(AdminDao)

	adminDao.insert(model.Admin{
		Name: "guest",
		Password: "guest",
	})

}

func TestSelectByNameAndPassword(t *testing.T) {

	adminDao := new(AdminDao)

	count := adminDao.selectByNameAndPassword("admin", "123")
	fmt.Println(count)

}

