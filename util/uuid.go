package util

import "github.com/gofrs/uuid"

func GenerateUUID() string {
	id := uuid.Must(uuid.NewV4())
	return id.String()
}
