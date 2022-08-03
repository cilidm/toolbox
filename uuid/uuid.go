package uuid

import (
	"github.com/google/uuid"
	"log"
)

// UUID Define alias
type UUID = uuid.UUID

// NewUUID Create uuid
func NewUUID() (UUID, error) {
	return uuid.NewRandom()
}

// MustUUID Create uuid(Throw panic if something goes wrong)
func MustUUID() UUID {
	v, err := NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return v
}

// MustString Create uuid
func MustString() string {
	return MustUUID().String()
}
