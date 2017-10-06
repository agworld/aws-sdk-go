// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotdataplaneiface provides an interface to enable mocking the AWS IoT Data Plane service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotdataplaneiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

// IoTDataPlaneAPI provides an interface to enable mocking the
// iotdataplane.IoTDataPlane service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT Data Plane.
//    func myFunc(svc iotdataplaneiface.IoTDataPlaneAPI) bool {
//        // Make svc.DeleteThingShadow request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iotdataplane.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTDataPlaneClient struct {
//        iotdataplaneiface.IoTDataPlaneAPI
//    }
//    func (m *mockIoTDataPlaneClient) DeleteThingShadow(input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTDataPlaneClient{}
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
type IoTDataPlaneAPI interface {
	DeleteThingShadow(*iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error)
	DeleteThingShadowWithContext(aws.Context, *iotdataplane.DeleteThingShadowInput, ...request.Option) (*iotdataplane.DeleteThingShadowOutput, error)
	DeleteThingShadowRequest(*iotdataplane.DeleteThingShadowInput) (*request.Request, *iotdataplane.DeleteThingShadowOutput)

	GetThingShadow(*iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error)
	GetThingShadowWithContext(aws.Context, *iotdataplane.GetThingShadowInput, ...request.Option) (*iotdataplane.GetThingShadowOutput, error)
	GetThingShadowRequest(*iotdataplane.GetThingShadowInput) (*request.Request, *iotdataplane.GetThingShadowOutput)

	Publish(*iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error)
	PublishWithContext(aws.Context, *iotdataplane.PublishInput, ...request.Option) (*iotdataplane.PublishOutput, error)
	PublishRequest(*iotdataplane.PublishInput) (*request.Request, *iotdataplane.PublishOutput)

	UpdateThingShadow(*iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error)
	UpdateThingShadowWithContext(aws.Context, *iotdataplane.UpdateThingShadowInput, ...request.Option) (*iotdataplane.UpdateThingShadowOutput, error)
	UpdateThingShadowRequest(*iotdataplane.UpdateThingShadowInput) (*request.Request, *iotdataplane.UpdateThingShadowOutput)
}

var _ IoTDataPlaneAPI = (*iotdataplane.IoTDataPlane)(nil)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
