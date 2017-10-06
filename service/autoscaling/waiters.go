// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilGroupExists uses the Auto Scaling API operation
// DescribeAutoScalingGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *AutoScaling) WaitUntilGroupExists(input *DescribeAutoScalingGroupsInput) error {
	return c.WaitUntilGroupExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilGroupExistsWithContext is an extended version of WaitUntilGroupExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AutoScaling) WaitUntilGroupExistsWithContext(ctx aws.Context, input *DescribeAutoScalingGroupsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilGroupExists",
		MaxAttempts: 10,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "length(AutoScalingGroups) > `0`",
				Expected: true,
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "length(AutoScalingGroups) > `0`",
				Expected: false,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeAutoScalingGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeAutoScalingGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilGroupInService uses the Auto Scaling API operation
// DescribeAutoScalingGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *AutoScaling) WaitUntilGroupInService(input *DescribeAutoScalingGroupsInput) error {
	return c.WaitUntilGroupInServiceWithContext(aws.BackgroundContext(), input)
}

// WaitUntilGroupInServiceWithContext is an extended version of WaitUntilGroupInService.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AutoScaling) WaitUntilGroupInServiceWithContext(ctx aws.Context, input *DescribeAutoScalingGroupsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilGroupInService",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "contains(AutoScalingGroups[].[length(Instances[?LifecycleState=='InService']) >= MinSize][], `false`)",
				Expected: false,
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "contains(AutoScalingGroups[].[length(Instances[?LifecycleState=='InService']) >= MinSize][], `false`)",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeAutoScalingGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeAutoScalingGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilGroupNotExists uses the Auto Scaling API operation
// DescribeAutoScalingGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *AutoScaling) WaitUntilGroupNotExists(input *DescribeAutoScalingGroupsInput) error {
	return c.WaitUntilGroupNotExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilGroupNotExistsWithContext is an extended version of WaitUntilGroupNotExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AutoScaling) WaitUntilGroupNotExistsWithContext(ctx aws.Context, input *DescribeAutoScalingGroupsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilGroupNotExists",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "length(AutoScalingGroups) > `0`",
				Expected: false,
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "length(AutoScalingGroups) > `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeAutoScalingGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeAutoScalingGroupsRequest(inCpy)
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
