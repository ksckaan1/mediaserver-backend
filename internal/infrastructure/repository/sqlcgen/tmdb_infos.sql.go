// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tmdb_infos.sql

package sqlcgen

import (
	"context"
)

const getTMDBInfo = `-- name: GetTMDBInfo :one
SELECT id, data FROM tmdb_infos WHERE id = ?
`

func (q *Queries) GetTMDBInfo(ctx context.Context, id string) (TmdbInfo, error) {
	row := q.db.QueryRowContext(ctx, getTMDBInfo, id)
	var i TmdbInfo
	err := row.Scan(&i.ID, &i.Data)
	return i, err
}

const setTMDBInfo = `-- name: SetTMDBInfo :exec
INSERT INTO tmdb_infos (id, data)
		VALUES(?1, ?2) ON CONFLICT (id)
		DO
		UPDATE
		SET
			data = ?2
`

type SetTMDBInfoParams struct {
	ID   string
	Data []byte
}

func (q *Queries) SetTMDBInfo(ctx context.Context, arg SetTMDBInfoParams) error {
	_, err := q.db.ExecContext(ctx, setTMDBInfo, arg.ID, arg.Data)
	return err
}
