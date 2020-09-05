package dao

import (
	"go-music/model"
	"go-music/utils"
)

type AdminDao struct {
	
}

func (dao *AdminDao) Insert(admin model.Admin) {

	utils.DB.MustExec("INSERT INTO admin(name, `password`) VALUES(?, ?)", admin.Name, admin.Password)

}

func (dao *AdminDao) SelectByNameAndPassword(name string, password string) int {

	var count int
	//utils.DB.Select(&count, "select count(1) from admin where name = ? and `password` = ? ", name, password)

	row := utils.DB.QueryRow("select count(1) from admin where name = ? and `password` = ? ", name, password)
	row.Scan(&count)
	return count
}

