// +build !appengine,!plan9

package request_test

import (
	"net"
	"os"
	"syscall"
)

var stubConnectionResetError = &net.OpError{Err: &os.SyscallError{Syscall: "read", Err: syscall.ECONNRESET}}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
