package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/TimKotowski/future-take-home/internal/appointment"
)

var (
	_ RouteRegister = &appointmentRouteRegister{}
)

type appointmentRouteRegister struct {
	controller appointment.AppointmentController
}

func NewAppointmentRouteRegister(db *sqlx.DB) RouteRegister {
	return &appointmentRouteRegister{
		controller: appointment.NewAppointmentController(
			appointment.NewAppointments(db),
		),
	}
}

func (c appointmentRouteRegister) RegisterRoutes(router *chi.Mux) {
	router.Get("/appointments/v1/slots/{trainerID}/{startSlot}/{endSlot}", c.controller.GetAppointmentsByTimeRange)
	router.Get("/appointments/v1/{trainerID}", c.controller.GetAppointmentsByTrainer)
	router.Post("/appointments/v1/slots/{trainerID}/{userID}", c.controller.CreateAppointments)
}
