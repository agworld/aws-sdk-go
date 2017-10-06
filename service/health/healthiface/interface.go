// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package healthiface provides an interface to enable mocking the AWS Health APIs and Notifications service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package healthiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/health"
)

// HealthAPI provides an interface to enable mocking the
// health.Health service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Health APIs and Notifications.
//    func myFunc(svc healthiface.HealthAPI) bool {
//        // Make svc.DescribeAffectedEntities request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := health.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockHealthClient struct {
//        healthiface.HealthAPI
//    }
//    func (m *mockHealthClient) DescribeAffectedEntities(input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockHealthClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type HealthAPI interface {
	DescribeAffectedEntities(*health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error)
	DescribeAffectedEntitiesWithContext(aws.Context, *health.DescribeAffectedEntitiesInput, ...request.Option) (*health.DescribeAffectedEntitiesOutput, error)
	DescribeAffectedEntitiesRequest(*health.DescribeAffectedEntitiesInput) (*request.Request, *health.DescribeAffectedEntitiesOutput)

	DescribeAffectedEntitiesPages(*health.DescribeAffectedEntitiesInput, func(*health.DescribeAffectedEntitiesOutput, bool) bool) error
	DescribeAffectedEntitiesPagesWithContext(aws.Context, *health.DescribeAffectedEntitiesInput, func(*health.DescribeAffectedEntitiesOutput, bool) bool, ...request.Option) error

	DescribeEntityAggregates(*health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error)
	DescribeEntityAggregatesWithContext(aws.Context, *health.DescribeEntityAggregatesInput, ...request.Option) (*health.DescribeEntityAggregatesOutput, error)
	DescribeEntityAggregatesRequest(*health.DescribeEntityAggregatesInput) (*request.Request, *health.DescribeEntityAggregatesOutput)

	DescribeEventAggregates(*health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error)
	DescribeEventAggregatesWithContext(aws.Context, *health.DescribeEventAggregatesInput, ...request.Option) (*health.DescribeEventAggregatesOutput, error)
	DescribeEventAggregatesRequest(*health.DescribeEventAggregatesInput) (*request.Request, *health.DescribeEventAggregatesOutput)

	DescribeEventAggregatesPages(*health.DescribeEventAggregatesInput, func(*health.DescribeEventAggregatesOutput, bool) bool) error
	DescribeEventAggregatesPagesWithContext(aws.Context, *health.DescribeEventAggregatesInput, func(*health.DescribeEventAggregatesOutput, bool) bool, ...request.Option) error

	DescribeEventDetails(*health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error)
	DescribeEventDetailsWithContext(aws.Context, *health.DescribeEventDetailsInput, ...request.Option) (*health.DescribeEventDetailsOutput, error)
	DescribeEventDetailsRequest(*health.DescribeEventDetailsInput) (*request.Request, *health.DescribeEventDetailsOutput)

	DescribeEventTypes(*health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error)
	DescribeEventTypesWithContext(aws.Context, *health.DescribeEventTypesInput, ...request.Option) (*health.DescribeEventTypesOutput, error)
	DescribeEventTypesRequest(*health.DescribeEventTypesInput) (*request.Request, *health.DescribeEventTypesOutput)

	DescribeEventTypesPages(*health.DescribeEventTypesInput, func(*health.DescribeEventTypesOutput, bool) bool) error
	DescribeEventTypesPagesWithContext(aws.Context, *health.DescribeEventTypesInput, func(*health.DescribeEventTypesOutput, bool) bool, ...request.Option) error

	DescribeEvents(*health.DescribeEventsInput) (*health.DescribeEventsOutput, error)
	DescribeEventsWithContext(aws.Context, *health.DescribeEventsInput, ...request.Option) (*health.DescribeEventsOutput, error)
	DescribeEventsRequest(*health.DescribeEventsInput) (*request.Request, *health.DescribeEventsOutput)

	DescribeEventsPages(*health.DescribeEventsInput, func(*health.DescribeEventsOutput, bool) bool) error
	DescribeEventsPagesWithContext(aws.Context, *health.DescribeEventsInput, func(*health.DescribeEventsOutput, bool) bool, ...request.Option) error
}

var _ HealthAPI = (*health.Health)(nil)
//Added a line for testing
//Adding another line for Git event testing part 2
