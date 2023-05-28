package utils

import (
	"github.com/gofrs/uuid"
	"log"
	"time"
	"unicode"
)

func StringToTime(str string) time.Time {
	t, err := time.Parse("15:04", str)
	if err != nil {
		log.Println(err)
		return time.Time{}
	}
	return t
}

func StringToUUID(str string) (*uuid.UUID, error) {
	u, err := uuid.FromString(str)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func UcFirst(str string) string {
	if str == "" {
		return ""
	}
	r := []rune(str)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
