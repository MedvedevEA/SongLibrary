package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	controllerDto "songLibrary/internal/controller/dto"
	outsideserverDto "songLibrary/internal/outsideserver/dto"
	repositoryDto "songLibrary/internal/repository/dto"

	"songLibrary/internal/model"
	"songLibrary/internal/outsideserver"
	"songLibrary/internal/store"
)

type Service struct {
	store         *store.Store
	outsideServer *outsideserver.OutsideServer
}

func New(store *store.Store, outsideServer *outsideserver.OutsideServer) *Service {
	return &Service{
		store,
		outsideServer,
	}
}
func (service *Service) AddSong(controllerReq *controllerDto.AddSongReq) (*controllerDto.AddSongRes, error) {
	//Запрос на внешний api сервер
	outsideServerRes, err := service.outsideServer.GetInfo(&outsideserverDto.GetInfoReq{
		Group: controllerReq.Group,
		Song:  controllerReq.Name,
	})
	if err != nil {
		return nil, err
	}
	//Создание новой записи на основании данных полученных от внешнего сервера
	var outsideServerReleaseDate *time.Time = nil
	if outsideServerRes.ReleaseDate != "" {
		date, err := time.Parse("02-01-2006", outsideServerRes.ReleaseDate)
		if err != nil {
			return nil, fmt.Errorf("AddSongService: %s", err)
		}
		outsideServerReleaseDate = &date
	}

	repositoryRes, err := service.store.AddSong(&repositoryDto.AddSongReq{
		Group:       controllerReq.Group,
		Name:        controllerReq.Name,
		ReleaseDate: outsideServerReleaseDate,
		Text:        strings.Split(outsideServerRes.Text, "\n\n"),
		Link:        &outsideServerRes.Link,
	})
	if err != nil {
		return nil, err
	}

	var releaseDate *string = nil
	if repositoryRes.ReleaseDate != nil {
		date := repositoryRes.ReleaseDate.Format("02-01-2006")
		releaseDate = &date
	}
	return &controllerDto.AddSongRes{
		SongId:      repositoryRes.SongId,
		Group:       repositoryRes.Group,
		Name:        repositoryRes.Name,
		ReleaseDate: releaseDate,
		Text:        strings.Join(repositoryRes.Text, "\n\n"),
		Link:        repositoryRes.Link,
	}, nil
}
func (service *Service) GetInfo(controllerReq *controllerDto.GetInfoReq) (*controllerDto.GetInfoRes, error) {
	repositoryRes, err := service.store.GetInfo(&repositoryDto.GetInfoReq{
		Group: controllerReq.Group,
		Name:  controllerReq.Song,
	})
	if err != nil {
		return nil, err
	}
	var releaseDate *string = nil
	if repositoryRes.ReleaseDate != nil {
		formatData := repositoryRes.ReleaseDate.Format("02-01-2006")
		releaseDate = &formatData
	}
	return &controllerDto.GetInfoRes{
		ReleaseDate: releaseDate,
		Text:        strings.Join(repositoryRes.Text, "\n\n"),
		Link:        repositoryRes.Link,
	}, nil
}
func (service *Service) GetSong(controllerReq *uuid.UUID) (*controllerDto.GetSongRes, error) {
	repositoryRes, err := service.store.GetSong(controllerReq)
	if err != nil {
		return nil, err
	}
	var releaseDate *string = nil
	if repositoryRes.ReleaseDate != nil {
		formatData := repositoryRes.ReleaseDate.Format("02-01-2006")
		releaseDate = &formatData
	}
	return &controllerDto.GetSongRes{
		SongId:      repositoryRes.SongId,
		Group:       repositoryRes.Group,
		Name:        repositoryRes.Name,
		ReleaseDate: releaseDate,
		Text:        strings.Join(repositoryRes.Text, "\n\n"),
		Link:        repositoryRes.Link,
	}, nil

}
func (service *Service) GetSongText(controllerReq *controllerDto.GetSongTextReq) ([]*controllerDto.GetSongTextRes, error) {
	offset := 0
	if controllerReq.Offset != nil {
		offset = *controllerReq.Offset
	}
	limit := 5
	if controllerReq.Limit != nil {
		limit = *controllerReq.Limit
	}
	repositoryRes, err := service.store.GetSongText(&repositoryDto.GetSongTextReq{
		SongId: controllerReq.SongId,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	controllerRes := []*controllerDto.GetSongTextRes{}
	for _, value := range *repositoryRes {
		controllerRes = append(controllerRes, &controllerDto.GetSongTextRes{Verse: value})
	}
	return controllerRes, nil
}
func (service *Service) GetSongs(controllerReq *controllerDto.GetSongsReq) ([]*controllerDto.GetSongsRes, error) {
	offset := 0
	if controllerReq.Offset != nil {
		offset = *controllerReq.Offset
	}
	limit := 10
	if controllerReq.Limit != nil {
		limit = *controllerReq.Limit
	}
	var releaseDate *time.Time = nil
	if controllerReq.ReleaseDate != nil {
		releaseDate = &controllerReq.ReleaseDate.Time
	}
	repositoryRes, err := service.store.GetSongs(&repositoryDto.GetSongsReq{
		Offset:      offset,
		Limit:       limit,
		Group:       controllerReq.Group,
		Name:        controllerReq.Name,
		ReleaseDate: releaseDate,
		Text:        controllerReq.Text,
		Link:        controllerReq.Link,
	})
	if err != nil {
		return nil, err
	}
	controllerRes := []*controllerDto.GetSongsRes{}
	for _, value := range repositoryRes {
		var releaseDate *string = nil
		if value.ReleaseDate != nil {
			formatData := value.ReleaseDate.Format("02-01-2006")
			releaseDate = &formatData
		}
		controllerRes = append(controllerRes, &controllerDto.GetSongsRes{
			SongId:      value.SongId,
			Group:       value.Group,
			Name:        value.Name,
			ReleaseDate: releaseDate,
			Text:        strings.Join(value.Text, "\n\n"),
			Link:        value.Link,
		})

	}
	return controllerRes, nil
}
func (service *Service) UpdateSong(controllerReq *controllerDto.UpdateSongReq) error {
	var releaseDate *time.Time = nil
	if controllerReq.ReleaseDate != nil {
		releaseDate = &controllerReq.ReleaseDate.Time
	}
	return service.store.UpdateSong(&model.Song{
		SongId:      controllerReq.SongId,
		Group:       controllerReq.Group,
		Name:        controllerReq.Name,
		ReleaseDate: releaseDate,
		Text:        strings.Split(controllerReq.Text, "\n\n"),
		Link:        controllerReq.Link,
	})
}
func (service *Service) RemoveSong(req *uuid.UUID) error {
	return service.store.RemoveSong(req)

}
