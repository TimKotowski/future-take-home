package utils

import "time"

func ParseStartEndSlots(startSlot, endSlot string) (time.Time, time.Time, error) {
	start, err := time.Parse(time.RFC3339, startSlot)
	if err != nil {
		return time.Time{}, time.Time{}, nil
	}
	end, err := time.Parse(time.RFC3339, endSlot)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return start, end, nil
}
