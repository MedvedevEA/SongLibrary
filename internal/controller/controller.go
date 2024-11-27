package controller

import (
	"database/sql"
	"errors"
	"fmt"

	"songLibrary/internal/controller/dto"
	"songLibrary/internal/service"
	"songLibrary/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	router  *gin.Engine
	service *service.Service
}

func InitController(router *gin.Engine, service *service.Service) {
	controller := &Controller{

		router,
		service,
	}
	router.GET("songs", controller.getSongs)
	router.GET("songs/:song_id", controller.getSong)
	router.GET("songs/:song_id/text", controller.getSongText)
	router.POST("songs", controller.addSong)
	router.PUT("songs", controller.updateSong)
	router.DELETE("songs/:song_id", controller.removeSong)

	router.GET("info", controller.getInfo)

}
func (controller *Controller) addSong(ctx *gin.Context) {
	req := new(dto.AddSongReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	res, err := controller.service.AddSong(req)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
func (controller *Controller) getInfo(ctx *gin.Context) {
	group := helpers.ToString(ctx.GetQuery("group"))
	if group == nil {
		ctx.JSON(400, gin.H{"message": "Controller: parameter missing 'group'"})
		return
	}
	song := helpers.ToString(ctx.GetQuery("song"))
	if song == nil {
		ctx.JSON(400, gin.H{"message": "Controller: parameter missing 'song'"})
		return
	}
	res, err := controller.service.GetInfo(&dto.GetInfoReq{
		Group: *group,
		Song:  *song,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
func (controller *Controller) getSong(ctx *gin.Context) {
	req, err := helpers.ToUuid(ctx.Param("song_id"), true)
	if err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	res, err := controller.service.GetSong(req)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(404, gin.H{"message": "Song not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
func (controller *Controller) getSongText(ctx *gin.Context) {
	songId, err := helpers.ToUuid(ctx.Param("song_id"), true)
	if err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	offset, err := helpers.ToInt(ctx.GetQuery("offset"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	limit, err := helpers.ToInt(ctx.GetQuery("limit"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	res, err := controller.service.GetSongText(&dto.GetSongTextReq{
		SongId: songId,
		Offset: offset,
		Limit:  limit,
	})
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(404, gin.H{"message": "Song not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
func (controller *Controller) getSongs(ctx *gin.Context) {
	offset, err := helpers.ToInt(ctx.GetQuery("offset"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	limit, err := helpers.ToInt(ctx.GetQuery("limit"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	releaseDate, err := helpers.ToDate(ctx.GetQuery("release_date"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	res, err := controller.service.GetSongs(&dto.GetSongsReq{
		Offset:      offset,
		Limit:       limit,
		Group:       helpers.ToString(ctx.GetQuery("group")),
		Name:        helpers.ToString(ctx.GetQuery("name")),
		ReleaseDate: releaseDate,
		Text:        helpers.ToString(ctx.GetQuery("text")),
		Link:        helpers.ToString(ctx.GetQuery("link")),
	})
	if err != nil {

		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, res)

}
func (controller *Controller) updateSong(ctx *gin.Context) {
	req := new(dto.UpdateSongReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	err := controller.service.UpdateSong(req)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(404, gin.H{"message": "Song not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(204)
}
func (controller *Controller) removeSong(ctx *gin.Context) {
	req, err := helpers.ToUuid(ctx.Param("song_id"), true)
	if err != nil {
		ctx.JSON(400, gin.H{"message": fmt.Sprintf("Controller: %s", err)})
		return
	}
	err = controller.service.RemoveSong(req)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(404, gin.H{"message": "Song not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(204)
}
