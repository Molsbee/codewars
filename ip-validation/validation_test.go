package ip_validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidation(t *testing.T) {
	assert := assert.New(t)
	assert.False(is_valid_ip("1.2.3"))
	assert.False(is_valid_ip("1.2.3.4.5"))
	assert.False(is_valid_ip("123.456.78.90"))
	assert.False(is_valid_ip("123.045.067.089"))
	assert.False(is_valid_ip("300.300.300.300"))
	assert.True(is_valid_ip("10.121.12.15"))
	assert.True(is_valid_ip("1.2.3.4"))
}
