package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserUpdate struct {
	ID    int64  `db:"id"`
	Name  string `db:"username"`
	Email string `db:"email"`
	Role  int    `db:"role"`
}
