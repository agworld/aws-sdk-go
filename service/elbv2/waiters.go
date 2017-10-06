// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elbv2

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilLoadBalancerAvailable uses the Elastic Load Balancing v2 API operation
// DescribeLoadBalancers to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ELBV2) WaitUntilLoadBalancerAvailable(input *DescribeLoadBalancersInput) error {
	return c.WaitUntilLoadBalancerAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilLoadBalancerAvailableWithContext is an extended version of WaitUntilLoadBalancerAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ELBV2) WaitUntilLoadBalancerAvailableWithContext(ctx aws.Context, input *DescribeLoadBalancersInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilLoadBalancerAvailable",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "LoadBalancers[].State.Code",
				Expected: "active",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "LoadBalancers[].State.Code",
				Expected: "provisioning",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "LoadBalancerNotFound",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeLoadBalancersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeLoadBalancersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilLoadBalancerExists uses the Elastic Load Balancing v2 API operation
// DescribeLoadBalancers to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ELBV2) WaitUntilLoadBalancerExists(input *DescribeLoadBalancersInput) error {
	return c.WaitUntilLoadBalancerExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilLoadBalancerExistsWithContext is an extended version of WaitUntilLoadBalancerExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ELBV2) WaitUntilLoadBalancerExistsWithContext(ctx aws.Context, input *DescribeLoadBalancersInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilLoadBalancerExists",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "LoadBalancerNotFound",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeLoadBalancersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeLoadBalancersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilLoadBalancersDeleted uses the Elastic Load Balancing v2 API operation
// DescribeLoadBalancers to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ELBV2) WaitUntilLoadBalancersDeleted(input *DescribeLoadBalancersInput) error {
	return c.WaitUntilLoadBalancersDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilLoadBalancersDeletedWithContext is an extended version of WaitUntilLoadBalancersDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ELBV2) WaitUntilLoadBalancersDeletedWithContext(ctx aws.Context, input *DescribeLoadBalancersInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilLoadBalancersDeleted",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "LoadBalancers[].State.Code",
				Expected: "active",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "LoadBalancerNotFound",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeLoadBalancersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeLoadBalancersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilTargetDeregistered uses the Elastic Load Balancing v2 API operation
// DescribeTargetHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ELBV2) WaitUntilTargetDeregistered(input *DescribeTargetHealthInput) error {
	return c.WaitUntilTargetDeregisteredWithContext(aws.BackgroundContext(), input)
}

// WaitUntilTargetDeregisteredWithContext is an extended version of WaitUntilTargetDeregistered.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ELBV2) WaitUntilTargetDeregisteredWithContext(ctx aws.Context, input *DescribeTargetHealthInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilTargetDeregistered",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "InvalidTarget",
			},
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "TargetHealthDescriptions[].TargetHealth.State",
				Expected: "unused",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeTargetHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeTargetHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilTargetInService uses the Elastic Load Balancing v2 API operation
// DescribeTargetHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ELBV2) WaitUntilTargetInService(input *DescribeTargetHealthInput) error {
	return c.WaitUntilTargetInServiceWithContext(aws.BackgroundContext(), input)
}

// WaitUntilTargetInServiceWithContext is an extended version of WaitUntilTargetInService.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ELBV2) WaitUntilTargetInServiceWithContext(ctx aws.Context, input *DescribeTargetHealthInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilTargetInService",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "TargetHealthDescriptions[].TargetHealth.State",
				Expected: "healthy",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "InvalidInstance",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeTargetHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeTargetHealthRequest(inCpy)
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
