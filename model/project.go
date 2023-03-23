package model

import (
	"github.com/google/uuid"
)

type Project struct {
	Base
	ID   uuid.UUID
	Name string
}
