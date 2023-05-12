package api

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func FeedTimeFormatValidator(fl validator.FieldLevel) bool {
	timePattern := "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$"
	regex := regexp.MustCompile(timePattern)
	timeStr := fl.Field().String()
	return regex.MatchString(timeStr)
}
