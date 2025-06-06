// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package database

import (
	"context"
)

const getPageViews = `-- name: GetPageViews :one
SELECT page_views FROM stats LIMIT 1
`

func (q *Queries) GetPageViews(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, getPageViews)
	var page_views int32
	err := row.Scan(&page_views)
	return page_views, err
}

const incrementPageViews = `-- name: IncrementPageViews :exec
UPDATE stats SET page_views = page_views + 1 WHERE TRUE
`

func (q *Queries) IncrementPageViews(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, incrementPageViews)
	return err
}
