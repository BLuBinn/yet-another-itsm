package utils

import (
	"github.com/google/uuid"
)

const ErrorFormat = "%s: %w"

func GetStringValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func ParseUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}
