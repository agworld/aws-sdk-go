// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lightsailiface provides an interface to enable mocking the Amazon Lightsail service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lightsailiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lightsail"
)

// LightsailAPI provides an interface to enable mocking the
// lightsail.Lightsail service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Lightsail.
//    func myFunc(svc lightsailiface.LightsailAPI) bool {
//        // Make svc.AllocateStaticIp request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lightsail.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLightsailClient struct {
//        lightsailiface.LightsailAPI
//    }
//    func (m *mockLightsailClient) AllocateStaticIp(input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLightsailClient{}
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
type LightsailAPI interface {
	AllocateStaticIp(*lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error)
	AllocateStaticIpWithContext(aws.Context, *lightsail.AllocateStaticIpInput, ...request.Option) (*lightsail.AllocateStaticIpOutput, error)
	AllocateStaticIpRequest(*lightsail.AllocateStaticIpInput) (*request.Request, *lightsail.AllocateStaticIpOutput)

	AttachStaticIp(*lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error)
	AttachStaticIpWithContext(aws.Context, *lightsail.AttachStaticIpInput, ...request.Option) (*lightsail.AttachStaticIpOutput, error)
	AttachStaticIpRequest(*lightsail.AttachStaticIpInput) (*request.Request, *lightsail.AttachStaticIpOutput)

	CloseInstancePublicPorts(*lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error)
	CloseInstancePublicPortsWithContext(aws.Context, *lightsail.CloseInstancePublicPortsInput, ...request.Option) (*lightsail.CloseInstancePublicPortsOutput, error)
	CloseInstancePublicPortsRequest(*lightsail.CloseInstancePublicPortsInput) (*request.Request, *lightsail.CloseInstancePublicPortsOutput)

	CreateDomain(*lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error)
	CreateDomainWithContext(aws.Context, *lightsail.CreateDomainInput, ...request.Option) (*lightsail.CreateDomainOutput, error)
	CreateDomainRequest(*lightsail.CreateDomainInput) (*request.Request, *lightsail.CreateDomainOutput)

	CreateDomainEntry(*lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error)
	CreateDomainEntryWithContext(aws.Context, *lightsail.CreateDomainEntryInput, ...request.Option) (*lightsail.CreateDomainEntryOutput, error)
	CreateDomainEntryRequest(*lightsail.CreateDomainEntryInput) (*request.Request, *lightsail.CreateDomainEntryOutput)

	CreateInstanceSnapshot(*lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error)
	CreateInstanceSnapshotWithContext(aws.Context, *lightsail.CreateInstanceSnapshotInput, ...request.Option) (*lightsail.CreateInstanceSnapshotOutput, error)
	CreateInstanceSnapshotRequest(*lightsail.CreateInstanceSnapshotInput) (*request.Request, *lightsail.CreateInstanceSnapshotOutput)

	CreateInstances(*lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error)
	CreateInstancesWithContext(aws.Context, *lightsail.CreateInstancesInput, ...request.Option) (*lightsail.CreateInstancesOutput, error)
	CreateInstancesRequest(*lightsail.CreateInstancesInput) (*request.Request, *lightsail.CreateInstancesOutput)

	CreateInstancesFromSnapshot(*lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error)
	CreateInstancesFromSnapshotWithContext(aws.Context, *lightsail.CreateInstancesFromSnapshotInput, ...request.Option) (*lightsail.CreateInstancesFromSnapshotOutput, error)
	CreateInstancesFromSnapshotRequest(*lightsail.CreateInstancesFromSnapshotInput) (*request.Request, *lightsail.CreateInstancesFromSnapshotOutput)

	CreateKeyPair(*lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error)
	CreateKeyPairWithContext(aws.Context, *lightsail.CreateKeyPairInput, ...request.Option) (*lightsail.CreateKeyPairOutput, error)
	CreateKeyPairRequest(*lightsail.CreateKeyPairInput) (*request.Request, *lightsail.CreateKeyPairOutput)

	DeleteDomain(*lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error)
	DeleteDomainWithContext(aws.Context, *lightsail.DeleteDomainInput, ...request.Option) (*lightsail.DeleteDomainOutput, error)
	DeleteDomainRequest(*lightsail.DeleteDomainInput) (*request.Request, *lightsail.DeleteDomainOutput)

	DeleteDomainEntry(*lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error)
	DeleteDomainEntryWithContext(aws.Context, *lightsail.DeleteDomainEntryInput, ...request.Option) (*lightsail.DeleteDomainEntryOutput, error)
	DeleteDomainEntryRequest(*lightsail.DeleteDomainEntryInput) (*request.Request, *lightsail.DeleteDomainEntryOutput)

	DeleteInstance(*lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error)
	DeleteInstanceWithContext(aws.Context, *lightsail.DeleteInstanceInput, ...request.Option) (*lightsail.DeleteInstanceOutput, error)
	DeleteInstanceRequest(*lightsail.DeleteInstanceInput) (*request.Request, *lightsail.DeleteInstanceOutput)

	DeleteInstanceSnapshot(*lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error)
	DeleteInstanceSnapshotWithContext(aws.Context, *lightsail.DeleteInstanceSnapshotInput, ...request.Option) (*lightsail.DeleteInstanceSnapshotOutput, error)
	DeleteInstanceSnapshotRequest(*lightsail.DeleteInstanceSnapshotInput) (*request.Request, *lightsail.DeleteInstanceSnapshotOutput)

	DeleteKeyPair(*lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error)
	DeleteKeyPairWithContext(aws.Context, *lightsail.DeleteKeyPairInput, ...request.Option) (*lightsail.DeleteKeyPairOutput, error)
	DeleteKeyPairRequest(*lightsail.DeleteKeyPairInput) (*request.Request, *lightsail.DeleteKeyPairOutput)

	DetachStaticIp(*lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error)
	DetachStaticIpWithContext(aws.Context, *lightsail.DetachStaticIpInput, ...request.Option) (*lightsail.DetachStaticIpOutput, error)
	DetachStaticIpRequest(*lightsail.DetachStaticIpInput) (*request.Request, *lightsail.DetachStaticIpOutput)

	DownloadDefaultKeyPair(*lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error)
	DownloadDefaultKeyPairWithContext(aws.Context, *lightsail.DownloadDefaultKeyPairInput, ...request.Option) (*lightsail.DownloadDefaultKeyPairOutput, error)
	DownloadDefaultKeyPairRequest(*lightsail.DownloadDefaultKeyPairInput) (*request.Request, *lightsail.DownloadDefaultKeyPairOutput)

	GetActiveNames(*lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error)
	GetActiveNamesWithContext(aws.Context, *lightsail.GetActiveNamesInput, ...request.Option) (*lightsail.GetActiveNamesOutput, error)
	GetActiveNamesRequest(*lightsail.GetActiveNamesInput) (*request.Request, *lightsail.GetActiveNamesOutput)

	GetBlueprints(*lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error)
	GetBlueprintsWithContext(aws.Context, *lightsail.GetBlueprintsInput, ...request.Option) (*lightsail.GetBlueprintsOutput, error)
	GetBlueprintsRequest(*lightsail.GetBlueprintsInput) (*request.Request, *lightsail.GetBlueprintsOutput)

	GetBundles(*lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error)
	GetBundlesWithContext(aws.Context, *lightsail.GetBundlesInput, ...request.Option) (*lightsail.GetBundlesOutput, error)
	GetBundlesRequest(*lightsail.GetBundlesInput) (*request.Request, *lightsail.GetBundlesOutput)

	GetDomain(*lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error)
	GetDomainWithContext(aws.Context, *lightsail.GetDomainInput, ...request.Option) (*lightsail.GetDomainOutput, error)
	GetDomainRequest(*lightsail.GetDomainInput) (*request.Request, *lightsail.GetDomainOutput)

	GetDomains(*lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error)
	GetDomainsWithContext(aws.Context, *lightsail.GetDomainsInput, ...request.Option) (*lightsail.GetDomainsOutput, error)
	GetDomainsRequest(*lightsail.GetDomainsInput) (*request.Request, *lightsail.GetDomainsOutput)

	GetInstance(*lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error)
	GetInstanceWithContext(aws.Context, *lightsail.GetInstanceInput, ...request.Option) (*lightsail.GetInstanceOutput, error)
	GetInstanceRequest(*lightsail.GetInstanceInput) (*request.Request, *lightsail.GetInstanceOutput)

	GetInstanceAccessDetails(*lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstanceAccessDetailsWithContext(aws.Context, *lightsail.GetInstanceAccessDetailsInput, ...request.Option) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstanceAccessDetailsRequest(*lightsail.GetInstanceAccessDetailsInput) (*request.Request, *lightsail.GetInstanceAccessDetailsOutput)

	GetInstanceMetricData(*lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error)
	GetInstanceMetricDataWithContext(aws.Context, *lightsail.GetInstanceMetricDataInput, ...request.Option) (*lightsail.GetInstanceMetricDataOutput, error)
	GetInstanceMetricDataRequest(*lightsail.GetInstanceMetricDataInput) (*request.Request, *lightsail.GetInstanceMetricDataOutput)

	GetInstancePortStates(*lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstancePortStatesWithContext(aws.Context, *lightsail.GetInstancePortStatesInput, ...request.Option) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstancePortStatesRequest(*lightsail.GetInstancePortStatesInput) (*request.Request, *lightsail.GetInstancePortStatesOutput)

	GetInstanceSnapshot(*lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error)
	GetInstanceSnapshotWithContext(aws.Context, *lightsail.GetInstanceSnapshotInput, ...request.Option) (*lightsail.GetInstanceSnapshotOutput, error)
	GetInstanceSnapshotRequest(*lightsail.GetInstanceSnapshotInput) (*request.Request, *lightsail.GetInstanceSnapshotOutput)

	GetInstanceSnapshots(*lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetInstanceSnapshotsWithContext(aws.Context, *lightsail.GetInstanceSnapshotsInput, ...request.Option) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetInstanceSnapshotsRequest(*lightsail.GetInstanceSnapshotsInput) (*request.Request, *lightsail.GetInstanceSnapshotsOutput)

	GetInstanceState(*lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error)
	GetInstanceStateWithContext(aws.Context, *lightsail.GetInstanceStateInput, ...request.Option) (*lightsail.GetInstanceStateOutput, error)
	GetInstanceStateRequest(*lightsail.GetInstanceStateInput) (*request.Request, *lightsail.GetInstanceStateOutput)

	GetInstances(*lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error)
	GetInstancesWithContext(aws.Context, *lightsail.GetInstancesInput, ...request.Option) (*lightsail.GetInstancesOutput, error)
	GetInstancesRequest(*lightsail.GetInstancesInput) (*request.Request, *lightsail.GetInstancesOutput)

	GetKeyPair(*lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error)
	GetKeyPairWithContext(aws.Context, *lightsail.GetKeyPairInput, ...request.Option) (*lightsail.GetKeyPairOutput, error)
	GetKeyPairRequest(*lightsail.GetKeyPairInput) (*request.Request, *lightsail.GetKeyPairOutput)

	GetKeyPairs(*lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error)
	GetKeyPairsWithContext(aws.Context, *lightsail.GetKeyPairsInput, ...request.Option) (*lightsail.GetKeyPairsOutput, error)
	GetKeyPairsRequest(*lightsail.GetKeyPairsInput) (*request.Request, *lightsail.GetKeyPairsOutput)

	GetOperation(*lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error)
	GetOperationWithContext(aws.Context, *lightsail.GetOperationInput, ...request.Option) (*lightsail.GetOperationOutput, error)
	GetOperationRequest(*lightsail.GetOperationInput) (*request.Request, *lightsail.GetOperationOutput)

	GetOperations(*lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error)
	GetOperationsWithContext(aws.Context, *lightsail.GetOperationsInput, ...request.Option) (*lightsail.GetOperationsOutput, error)
	GetOperationsRequest(*lightsail.GetOperationsInput) (*request.Request, *lightsail.GetOperationsOutput)

	GetOperationsForResource(*lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error)
	GetOperationsForResourceWithContext(aws.Context, *lightsail.GetOperationsForResourceInput, ...request.Option) (*lightsail.GetOperationsForResourceOutput, error)
	GetOperationsForResourceRequest(*lightsail.GetOperationsForResourceInput) (*request.Request, *lightsail.GetOperationsForResourceOutput)

	GetRegions(*lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error)
	GetRegionsWithContext(aws.Context, *lightsail.GetRegionsInput, ...request.Option) (*lightsail.GetRegionsOutput, error)
	GetRegionsRequest(*lightsail.GetRegionsInput) (*request.Request, *lightsail.GetRegionsOutput)

	GetStaticIp(*lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error)
	GetStaticIpWithContext(aws.Context, *lightsail.GetStaticIpInput, ...request.Option) (*lightsail.GetStaticIpOutput, error)
	GetStaticIpRequest(*lightsail.GetStaticIpInput) (*request.Request, *lightsail.GetStaticIpOutput)

	GetStaticIps(*lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error)
	GetStaticIpsWithContext(aws.Context, *lightsail.GetStaticIpsInput, ...request.Option) (*lightsail.GetStaticIpsOutput, error)
	GetStaticIpsRequest(*lightsail.GetStaticIpsInput) (*request.Request, *lightsail.GetStaticIpsOutput)

	ImportKeyPair(*lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error)
	ImportKeyPairWithContext(aws.Context, *lightsail.ImportKeyPairInput, ...request.Option) (*lightsail.ImportKeyPairOutput, error)
	ImportKeyPairRequest(*lightsail.ImportKeyPairInput) (*request.Request, *lightsail.ImportKeyPairOutput)

	IsVpcPeered(*lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error)
	IsVpcPeeredWithContext(aws.Context, *lightsail.IsVpcPeeredInput, ...request.Option) (*lightsail.IsVpcPeeredOutput, error)
	IsVpcPeeredRequest(*lightsail.IsVpcPeeredInput) (*request.Request, *lightsail.IsVpcPeeredOutput)

	OpenInstancePublicPorts(*lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error)
	OpenInstancePublicPortsWithContext(aws.Context, *lightsail.OpenInstancePublicPortsInput, ...request.Option) (*lightsail.OpenInstancePublicPortsOutput, error)
	OpenInstancePublicPortsRequest(*lightsail.OpenInstancePublicPortsInput) (*request.Request, *lightsail.OpenInstancePublicPortsOutput)

	PeerVpc(*lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error)
	PeerVpcWithContext(aws.Context, *lightsail.PeerVpcInput, ...request.Option) (*lightsail.PeerVpcOutput, error)
	PeerVpcRequest(*lightsail.PeerVpcInput) (*request.Request, *lightsail.PeerVpcOutput)

	PutInstancePublicPorts(*lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error)
	PutInstancePublicPortsWithContext(aws.Context, *lightsail.PutInstancePublicPortsInput, ...request.Option) (*lightsail.PutInstancePublicPortsOutput, error)
	PutInstancePublicPortsRequest(*lightsail.PutInstancePublicPortsInput) (*request.Request, *lightsail.PutInstancePublicPortsOutput)

	RebootInstance(*lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error)
	RebootInstanceWithContext(aws.Context, *lightsail.RebootInstanceInput, ...request.Option) (*lightsail.RebootInstanceOutput, error)
	RebootInstanceRequest(*lightsail.RebootInstanceInput) (*request.Request, *lightsail.RebootInstanceOutput)

	ReleaseStaticIp(*lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error)
	ReleaseStaticIpWithContext(aws.Context, *lightsail.ReleaseStaticIpInput, ...request.Option) (*lightsail.ReleaseStaticIpOutput, error)
	ReleaseStaticIpRequest(*lightsail.ReleaseStaticIpInput) (*request.Request, *lightsail.ReleaseStaticIpOutput)

	StartInstance(*lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error)
	StartInstanceWithContext(aws.Context, *lightsail.StartInstanceInput, ...request.Option) (*lightsail.StartInstanceOutput, error)
	StartInstanceRequest(*lightsail.StartInstanceInput) (*request.Request, *lightsail.StartInstanceOutput)

	StopInstance(*lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error)
	StopInstanceWithContext(aws.Context, *lightsail.StopInstanceInput, ...request.Option) (*lightsail.StopInstanceOutput, error)
	StopInstanceRequest(*lightsail.StopInstanceInput) (*request.Request, *lightsail.StopInstanceOutput)

	UnpeerVpc(*lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error)
	UnpeerVpcWithContext(aws.Context, *lightsail.UnpeerVpcInput, ...request.Option) (*lightsail.UnpeerVpcOutput, error)
	UnpeerVpcRequest(*lightsail.UnpeerVpcInput) (*request.Request, *lightsail.UnpeerVpcOutput)

	UpdateDomainEntry(*lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error)
	UpdateDomainEntryWithContext(aws.Context, *lightsail.UpdateDomainEntryInput, ...request.Option) (*lightsail.UpdateDomainEntryOutput, error)
	UpdateDomainEntryRequest(*lightsail.UpdateDomainEntryInput) (*request.Request, *lightsail.UpdateDomainEntryOutput)
}

var _ LightsailAPI = (*lightsail.Lightsail)(nil)
//Added a line for testing
