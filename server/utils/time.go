package utils

import (
	"time"
)

func FormatTime(timeStr string) string {
	// Parse the input time string in RFC3339 format
	t, _ := time.Parse(time.RFC3339, timeStr)

	// Format to desired layout: "15:04 02/01/2006"
	formattedTime := t.Format("15:04 02/01/2006")
	return formattedTime
}
