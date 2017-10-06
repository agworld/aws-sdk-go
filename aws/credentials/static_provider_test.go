package credentials

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaticProviderGet(t *testing.T) {
	s := StaticProvider{
		Value: Value{
			AccessKeyID:     "AKID",
			SecretAccessKey: "SECRET",
			SessionToken:    "",
		},
	}

	creds, err := s.Retrieve()
	assert.Nil(t, err, "Expect no error")
	assert.Equal(t, "AKID", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "SECRET", creds.SecretAccessKey, "Expect secret access key to match")
	assert.Empty(t, creds.SessionToken, "Expect no session token")
}

func TestStaticProviderIsExpired(t *testing.T) {
	s := StaticProvider{
		Value: Value{
			AccessKeyID:     "AKID",
			SecretAccessKey: "SECRET",
			SessionToken:    "",
		},
	}

	assert.False(t, s.IsExpired(), "Expect static credentials to never expire")
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
