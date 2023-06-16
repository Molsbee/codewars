package human_readable_duration_format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal("1 second", FormatDuration(1))
	a.Equal("1 minute and 2 seconds", FormatDuration(62))
	a.Equal("2 minutes", FormatDuration(120))
	a.Equal("1 hour", FormatDuration(3600))
	a.Equal("1 hour, 1 minute and 2 seconds", FormatDuration(3662))
	a.Equal("1 hour and 2 seconds", FormatDuration(3602))
	a.Equal("1 hour and 1 minute", FormatDuration(3660))
	a.Equal("1 day", FormatDuration(86400))
	a.Equal("1 year and 1 day", FormatDuration(31622400))
}
