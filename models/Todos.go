package models

import (
	"database/sql"
)

type Todos struct {
	ID        int          `db:"id"`
	Title     string       `db:"title"`
	Completed sql.NullBool `db:"completed"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
