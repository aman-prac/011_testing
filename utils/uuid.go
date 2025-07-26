package uniqueIdentity

import "github.com/google/uuid"

func NewUniqueID() string {
	return generateUUID()
}

func generateUUID() string {
	return uuid.NewString()
}
