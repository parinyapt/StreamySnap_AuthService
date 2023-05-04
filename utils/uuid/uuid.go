package utilsUUID

import "github.com/google/uuid"

func GenerateUUIDv4() string {
	return uuid.New().String()
}