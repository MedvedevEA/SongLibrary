package repository

import (
	"songLibrary/internal/model"
	"songLibrary/internal/repository/dto"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Repository interface {
	AddSong(*dto.AddSongReq) (*model.Song, error)
	GetInfo(*dto.GetInfoReq) (*model.Song, error)
	GetSong(*uuid.UUID) (*model.Song, error)
	GetSongText(*dto.GetSongTextReq) (*pq.StringArray, error)
	GetSongs(*dto.GetSongsReq) ([]*model.Song, error)
	UpdateSong(*model.Song) error
	RemoveSong(*uuid.UUID) error
}
