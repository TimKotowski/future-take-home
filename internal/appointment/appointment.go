package appointment

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/TimKotowski/future-take-home/internal/entities"
	"github.com/TimKotowski/future-take-home/internal/utils"
)

var (
	_ Appointments = &appointment{}
)

type appointment struct {
	repository AppointmentsRepository
}

type Appointments interface {
	GetAppointmentsByTrainer(trainerId int) ([]entities.Appointment, error)
	GetAppointmentsByTimeRange(trainerId int, startSlot, endSlot string) ([]entities.Appointment, error)
	CreateAppointments(trainerId, userid int, startSlot, endSlot, status string) (*entities.Appointment, error)
}

func NewAppointments(db *sqlx.DB) Appointments {
	return &appointment{
		repository: NewAppointmentsRepository(db),
	}
}

func (a appointment) GetAppointmentsByTrainer(trainerId int) ([]entities.Appointment, error) {
	return a.repository.GetAppointmentsByTrainer(trainerId)
}

func (a appointment) GetAppointmentsByTimeRange(
	trainerId int,
	startSlot,
	endSlot string,
) ([]entities.Appointment, error) {
	start, end, err := utils.ParseStartEndSlots(startSlot, endSlot)
	if err != nil {
		return nil, err
	}
	startSlot = start.UTC().Format(time.RFC3339)
	endSlot = end.UTC().Format(time.RFC3339)

	return a.repository.GetAppointmentsByTimeRange(trainerId, startSlot, endSlot)
}

func (a appointment) CreateAppointments(
	trainerId,
	userId int,
	startSlot,
	endSlot string,
	status string,
) (*entities.Appointment, error) {
	start, end, err := utils.ParseStartEndSlots(startSlot, endSlot)
	if err != nil {
		return nil, err
	}
	timeChecker, err := NewTimeChecker(start, end)
	if err != nil {
		return nil, err
	}
	if err := timeChecker.isPassedBusinessHours(); err != nil {
		return nil, err
	}

	if err := timeChecker.isScheduledAtAppropriateMinutes(); err != nil {
		return nil, err
	}

	if err := timeChecker.isScheduledAtAppropriateTimeBlocks(); err != nil {
		return nil, err
	}

	return a.repository.CreateAppointments(trainerId, userId, startSlot, endSlot, status)
}
