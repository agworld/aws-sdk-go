// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilIdentityExists uses the Amazon SES API operation
// GetIdentityVerificationAttributes to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *SES) WaitUntilIdentityExists(input *GetIdentityVerificationAttributesInput) error {
	return c.WaitUntilIdentityExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilIdentityExistsWithContext is an extended version of WaitUntilIdentityExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SES) WaitUntilIdentityExistsWithContext(ctx aws.Context, input *GetIdentityVerificationAttributesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilIdentityExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "VerificationAttributes.*.VerificationStatus",
				Expected: "Success",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetIdentityVerificationAttributesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetIdentityVerificationAttributesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
