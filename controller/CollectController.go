package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/dao"
	"go-music/model"
	"net/http"
	"strconv"
	"time"
)

func AddCollection(context *gin.Context) {

	collectDao := new(dao.CollectionDao)

	userId := context.Query("userId")
	typeStr := context.Query("type")
	songId := context.Query("songId")
	songListId := context.Query("songListId")
	//name := context.Query("name")

	if songId == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "收藏歌曲为空",
		})
	}

	_userId, _ := strconv.Atoi(userId)
	_songId, _ := strconv.Atoi(songId)
	_type, _ := strconv.Atoi(typeStr)
	_songListId, _ := strconv.Atoi(songListId)


	count := collectDao.SelectByUserIdAndSongId(_userId, _songId)
	if count > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "已收藏",
		})
	}

	collcet := model.Collect{}
	collcet.UserId = _userId
	collcet.Type = _type
	if _type == 0 {
		collcet.SongId = _songId
	} else if _type == 1 {
		collcet.SongListId = _songListId
	}
	collcet.CreateTime = time.Now()

	row := collectDao.Insert(collcet)
	if row > 1 {
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "收藏成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "收藏失败",
		})
	}

}
