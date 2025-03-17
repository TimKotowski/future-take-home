package entities

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type AppointmentStatus = string

const (
	ACTIVE    AppointmentStatus = "ACTIVE"
	CANCELLED AppointmentStatus = "CANCELLED"
	COMPLETED AppointmentStatus = "COMPLETED"
)

type Appointment struct {
	Id        uuid.UUID          `db:"id"`
	TrainerId int                `db:"trainer_id"`
	UserId    int                `db:"user_id"`
	StartSlot pgtype.Timestamptz `db:"start_slot"`
	EndSlot   pgtype.Timestamptz `db:"end_slot"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at"`
	Status    AppointmentStatus  `db:"status"`
}
