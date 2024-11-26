package helpers

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

func ToInt(value string, isValue bool) (*int, error) {
	if !isValue {
		return nil, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}
func ToString(value string, isValue bool) *string {
	if !isValue {
		return nil
	}
	return &value
}
func ToDateDDMMYYYY(value string, isValue bool) (*DateDDMMYYYY, error) {
	if !isValue {
		return nil, nil
	}
	dateValue, err := time.Parse("02-01-2006", value)
	if err != nil {
		return nil, err
	}
	return &DateDDMMYYYY{Time: dateValue}, nil
}
func ToUuid(value string, isValue bool) (*uuid.UUID, error) {
	if !isValue {
		return nil, nil
	}
	uuidValue, err := uuid.Parse(value)
	if err != nil {
		return nil, err
	}
	return &uuidValue, nil
}
