package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Helper function to convert pgtype.Timestamptz to time.Time
func ConvertPgTimestamp(pgTime pgtype.Timestamptz) time.Time {
	if pgTime.Valid {
		return pgTime.Time
	}
	return time.Time{} // Return zero time for invalid timestamps
}
