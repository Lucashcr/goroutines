package entities

import "github.com/google/uuid"

type Task struct {
	Id    uuid.UUID
	Value int
}
