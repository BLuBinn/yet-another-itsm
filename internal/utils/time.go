package utils

import (
	"time"

	"yet-another-itsm/internal/constants"

	"github.com/jackc/pgx/v5/pgtype"
)

// Helper function to convert pgtype.Timestamptz to time.Time
func ConvertPgTimestamp(pgTime pgtype.Timestamptz) time.Time {
	if pgTime.Valid {
		return pgTime.Time
	}
	return time.Time{} // Return zero time for invalid timestamps
}

// FormatTime converts a time.Time to ISO8601 string format.
func FormatTime(t time.Time) string {
	return t.Format(constants.Iso8601Format)
}
