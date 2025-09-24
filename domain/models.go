package domain

import "github.com/google/uuid"

type Message struct {
	id   int
	text string
	user bool
}

type Chat struct {
	id       uuid.UUID
	name     string
	messages []Message
}
