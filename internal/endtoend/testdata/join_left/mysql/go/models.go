// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package querytest

import (
	"database/sql"
	"time"
)

type Author struct {
	ID       int32
	Name     string
	ParentID sql.NullInt32
}

type City struct {
	CityID  int32
	MayorID int32
}

type Mayor struct {
	MayorID  int32
	FullName string
}

type Medium struct {
	MediaID        int32
	MediaCreatedAt time.Time
	MediaHash      string
	MediaDirectory string
	MediaAuthorID  int32
	MediaWidth     int32
	MediaHeight    int32
}

type SuperAuthor struct {
	SuperID       int32
	SuperName     string
	SuperParentID sql.NullInt32
}

type User struct {
	UserID int32
	CityID sql.NullInt32
}

type Users2 struct {
	UserID          int32
	UserNickname    string
	UserEmail       string
	UserDisplayName string
	UserPassword    sql.NullString
	UserGoogleID    sql.NullString
	UserAppleID     sql.NullString
	UserBio         string
	UserCreatedAt   time.Time
	UserAvatarID    sql.NullInt32
}
