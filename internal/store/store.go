package store

import (
	"database/sql"
	"errors"
	"songLibrary/internal/model"
	"songLibrary/internal/pkg/helpers"
	"songLibrary/internal/pkg/servererrors"
	"songLibrary/internal/repository/dto"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db,
	}
}
func (s *Store) AddGroup(req *dto.AddGroup) (*model.Group, error) {
	return helpers.StoreJsonRequest[dto.AddGroup, model.Group](s.db, "public.add_group", req)
}
func (s *Store) GetGroup(req *dto.GetGroup) (*model.Group, error) {
	res, err := helpers.StoreJsonRequest[dto.GetGroup, model.Group](s.db, "public.get_group", req)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, servererrors.ErrorRecordNotFound
	}
	return res, err
}
func (s *Store) GetGroups(req *dto.GetGroups) (*model.Pagination[model.Group], error) {
	return helpers.StoreJsonRequest[dto.GetGroups, model.Pagination[model.Group]](s.db, "public.get_groups", req)
}
func (s *Store) UpdateGroup(req *dto.UpdateGroup) error {
	_, err := helpers.StoreJsonRequest[dto.UpdateGroup, model.Group](s.db, "public.update_group", req)
	if errors.Is(err, sql.ErrNoRows) {
		return servererrors.ErrorRecordNotFound
	}
	return err
}
func (s *Store) RemoveGroup(req *dto.RemoveGroup) error {
	_, err := helpers.StoreJsonRequest[dto.RemoveGroup, model.Group](s.db, "public.remove_group", req)
	if errors.Is(err, sql.ErrNoRows) {
		return servererrors.ErrorRecordNotFound
	}
	return err
}

func (s *Store) AddSong(req *dto.AddSong) (*model.Song, error) {
	return helpers.StoreJsonRequest[dto.AddSong, model.Song](s.db, "public.add_song", req)
}
func (s *Store) GetSong(req *dto.GetSong) (*model.Song, error) {
	res, err := helpers.StoreJsonRequest[dto.GetSong, model.Song](s.db, "public.get_song", req)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, servererrors.ErrorRecordNotFound
	}
	return res, err
}
func (s *Store) GetSongText(req *dto.GetSongText) (*model.Pagination[model.Verse], error) {
	res, err := helpers.StoreJsonRequest[dto.GetSongText, model.Pagination[model.Verse]](s.db, "public.get_song_text", req)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, servererrors.ErrorRecordNotFound
	}
	return res, err
}
func (s *Store) GetSongs(req *dto.GetSongs) (*model.Pagination[model.Song], error) {
	return helpers.StoreJsonRequest[dto.GetSongs, model.Pagination[model.Song]](s.db, "public.get_songs", req)
}
func (s *Store) UpdateSong(req *dto.UpdateSong) error {
	_, err := helpers.StoreJsonRequest[dto.UpdateSong, model.Song](s.db, "public.update_song", req)
	if errors.Is(err, sql.ErrNoRows) {
		return servererrors.ErrorRecordNotFound
	}
	return err
}
func (s *Store) RemoveSong(req *dto.RemoveSong) error {
	_, err := helpers.StoreJsonRequest[dto.RemoveSong, model.Song](s.db, "public.remove_song", req)
	if errors.Is(err, sql.ErrNoRows) {
		return servererrors.ErrorRecordNotFound
	}
	return err
}
