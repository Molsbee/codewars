package encrypt_this

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptThis(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("65 119esi 111dl 111lw 108dvei 105n 97n 111ka", EncryptThis("A wise old owl lived in an oak"))
	assert.Equal("84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp", EncryptThis("The more he saw the less he spoke"))
	assert.Equal("84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare", EncryptThis("The less he spoke the more he heard"))
	assert.Equal("87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri", EncryptThis("Why can we not all be like that wise old bird"))
	assert.Equal("84kanh 121uo 80roti 102ro 97ll 121ruo 104ple", EncryptThis("Thank you Piotr for all your help"))
	assert.Equal("65", EncryptThis("A"))
}
