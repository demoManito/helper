package conversion

import (
	"testing"

	"github.com/atotto/clipboard"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert := assert.New(t)

	upperLowerRun(nil, []string{"ul", "u", "user_id"})
	str, _ := clipboard.ReadAll()
	assert.Equal(str, "UserId")

	upperLowerRun(nil, []string{"ul", "u", "user-id"})
	str, _ = clipboard.ReadAll()
	assert.Equal(str, "UserId")

	upperLowerRun(nil, []string{"ul", "l", "UserID"})
	str, _ = clipboard.ReadAll()
	assert.Equal(str, "user_id")
}
