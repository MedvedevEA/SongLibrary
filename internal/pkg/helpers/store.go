package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func StoreJsonRequest[Req any, Res any](db *sql.DB, funcName string, req *Req) (*Res, error) {
	j, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("Select * FROM %s($1)", funcName)
	row := db.QueryRow(sql, j)
	if err := row.Scan(&j); err != nil {
		return nil, err
	}
	res := new(Res)
	err = json.Unmarshal(j, res)
	return res, err
}
