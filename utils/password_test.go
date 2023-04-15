package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	ps1, err := HashPassword("123..COM")
	assert.NoError(t, err)
	t.Logf("生成的密码样例：%s", ps1)
	ps2, err := HashPassword("123..COM")
	assert.NoError(t, err)
	ps3, err := HashPassword("123..com")
	assert.NoError(t, err)

	assert.NotEqual(t, ps1, ps3)
	assert.Equal(t, ps1, ps2)
}
