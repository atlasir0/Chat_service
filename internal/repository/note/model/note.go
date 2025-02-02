package modelRepo

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Role      int          `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}


type UserFilter struct {
	ID   *int64
	Name *string
}
