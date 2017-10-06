// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package directoryserviceiface provides an interface to enable mocking the AWS Directory Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package directoryserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/directoryservice"
)

// DirectoryServiceAPI provides an interface to enable mocking the
// directoryservice.DirectoryService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Directory Service.
//    func myFunc(svc directoryserviceiface.DirectoryServiceAPI) bool {
//        // Make svc.AddIpRoutes request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := directoryservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDirectoryServiceClient struct {
//        directoryserviceiface.DirectoryServiceAPI
//    }
//    func (m *mockDirectoryServiceClient) AddIpRoutes(input *directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDirectoryServiceClient{}
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
type DirectoryServiceAPI interface {
	AddIpRoutes(*directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error)
	AddIpRoutesWithContext(aws.Context, *directoryservice.AddIpRoutesInput, ...request.Option) (*directoryservice.AddIpRoutesOutput, error)
	AddIpRoutesRequest(*directoryservice.AddIpRoutesInput) (*request.Request, *directoryservice.AddIpRoutesOutput)

	AddTagsToResource(*directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error)
	AddTagsToResourceWithContext(aws.Context, *directoryservice.AddTagsToResourceInput, ...request.Option) (*directoryservice.AddTagsToResourceOutput, error)
	AddTagsToResourceRequest(*directoryservice.AddTagsToResourceInput) (*request.Request, *directoryservice.AddTagsToResourceOutput)

	CancelSchemaExtension(*directoryservice.CancelSchemaExtensionInput) (*directoryservice.CancelSchemaExtensionOutput, error)
	CancelSchemaExtensionWithContext(aws.Context, *directoryservice.CancelSchemaExtensionInput, ...request.Option) (*directoryservice.CancelSchemaExtensionOutput, error)
	CancelSchemaExtensionRequest(*directoryservice.CancelSchemaExtensionInput) (*request.Request, *directoryservice.CancelSchemaExtensionOutput)

	ConnectDirectory(*directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error)
	ConnectDirectoryWithContext(aws.Context, *directoryservice.ConnectDirectoryInput, ...request.Option) (*directoryservice.ConnectDirectoryOutput, error)
	ConnectDirectoryRequest(*directoryservice.ConnectDirectoryInput) (*request.Request, *directoryservice.ConnectDirectoryOutput)

	CreateAlias(*directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error)
	CreateAliasWithContext(aws.Context, *directoryservice.CreateAliasInput, ...request.Option) (*directoryservice.CreateAliasOutput, error)
	CreateAliasRequest(*directoryservice.CreateAliasInput) (*request.Request, *directoryservice.CreateAliasOutput)

	CreateComputer(*directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error)
	CreateComputerWithContext(aws.Context, *directoryservice.CreateComputerInput, ...request.Option) (*directoryservice.CreateComputerOutput, error)
	CreateComputerRequest(*directoryservice.CreateComputerInput) (*request.Request, *directoryservice.CreateComputerOutput)

	CreateConditionalForwarder(*directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error)
	CreateConditionalForwarderWithContext(aws.Context, *directoryservice.CreateConditionalForwarderInput, ...request.Option) (*directoryservice.CreateConditionalForwarderOutput, error)
	CreateConditionalForwarderRequest(*directoryservice.CreateConditionalForwarderInput) (*request.Request, *directoryservice.CreateConditionalForwarderOutput)

	CreateDirectory(*directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error)
	CreateDirectoryWithContext(aws.Context, *directoryservice.CreateDirectoryInput, ...request.Option) (*directoryservice.CreateDirectoryOutput, error)
	CreateDirectoryRequest(*directoryservice.CreateDirectoryInput) (*request.Request, *directoryservice.CreateDirectoryOutput)

	CreateMicrosoftAD(*directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error)
	CreateMicrosoftADWithContext(aws.Context, *directoryservice.CreateMicrosoftADInput, ...request.Option) (*directoryservice.CreateMicrosoftADOutput, error)
	CreateMicrosoftADRequest(*directoryservice.CreateMicrosoftADInput) (*request.Request, *directoryservice.CreateMicrosoftADOutput)

	CreateSnapshot(*directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error)
	CreateSnapshotWithContext(aws.Context, *directoryservice.CreateSnapshotInput, ...request.Option) (*directoryservice.CreateSnapshotOutput, error)
	CreateSnapshotRequest(*directoryservice.CreateSnapshotInput) (*request.Request, *directoryservice.CreateSnapshotOutput)

	CreateTrust(*directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error)
	CreateTrustWithContext(aws.Context, *directoryservice.CreateTrustInput, ...request.Option) (*directoryservice.CreateTrustOutput, error)
	CreateTrustRequest(*directoryservice.CreateTrustInput) (*request.Request, *directoryservice.CreateTrustOutput)

	DeleteConditionalForwarder(*directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error)
	DeleteConditionalForwarderWithContext(aws.Context, *directoryservice.DeleteConditionalForwarderInput, ...request.Option) (*directoryservice.DeleteConditionalForwarderOutput, error)
	DeleteConditionalForwarderRequest(*directoryservice.DeleteConditionalForwarderInput) (*request.Request, *directoryservice.DeleteConditionalForwarderOutput)

	DeleteDirectory(*directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error)
	DeleteDirectoryWithContext(aws.Context, *directoryservice.DeleteDirectoryInput, ...request.Option) (*directoryservice.DeleteDirectoryOutput, error)
	DeleteDirectoryRequest(*directoryservice.DeleteDirectoryInput) (*request.Request, *directoryservice.DeleteDirectoryOutput)

	DeleteSnapshot(*directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error)
	DeleteSnapshotWithContext(aws.Context, *directoryservice.DeleteSnapshotInput, ...request.Option) (*directoryservice.DeleteSnapshotOutput, error)
	DeleteSnapshotRequest(*directoryservice.DeleteSnapshotInput) (*request.Request, *directoryservice.DeleteSnapshotOutput)

	DeleteTrust(*directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error)
	DeleteTrustWithContext(aws.Context, *directoryservice.DeleteTrustInput, ...request.Option) (*directoryservice.DeleteTrustOutput, error)
	DeleteTrustRequest(*directoryservice.DeleteTrustInput) (*request.Request, *directoryservice.DeleteTrustOutput)

	DeregisterEventTopic(*directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error)
	DeregisterEventTopicWithContext(aws.Context, *directoryservice.DeregisterEventTopicInput, ...request.Option) (*directoryservice.DeregisterEventTopicOutput, error)
	DeregisterEventTopicRequest(*directoryservice.DeregisterEventTopicInput) (*request.Request, *directoryservice.DeregisterEventTopicOutput)

	DescribeConditionalForwarders(*directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error)
	DescribeConditionalForwardersWithContext(aws.Context, *directoryservice.DescribeConditionalForwardersInput, ...request.Option) (*directoryservice.DescribeConditionalForwardersOutput, error)
	DescribeConditionalForwardersRequest(*directoryservice.DescribeConditionalForwardersInput) (*request.Request, *directoryservice.DescribeConditionalForwardersOutput)

	DescribeDirectories(*directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error)
	DescribeDirectoriesWithContext(aws.Context, *directoryservice.DescribeDirectoriesInput, ...request.Option) (*directoryservice.DescribeDirectoriesOutput, error)
	DescribeDirectoriesRequest(*directoryservice.DescribeDirectoriesInput) (*request.Request, *directoryservice.DescribeDirectoriesOutput)

	DescribeDomainControllers(*directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error)
	DescribeDomainControllersWithContext(aws.Context, *directoryservice.DescribeDomainControllersInput, ...request.Option) (*directoryservice.DescribeDomainControllersOutput, error)
	DescribeDomainControllersRequest(*directoryservice.DescribeDomainControllersInput) (*request.Request, *directoryservice.DescribeDomainControllersOutput)

	DescribeDomainControllersPages(*directoryservice.DescribeDomainControllersInput, func(*directoryservice.DescribeDomainControllersOutput, bool) bool) error
	DescribeDomainControllersPagesWithContext(aws.Context, *directoryservice.DescribeDomainControllersInput, func(*directoryservice.DescribeDomainControllersOutput, bool) bool, ...request.Option) error

	DescribeEventTopics(*directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error)
	DescribeEventTopicsWithContext(aws.Context, *directoryservice.DescribeEventTopicsInput, ...request.Option) (*directoryservice.DescribeEventTopicsOutput, error)
	DescribeEventTopicsRequest(*directoryservice.DescribeEventTopicsInput) (*request.Request, *directoryservice.DescribeEventTopicsOutput)

	DescribeSnapshots(*directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error)
	DescribeSnapshotsWithContext(aws.Context, *directoryservice.DescribeSnapshotsInput, ...request.Option) (*directoryservice.DescribeSnapshotsOutput, error)
	DescribeSnapshotsRequest(*directoryservice.DescribeSnapshotsInput) (*request.Request, *directoryservice.DescribeSnapshotsOutput)

	DescribeTrusts(*directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error)
	DescribeTrustsWithContext(aws.Context, *directoryservice.DescribeTrustsInput, ...request.Option) (*directoryservice.DescribeTrustsOutput, error)
	DescribeTrustsRequest(*directoryservice.DescribeTrustsInput) (*request.Request, *directoryservice.DescribeTrustsOutput)

	DisableRadius(*directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error)
	DisableRadiusWithContext(aws.Context, *directoryservice.DisableRadiusInput, ...request.Option) (*directoryservice.DisableRadiusOutput, error)
	DisableRadiusRequest(*directoryservice.DisableRadiusInput) (*request.Request, *directoryservice.DisableRadiusOutput)

	DisableSso(*directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error)
	DisableSsoWithContext(aws.Context, *directoryservice.DisableSsoInput, ...request.Option) (*directoryservice.DisableSsoOutput, error)
	DisableSsoRequest(*directoryservice.DisableSsoInput) (*request.Request, *directoryservice.DisableSsoOutput)

	EnableRadius(*directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error)
	EnableRadiusWithContext(aws.Context, *directoryservice.EnableRadiusInput, ...request.Option) (*directoryservice.EnableRadiusOutput, error)
	EnableRadiusRequest(*directoryservice.EnableRadiusInput) (*request.Request, *directoryservice.EnableRadiusOutput)

	EnableSso(*directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error)
	EnableSsoWithContext(aws.Context, *directoryservice.EnableSsoInput, ...request.Option) (*directoryservice.EnableSsoOutput, error)
	EnableSsoRequest(*directoryservice.EnableSsoInput) (*request.Request, *directoryservice.EnableSsoOutput)

	GetDirectoryLimits(*directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error)
	GetDirectoryLimitsWithContext(aws.Context, *directoryservice.GetDirectoryLimitsInput, ...request.Option) (*directoryservice.GetDirectoryLimitsOutput, error)
	GetDirectoryLimitsRequest(*directoryservice.GetDirectoryLimitsInput) (*request.Request, *directoryservice.GetDirectoryLimitsOutput)

	GetSnapshotLimits(*directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error)
	GetSnapshotLimitsWithContext(aws.Context, *directoryservice.GetSnapshotLimitsInput, ...request.Option) (*directoryservice.GetSnapshotLimitsOutput, error)
	GetSnapshotLimitsRequest(*directoryservice.GetSnapshotLimitsInput) (*request.Request, *directoryservice.GetSnapshotLimitsOutput)

	ListIpRoutes(*directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error)
	ListIpRoutesWithContext(aws.Context, *directoryservice.ListIpRoutesInput, ...request.Option) (*directoryservice.ListIpRoutesOutput, error)
	ListIpRoutesRequest(*directoryservice.ListIpRoutesInput) (*request.Request, *directoryservice.ListIpRoutesOutput)

	ListSchemaExtensions(*directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error)
	ListSchemaExtensionsWithContext(aws.Context, *directoryservice.ListSchemaExtensionsInput, ...request.Option) (*directoryservice.ListSchemaExtensionsOutput, error)
	ListSchemaExtensionsRequest(*directoryservice.ListSchemaExtensionsInput) (*request.Request, *directoryservice.ListSchemaExtensionsOutput)

	ListTagsForResource(*directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *directoryservice.ListTagsForResourceInput, ...request.Option) (*directoryservice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*directoryservice.ListTagsForResourceInput) (*request.Request, *directoryservice.ListTagsForResourceOutput)

	RegisterEventTopic(*directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error)
	RegisterEventTopicWithContext(aws.Context, *directoryservice.RegisterEventTopicInput, ...request.Option) (*directoryservice.RegisterEventTopicOutput, error)
	RegisterEventTopicRequest(*directoryservice.RegisterEventTopicInput) (*request.Request, *directoryservice.RegisterEventTopicOutput)

	RemoveIpRoutes(*directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error)
	RemoveIpRoutesWithContext(aws.Context, *directoryservice.RemoveIpRoutesInput, ...request.Option) (*directoryservice.RemoveIpRoutesOutput, error)
	RemoveIpRoutesRequest(*directoryservice.RemoveIpRoutesInput) (*request.Request, *directoryservice.RemoveIpRoutesOutput)

	RemoveTagsFromResource(*directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceWithContext(aws.Context, *directoryservice.RemoveTagsFromResourceInput, ...request.Option) (*directoryservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceRequest(*directoryservice.RemoveTagsFromResourceInput) (*request.Request, *directoryservice.RemoveTagsFromResourceOutput)

	RestoreFromSnapshot(*directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error)
	RestoreFromSnapshotWithContext(aws.Context, *directoryservice.RestoreFromSnapshotInput, ...request.Option) (*directoryservice.RestoreFromSnapshotOutput, error)
	RestoreFromSnapshotRequest(*directoryservice.RestoreFromSnapshotInput) (*request.Request, *directoryservice.RestoreFromSnapshotOutput)

	StartSchemaExtension(*directoryservice.StartSchemaExtensionInput) (*directoryservice.StartSchemaExtensionOutput, error)
	StartSchemaExtensionWithContext(aws.Context, *directoryservice.StartSchemaExtensionInput, ...request.Option) (*directoryservice.StartSchemaExtensionOutput, error)
	StartSchemaExtensionRequest(*directoryservice.StartSchemaExtensionInput) (*request.Request, *directoryservice.StartSchemaExtensionOutput)

	UpdateConditionalForwarder(*directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error)
	UpdateConditionalForwarderWithContext(aws.Context, *directoryservice.UpdateConditionalForwarderInput, ...request.Option) (*directoryservice.UpdateConditionalForwarderOutput, error)
	UpdateConditionalForwarderRequest(*directoryservice.UpdateConditionalForwarderInput) (*request.Request, *directoryservice.UpdateConditionalForwarderOutput)

	UpdateNumberOfDomainControllers(*directoryservice.UpdateNumberOfDomainControllersInput) (*directoryservice.UpdateNumberOfDomainControllersOutput, error)
	UpdateNumberOfDomainControllersWithContext(aws.Context, *directoryservice.UpdateNumberOfDomainControllersInput, ...request.Option) (*directoryservice.UpdateNumberOfDomainControllersOutput, error)
	UpdateNumberOfDomainControllersRequest(*directoryservice.UpdateNumberOfDomainControllersInput) (*request.Request, *directoryservice.UpdateNumberOfDomainControllersOutput)

	UpdateRadius(*directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error)
	UpdateRadiusWithContext(aws.Context, *directoryservice.UpdateRadiusInput, ...request.Option) (*directoryservice.UpdateRadiusOutput, error)
	UpdateRadiusRequest(*directoryservice.UpdateRadiusInput) (*request.Request, *directoryservice.UpdateRadiusOutput)

	VerifyTrust(*directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error)
	VerifyTrustWithContext(aws.Context, *directoryservice.VerifyTrustInput, ...request.Option) (*directoryservice.VerifyTrustOutput, error)
	VerifyTrustRequest(*directoryservice.VerifyTrustInput) (*request.Request, *directoryservice.VerifyTrustOutput)
}

var _ DirectoryServiceAPI = (*directoryservice.DirectoryService)(nil)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
