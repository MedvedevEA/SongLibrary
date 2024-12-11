package dto

import (
	"songLibrary/internal/pkg/types"

	"github.com/google/uuid"
)

type AddSong struct {
	Group       string      `json:"group" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	ReleaseDate *types.Date `form:"release_date"`
	Text        *string     `form:"text"`
	Link        *string     `form:"link"`
}
type GetSong struct {
	SongId string `uri:"song_id" binding:"required,uuid4_rfc4122"`
}
type GetSongText struct {
	SongId string `uri:"song_id" binding:"required,uuid4_rfc4122"`
	Limit  *int   `form:"limit"`
	Offset *int   `form:"offset"`
}
type GetSongs struct {
	Limit       *int        `form:"limit"`
	Offset      *int        `form:"offset"`
	Group       *string     `form:"group"`
	Name        *string     `form:"name"`
	ReleaseDate *types.Date `form:"release_date"`
	Text        *string     `form:"text"`
	Link        *string     `form:"link"`
}

type UpdateSong struct {
	SongId        *uuid.UUID  `json:"song_id" binding:"required"`
	Group         *string     `json:"group"`
	Name          *string     `json:"name"`
	ReleaseDate   *types.Date `json:"release_date"`
	Text          *string     `json:"text"`
	Link          *string     `json:"link"`
	SetRelaseDate *bool       `json:"set_release_date"`
	SetText       *bool       `json:"set_text"`
	SetLink       *bool       `json:"set_link"`
}
type RemoveSong struct {
	SongId string `uri:"song_id" binding:"required,uuid4_rfc4122"`
}
