package dto

import (
	"songLibrary/pkg/helpers"

	"github.com/google/uuid"
)

type AddSongReq struct {
	Group string `json:"group"`
	Name  string `json:"name"`
}
type AddSongRes struct {
	SongId      *uuid.UUID    `json:"song_id"`
	Group       string        `json:"group"`
	Name        string        `json:"name"`
	ReleaseDate *helpers.Date `json:"release_date"`
	Text        string        `json:"text"`
	Link        *string       `json:"link"`
}
type GetInfoReq struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
type GetInfoRes struct {
	ReleaseDate *helpers.Date `json:"release_date"`
	Text        string        `json:"text"`
	Link        *string       `json:"link"`
}
type GetSongRes struct {
	SongId      *uuid.UUID    `json:"song_id"`
	Group       string        `json:"group"`
	Name        string        `json:"name"`
	ReleaseDate *helpers.Date `json:"release_date"`
	Text        string        `json:"text"`
	Link        *string       `json:"link"`
}
type GetSongTextReq struct {
	Limit  *int       `json:"limit"`
	Offset *int       `json:"offset"`
	SongId *uuid.UUID `json:"song_id"`
}
type GetSongTextRes struct {
	Verse string `json:"verse"`
}
type GetSongsReq struct {
	Limit       *int          `json:"limit"`
	Offset      *int          `json:"offset"`
	Group       *string       `json:"group"`
	Name        *string       `json:"name"`
	ReleaseDate *helpers.Date `json:"release_date"`
	Text        *string       `json:"text"`
	Link        *string       `json:"link"`
}
type GetSongsRes struct {
	SongId      *uuid.UUID    `json:"song_id"`
	Group       string        `json:"group"`
	Name        string        `json:"name"`
	ReleaseDate *helpers.Date `json:"release_date"`
	Text        string        `json:"text"`
	Link        *string       `json:"link"`
}
type UpdateSongReq struct {
	SongId      *uuid.UUID    `json:"song_id"`
	Group       string        `json:"group"`
	Name        string        `json:"name"`
	ReleaseDate *helpers.Date `json:"release_date"`
	Text        string        `json:"text"`
	Link        *string       `json:"link"`
}
