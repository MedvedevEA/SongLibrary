package repository

import (
	"songLibrary/internal/model"
	"songLibrary/internal/repository/dto"

	"github.com/google/uuid"
)

type Reposytory interface {
	AddGroup(name *dto.AddGroup) (*model.Group, error)
	GetGroup(groupId *dto.GetGroup) (*model.Group, error)
	GetGroups(dto *dto.GetGroups) (*model.Pagination[model.Group], error)
	UpdateGroup(dto *dto.UpdateGroup) (*model.Group, error)
	RemoveGroup(groupId *dto.RemoveGroup) (*model.Group, error)

	AddSong(dto *dto.AddSong) (*model.Song, error)
	GetSong(songId *dto.GetSong) (*model.Song, error)
	GetSongText(dto *dto.GetSongText) (*model.Pagination[model.Verse], error)
	GetSongs(dto *dto.GetSongs) (*model.Pagination[model.Song], error)
	UpdateSong(dto *dto.UpdateSong) (*model.Song, error)
	RemoveSong(songId *uuid.UUID) (*model.Song, error)
}
