package human_readable_duration_format

import "fmt"

const (
	YEAR   = 365 * DAY
	DAY    = 24 * HOUR
	HOUR   = 60 * MINUTE
	MINUTE = 60
)

func FormatDuration(seconds int64) string {
	years := seconds / YEAR
	seconds -= years * YEAR
	days := seconds / DAY
	seconds -= days * DAY
	hours := seconds / HOUR
	seconds -= hours * HOUR
	minutes := seconds / MINUTE
	seconds -= minutes * MINUTE

	t := []string{
		formatString("second", seconds),
		formatString("minute", minutes),
		formatString("hour", hours),
		formatString("day", days),
		formatString("year", years),
	}

	formattedTime := ""
	separator := " and "
	for _, v := range t {
		if len(v) != 0 {
			if len(formattedTime) > 0 {
				formattedTime = separator + formattedTime
				separator = ", "
			}

			formattedTime = v + formattedTime
		}
	}

	return formattedTime
}

func formatString(unit string, v int64) string {
	if v != 0 {
		if v > 1 {
			unit += "s"
		}
		return fmt.Sprintf("%d %s", v, unit)
	}
	return ""
}
