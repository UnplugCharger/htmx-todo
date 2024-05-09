package utils

import "github.com/jackc/pgx/v5/pgtype"

// Convert pgtype.Timestamp to string
func ConvertTimestampToString(timestamp pgtype.Timestamp) string {
	return timestamp.Time.String()
}
