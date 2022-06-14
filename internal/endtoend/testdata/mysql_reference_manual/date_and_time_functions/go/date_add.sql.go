// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: date_add.sql

package date_and_time_functions

import (
	"context"
	"time"
)

const dateAddDayHour = `-- name: DateAddDayHour :one
SELECT DATE_ADD('1900-01-01 00:00:00',
                INTERVAL '-1 10' DAY_HOUR)
`

func (q *Queries) DateAddDayHour(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, dateAddDayHour)
	var date_add time.Time
	err := row.Scan(&date_add)
	return date_add, err
}

const dateAddMinuteSecond = `-- name: DateAddMinuteSecond :one
SELECT DATE_ADD('2100-12-31 23:59:59',
                INTERVAL '1:1' MINUTE_SECOND)
`

func (q *Queries) DateAddMinuteSecond(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, dateAddMinuteSecond)
	var date_add time.Time
	err := row.Scan(&date_add)
	return date_add, err
}

const dateAddOneDay = `-- name: DateAddOneDay :one

SELECT DATE_ADD('2018-05-01',INTERVAL 1 DAY)
`

// https://dev.mysql.com/doc/refman/8.0/en/date-and-time-functions.html#function_date-add
func (q *Queries) DateAddOneDay(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, dateAddOneDay)
	var date_add time.Time
	err := row.Scan(&date_add)
	return date_add, err
}

const dateAddOneSecond = `-- name: DateAddOneSecond :one
SELECT DATE_ADD('2020-12-31 23:59:59',
                INTERVAL 1 SECOND)
`

func (q *Queries) DateAddOneSecond(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, dateAddOneSecond)
	var date_add time.Time
	err := row.Scan(&date_add)
	return date_add, err
}

const dateAddSecondMicrosecond = `-- name: DateAddSecondMicrosecond :one
SELECT DATE_ADD('1992-12-31 23:59:59.000002',
           INTERVAL '1.999999' SECOND_MICROSECOND)
`

func (q *Queries) DateAddSecondMicrosecond(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, dateAddSecondMicrosecond)
	var date_add time.Time
	err := row.Scan(&date_add)
	return date_add, err
}

const dateAddTimestampOneSecond = `-- name: DateAddTimestampOneSecond :one
SELECT DATE_ADD('2018-12-31 23:59:59',
                INTERVAL 1 DAY)
`

func (q *Queries) DateAddTimestampOneSecond(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, dateAddTimestampOneSecond)
	var date_add time.Time
	err := row.Scan(&date_add)
	return date_add, err
}
