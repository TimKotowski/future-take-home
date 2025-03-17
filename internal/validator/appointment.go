package validator

import (
	"errors"
	"fmt"

	"github.com/TimKotowski/future-take-home/internal/utils"
)

var (
	_ Validator[AppointmentsTimeRangeValidator] = &AppointmentsTimeRangeValidator{}
)

type AppointmentsTimeRangeValidator struct {
	StartSlot string `json:"start_slot"`
	EndSlot   string `json:"end_slot"`
}

func (v AppointmentsTimeRangeValidator) Validate() error {
	start, end, err := utils.ParseStartEndSlots(v.StartSlot, v.EndSlot)
	if err != nil {
		return err
	}
	if start.After(end) {
		return errors.New(fmt.Sprintf("%s, start slot is after end slot. Not allowed", errValidation))
	}

	return nil
}
