package appointment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/TimKotowski/future-take-home/internal/validator"
)

// AppointmentController embeds Controller to inherit shared dependencies.
type AppointmentController struct {
	Appointment Appointments
}

func NewAppointmentController(a Appointments) AppointmentController {
	return AppointmentController{
		Appointment: a,
	}
}

type getAppointmentsByTimeRangeReqBody struct {
	StartSlot string `json:"start_slot"`
	EndSlot   string `json:"end_slot"`
}

type createAppointmentsReqBody struct {
	StartSlot string `json:"start_slot"`
	EndSlot   string `json:"end_slot"`
	Status    string `json:"status"`
}

func (c *AppointmentController) GetAppointmentsByTrainer(w http.ResponseWriter, r *http.Request) {
	trainerIdParam := chi.URLParam(r, "trainerID")
	trainerId, err := strconv.Atoi(trainerIdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID for trainer, make sure its a valid numerical value"))
		return
	}
	appointments, err := c.Appointment.GetAppointmentsByTrainer(trainerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	results, err := json.Marshal(appointments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(results)
}

func (c *AppointmentController) GetAppointmentsByTimeRange(w http.ResponseWriter, r *http.Request) {
	trainerIdParam := chi.URLParam(r, "trainerID")
	trainerId, err := strconv.Atoi(trainerIdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID for trainer, make sure its a valid numerical value"))
		return
	}

	startSlot := chi.URLParam(r, "startSlot")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID for trainer, make sure its a valid numerical value"))
		return
	}

	endSlot := chi.URLParam(r, "endSlot")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID for trainer, make sure its a valid numerical value"))
		return
	}

	v := validator.AppointmentsTimeRangeValidator{
		StartSlot: startSlot,
		EndSlot:   endSlot,
	}
	if err := v.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	appointments, err := c.Appointment.GetAppointmentsByTimeRange(trainerId, v.StartSlot, v.EndSlot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	results, err := json.Marshal(appointments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(results)
}

func (c *AppointmentController) CreateAppointments(w http.ResponseWriter, r *http.Request) {
	trainerIdParam := chi.URLParam(r, "trainerID")
	trainerId, err := strconv.Atoi(trainerIdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID for trainer, make sure its a valid numerical value"))
		return
	}

	userIdParam := chi.URLParam(r, "userID")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID for userID, make sure its a valid numerical value"))
		return
	}

	var reqBody createAppointmentsReqBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse request body"))
		return
	}
	v := validator.AppointmentsTimeRangeValidator{
		StartSlot: reqBody.StartSlot,
		EndSlot:   reqBody.EndSlot,
	}
	if err := v.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	appointment, err := c.Appointment.CreateAppointments(
		trainerId,
		userId,
		reqBody.StartSlot,
		reqBody.EndSlot,
		reqBody.Status,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	results, err := json.Marshal(appointment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(results)
}
