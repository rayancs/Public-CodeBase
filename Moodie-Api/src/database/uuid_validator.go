package database

import "github.com/google/uuid"
// can not place the uuid validatio in the service layer as this would make a circular flow that is not allowed in golang
func IsValidUUID(u string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(u)
	if err != nil {
		return uuid.Nil, err
	}
	return parsedUUID, err
}
