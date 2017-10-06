// +build !go1.6

package request_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

var errTimeout = awserr.New("foo", "bar", errors.New("net/http: request canceled Timeout"))
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
