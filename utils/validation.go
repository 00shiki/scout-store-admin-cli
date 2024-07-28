package utils

import (
	"errors"
	"regexp"
)

func ValidateDate(date string) error {
	regex, err := regexp.Compile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`)
	if err != nil {
		return err
	}
	valid := regex.MatchString(date)
	if !valid {
		return errors.New("invalid date format")
	}
	return nil
}
