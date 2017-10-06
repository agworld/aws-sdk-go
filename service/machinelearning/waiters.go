// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilBatchPredictionAvailable uses the Amazon Machine Learning API operation
// DescribeBatchPredictions to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MachineLearning) WaitUntilBatchPredictionAvailable(input *DescribeBatchPredictionsInput) error {
	return c.WaitUntilBatchPredictionAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilBatchPredictionAvailableWithContext is an extended version of WaitUntilBatchPredictionAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MachineLearning) WaitUntilBatchPredictionAvailableWithContext(ctx aws.Context, input *DescribeBatchPredictionsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilBatchPredictionAvailable",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Results[].Status",
				Expected: "COMPLETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Results[].Status",
				Expected: "FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeBatchPredictionsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeBatchPredictionsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilDataSourceAvailable uses the Amazon Machine Learning API operation
// DescribeDataSources to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MachineLearning) WaitUntilDataSourceAvailable(input *DescribeDataSourcesInput) error {
	return c.WaitUntilDataSourceAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDataSourceAvailableWithContext is an extended version of WaitUntilDataSourceAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MachineLearning) WaitUntilDataSourceAvailableWithContext(ctx aws.Context, input *DescribeDataSourcesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilDataSourceAvailable",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Results[].Status",
				Expected: "COMPLETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Results[].Status",
				Expected: "FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeDataSourcesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeDataSourcesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilEvaluationAvailable uses the Amazon Machine Learning API operation
// DescribeEvaluations to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MachineLearning) WaitUntilEvaluationAvailable(input *DescribeEvaluationsInput) error {
	return c.WaitUntilEvaluationAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilEvaluationAvailableWithContext is an extended version of WaitUntilEvaluationAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MachineLearning) WaitUntilEvaluationAvailableWithContext(ctx aws.Context, input *DescribeEvaluationsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilEvaluationAvailable",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Results[].Status",
				Expected: "COMPLETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Results[].Status",
				Expected: "FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeEvaluationsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeEvaluationsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilMLModelAvailable uses the Amazon Machine Learning API operation
// DescribeMLModels to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MachineLearning) WaitUntilMLModelAvailable(input *DescribeMLModelsInput) error {
	return c.WaitUntilMLModelAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilMLModelAvailableWithContext is an extended version of WaitUntilMLModelAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MachineLearning) WaitUntilMLModelAvailableWithContext(ctx aws.Context, input *DescribeMLModelsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilMLModelAvailable",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Results[].Status",
				Expected: "COMPLETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Results[].Status",
				Expected: "FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeMLModelsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeMLModelsRequest(inCpy)
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
