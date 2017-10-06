// +build appengine plan9

package request

import (
	"strings"
)

func isErrConnectionReset(err error) bool {
	return strings.Contains(err.Error(), "connection reset")
}
//Added a line for testing
//Adding another line for Git event testing part 2
