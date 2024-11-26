package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type AddSongReq struct {
	Group       string         `db:"group"`
	Name        string         `db:"name"`
	ReleaseDate *time.Time     `db:"release_date"`
	Text        pq.StringArray `db:"text"`
	Link        *string        `db:"link"`
}
type GetInfoReq struct {
	Group string `db:"group"`
	Name  string `db:"name"`
}
type GetSongsReq struct {
	Limit       int        `db:"limit"`
	Offset      int        `db:"offset"`
	Group       *string    `db:"group"`
	Name        *string    `db:"name"`
	ReleaseDate *time.Time `db:"release_date"`
	Text        *string    `db:"text"`
	Link        *string    `db:"link"`
}
type GetSongTextReq struct {
	SongId *uuid.UUID `db:"song_id"`
	Limit  int        `db:"limit"`
	Offset int        `db:"offset"`
}
