package uuid

import "github.com/google/uuid"

func Gen() uuid.UUID {
	return uuid.New()
}

func GenStr() string {
	return uuid.New().String()
}
