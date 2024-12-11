package controller

import (
	"errors"
	"fmt"

	controllerDto "songLibrary/internal/controller/dto"
	"songLibrary/internal/pkg/servererrors"
	repositoryDto "songLibrary/internal/repository/dto"
	"songLibrary/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	router  *gin.Engine
	service *service.Service
}

func New(router *gin.Engine, service *service.Service) *Controller {
	c := &Controller{
		router,
		service,
	}

	router.POST("groups", c.addGroup)
	router.GET("groups/:group_id", c.getGroup)
	router.GET("groups", c.getGroups)
	router.PUT("groups", c.updateGroup)
	router.DELETE("groups/:group_id", c.removeGroup)

	router.POST("songs", c.addSong)
	router.GET("songs/:song_id", c.getSong)
	router.GET("songs", c.getSongs)
	router.GET("songs/:song_id/text", c.getSongText)
	router.PUT("songs", c.updateSong)
	router.DELETE("songs/:song_id", c.removeSong)

	return c
}
func (c *Controller) addGroup(ctx *gin.Context) {
	req := new(controllerDto.AddGroup)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	res, err := c.service.AddGroup(&repositoryDto.AddGroup{
		Name: req.Name,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.JSON(200, res)
}
func (c *Controller) getGroup(ctx *gin.Context) {
	req := new(controllerDto.GetGroup)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	res, err := c.service.GetGroup(&repositoryDto.GetGroup{
		GroupId: &groupId,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

func (c *Controller) getGroups(ctx *gin.Context) {
	req := new(controllerDto.GetGroups)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	res, err := c.service.GetGroups(&repositoryDto.GetGroups{
		Offset: req.Offset,
		Limit:  req.Limit,
		Name:   req.Name,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.JSON(200, res)
}

func (c *Controller) updateGroup(ctx *gin.Context) {
	req := new(controllerDto.UpdateGroup)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = c.service.UpdateGroup(&repositoryDto.UpdateGroup{
		GroupId: &groupId,
		Name:    req.Name,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.Status(204)
}

func (c *Controller) removeGroup(ctx *gin.Context) {
	req := new(controllerDto.RemoveGroup)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = c.service.RemoveGroup(&repositoryDto.RemoveGroup{
		GroupId: &groupId,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.Status(204)
}

func (c *Controller) addSong(ctx *gin.Context) {
	req := new(controllerDto.AddSong)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	res, err := c.service.AddSong(&repositoryDto.AddSong{
		Group: req.Group,
		Name:  req.Name,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.JSON(200, res)
}
func (c *Controller) getSong(ctx *gin.Context) {
	req := new(controllerDto.GetSong)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	songId, err := uuid.Parse(req.SongId)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	res, err := c.service.GetSong(&repositoryDto.GetSong{
		SongId: &songId,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.JSON(200, res)
}
func (c *Controller) getSongText(ctx *gin.Context) {
	req := new(controllerDto.GetSongText)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err := ctx.BindQuery(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	songId, err := uuid.Parse(req.SongId)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	res, err := c.service.GetSongText(&repositoryDto.GetSongText{
		SongId: &songId,
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.JSON(200, res)
}
func (c *Controller) getSongs(ctx *gin.Context) {
	req := new(controllerDto.GetSongs)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	res, err := c.service.GetSongs(&repositoryDto.GetSongs{
		Offset:      req.Offset,
		Limit:       req.Limit,
		Group:       req.Group,
		Name:        req.Name,
		ReleaseDate: req.ReleaseDate,
		Text:        req.Text,
		Link:        req.Link,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		fmt.Println(err.Error())
		return
	}
	ctx.JSON(200, res)
}

func (c *Controller) updateSong(ctx *gin.Context) {
	req := new(controllerDto.UpdateSong)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err := c.service.UpdateSong(&repositoryDto.UpdateSong{
		SongId:      req.SongId,
		Group:       req.Group,
		Name:        req.Name,
		ReleaseDate: req.ReleaseDate,
		Text:        req.Text,
		Link:        req.Link,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.Status(204)
}
func (c *Controller) removeSong(ctx *gin.Context) {
	req := new(controllerDto.RemoveSong)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	songId, err := uuid.Parse(req.SongId)
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = c.service.RemoveSong(&repositoryDto.RemoveSong{
		SongId: &songId,
	})

	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"message": servererrors.ErrorInternal.Error()})
		return
	}
	ctx.Status(204)
}
