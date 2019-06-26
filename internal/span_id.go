package internal

import "github.com/gofrs/uuid"

func NewSpanID() [8]byte {
	u, _ := uuid.NewV4()
	var id [8]byte

	copy(id[:], u[8:])

	return id
}
