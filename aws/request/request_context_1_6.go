// +build !go1.7

package request

import "github.com/aws/aws-sdk-go/aws"

// setContext updates the Request to use the passed in context for cancellation.
// Context will also be used for request retry delay.
//
// Creates shallow copy of the http.Request with the WithContext method.
func setRequestContext(r *Request, ctx aws.Context) {
	r.context = ctx
	r.HTTPRequest.Cancel = ctx.Done()
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
