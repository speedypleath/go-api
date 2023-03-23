package model

import (
	"github.com/google/uuid"
)

type Todo struct {
	Base
	ID   uuid.UUID
	Text string
	Done bool
	User *User
}
