package utils

import (
	"github.com/gofrs/uuid"
	"log"
	"time"
)

func StringToTime(str string) time.Time {
	t, err := time.Parse("15:04", str)
	if err != nil {
		log.Println(err)
		return time.Time{}
	}
	return t
}

func StringToUUID(str string) uuid.UUID {
	u, err := uuid.FromString(str)
	if err != nil {
		log.Println(err)
		return uuid.UUID{}
	}
	return u
}
