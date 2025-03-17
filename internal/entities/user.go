package entities

import (
	"database/sql"
	"math/big"
)

type User struct {
	Id        big.Int        `db:"id" json:"id"`
	FirstName string         `db:"first_name" json:"first_name"`
	LastName  sql.NullString `db:"last_name" json:"last_name"`
	Email     int            `db:"email" json:"email"`
}
