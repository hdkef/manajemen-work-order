package services

import (
	"errors"
	"net/mail"
)

func IsNotEmpty(value ...string) error {
	for _, val := range value {
		if val == "" {
			return errors.New("EMPTY")
		}
	}
	return nil
}

func IsEmail(value string) error {
	_, err := mail.ParseAddress(value)
	return err
}
