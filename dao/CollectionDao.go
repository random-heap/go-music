package dao

import (
	"go-music/model"
	"go-music/utils"
)

type CollectionDao struct {
	
}

func (dao *CollectionDao) Insert(collect model.Collect) int64 {
	result := utils.DB.MustExec("INSERT INTO  collection (user_id, type, song_id, song_list_id, create_time) " +
		" VALUES(?, ?, ?, ?, ?)", collect.UserId, collect.Type, collect.SongId, collect.SongListId, collect.CreateTime)

	row, _ := result.RowsAffected()
	return row
}

func (dao *CollectionDao) SelectByUserIdAndSongId(userId int, songId int) int {
	var count int
	row := utils.DB.QueryRow("select count(1) from collect where user_id = ? and song_id = ? ", userId, songId)
	row.Scan(&count)
	return count
}
