// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotiface provides an interface to enable mocking the AWS IoT service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iot"
)

// IoTAPI provides an interface to enable mocking the
// iot.IoT service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT.
//    func myFunc(svc iotiface.IoTAPI) bool {
//        // Make svc.AcceptCertificateTransfer request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iot.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTClient struct {
//        iotiface.IoTAPI
//    }
//    func (m *mockIoTClient) AcceptCertificateTransfer(input *iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTClient{}
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
type IoTAPI interface {
	AcceptCertificateTransfer(*iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error)
	AcceptCertificateTransferWithContext(aws.Context, *iot.AcceptCertificateTransferInput, ...request.Option) (*iot.AcceptCertificateTransferOutput, error)
	AcceptCertificateTransferRequest(*iot.AcceptCertificateTransferInput) (*request.Request, *iot.AcceptCertificateTransferOutput)

	AttachPrincipalPolicy(*iot.AttachPrincipalPolicyInput) (*iot.AttachPrincipalPolicyOutput, error)
	AttachPrincipalPolicyWithContext(aws.Context, *iot.AttachPrincipalPolicyInput, ...request.Option) (*iot.AttachPrincipalPolicyOutput, error)
	AttachPrincipalPolicyRequest(*iot.AttachPrincipalPolicyInput) (*request.Request, *iot.AttachPrincipalPolicyOutput)

	AttachThingPrincipal(*iot.AttachThingPrincipalInput) (*iot.AttachThingPrincipalOutput, error)
	AttachThingPrincipalWithContext(aws.Context, *iot.AttachThingPrincipalInput, ...request.Option) (*iot.AttachThingPrincipalOutput, error)
	AttachThingPrincipalRequest(*iot.AttachThingPrincipalInput) (*request.Request, *iot.AttachThingPrincipalOutput)

	CancelCertificateTransfer(*iot.CancelCertificateTransferInput) (*iot.CancelCertificateTransferOutput, error)
	CancelCertificateTransferWithContext(aws.Context, *iot.CancelCertificateTransferInput, ...request.Option) (*iot.CancelCertificateTransferOutput, error)
	CancelCertificateTransferRequest(*iot.CancelCertificateTransferInput) (*request.Request, *iot.CancelCertificateTransferOutput)

	CreateCertificateFromCsr(*iot.CreateCertificateFromCsrInput) (*iot.CreateCertificateFromCsrOutput, error)
	CreateCertificateFromCsrWithContext(aws.Context, *iot.CreateCertificateFromCsrInput, ...request.Option) (*iot.CreateCertificateFromCsrOutput, error)
	CreateCertificateFromCsrRequest(*iot.CreateCertificateFromCsrInput) (*request.Request, *iot.CreateCertificateFromCsrOutput)

	CreateKeysAndCertificate(*iot.CreateKeysAndCertificateInput) (*iot.CreateKeysAndCertificateOutput, error)
	CreateKeysAndCertificateWithContext(aws.Context, *iot.CreateKeysAndCertificateInput, ...request.Option) (*iot.CreateKeysAndCertificateOutput, error)
	CreateKeysAndCertificateRequest(*iot.CreateKeysAndCertificateInput) (*request.Request, *iot.CreateKeysAndCertificateOutput)

	CreatePolicy(*iot.CreatePolicyInput) (*iot.CreatePolicyOutput, error)
	CreatePolicyWithContext(aws.Context, *iot.CreatePolicyInput, ...request.Option) (*iot.CreatePolicyOutput, error)
	CreatePolicyRequest(*iot.CreatePolicyInput) (*request.Request, *iot.CreatePolicyOutput)

	CreatePolicyVersion(*iot.CreatePolicyVersionInput) (*iot.CreatePolicyVersionOutput, error)
	CreatePolicyVersionWithContext(aws.Context, *iot.CreatePolicyVersionInput, ...request.Option) (*iot.CreatePolicyVersionOutput, error)
	CreatePolicyVersionRequest(*iot.CreatePolicyVersionInput) (*request.Request, *iot.CreatePolicyVersionOutput)

	CreateThing(*iot.CreateThingInput) (*iot.CreateThingOutput, error)
	CreateThingWithContext(aws.Context, *iot.CreateThingInput, ...request.Option) (*iot.CreateThingOutput, error)
	CreateThingRequest(*iot.CreateThingInput) (*request.Request, *iot.CreateThingOutput)

	CreateThingType(*iot.CreateThingTypeInput) (*iot.CreateThingTypeOutput, error)
	CreateThingTypeWithContext(aws.Context, *iot.CreateThingTypeInput, ...request.Option) (*iot.CreateThingTypeOutput, error)
	CreateThingTypeRequest(*iot.CreateThingTypeInput) (*request.Request, *iot.CreateThingTypeOutput)

	CreateTopicRule(*iot.CreateTopicRuleInput) (*iot.CreateTopicRuleOutput, error)
	CreateTopicRuleWithContext(aws.Context, *iot.CreateTopicRuleInput, ...request.Option) (*iot.CreateTopicRuleOutput, error)
	CreateTopicRuleRequest(*iot.CreateTopicRuleInput) (*request.Request, *iot.CreateTopicRuleOutput)

	DeleteCACertificate(*iot.DeleteCACertificateInput) (*iot.DeleteCACertificateOutput, error)
	DeleteCACertificateWithContext(aws.Context, *iot.DeleteCACertificateInput, ...request.Option) (*iot.DeleteCACertificateOutput, error)
	DeleteCACertificateRequest(*iot.DeleteCACertificateInput) (*request.Request, *iot.DeleteCACertificateOutput)

	DeleteCertificate(*iot.DeleteCertificateInput) (*iot.DeleteCertificateOutput, error)
	DeleteCertificateWithContext(aws.Context, *iot.DeleteCertificateInput, ...request.Option) (*iot.DeleteCertificateOutput, error)
	DeleteCertificateRequest(*iot.DeleteCertificateInput) (*request.Request, *iot.DeleteCertificateOutput)

	DeletePolicy(*iot.DeletePolicyInput) (*iot.DeletePolicyOutput, error)
	DeletePolicyWithContext(aws.Context, *iot.DeletePolicyInput, ...request.Option) (*iot.DeletePolicyOutput, error)
	DeletePolicyRequest(*iot.DeletePolicyInput) (*request.Request, *iot.DeletePolicyOutput)

	DeletePolicyVersion(*iot.DeletePolicyVersionInput) (*iot.DeletePolicyVersionOutput, error)
	DeletePolicyVersionWithContext(aws.Context, *iot.DeletePolicyVersionInput, ...request.Option) (*iot.DeletePolicyVersionOutput, error)
	DeletePolicyVersionRequest(*iot.DeletePolicyVersionInput) (*request.Request, *iot.DeletePolicyVersionOutput)

	DeleteRegistrationCode(*iot.DeleteRegistrationCodeInput) (*iot.DeleteRegistrationCodeOutput, error)
	DeleteRegistrationCodeWithContext(aws.Context, *iot.DeleteRegistrationCodeInput, ...request.Option) (*iot.DeleteRegistrationCodeOutput, error)
	DeleteRegistrationCodeRequest(*iot.DeleteRegistrationCodeInput) (*request.Request, *iot.DeleteRegistrationCodeOutput)

	DeleteThing(*iot.DeleteThingInput) (*iot.DeleteThingOutput, error)
	DeleteThingWithContext(aws.Context, *iot.DeleteThingInput, ...request.Option) (*iot.DeleteThingOutput, error)
	DeleteThingRequest(*iot.DeleteThingInput) (*request.Request, *iot.DeleteThingOutput)

	DeleteThingType(*iot.DeleteThingTypeInput) (*iot.DeleteThingTypeOutput, error)
	DeleteThingTypeWithContext(aws.Context, *iot.DeleteThingTypeInput, ...request.Option) (*iot.DeleteThingTypeOutput, error)
	DeleteThingTypeRequest(*iot.DeleteThingTypeInput) (*request.Request, *iot.DeleteThingTypeOutput)

	DeleteTopicRule(*iot.DeleteTopicRuleInput) (*iot.DeleteTopicRuleOutput, error)
	DeleteTopicRuleWithContext(aws.Context, *iot.DeleteTopicRuleInput, ...request.Option) (*iot.DeleteTopicRuleOutput, error)
	DeleteTopicRuleRequest(*iot.DeleteTopicRuleInput) (*request.Request, *iot.DeleteTopicRuleOutput)

	DeprecateThingType(*iot.DeprecateThingTypeInput) (*iot.DeprecateThingTypeOutput, error)
	DeprecateThingTypeWithContext(aws.Context, *iot.DeprecateThingTypeInput, ...request.Option) (*iot.DeprecateThingTypeOutput, error)
	DeprecateThingTypeRequest(*iot.DeprecateThingTypeInput) (*request.Request, *iot.DeprecateThingTypeOutput)

	DescribeCACertificate(*iot.DescribeCACertificateInput) (*iot.DescribeCACertificateOutput, error)
	DescribeCACertificateWithContext(aws.Context, *iot.DescribeCACertificateInput, ...request.Option) (*iot.DescribeCACertificateOutput, error)
	DescribeCACertificateRequest(*iot.DescribeCACertificateInput) (*request.Request, *iot.DescribeCACertificateOutput)

	DescribeCertificate(*iot.DescribeCertificateInput) (*iot.DescribeCertificateOutput, error)
	DescribeCertificateWithContext(aws.Context, *iot.DescribeCertificateInput, ...request.Option) (*iot.DescribeCertificateOutput, error)
	DescribeCertificateRequest(*iot.DescribeCertificateInput) (*request.Request, *iot.DescribeCertificateOutput)

	DescribeEndpoint(*iot.DescribeEndpointInput) (*iot.DescribeEndpointOutput, error)
	DescribeEndpointWithContext(aws.Context, *iot.DescribeEndpointInput, ...request.Option) (*iot.DescribeEndpointOutput, error)
	DescribeEndpointRequest(*iot.DescribeEndpointInput) (*request.Request, *iot.DescribeEndpointOutput)

	DescribeThing(*iot.DescribeThingInput) (*iot.DescribeThingOutput, error)
	DescribeThingWithContext(aws.Context, *iot.DescribeThingInput, ...request.Option) (*iot.DescribeThingOutput, error)
	DescribeThingRequest(*iot.DescribeThingInput) (*request.Request, *iot.DescribeThingOutput)

	DescribeThingType(*iot.DescribeThingTypeInput) (*iot.DescribeThingTypeOutput, error)
	DescribeThingTypeWithContext(aws.Context, *iot.DescribeThingTypeInput, ...request.Option) (*iot.DescribeThingTypeOutput, error)
	DescribeThingTypeRequest(*iot.DescribeThingTypeInput) (*request.Request, *iot.DescribeThingTypeOutput)

	DetachPrincipalPolicy(*iot.DetachPrincipalPolicyInput) (*iot.DetachPrincipalPolicyOutput, error)
	DetachPrincipalPolicyWithContext(aws.Context, *iot.DetachPrincipalPolicyInput, ...request.Option) (*iot.DetachPrincipalPolicyOutput, error)
	DetachPrincipalPolicyRequest(*iot.DetachPrincipalPolicyInput) (*request.Request, *iot.DetachPrincipalPolicyOutput)

	DetachThingPrincipal(*iot.DetachThingPrincipalInput) (*iot.DetachThingPrincipalOutput, error)
	DetachThingPrincipalWithContext(aws.Context, *iot.DetachThingPrincipalInput, ...request.Option) (*iot.DetachThingPrincipalOutput, error)
	DetachThingPrincipalRequest(*iot.DetachThingPrincipalInput) (*request.Request, *iot.DetachThingPrincipalOutput)

	DisableTopicRule(*iot.DisableTopicRuleInput) (*iot.DisableTopicRuleOutput, error)
	DisableTopicRuleWithContext(aws.Context, *iot.DisableTopicRuleInput, ...request.Option) (*iot.DisableTopicRuleOutput, error)
	DisableTopicRuleRequest(*iot.DisableTopicRuleInput) (*request.Request, *iot.DisableTopicRuleOutput)

	EnableTopicRule(*iot.EnableTopicRuleInput) (*iot.EnableTopicRuleOutput, error)
	EnableTopicRuleWithContext(aws.Context, *iot.EnableTopicRuleInput, ...request.Option) (*iot.EnableTopicRuleOutput, error)
	EnableTopicRuleRequest(*iot.EnableTopicRuleInput) (*request.Request, *iot.EnableTopicRuleOutput)

	GetLoggingOptions(*iot.GetLoggingOptionsInput) (*iot.GetLoggingOptionsOutput, error)
	GetLoggingOptionsWithContext(aws.Context, *iot.GetLoggingOptionsInput, ...request.Option) (*iot.GetLoggingOptionsOutput, error)
	GetLoggingOptionsRequest(*iot.GetLoggingOptionsInput) (*request.Request, *iot.GetLoggingOptionsOutput)

	GetPolicy(*iot.GetPolicyInput) (*iot.GetPolicyOutput, error)
	GetPolicyWithContext(aws.Context, *iot.GetPolicyInput, ...request.Option) (*iot.GetPolicyOutput, error)
	GetPolicyRequest(*iot.GetPolicyInput) (*request.Request, *iot.GetPolicyOutput)

	GetPolicyVersion(*iot.GetPolicyVersionInput) (*iot.GetPolicyVersionOutput, error)
	GetPolicyVersionWithContext(aws.Context, *iot.GetPolicyVersionInput, ...request.Option) (*iot.GetPolicyVersionOutput, error)
	GetPolicyVersionRequest(*iot.GetPolicyVersionInput) (*request.Request, *iot.GetPolicyVersionOutput)

	GetRegistrationCode(*iot.GetRegistrationCodeInput) (*iot.GetRegistrationCodeOutput, error)
	GetRegistrationCodeWithContext(aws.Context, *iot.GetRegistrationCodeInput, ...request.Option) (*iot.GetRegistrationCodeOutput, error)
	GetRegistrationCodeRequest(*iot.GetRegistrationCodeInput) (*request.Request, *iot.GetRegistrationCodeOutput)

	GetTopicRule(*iot.GetTopicRuleInput) (*iot.GetTopicRuleOutput, error)
	GetTopicRuleWithContext(aws.Context, *iot.GetTopicRuleInput, ...request.Option) (*iot.GetTopicRuleOutput, error)
	GetTopicRuleRequest(*iot.GetTopicRuleInput) (*request.Request, *iot.GetTopicRuleOutput)

	ListCACertificates(*iot.ListCACertificatesInput) (*iot.ListCACertificatesOutput, error)
	ListCACertificatesWithContext(aws.Context, *iot.ListCACertificatesInput, ...request.Option) (*iot.ListCACertificatesOutput, error)
	ListCACertificatesRequest(*iot.ListCACertificatesInput) (*request.Request, *iot.ListCACertificatesOutput)

	ListCertificates(*iot.ListCertificatesInput) (*iot.ListCertificatesOutput, error)
	ListCertificatesWithContext(aws.Context, *iot.ListCertificatesInput, ...request.Option) (*iot.ListCertificatesOutput, error)
	ListCertificatesRequest(*iot.ListCertificatesInput) (*request.Request, *iot.ListCertificatesOutput)

	ListCertificatesByCA(*iot.ListCertificatesByCAInput) (*iot.ListCertificatesByCAOutput, error)
	ListCertificatesByCAWithContext(aws.Context, *iot.ListCertificatesByCAInput, ...request.Option) (*iot.ListCertificatesByCAOutput, error)
	ListCertificatesByCARequest(*iot.ListCertificatesByCAInput) (*request.Request, *iot.ListCertificatesByCAOutput)

	ListOutgoingCertificates(*iot.ListOutgoingCertificatesInput) (*iot.ListOutgoingCertificatesOutput, error)
	ListOutgoingCertificatesWithContext(aws.Context, *iot.ListOutgoingCertificatesInput, ...request.Option) (*iot.ListOutgoingCertificatesOutput, error)
	ListOutgoingCertificatesRequest(*iot.ListOutgoingCertificatesInput) (*request.Request, *iot.ListOutgoingCertificatesOutput)

	ListPolicies(*iot.ListPoliciesInput) (*iot.ListPoliciesOutput, error)
	ListPoliciesWithContext(aws.Context, *iot.ListPoliciesInput, ...request.Option) (*iot.ListPoliciesOutput, error)
	ListPoliciesRequest(*iot.ListPoliciesInput) (*request.Request, *iot.ListPoliciesOutput)

	ListPolicyPrincipals(*iot.ListPolicyPrincipalsInput) (*iot.ListPolicyPrincipalsOutput, error)
	ListPolicyPrincipalsWithContext(aws.Context, *iot.ListPolicyPrincipalsInput, ...request.Option) (*iot.ListPolicyPrincipalsOutput, error)
	ListPolicyPrincipalsRequest(*iot.ListPolicyPrincipalsInput) (*request.Request, *iot.ListPolicyPrincipalsOutput)

	ListPolicyVersions(*iot.ListPolicyVersionsInput) (*iot.ListPolicyVersionsOutput, error)
	ListPolicyVersionsWithContext(aws.Context, *iot.ListPolicyVersionsInput, ...request.Option) (*iot.ListPolicyVersionsOutput, error)
	ListPolicyVersionsRequest(*iot.ListPolicyVersionsInput) (*request.Request, *iot.ListPolicyVersionsOutput)

	ListPrincipalPolicies(*iot.ListPrincipalPoliciesInput) (*iot.ListPrincipalPoliciesOutput, error)
	ListPrincipalPoliciesWithContext(aws.Context, *iot.ListPrincipalPoliciesInput, ...request.Option) (*iot.ListPrincipalPoliciesOutput, error)
	ListPrincipalPoliciesRequest(*iot.ListPrincipalPoliciesInput) (*request.Request, *iot.ListPrincipalPoliciesOutput)

	ListPrincipalThings(*iot.ListPrincipalThingsInput) (*iot.ListPrincipalThingsOutput, error)
	ListPrincipalThingsWithContext(aws.Context, *iot.ListPrincipalThingsInput, ...request.Option) (*iot.ListPrincipalThingsOutput, error)
	ListPrincipalThingsRequest(*iot.ListPrincipalThingsInput) (*request.Request, *iot.ListPrincipalThingsOutput)

	ListThingPrincipals(*iot.ListThingPrincipalsInput) (*iot.ListThingPrincipalsOutput, error)
	ListThingPrincipalsWithContext(aws.Context, *iot.ListThingPrincipalsInput, ...request.Option) (*iot.ListThingPrincipalsOutput, error)
	ListThingPrincipalsRequest(*iot.ListThingPrincipalsInput) (*request.Request, *iot.ListThingPrincipalsOutput)

	ListThingTypes(*iot.ListThingTypesInput) (*iot.ListThingTypesOutput, error)
	ListThingTypesWithContext(aws.Context, *iot.ListThingTypesInput, ...request.Option) (*iot.ListThingTypesOutput, error)
	ListThingTypesRequest(*iot.ListThingTypesInput) (*request.Request, *iot.ListThingTypesOutput)

	ListThings(*iot.ListThingsInput) (*iot.ListThingsOutput, error)
	ListThingsWithContext(aws.Context, *iot.ListThingsInput, ...request.Option) (*iot.ListThingsOutput, error)
	ListThingsRequest(*iot.ListThingsInput) (*request.Request, *iot.ListThingsOutput)

	ListTopicRules(*iot.ListTopicRulesInput) (*iot.ListTopicRulesOutput, error)
	ListTopicRulesWithContext(aws.Context, *iot.ListTopicRulesInput, ...request.Option) (*iot.ListTopicRulesOutput, error)
	ListTopicRulesRequest(*iot.ListTopicRulesInput) (*request.Request, *iot.ListTopicRulesOutput)

	RegisterCACertificate(*iot.RegisterCACertificateInput) (*iot.RegisterCACertificateOutput, error)
	RegisterCACertificateWithContext(aws.Context, *iot.RegisterCACertificateInput, ...request.Option) (*iot.RegisterCACertificateOutput, error)
	RegisterCACertificateRequest(*iot.RegisterCACertificateInput) (*request.Request, *iot.RegisterCACertificateOutput)

	RegisterCertificate(*iot.RegisterCertificateInput) (*iot.RegisterCertificateOutput, error)
	RegisterCertificateWithContext(aws.Context, *iot.RegisterCertificateInput, ...request.Option) (*iot.RegisterCertificateOutput, error)
	RegisterCertificateRequest(*iot.RegisterCertificateInput) (*request.Request, *iot.RegisterCertificateOutput)

	RejectCertificateTransfer(*iot.RejectCertificateTransferInput) (*iot.RejectCertificateTransferOutput, error)
	RejectCertificateTransferWithContext(aws.Context, *iot.RejectCertificateTransferInput, ...request.Option) (*iot.RejectCertificateTransferOutput, error)
	RejectCertificateTransferRequest(*iot.RejectCertificateTransferInput) (*request.Request, *iot.RejectCertificateTransferOutput)

	ReplaceTopicRule(*iot.ReplaceTopicRuleInput) (*iot.ReplaceTopicRuleOutput, error)
	ReplaceTopicRuleWithContext(aws.Context, *iot.ReplaceTopicRuleInput, ...request.Option) (*iot.ReplaceTopicRuleOutput, error)
	ReplaceTopicRuleRequest(*iot.ReplaceTopicRuleInput) (*request.Request, *iot.ReplaceTopicRuleOutput)

	SetDefaultPolicyVersion(*iot.SetDefaultPolicyVersionInput) (*iot.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionWithContext(aws.Context, *iot.SetDefaultPolicyVersionInput, ...request.Option) (*iot.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionRequest(*iot.SetDefaultPolicyVersionInput) (*request.Request, *iot.SetDefaultPolicyVersionOutput)

	SetLoggingOptions(*iot.SetLoggingOptionsInput) (*iot.SetLoggingOptionsOutput, error)
	SetLoggingOptionsWithContext(aws.Context, *iot.SetLoggingOptionsInput, ...request.Option) (*iot.SetLoggingOptionsOutput, error)
	SetLoggingOptionsRequest(*iot.SetLoggingOptionsInput) (*request.Request, *iot.SetLoggingOptionsOutput)

	TransferCertificate(*iot.TransferCertificateInput) (*iot.TransferCertificateOutput, error)
	TransferCertificateWithContext(aws.Context, *iot.TransferCertificateInput, ...request.Option) (*iot.TransferCertificateOutput, error)
	TransferCertificateRequest(*iot.TransferCertificateInput) (*request.Request, *iot.TransferCertificateOutput)

	UpdateCACertificate(*iot.UpdateCACertificateInput) (*iot.UpdateCACertificateOutput, error)
	UpdateCACertificateWithContext(aws.Context, *iot.UpdateCACertificateInput, ...request.Option) (*iot.UpdateCACertificateOutput, error)
	UpdateCACertificateRequest(*iot.UpdateCACertificateInput) (*request.Request, *iot.UpdateCACertificateOutput)

	UpdateCertificate(*iot.UpdateCertificateInput) (*iot.UpdateCertificateOutput, error)
	UpdateCertificateWithContext(aws.Context, *iot.UpdateCertificateInput, ...request.Option) (*iot.UpdateCertificateOutput, error)
	UpdateCertificateRequest(*iot.UpdateCertificateInput) (*request.Request, *iot.UpdateCertificateOutput)

	UpdateThing(*iot.UpdateThingInput) (*iot.UpdateThingOutput, error)
	UpdateThingWithContext(aws.Context, *iot.UpdateThingInput, ...request.Option) (*iot.UpdateThingOutput, error)
	UpdateThingRequest(*iot.UpdateThingInput) (*request.Request, *iot.UpdateThingOutput)
}

var _ IoTAPI = (*iot.IoT)(nil)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
