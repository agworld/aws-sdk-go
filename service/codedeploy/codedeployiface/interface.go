// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codedeployiface provides an interface to enable mocking the AWS CodeDeploy service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codedeployiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codedeploy"
)

// CodeDeployAPI provides an interface to enable mocking the
// codedeploy.CodeDeploy service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodeDeploy.
//    func myFunc(svc codedeployiface.CodeDeployAPI) bool {
//        // Make svc.AddTagsToOnPremisesInstances request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codedeploy.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeDeployClient struct {
//        codedeployiface.CodeDeployAPI
//    }
//    func (m *mockCodeDeployClient) AddTagsToOnPremisesInstances(input *codedeploy.AddTagsToOnPremisesInstancesInput) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeDeployClient{}
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
type CodeDeployAPI interface {
	AddTagsToOnPremisesInstances(*codedeploy.AddTagsToOnPremisesInstancesInput) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error)
	AddTagsToOnPremisesInstancesWithContext(aws.Context, *codedeploy.AddTagsToOnPremisesInstancesInput, ...request.Option) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error)
	AddTagsToOnPremisesInstancesRequest(*codedeploy.AddTagsToOnPremisesInstancesInput) (*request.Request, *codedeploy.AddTagsToOnPremisesInstancesOutput)

	BatchGetApplicationRevisions(*codedeploy.BatchGetApplicationRevisionsInput) (*codedeploy.BatchGetApplicationRevisionsOutput, error)
	BatchGetApplicationRevisionsWithContext(aws.Context, *codedeploy.BatchGetApplicationRevisionsInput, ...request.Option) (*codedeploy.BatchGetApplicationRevisionsOutput, error)
	BatchGetApplicationRevisionsRequest(*codedeploy.BatchGetApplicationRevisionsInput) (*request.Request, *codedeploy.BatchGetApplicationRevisionsOutput)

	BatchGetApplications(*codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error)
	BatchGetApplicationsWithContext(aws.Context, *codedeploy.BatchGetApplicationsInput, ...request.Option) (*codedeploy.BatchGetApplicationsOutput, error)
	BatchGetApplicationsRequest(*codedeploy.BatchGetApplicationsInput) (*request.Request, *codedeploy.BatchGetApplicationsOutput)

	BatchGetDeploymentGroups(*codedeploy.BatchGetDeploymentGroupsInput) (*codedeploy.BatchGetDeploymentGroupsOutput, error)
	BatchGetDeploymentGroupsWithContext(aws.Context, *codedeploy.BatchGetDeploymentGroupsInput, ...request.Option) (*codedeploy.BatchGetDeploymentGroupsOutput, error)
	BatchGetDeploymentGroupsRequest(*codedeploy.BatchGetDeploymentGroupsInput) (*request.Request, *codedeploy.BatchGetDeploymentGroupsOutput)

	BatchGetDeploymentInstances(*codedeploy.BatchGetDeploymentInstancesInput) (*codedeploy.BatchGetDeploymentInstancesOutput, error)
	BatchGetDeploymentInstancesWithContext(aws.Context, *codedeploy.BatchGetDeploymentInstancesInput, ...request.Option) (*codedeploy.BatchGetDeploymentInstancesOutput, error)
	BatchGetDeploymentInstancesRequest(*codedeploy.BatchGetDeploymentInstancesInput) (*request.Request, *codedeploy.BatchGetDeploymentInstancesOutput)

	BatchGetDeployments(*codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error)
	BatchGetDeploymentsWithContext(aws.Context, *codedeploy.BatchGetDeploymentsInput, ...request.Option) (*codedeploy.BatchGetDeploymentsOutput, error)
	BatchGetDeploymentsRequest(*codedeploy.BatchGetDeploymentsInput) (*request.Request, *codedeploy.BatchGetDeploymentsOutput)

	BatchGetOnPremisesInstances(*codedeploy.BatchGetOnPremisesInstancesInput) (*codedeploy.BatchGetOnPremisesInstancesOutput, error)
	BatchGetOnPremisesInstancesWithContext(aws.Context, *codedeploy.BatchGetOnPremisesInstancesInput, ...request.Option) (*codedeploy.BatchGetOnPremisesInstancesOutput, error)
	BatchGetOnPremisesInstancesRequest(*codedeploy.BatchGetOnPremisesInstancesInput) (*request.Request, *codedeploy.BatchGetOnPremisesInstancesOutput)

	ContinueDeployment(*codedeploy.ContinueDeploymentInput) (*codedeploy.ContinueDeploymentOutput, error)
	ContinueDeploymentWithContext(aws.Context, *codedeploy.ContinueDeploymentInput, ...request.Option) (*codedeploy.ContinueDeploymentOutput, error)
	ContinueDeploymentRequest(*codedeploy.ContinueDeploymentInput) (*request.Request, *codedeploy.ContinueDeploymentOutput)

	CreateApplication(*codedeploy.CreateApplicationInput) (*codedeploy.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *codedeploy.CreateApplicationInput, ...request.Option) (*codedeploy.CreateApplicationOutput, error)
	CreateApplicationRequest(*codedeploy.CreateApplicationInput) (*request.Request, *codedeploy.CreateApplicationOutput)

	CreateDeployment(*codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error)
	CreateDeploymentWithContext(aws.Context, *codedeploy.CreateDeploymentInput, ...request.Option) (*codedeploy.CreateDeploymentOutput, error)
	CreateDeploymentRequest(*codedeploy.CreateDeploymentInput) (*request.Request, *codedeploy.CreateDeploymentOutput)

	CreateDeploymentConfig(*codedeploy.CreateDeploymentConfigInput) (*codedeploy.CreateDeploymentConfigOutput, error)
	CreateDeploymentConfigWithContext(aws.Context, *codedeploy.CreateDeploymentConfigInput, ...request.Option) (*codedeploy.CreateDeploymentConfigOutput, error)
	CreateDeploymentConfigRequest(*codedeploy.CreateDeploymentConfigInput) (*request.Request, *codedeploy.CreateDeploymentConfigOutput)

	CreateDeploymentGroup(*codedeploy.CreateDeploymentGroupInput) (*codedeploy.CreateDeploymentGroupOutput, error)
	CreateDeploymentGroupWithContext(aws.Context, *codedeploy.CreateDeploymentGroupInput, ...request.Option) (*codedeploy.CreateDeploymentGroupOutput, error)
	CreateDeploymentGroupRequest(*codedeploy.CreateDeploymentGroupInput) (*request.Request, *codedeploy.CreateDeploymentGroupOutput)

	DeleteApplication(*codedeploy.DeleteApplicationInput) (*codedeploy.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *codedeploy.DeleteApplicationInput, ...request.Option) (*codedeploy.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*codedeploy.DeleteApplicationInput) (*request.Request, *codedeploy.DeleteApplicationOutput)

	DeleteDeploymentConfig(*codedeploy.DeleteDeploymentConfigInput) (*codedeploy.DeleteDeploymentConfigOutput, error)
	DeleteDeploymentConfigWithContext(aws.Context, *codedeploy.DeleteDeploymentConfigInput, ...request.Option) (*codedeploy.DeleteDeploymentConfigOutput, error)
	DeleteDeploymentConfigRequest(*codedeploy.DeleteDeploymentConfigInput) (*request.Request, *codedeploy.DeleteDeploymentConfigOutput)

	DeleteDeploymentGroup(*codedeploy.DeleteDeploymentGroupInput) (*codedeploy.DeleteDeploymentGroupOutput, error)
	DeleteDeploymentGroupWithContext(aws.Context, *codedeploy.DeleteDeploymentGroupInput, ...request.Option) (*codedeploy.DeleteDeploymentGroupOutput, error)
	DeleteDeploymentGroupRequest(*codedeploy.DeleteDeploymentGroupInput) (*request.Request, *codedeploy.DeleteDeploymentGroupOutput)

	DeregisterOnPremisesInstance(*codedeploy.DeregisterOnPremisesInstanceInput) (*codedeploy.DeregisterOnPremisesInstanceOutput, error)
	DeregisterOnPremisesInstanceWithContext(aws.Context, *codedeploy.DeregisterOnPremisesInstanceInput, ...request.Option) (*codedeploy.DeregisterOnPremisesInstanceOutput, error)
	DeregisterOnPremisesInstanceRequest(*codedeploy.DeregisterOnPremisesInstanceInput) (*request.Request, *codedeploy.DeregisterOnPremisesInstanceOutput)

	GetApplication(*codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error)
	GetApplicationWithContext(aws.Context, *codedeploy.GetApplicationInput, ...request.Option) (*codedeploy.GetApplicationOutput, error)
	GetApplicationRequest(*codedeploy.GetApplicationInput) (*request.Request, *codedeploy.GetApplicationOutput)

	GetApplicationRevision(*codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error)
	GetApplicationRevisionWithContext(aws.Context, *codedeploy.GetApplicationRevisionInput, ...request.Option) (*codedeploy.GetApplicationRevisionOutput, error)
	GetApplicationRevisionRequest(*codedeploy.GetApplicationRevisionInput) (*request.Request, *codedeploy.GetApplicationRevisionOutput)

	GetDeployment(*codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error)
	GetDeploymentWithContext(aws.Context, *codedeploy.GetDeploymentInput, ...request.Option) (*codedeploy.GetDeploymentOutput, error)
	GetDeploymentRequest(*codedeploy.GetDeploymentInput) (*request.Request, *codedeploy.GetDeploymentOutput)

	GetDeploymentConfig(*codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error)
	GetDeploymentConfigWithContext(aws.Context, *codedeploy.GetDeploymentConfigInput, ...request.Option) (*codedeploy.GetDeploymentConfigOutput, error)
	GetDeploymentConfigRequest(*codedeploy.GetDeploymentConfigInput) (*request.Request, *codedeploy.GetDeploymentConfigOutput)

	GetDeploymentGroup(*codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error)
	GetDeploymentGroupWithContext(aws.Context, *codedeploy.GetDeploymentGroupInput, ...request.Option) (*codedeploy.GetDeploymentGroupOutput, error)
	GetDeploymentGroupRequest(*codedeploy.GetDeploymentGroupInput) (*request.Request, *codedeploy.GetDeploymentGroupOutput)

	GetDeploymentInstance(*codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error)
	GetDeploymentInstanceWithContext(aws.Context, *codedeploy.GetDeploymentInstanceInput, ...request.Option) (*codedeploy.GetDeploymentInstanceOutput, error)
	GetDeploymentInstanceRequest(*codedeploy.GetDeploymentInstanceInput) (*request.Request, *codedeploy.GetDeploymentInstanceOutput)

	GetOnPremisesInstance(*codedeploy.GetOnPremisesInstanceInput) (*codedeploy.GetOnPremisesInstanceOutput, error)
	GetOnPremisesInstanceWithContext(aws.Context, *codedeploy.GetOnPremisesInstanceInput, ...request.Option) (*codedeploy.GetOnPremisesInstanceOutput, error)
	GetOnPremisesInstanceRequest(*codedeploy.GetOnPremisesInstanceInput) (*request.Request, *codedeploy.GetOnPremisesInstanceOutput)

	ListApplicationRevisions(*codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error)
	ListApplicationRevisionsWithContext(aws.Context, *codedeploy.ListApplicationRevisionsInput, ...request.Option) (*codedeploy.ListApplicationRevisionsOutput, error)
	ListApplicationRevisionsRequest(*codedeploy.ListApplicationRevisionsInput) (*request.Request, *codedeploy.ListApplicationRevisionsOutput)

	ListApplicationRevisionsPages(*codedeploy.ListApplicationRevisionsInput, func(*codedeploy.ListApplicationRevisionsOutput, bool) bool) error
	ListApplicationRevisionsPagesWithContext(aws.Context, *codedeploy.ListApplicationRevisionsInput, func(*codedeploy.ListApplicationRevisionsOutput, bool) bool, ...request.Option) error

	ListApplications(*codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error)
	ListApplicationsWithContext(aws.Context, *codedeploy.ListApplicationsInput, ...request.Option) (*codedeploy.ListApplicationsOutput, error)
	ListApplicationsRequest(*codedeploy.ListApplicationsInput) (*request.Request, *codedeploy.ListApplicationsOutput)

	ListApplicationsPages(*codedeploy.ListApplicationsInput, func(*codedeploy.ListApplicationsOutput, bool) bool) error
	ListApplicationsPagesWithContext(aws.Context, *codedeploy.ListApplicationsInput, func(*codedeploy.ListApplicationsOutput, bool) bool, ...request.Option) error

	ListDeploymentConfigs(*codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error)
	ListDeploymentConfigsWithContext(aws.Context, *codedeploy.ListDeploymentConfigsInput, ...request.Option) (*codedeploy.ListDeploymentConfigsOutput, error)
	ListDeploymentConfigsRequest(*codedeploy.ListDeploymentConfigsInput) (*request.Request, *codedeploy.ListDeploymentConfigsOutput)

	ListDeploymentConfigsPages(*codedeploy.ListDeploymentConfigsInput, func(*codedeploy.ListDeploymentConfigsOutput, bool) bool) error
	ListDeploymentConfigsPagesWithContext(aws.Context, *codedeploy.ListDeploymentConfigsInput, func(*codedeploy.ListDeploymentConfigsOutput, bool) bool, ...request.Option) error

	ListDeploymentGroups(*codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error)
	ListDeploymentGroupsWithContext(aws.Context, *codedeploy.ListDeploymentGroupsInput, ...request.Option) (*codedeploy.ListDeploymentGroupsOutput, error)
	ListDeploymentGroupsRequest(*codedeploy.ListDeploymentGroupsInput) (*request.Request, *codedeploy.ListDeploymentGroupsOutput)

	ListDeploymentGroupsPages(*codedeploy.ListDeploymentGroupsInput, func(*codedeploy.ListDeploymentGroupsOutput, bool) bool) error
	ListDeploymentGroupsPagesWithContext(aws.Context, *codedeploy.ListDeploymentGroupsInput, func(*codedeploy.ListDeploymentGroupsOutput, bool) bool, ...request.Option) error

	ListDeploymentInstances(*codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error)
	ListDeploymentInstancesWithContext(aws.Context, *codedeploy.ListDeploymentInstancesInput, ...request.Option) (*codedeploy.ListDeploymentInstancesOutput, error)
	ListDeploymentInstancesRequest(*codedeploy.ListDeploymentInstancesInput) (*request.Request, *codedeploy.ListDeploymentInstancesOutput)

	ListDeploymentInstancesPages(*codedeploy.ListDeploymentInstancesInput, func(*codedeploy.ListDeploymentInstancesOutput, bool) bool) error
	ListDeploymentInstancesPagesWithContext(aws.Context, *codedeploy.ListDeploymentInstancesInput, func(*codedeploy.ListDeploymentInstancesOutput, bool) bool, ...request.Option) error

	ListDeployments(*codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error)
	ListDeploymentsWithContext(aws.Context, *codedeploy.ListDeploymentsInput, ...request.Option) (*codedeploy.ListDeploymentsOutput, error)
	ListDeploymentsRequest(*codedeploy.ListDeploymentsInput) (*request.Request, *codedeploy.ListDeploymentsOutput)

	ListDeploymentsPages(*codedeploy.ListDeploymentsInput, func(*codedeploy.ListDeploymentsOutput, bool) bool) error
	ListDeploymentsPagesWithContext(aws.Context, *codedeploy.ListDeploymentsInput, func(*codedeploy.ListDeploymentsOutput, bool) bool, ...request.Option) error

	ListGitHubAccountTokenNames(*codedeploy.ListGitHubAccountTokenNamesInput) (*codedeploy.ListGitHubAccountTokenNamesOutput, error)
	ListGitHubAccountTokenNamesWithContext(aws.Context, *codedeploy.ListGitHubAccountTokenNamesInput, ...request.Option) (*codedeploy.ListGitHubAccountTokenNamesOutput, error)
	ListGitHubAccountTokenNamesRequest(*codedeploy.ListGitHubAccountTokenNamesInput) (*request.Request, *codedeploy.ListGitHubAccountTokenNamesOutput)

	ListOnPremisesInstances(*codedeploy.ListOnPremisesInstancesInput) (*codedeploy.ListOnPremisesInstancesOutput, error)
	ListOnPremisesInstancesWithContext(aws.Context, *codedeploy.ListOnPremisesInstancesInput, ...request.Option) (*codedeploy.ListOnPremisesInstancesOutput, error)
	ListOnPremisesInstancesRequest(*codedeploy.ListOnPremisesInstancesInput) (*request.Request, *codedeploy.ListOnPremisesInstancesOutput)

	RegisterApplicationRevision(*codedeploy.RegisterApplicationRevisionInput) (*codedeploy.RegisterApplicationRevisionOutput, error)
	RegisterApplicationRevisionWithContext(aws.Context, *codedeploy.RegisterApplicationRevisionInput, ...request.Option) (*codedeploy.RegisterApplicationRevisionOutput, error)
	RegisterApplicationRevisionRequest(*codedeploy.RegisterApplicationRevisionInput) (*request.Request, *codedeploy.RegisterApplicationRevisionOutput)

	RegisterOnPremisesInstance(*codedeploy.RegisterOnPremisesInstanceInput) (*codedeploy.RegisterOnPremisesInstanceOutput, error)
	RegisterOnPremisesInstanceWithContext(aws.Context, *codedeploy.RegisterOnPremisesInstanceInput, ...request.Option) (*codedeploy.RegisterOnPremisesInstanceOutput, error)
	RegisterOnPremisesInstanceRequest(*codedeploy.RegisterOnPremisesInstanceInput) (*request.Request, *codedeploy.RegisterOnPremisesInstanceOutput)

	RemoveTagsFromOnPremisesInstances(*codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error)
	RemoveTagsFromOnPremisesInstancesWithContext(aws.Context, *codedeploy.RemoveTagsFromOnPremisesInstancesInput, ...request.Option) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error)
	RemoveTagsFromOnPremisesInstancesRequest(*codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*request.Request, *codedeploy.RemoveTagsFromOnPremisesInstancesOutput)

	SkipWaitTimeForInstanceTermination(*codedeploy.SkipWaitTimeForInstanceTerminationInput) (*codedeploy.SkipWaitTimeForInstanceTerminationOutput, error)
	SkipWaitTimeForInstanceTerminationWithContext(aws.Context, *codedeploy.SkipWaitTimeForInstanceTerminationInput, ...request.Option) (*codedeploy.SkipWaitTimeForInstanceTerminationOutput, error)
	SkipWaitTimeForInstanceTerminationRequest(*codedeploy.SkipWaitTimeForInstanceTerminationInput) (*request.Request, *codedeploy.SkipWaitTimeForInstanceTerminationOutput)

	StopDeployment(*codedeploy.StopDeploymentInput) (*codedeploy.StopDeploymentOutput, error)
	StopDeploymentWithContext(aws.Context, *codedeploy.StopDeploymentInput, ...request.Option) (*codedeploy.StopDeploymentOutput, error)
	StopDeploymentRequest(*codedeploy.StopDeploymentInput) (*request.Request, *codedeploy.StopDeploymentOutput)

	UpdateApplication(*codedeploy.UpdateApplicationInput) (*codedeploy.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *codedeploy.UpdateApplicationInput, ...request.Option) (*codedeploy.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*codedeploy.UpdateApplicationInput) (*request.Request, *codedeploy.UpdateApplicationOutput)

	UpdateDeploymentGroup(*codedeploy.UpdateDeploymentGroupInput) (*codedeploy.UpdateDeploymentGroupOutput, error)
	UpdateDeploymentGroupWithContext(aws.Context, *codedeploy.UpdateDeploymentGroupInput, ...request.Option) (*codedeploy.UpdateDeploymentGroupOutput, error)
	UpdateDeploymentGroupRequest(*codedeploy.UpdateDeploymentGroupInput) (*request.Request, *codedeploy.UpdateDeploymentGroupOutput)

	WaitUntilDeploymentSuccessful(*codedeploy.GetDeploymentInput) error
	WaitUntilDeploymentSuccessfulWithContext(aws.Context, *codedeploy.GetDeploymentInput, ...request.WaiterOption) error
}

var _ CodeDeployAPI = (*codedeploy.CodeDeploy)(nil)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
