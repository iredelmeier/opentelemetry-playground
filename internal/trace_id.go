package internal

import "github.com/gofrs/uuid"

func NewTraceID() [16]byte {
	u, _ := uuid.NewV4()

	return u
}
