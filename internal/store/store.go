package store

import (
	"database/sql"
	"fmt"
	"songLibrary/internal/model"
	"songLibrary/internal/repository/dto"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db,
	}
}
func (store *Store) AddSong(req *dto.AddSongReq) (*model.Song, error) {
	row := store.db.QueryRow("SELECT * FROM public.add_song($1,$2,$3,$4,$5)", req.Group, req.Name, req.ReleaseDate, req.Text, req.Link)
	res := new(model.Song)
	err := row.Scan(&res.SongId, &res.Group, &res.Name, &res.ReleaseDate, &res.Text, &res.Link)
	return res, err
}
func (store *Store) GetInfo(req *dto.GetInfoReq) (*model.Song, error) {
	row := store.db.QueryRow("SELECT * FROM public.get_info($1,$2)", req.Group, req.Name)
	res := new(model.Song)
	err := row.Scan(&res.SongId, &res.Group, &res.Name, &res.ReleaseDate, &res.Text, &res.Link)
	return res, err
}
func (store *Store) GetSong(req *uuid.UUID) (*model.Song, error) {
	row := store.db.QueryRow("SELECT * FROM public.get_song($1)", req)
	res := new(model.Song)
	err := row.Scan(&res.SongId, &res.Group, &res.Name, &res.ReleaseDate, &res.Text, &res.Link)
	return res, err
}
func (store *Store) GetSongText(req *dto.GetSongTextReq) (*pq.StringArray, error) {
	row := store.db.QueryRow("SELECT * FROM public.get_song_text($1,$2,$3)", req.SongId, req.Offset, req.Limit)
	res := new(pq.StringArray)
	err := row.Scan(&res)
	if res == nil {
		return nil, sql.ErrNoRows
	}
	return res, err
}
func (store *Store) GetSongs(req *dto.GetSongsReq) ([]*model.Song, error) {
	rows, err := store.db.Query("SELECT * FROM public.get_songs($1,$2,$3,$4,$5,$6,$7)", req.Offset, req.Limit, req.Group, req.Name, req.ReleaseDate, req.Text, req.Link)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []*model.Song{}
	for rows.Next() {
		row := new(model.Song)
		if err := rows.Scan(&row.SongId, &row.Group, &row.Name, &row.ReleaseDate, &row.Text, &row.Link); err != nil {
			return nil, fmt.Errorf("Store: %s", err)
		}
		res = append(res, row)
	}
	return res, nil
}

func (store *Store) UpdateSong(req *model.Song) error {
	row := store.db.QueryRow("SELECT * FROM public.update_song($1,$2,$3,$4,$5,$6)", req.SongId, req.Group, req.Name, req.ReleaseDate, req.Text, req.Link)
	res := new(uuid.UUID)
	return row.Scan(&res)
}

func (store *Store) RemoveSong(req *uuid.UUID) error {
	row := store.db.QueryRow("SELECT * FROM public.remove_song($1)", req)
	res := new(uuid.UUID)
	return row.Scan(&res)
}
