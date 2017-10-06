// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mobileiface provides an interface to enable mocking the AWS Mobile service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mobileiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mobile"
)

// MobileAPI provides an interface to enable mocking the
// mobile.Mobile service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Mobile.
//    func myFunc(svc mobileiface.MobileAPI) bool {
//        // Make svc.CreateProject request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mobile.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMobileClient struct {
//        mobileiface.MobileAPI
//    }
//    func (m *mockMobileClient) CreateProject(input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMobileClient{}
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
type MobileAPI interface {
	CreateProject(*mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error)
	CreateProjectWithContext(aws.Context, *mobile.CreateProjectInput, ...request.Option) (*mobile.CreateProjectOutput, error)
	CreateProjectRequest(*mobile.CreateProjectInput) (*request.Request, *mobile.CreateProjectOutput)

	DeleteProject(*mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error)
	DeleteProjectWithContext(aws.Context, *mobile.DeleteProjectInput, ...request.Option) (*mobile.DeleteProjectOutput, error)
	DeleteProjectRequest(*mobile.DeleteProjectInput) (*request.Request, *mobile.DeleteProjectOutput)

	DescribeBundle(*mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error)
	DescribeBundleWithContext(aws.Context, *mobile.DescribeBundleInput, ...request.Option) (*mobile.DescribeBundleOutput, error)
	DescribeBundleRequest(*mobile.DescribeBundleInput) (*request.Request, *mobile.DescribeBundleOutput)

	DescribeProject(*mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error)
	DescribeProjectWithContext(aws.Context, *mobile.DescribeProjectInput, ...request.Option) (*mobile.DescribeProjectOutput, error)
	DescribeProjectRequest(*mobile.DescribeProjectInput) (*request.Request, *mobile.DescribeProjectOutput)

	ExportBundle(*mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error)
	ExportBundleWithContext(aws.Context, *mobile.ExportBundleInput, ...request.Option) (*mobile.ExportBundleOutput, error)
	ExportBundleRequest(*mobile.ExportBundleInput) (*request.Request, *mobile.ExportBundleOutput)

	ExportProject(*mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error)
	ExportProjectWithContext(aws.Context, *mobile.ExportProjectInput, ...request.Option) (*mobile.ExportProjectOutput, error)
	ExportProjectRequest(*mobile.ExportProjectInput) (*request.Request, *mobile.ExportProjectOutput)

	ListBundles(*mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error)
	ListBundlesWithContext(aws.Context, *mobile.ListBundlesInput, ...request.Option) (*mobile.ListBundlesOutput, error)
	ListBundlesRequest(*mobile.ListBundlesInput) (*request.Request, *mobile.ListBundlesOutput)

	ListBundlesPages(*mobile.ListBundlesInput, func(*mobile.ListBundlesOutput, bool) bool) error
	ListBundlesPagesWithContext(aws.Context, *mobile.ListBundlesInput, func(*mobile.ListBundlesOutput, bool) bool, ...request.Option) error

	ListProjects(*mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error)
	ListProjectsWithContext(aws.Context, *mobile.ListProjectsInput, ...request.Option) (*mobile.ListProjectsOutput, error)
	ListProjectsRequest(*mobile.ListProjectsInput) (*request.Request, *mobile.ListProjectsOutput)

	ListProjectsPages(*mobile.ListProjectsInput, func(*mobile.ListProjectsOutput, bool) bool) error
	ListProjectsPagesWithContext(aws.Context, *mobile.ListProjectsInput, func(*mobile.ListProjectsOutput, bool) bool, ...request.Option) error

	UpdateProject(*mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error)
	UpdateProjectWithContext(aws.Context, *mobile.UpdateProjectInput, ...request.Option) (*mobile.UpdateProjectOutput, error)
	UpdateProjectRequest(*mobile.UpdateProjectInput) (*request.Request, *mobile.UpdateProjectOutput)
}

var _ MobileAPI = (*mobile.Mobile)(nil)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
