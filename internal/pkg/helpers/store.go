package helpers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"songLibrary/internal/pkg/servererrors"

	log "github.com/sirupsen/logrus"
)

func StoreJsonRequest[Req any, Res any](db *sql.DB, funcName string, req *Req) (*Res, error) {
	j, err := json.Marshal(req)
	if err != nil {
		log.Errorf("store: %s: %s", funcName, err)
		return nil, servererrors.ErrorInternal
	}
	query := fmt.Sprintf("Select * FROM public.%s($1)", funcName)
	row := db.QueryRow(query, j)
	if err := row.Scan(&j); err != nil {
		log.Errorf("store: %s: %s", funcName, err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, servererrors.ErrorRecordNotFound
		}
		return nil, servererrors.ErrorInternal
	}
	res := new(Res)
	err = json.Unmarshal(j, res)
	if err != nil {
		log.Errorf("store: %s: %s", funcName, err)
	}
	return res, err
}
