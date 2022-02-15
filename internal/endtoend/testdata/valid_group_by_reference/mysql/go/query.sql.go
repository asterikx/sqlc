// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const listAuthors = `-- name: ListAuthors :many
SELECT   id, name as full_name, bio
FROM     authors
GROUP BY full_name
`

type ListAuthorsRow struct {
	ID       int64
	FullName string
	Bio      sql.NullString
}

func (q *Queries) ListAuthors(ctx context.Context) ([]ListAuthorsRow, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAuthorsRow
	for rows.Next() {
		var i ListAuthorsRow
		if err := rows.Scan(&i.ID, &i.FullName, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthorsIdenticalAlias = `-- name: ListAuthorsIdenticalAlias :many
SELECT   id, name as name, bio
FROM     authors
GROUP BY name
`

func (q *Queries) ListAuthorsIdenticalAlias(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsIdenticalAlias)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMetrics = `-- name: ListMetrics :many
SELECT time_bucket('15 days', time) AS bucket, city_name, AVG(temp_c)
FROM weather_metrics
WHERE DATE_SUB(NOW(), INTERVAL 6 MONTH)
GROUP BY bucket, city_name
ORDER BY bucket DESC
`

type ListMetricsRow struct {
	Bucket   interface{}
	CityName sql.NullString
	Avg      interface{}
}

func (q *Queries) ListMetrics(ctx context.Context) ([]ListMetricsRow, error) {
	rows, err := q.db.QueryContext(ctx, listMetrics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMetricsRow
	for rows.Next() {
		var i ListMetricsRow
		if err := rows.Scan(&i.Bucket, &i.CityName, &i.Avg); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
