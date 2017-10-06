// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elastictranscoderiface provides an interface to enable mocking the Amazon Elastic Transcoder service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elastictranscoderiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
)

// ElasticTranscoderAPI provides an interface to enable mocking the
// elastictranscoder.ElasticTranscoder service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elastic Transcoder.
//    func myFunc(svc elastictranscoderiface.ElasticTranscoderAPI) bool {
//        // Make svc.CancelJob request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := elastictranscoder.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockElasticTranscoderClient struct {
//        elastictranscoderiface.ElasticTranscoderAPI
//    }
//    func (m *mockElasticTranscoderClient) CancelJob(input *elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockElasticTranscoderClient{}
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
type ElasticTranscoderAPI interface {
	CancelJob(*elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error)
	CancelJobWithContext(aws.Context, *elastictranscoder.CancelJobInput, ...request.Option) (*elastictranscoder.CancelJobOutput, error)
	CancelJobRequest(*elastictranscoder.CancelJobInput) (*request.Request, *elastictranscoder.CancelJobOutput)

	CreateJob(*elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error)
	CreateJobWithContext(aws.Context, *elastictranscoder.CreateJobInput, ...request.Option) (*elastictranscoder.CreateJobResponse, error)
	CreateJobRequest(*elastictranscoder.CreateJobInput) (*request.Request, *elastictranscoder.CreateJobResponse)

	CreatePipeline(*elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error)
	CreatePipelineWithContext(aws.Context, *elastictranscoder.CreatePipelineInput, ...request.Option) (*elastictranscoder.CreatePipelineOutput, error)
	CreatePipelineRequest(*elastictranscoder.CreatePipelineInput) (*request.Request, *elastictranscoder.CreatePipelineOutput)

	CreatePreset(*elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error)
	CreatePresetWithContext(aws.Context, *elastictranscoder.CreatePresetInput, ...request.Option) (*elastictranscoder.CreatePresetOutput, error)
	CreatePresetRequest(*elastictranscoder.CreatePresetInput) (*request.Request, *elastictranscoder.CreatePresetOutput)

	DeletePipeline(*elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error)
	DeletePipelineWithContext(aws.Context, *elastictranscoder.DeletePipelineInput, ...request.Option) (*elastictranscoder.DeletePipelineOutput, error)
	DeletePipelineRequest(*elastictranscoder.DeletePipelineInput) (*request.Request, *elastictranscoder.DeletePipelineOutput)

	DeletePreset(*elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error)
	DeletePresetWithContext(aws.Context, *elastictranscoder.DeletePresetInput, ...request.Option) (*elastictranscoder.DeletePresetOutput, error)
	DeletePresetRequest(*elastictranscoder.DeletePresetInput) (*request.Request, *elastictranscoder.DeletePresetOutput)

	ListJobsByPipeline(*elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error)
	ListJobsByPipelineWithContext(aws.Context, *elastictranscoder.ListJobsByPipelineInput, ...request.Option) (*elastictranscoder.ListJobsByPipelineOutput, error)
	ListJobsByPipelineRequest(*elastictranscoder.ListJobsByPipelineInput) (*request.Request, *elastictranscoder.ListJobsByPipelineOutput)

	ListJobsByPipelinePages(*elastictranscoder.ListJobsByPipelineInput, func(*elastictranscoder.ListJobsByPipelineOutput, bool) bool) error
	ListJobsByPipelinePagesWithContext(aws.Context, *elastictranscoder.ListJobsByPipelineInput, func(*elastictranscoder.ListJobsByPipelineOutput, bool) bool, ...request.Option) error

	ListJobsByStatus(*elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error)
	ListJobsByStatusWithContext(aws.Context, *elastictranscoder.ListJobsByStatusInput, ...request.Option) (*elastictranscoder.ListJobsByStatusOutput, error)
	ListJobsByStatusRequest(*elastictranscoder.ListJobsByStatusInput) (*request.Request, *elastictranscoder.ListJobsByStatusOutput)

	ListJobsByStatusPages(*elastictranscoder.ListJobsByStatusInput, func(*elastictranscoder.ListJobsByStatusOutput, bool) bool) error
	ListJobsByStatusPagesWithContext(aws.Context, *elastictranscoder.ListJobsByStatusInput, func(*elastictranscoder.ListJobsByStatusOutput, bool) bool, ...request.Option) error

	ListPipelines(*elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error)
	ListPipelinesWithContext(aws.Context, *elastictranscoder.ListPipelinesInput, ...request.Option) (*elastictranscoder.ListPipelinesOutput, error)
	ListPipelinesRequest(*elastictranscoder.ListPipelinesInput) (*request.Request, *elastictranscoder.ListPipelinesOutput)

	ListPipelinesPages(*elastictranscoder.ListPipelinesInput, func(*elastictranscoder.ListPipelinesOutput, bool) bool) error
	ListPipelinesPagesWithContext(aws.Context, *elastictranscoder.ListPipelinesInput, func(*elastictranscoder.ListPipelinesOutput, bool) bool, ...request.Option) error

	ListPresets(*elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error)
	ListPresetsWithContext(aws.Context, *elastictranscoder.ListPresetsInput, ...request.Option) (*elastictranscoder.ListPresetsOutput, error)
	ListPresetsRequest(*elastictranscoder.ListPresetsInput) (*request.Request, *elastictranscoder.ListPresetsOutput)

	ListPresetsPages(*elastictranscoder.ListPresetsInput, func(*elastictranscoder.ListPresetsOutput, bool) bool) error
	ListPresetsPagesWithContext(aws.Context, *elastictranscoder.ListPresetsInput, func(*elastictranscoder.ListPresetsOutput, bool) bool, ...request.Option) error

	ReadJob(*elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error)
	ReadJobWithContext(aws.Context, *elastictranscoder.ReadJobInput, ...request.Option) (*elastictranscoder.ReadJobOutput, error)
	ReadJobRequest(*elastictranscoder.ReadJobInput) (*request.Request, *elastictranscoder.ReadJobOutput)

	ReadPipeline(*elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error)
	ReadPipelineWithContext(aws.Context, *elastictranscoder.ReadPipelineInput, ...request.Option) (*elastictranscoder.ReadPipelineOutput, error)
	ReadPipelineRequest(*elastictranscoder.ReadPipelineInput) (*request.Request, *elastictranscoder.ReadPipelineOutput)

	ReadPreset(*elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error)
	ReadPresetWithContext(aws.Context, *elastictranscoder.ReadPresetInput, ...request.Option) (*elastictranscoder.ReadPresetOutput, error)
	ReadPresetRequest(*elastictranscoder.ReadPresetInput) (*request.Request, *elastictranscoder.ReadPresetOutput)

	TestRole(*elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error)
	TestRoleWithContext(aws.Context, *elastictranscoder.TestRoleInput, ...request.Option) (*elastictranscoder.TestRoleOutput, error)
	TestRoleRequest(*elastictranscoder.TestRoleInput) (*request.Request, *elastictranscoder.TestRoleOutput)

	UpdatePipeline(*elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error)
	UpdatePipelineWithContext(aws.Context, *elastictranscoder.UpdatePipelineInput, ...request.Option) (*elastictranscoder.UpdatePipelineOutput, error)
	UpdatePipelineRequest(*elastictranscoder.UpdatePipelineInput) (*request.Request, *elastictranscoder.UpdatePipelineOutput)

	UpdatePipelineNotifications(*elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error)
	UpdatePipelineNotificationsWithContext(aws.Context, *elastictranscoder.UpdatePipelineNotificationsInput, ...request.Option) (*elastictranscoder.UpdatePipelineNotificationsOutput, error)
	UpdatePipelineNotificationsRequest(*elastictranscoder.UpdatePipelineNotificationsInput) (*request.Request, *elastictranscoder.UpdatePipelineNotificationsOutput)

	UpdatePipelineStatus(*elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error)
	UpdatePipelineStatusWithContext(aws.Context, *elastictranscoder.UpdatePipelineStatusInput, ...request.Option) (*elastictranscoder.UpdatePipelineStatusOutput, error)
	UpdatePipelineStatusRequest(*elastictranscoder.UpdatePipelineStatusInput) (*request.Request, *elastictranscoder.UpdatePipelineStatusOutput)

	WaitUntilJobComplete(*elastictranscoder.ReadJobInput) error
	WaitUntilJobCompleteWithContext(aws.Context, *elastictranscoder.ReadJobInput, ...request.WaiterOption) error
}

var _ ElasticTranscoderAPI = (*elastictranscoder.ElasticTranscoder)(nil)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
