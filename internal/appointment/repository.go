package appointment

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jmoiron/sqlx"

	"github.com/TimKotowski/future-take-home/internal/database"
	"github.com/TimKotowski/future-take-home/internal/entities"
	"github.com/TimKotowski/future-take-home/internal/queries"
)

var (
	_ AppointmentsRepository = &appointmentsRepository{}
)

type appointmentsRepository struct {
	db *sqlx.DB
}

type AppointmentsRepository interface {
	GetAppointmentsByTrainer(trainerID int) ([]entities.Appointment, error)
	CreateAppointments(trainerId, userid int, startSlot, endSlot, status string) (*entities.Appointment, error)
	GetAppointmentsByTimeRange(trainerId int, startSlot, endSlot string) ([]entities.Appointment, error)
}

func NewAppointmentsRepository(db *sqlx.DB) AppointmentsRepository {
	return appointmentsRepository{
		db: db,
	}
}

func (a appointmentsRepository) GetAppointmentsByTrainer(trainerID int) ([]entities.Appointment, error) {
	var appointments []entities.Appointment
	err := a.db.Select(&appointments, queries.GetAppointmentByTrainerQuery, trainerID, entities.ACTIVE)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (a appointmentsRepository) GetAppointmentsByTimeRange(
	trainerId int,
	startSlot,
	endSlot string,
) ([]entities.Appointment, error) {
	var appointments []entities.Appointment

	err := a.db.Select(
		&appointments,
		queries.GetAppointmentByTrainerAndStartSlotAndEndSlotQuery,
		trainerId,
		startSlot,
		endSlot,
		entities.ACTIVE,
	)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (a appointmentsRepository) CreateAppointments(
	trainerId,
	userId int,
	startSlot,
	endSlot,
	status string,
) (*entities.Appointment, error) {
	var appointment entities.Appointment
	result := a.db.QueryRow(queries.InsertAppointmentQuery, uuid.New(), trainerId, userId, startSlot, endSlot, status)
	if err := result.Scan(
		&appointment.Id,
		&appointment.TrainerId,
		&appointment.UserId,
		&appointment.StartSlot,
		&appointment.EndSlot,
		&appointment.Status,
		&appointment.CreatedAt,
		&appointment.UpdatedAt,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == database.ExclusionViolation {
				return nil, fmt.Errorf("appointment cannot overlap with other appointments")
			}
		}
		return nil, err
	}

	return &appointment, nil
}
