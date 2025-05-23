// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
	"time"
)

type FlywaySchemaHistory struct {
	InstalledRank int32
	Version       sql.NullString
	Description   string
	Type          string
	Script        string
	Checksum      sql.NullInt32
	InstalledBy   string
	InstalledOn   time.Time
	ExecutionTime int32
	Success       bool
}

type Stat struct {
	PageViews int32
}
