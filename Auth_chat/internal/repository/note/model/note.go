package modelRepo

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

type UserInfo struct {
	Id        int64
	Name      string
	Email     string
	Role      int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
