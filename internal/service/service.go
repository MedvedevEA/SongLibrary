package service

import (
	"songLibrary/internal/model"
	"songLibrary/internal/outsideapi"
	outsideapiDto "songLibrary/internal/outsideapi/dto"
	repositoryDto "songLibrary/internal/repository/dto"
	"songLibrary/internal/store"
)

type Service struct {
	store      *store.Store
	outsideApi *outsideapi.OutsideApi
}

func New(store *store.Store, outsideApi *outsideapi.OutsideApi) *Service {
	return &Service{
		store:      store,
		outsideApi: outsideApi,
	}
}

func (s *Service) AddGroup(req *repositoryDto.AddGroup) (*model.Group, error) {
	return s.store.AddGroup(req)
}
func (s *Service) GetGroup(req *repositoryDto.GetGroup) (*model.Group, error) {
	return s.store.GetGroup(req)
}
func (s *Service) GetGroups(req *repositoryDto.GetGroups) (*model.Pagination[model.Group], error) {
	return s.store.GetGroups(req)
}
func (s *Service) UpdateGroup(req *repositoryDto.UpdateGroup) error {
	return s.store.UpdateGroup(req)
}
func (s *Service) RemoveGroup(req *repositoryDto.RemoveGroup) error {
	return s.store.RemoveGroup(req)
}

func (s *Service) AddSong(req *repositoryDto.AddSong) (*model.Song, error) {
	//Запрос на внешний api сервер
	outsideApiRes, err := s.outsideApi.GetInfo(&outsideapiDto.GetInfoReq{
		Group: req.Group,
		Song:  req.Name,
	})
	if err == nil {
		return s.store.AddSong(&repositoryDto.AddSong{
			Group:       req.Group,
			Name:        req.Name,
			ReleaseDate: outsideApiRes.ReleaseDate,
			Text:        outsideApiRes.Text,
			Link:        outsideApiRes.Link,
		})

	}

	return s.store.AddSong(req)
}
func (s *Service) GetSong(req *repositoryDto.GetSong) (*model.Song, error) {
	return s.store.GetSong(req)
}
func (s *Service) GetSongText(req *repositoryDto.GetSongText) (*model.Pagination[model.Verse], error) {
	return s.store.GetSongText(req)
}
func (s *Service) GetSongs(req *repositoryDto.GetSongs) (*model.Pagination[model.Song], error) {
	return s.store.GetSongs(req)
}
func (s *Service) UpdateSong(req *repositoryDto.UpdateSong) error {
	return s.store.UpdateSong(req)
}
func (service *Service) RemoveSong(req *repositoryDto.RemoveSong) error {
	return service.store.RemoveSong(req)
}
