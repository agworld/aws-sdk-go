package s3crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
)

func TestEncodeMaterialDescription(t *testing.T) {
	md := MaterialDescription{}
	md["foo"] = aws.String("bar")
	b, err := md.encodeDescription()
	expected := `{"foo":"bar"}`
	assert.NoError(t, err)
	assert.Equal(t, expected, string(b))
}
func TestDecodeMaterialDescription(t *testing.T) {
	md := MaterialDescription{}
	json := `{"foo":"bar"}`
	err := md.decodeDescription([]byte(json))
	expected := MaterialDescription{
		"foo": aws.String("bar"),
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, md)
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
