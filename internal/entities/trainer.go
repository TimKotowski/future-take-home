package entities

import (
	"database/sql"
	"math/big"
)

type Trainer struct {
	Id                big.Int        `db:"id" json:"id"`
	FirstName         string         `db:"first_name" json:"first_name"`
	LastName          sql.NullString `db:"last_name" json:"last_name"`
	Specialization    sql.NullString `db:"specialization" json:"specialization"`
	YearsOfExperience int            `db:"years_of_experience" json:"years_of_experience"`
}
