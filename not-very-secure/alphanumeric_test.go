package not_very_secure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlphanumeric(t *testing.T) {
	assert := assert.New(t)
	assert.False(alphanumeric(" "))
	assert.False(alphanumeric(".*?"))
	assert.True(alphanumeric("a"))
	assert.False(alphanumeric("hello world_"))
	assert.True(alphanumeric("Mazinkaiser"))
	assert.True(alphanumeric("PassW0rd"))
	assert.False(alphanumeric("     "))
	assert.False(alphanumeric(""))
	assert.False(alphanumeric("\n\t\n"))
	assert.False(alphanumeric("ciao\n$$_"))
	assert.False(alphanumeric("__ * __"))
	assert.False(alphanumeric("&)))((("))
	assert.True(alphanumeric("43534h56jmTHHF3K"))
}
