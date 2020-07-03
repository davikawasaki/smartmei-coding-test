package dateformat

import "time"

func GetNowFormattedDate(format string) string {
	if format == "" {
		format = "2006-01-02"
	}
	return time.Now().Format(format)
}
