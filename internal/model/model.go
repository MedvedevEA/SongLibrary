package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Song struct {
	SongId      *uuid.UUID     `db:"song_id"`
	Group       string         `db:"group"`
	Name        string         `db:"name"`
	ReleaseDate *time.Time     `db:"release_date"`
	Text        pq.StringArray `db:"text"`
	Link        *string        `db:"link"`
}
