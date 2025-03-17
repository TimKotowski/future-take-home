package appointment

import (
	"errors"
	"time"
)

type TimeChecker struct {
	openingTime time.Time
	closingTime time.Time
	location    *time.Location
	times       timeTuple
}

type timeTuple struct {
	startSlot time.Time
	endSlot   time.Time
}

var (
	NINE_AM = "09:00"
	FIVE_PM = "17:00"
)

func NewTimeChecker(startSlot, endSlot time.Time) (*TimeChecker, error) {
	openingTime, err := time.Parse("15:04", NINE_AM)
	if err != nil {
		return nil, err
	}

	closingTime, err := time.Parse("15:04", FIVE_PM)
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return nil, err
	}

	return &TimeChecker{
		openingTime: openingTime,
		closingTime: closingTime,
		location:    loc,
		times: timeTuple{
			startSlot: startSlot,
			endSlot:   endSlot,
		},
	}, nil
}

// Schedules only during business hours.
func (t *TimeChecker) isPassedBusinessHours() error {
	start := t.times.startSlot.In(t.location)
	end := t.times.endSlot.In(t.location)

	startTime, err := time.Parse(time.TimeOnly, start.Format(time.TimeOnly))
	if err != nil {
		return err
	}

	endTime, err := time.Parse(time.TimeOnly, end.Format(time.TimeOnly))
	if err != nil {
		return err
	}

	if startTime.Before(t.openingTime) || startTime.After(t.closingTime) {
		return errors.New("appointment start time must be scheduled during business hours")
	}

	if endTime.Before(t.openingTime) || endTime.After(t.closingTime) {
		return errors.New("appointments end time must be scheduled during business hours")
	}

	return nil
}

// Should be scheduled at :00, :30
func (t *TimeChecker) isScheduledAtAppropriateMinutes() error {
	if t.times.startSlot.Minute() != 30 && t.times.startSlot.Minute() != 0 {
		return errors.New("appointments start slot must be scheduled on the hour (:00) or half hour (:30) during business hours")
	}
	if t.times.endSlot.Minute() != 30 && t.times.endSlot.Minute() != 0 {
		return errors.New("appointments end slot must be scheduled on the hour (:00) or half hour (:30) during business hours")
	}

	return nil
}

// All appointments are 30 minutes long
func (t *TimeChecker) isScheduledAtAppropriateTimeBlocks() error {
	appointmentLength := t.times.endSlot.Sub(t.times.startSlot)
	if appointmentLength.Minutes() < 30 || appointmentLength.Minutes() > 30 {
		return errors.New("appointments must be 30 minutes long")
	}

	return nil
}
