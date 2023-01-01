package util

import "github.com/google/uuid"

type uuidGenerator struct {
}

func NewUUIDGenerator() *uuidGenerator {
	return &uuidGenerator{}
}
func (u *uuidGenerator) GenerateUUID() string {
	return uuid.New().String()
}
