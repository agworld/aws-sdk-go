// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appstreamiface provides an interface to enable mocking the Amazon AppStream service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appstreamiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appstream"
)

// AppStreamAPI provides an interface to enable mocking the
// appstream.AppStream service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon AppStream.
//    func myFunc(svc appstreamiface.AppStreamAPI) bool {
//        // Make svc.AssociateFleet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := appstream.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAppStreamClient struct {
//        appstreamiface.AppStreamAPI
//    }
//    func (m *mockAppStreamClient) AssociateFleet(input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAppStreamClient{}
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
type AppStreamAPI interface {
	AssociateFleet(*appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error)
	AssociateFleetWithContext(aws.Context, *appstream.AssociateFleetInput, ...request.Option) (*appstream.AssociateFleetOutput, error)
	AssociateFleetRequest(*appstream.AssociateFleetInput) (*request.Request, *appstream.AssociateFleetOutput)

	CreateDirectoryConfig(*appstream.CreateDirectoryConfigInput) (*appstream.CreateDirectoryConfigOutput, error)
	CreateDirectoryConfigWithContext(aws.Context, *appstream.CreateDirectoryConfigInput, ...request.Option) (*appstream.CreateDirectoryConfigOutput, error)
	CreateDirectoryConfigRequest(*appstream.CreateDirectoryConfigInput) (*request.Request, *appstream.CreateDirectoryConfigOutput)

	CreateFleet(*appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error)
	CreateFleetWithContext(aws.Context, *appstream.CreateFleetInput, ...request.Option) (*appstream.CreateFleetOutput, error)
	CreateFleetRequest(*appstream.CreateFleetInput) (*request.Request, *appstream.CreateFleetOutput)

	CreateImageBuilder(*appstream.CreateImageBuilderInput) (*appstream.CreateImageBuilderOutput, error)
	CreateImageBuilderWithContext(aws.Context, *appstream.CreateImageBuilderInput, ...request.Option) (*appstream.CreateImageBuilderOutput, error)
	CreateImageBuilderRequest(*appstream.CreateImageBuilderInput) (*request.Request, *appstream.CreateImageBuilderOutput)

	CreateImageBuilderStreamingURL(*appstream.CreateImageBuilderStreamingURLInput) (*appstream.CreateImageBuilderStreamingURLOutput, error)
	CreateImageBuilderStreamingURLWithContext(aws.Context, *appstream.CreateImageBuilderStreamingURLInput, ...request.Option) (*appstream.CreateImageBuilderStreamingURLOutput, error)
	CreateImageBuilderStreamingURLRequest(*appstream.CreateImageBuilderStreamingURLInput) (*request.Request, *appstream.CreateImageBuilderStreamingURLOutput)

	CreateStack(*appstream.CreateStackInput) (*appstream.CreateStackOutput, error)
	CreateStackWithContext(aws.Context, *appstream.CreateStackInput, ...request.Option) (*appstream.CreateStackOutput, error)
	CreateStackRequest(*appstream.CreateStackInput) (*request.Request, *appstream.CreateStackOutput)

	CreateStreamingURL(*appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error)
	CreateStreamingURLWithContext(aws.Context, *appstream.CreateStreamingURLInput, ...request.Option) (*appstream.CreateStreamingURLOutput, error)
	CreateStreamingURLRequest(*appstream.CreateStreamingURLInput) (*request.Request, *appstream.CreateStreamingURLOutput)

	DeleteDirectoryConfig(*appstream.DeleteDirectoryConfigInput) (*appstream.DeleteDirectoryConfigOutput, error)
	DeleteDirectoryConfigWithContext(aws.Context, *appstream.DeleteDirectoryConfigInput, ...request.Option) (*appstream.DeleteDirectoryConfigOutput, error)
	DeleteDirectoryConfigRequest(*appstream.DeleteDirectoryConfigInput) (*request.Request, *appstream.DeleteDirectoryConfigOutput)

	DeleteFleet(*appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error)
	DeleteFleetWithContext(aws.Context, *appstream.DeleteFleetInput, ...request.Option) (*appstream.DeleteFleetOutput, error)
	DeleteFleetRequest(*appstream.DeleteFleetInput) (*request.Request, *appstream.DeleteFleetOutput)

	DeleteImage(*appstream.DeleteImageInput) (*appstream.DeleteImageOutput, error)
	DeleteImageWithContext(aws.Context, *appstream.DeleteImageInput, ...request.Option) (*appstream.DeleteImageOutput, error)
	DeleteImageRequest(*appstream.DeleteImageInput) (*request.Request, *appstream.DeleteImageOutput)

	DeleteImageBuilder(*appstream.DeleteImageBuilderInput) (*appstream.DeleteImageBuilderOutput, error)
	DeleteImageBuilderWithContext(aws.Context, *appstream.DeleteImageBuilderInput, ...request.Option) (*appstream.DeleteImageBuilderOutput, error)
	DeleteImageBuilderRequest(*appstream.DeleteImageBuilderInput) (*request.Request, *appstream.DeleteImageBuilderOutput)

	DeleteStack(*appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error)
	DeleteStackWithContext(aws.Context, *appstream.DeleteStackInput, ...request.Option) (*appstream.DeleteStackOutput, error)
	DeleteStackRequest(*appstream.DeleteStackInput) (*request.Request, *appstream.DeleteStackOutput)

	DescribeDirectoryConfigs(*appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error)
	DescribeDirectoryConfigsWithContext(aws.Context, *appstream.DescribeDirectoryConfigsInput, ...request.Option) (*appstream.DescribeDirectoryConfigsOutput, error)
	DescribeDirectoryConfigsRequest(*appstream.DescribeDirectoryConfigsInput) (*request.Request, *appstream.DescribeDirectoryConfigsOutput)

	DescribeFleets(*appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error)
	DescribeFleetsWithContext(aws.Context, *appstream.DescribeFleetsInput, ...request.Option) (*appstream.DescribeFleetsOutput, error)
	DescribeFleetsRequest(*appstream.DescribeFleetsInput) (*request.Request, *appstream.DescribeFleetsOutput)

	DescribeImageBuilders(*appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error)
	DescribeImageBuildersWithContext(aws.Context, *appstream.DescribeImageBuildersInput, ...request.Option) (*appstream.DescribeImageBuildersOutput, error)
	DescribeImageBuildersRequest(*appstream.DescribeImageBuildersInput) (*request.Request, *appstream.DescribeImageBuildersOutput)

	DescribeImages(*appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error)
	DescribeImagesWithContext(aws.Context, *appstream.DescribeImagesInput, ...request.Option) (*appstream.DescribeImagesOutput, error)
	DescribeImagesRequest(*appstream.DescribeImagesInput) (*request.Request, *appstream.DescribeImagesOutput)

	DescribeSessions(*appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error)
	DescribeSessionsWithContext(aws.Context, *appstream.DescribeSessionsInput, ...request.Option) (*appstream.DescribeSessionsOutput, error)
	DescribeSessionsRequest(*appstream.DescribeSessionsInput) (*request.Request, *appstream.DescribeSessionsOutput)

	DescribeStacks(*appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error)
	DescribeStacksWithContext(aws.Context, *appstream.DescribeStacksInput, ...request.Option) (*appstream.DescribeStacksOutput, error)
	DescribeStacksRequest(*appstream.DescribeStacksInput) (*request.Request, *appstream.DescribeStacksOutput)

	DisassociateFleet(*appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error)
	DisassociateFleetWithContext(aws.Context, *appstream.DisassociateFleetInput, ...request.Option) (*appstream.DisassociateFleetOutput, error)
	DisassociateFleetRequest(*appstream.DisassociateFleetInput) (*request.Request, *appstream.DisassociateFleetOutput)

	ExpireSession(*appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error)
	ExpireSessionWithContext(aws.Context, *appstream.ExpireSessionInput, ...request.Option) (*appstream.ExpireSessionOutput, error)
	ExpireSessionRequest(*appstream.ExpireSessionInput) (*request.Request, *appstream.ExpireSessionOutput)

	ListAssociatedFleets(*appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error)
	ListAssociatedFleetsWithContext(aws.Context, *appstream.ListAssociatedFleetsInput, ...request.Option) (*appstream.ListAssociatedFleetsOutput, error)
	ListAssociatedFleetsRequest(*appstream.ListAssociatedFleetsInput) (*request.Request, *appstream.ListAssociatedFleetsOutput)

	ListAssociatedStacks(*appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error)
	ListAssociatedStacksWithContext(aws.Context, *appstream.ListAssociatedStacksInput, ...request.Option) (*appstream.ListAssociatedStacksOutput, error)
	ListAssociatedStacksRequest(*appstream.ListAssociatedStacksInput) (*request.Request, *appstream.ListAssociatedStacksOutput)

	StartFleet(*appstream.StartFleetInput) (*appstream.StartFleetOutput, error)
	StartFleetWithContext(aws.Context, *appstream.StartFleetInput, ...request.Option) (*appstream.StartFleetOutput, error)
	StartFleetRequest(*appstream.StartFleetInput) (*request.Request, *appstream.StartFleetOutput)

	StartImageBuilder(*appstream.StartImageBuilderInput) (*appstream.StartImageBuilderOutput, error)
	StartImageBuilderWithContext(aws.Context, *appstream.StartImageBuilderInput, ...request.Option) (*appstream.StartImageBuilderOutput, error)
	StartImageBuilderRequest(*appstream.StartImageBuilderInput) (*request.Request, *appstream.StartImageBuilderOutput)

	StopFleet(*appstream.StopFleetInput) (*appstream.StopFleetOutput, error)
	StopFleetWithContext(aws.Context, *appstream.StopFleetInput, ...request.Option) (*appstream.StopFleetOutput, error)
	StopFleetRequest(*appstream.StopFleetInput) (*request.Request, *appstream.StopFleetOutput)

	StopImageBuilder(*appstream.StopImageBuilderInput) (*appstream.StopImageBuilderOutput, error)
	StopImageBuilderWithContext(aws.Context, *appstream.StopImageBuilderInput, ...request.Option) (*appstream.StopImageBuilderOutput, error)
	StopImageBuilderRequest(*appstream.StopImageBuilderInput) (*request.Request, *appstream.StopImageBuilderOutput)

	UpdateDirectoryConfig(*appstream.UpdateDirectoryConfigInput) (*appstream.UpdateDirectoryConfigOutput, error)
	UpdateDirectoryConfigWithContext(aws.Context, *appstream.UpdateDirectoryConfigInput, ...request.Option) (*appstream.UpdateDirectoryConfigOutput, error)
	UpdateDirectoryConfigRequest(*appstream.UpdateDirectoryConfigInput) (*request.Request, *appstream.UpdateDirectoryConfigOutput)

	UpdateFleet(*appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error)
	UpdateFleetWithContext(aws.Context, *appstream.UpdateFleetInput, ...request.Option) (*appstream.UpdateFleetOutput, error)
	UpdateFleetRequest(*appstream.UpdateFleetInput) (*request.Request, *appstream.UpdateFleetOutput)

	UpdateStack(*appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error)
	UpdateStackWithContext(aws.Context, *appstream.UpdateStackInput, ...request.Option) (*appstream.UpdateStackOutput, error)
	UpdateStackRequest(*appstream.UpdateStackInput) (*request.Request, *appstream.UpdateStackOutput)

	WaitUntilFleetStarted(*appstream.DescribeFleetsInput) error
	WaitUntilFleetStartedWithContext(aws.Context, *appstream.DescribeFleetsInput, ...request.WaiterOption) error

	WaitUntilFleetStopped(*appstream.DescribeFleetsInput) error
	WaitUntilFleetStoppedWithContext(aws.Context, *appstream.DescribeFleetsInput, ...request.WaiterOption) error
}

var _ AppStreamAPI = (*appstream.AppStream)(nil)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
