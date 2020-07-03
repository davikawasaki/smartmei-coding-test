package stringformat

import (
	"fmt"
	dateformat "searchprice/utils/date"
	"strings"
)

func LogMessage(text string, status string) string {
	if status == "" {
		status = "INFO"
	}
	// time.Now().Format("02/01/2006 15:04:05")
	return fmt.Sprintf("[%s] %s | %s", status, dateformat.GetNowFormattedDate("02/01/2006 15:04:05"), text)
}

func ReplaceStringSubvals(val string, replacements map[string]string) string {
	if len(val) == 0 {
		return val
	}

	if len(replacements) == 0 {
		replacements = map[string]string{
			" ":  "",
			"\n": "",
		}
	}

	for from, to := range replacements {
		val = strings.Replace(val, from, to, -1)
	}

	return val
}
