package models

import (
	"database/sql"
)

type GooseDbVersion struct {
	ID        int          `db:"id"`
	VersionID int          `db:"version_id"`
	IsApplied bool         `db:"is_applied"`
	Tstamp    sql.NullTime `db:"tstamp"`
}
