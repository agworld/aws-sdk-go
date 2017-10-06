package s3crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBytes(t *testing.T) {
	b := generateBytes(5)
	assert.Equal(t, 5, len(b))
	b = generateBytes(0)
	assert.Equal(t, 0, len(b))
	b = generateBytes(1024)
	assert.Equal(t, 1024, len(b))
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
