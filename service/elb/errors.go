// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elb

const (

	// ErrCodeAccessPointNotFoundException for service response error code
	// "LoadBalancerNotFound".
	//
	// The specified load balancer does not exist.
	ErrCodeAccessPointNotFoundException = "LoadBalancerNotFound"

	// ErrCodeCertificateNotFoundException for service response error code
	// "CertificateNotFound".
	//
	// The specified ARN does not refer to a valid SSL certificate in AWS Identity
	// and Access Management (IAM) or AWS Certificate Manager (ACM). Note that if
	// you recently uploaded the certificate to IAM, this error might indicate that
	// the certificate is not fully available yet.
	ErrCodeCertificateNotFoundException = "CertificateNotFound"

	// ErrCodeDependencyThrottleException for service response error code
	// "DependencyThrottle".
	ErrCodeDependencyThrottleException = "DependencyThrottle"

	// ErrCodeDuplicateAccessPointNameException for service response error code
	// "DuplicateLoadBalancerName".
	//
	// The specified load balancer name already exists for this account.
	ErrCodeDuplicateAccessPointNameException = "DuplicateLoadBalancerName"

	// ErrCodeDuplicateListenerException for service response error code
	// "DuplicateListener".
	//
	// A listener already exists for the specified load balancer name and port,
	// but with a different instance port, protocol, or SSL certificate.
	ErrCodeDuplicateListenerException = "DuplicateListener"

	// ErrCodeDuplicatePolicyNameException for service response error code
	// "DuplicatePolicyName".
	//
	// A policy with the specified name already exists for this load balancer.
	ErrCodeDuplicatePolicyNameException = "DuplicatePolicyName"

	// ErrCodeDuplicateTagKeysException for service response error code
	// "DuplicateTagKeys".
	//
	// A tag key was specified more than once.
	ErrCodeDuplicateTagKeysException = "DuplicateTagKeys"

	// ErrCodeInvalidConfigurationRequestException for service response error code
	// "InvalidConfigurationRequest".
	//
	// The requested configuration change is not valid.
	ErrCodeInvalidConfigurationRequestException = "InvalidConfigurationRequest"

	// ErrCodeInvalidEndPointException for service response error code
	// "InvalidInstance".
	//
	// The specified endpoint is not valid.
	ErrCodeInvalidEndPointException = "InvalidInstance"

	// ErrCodeInvalidSchemeException for service response error code
	// "InvalidScheme".
	//
	// The specified value for the schema is not valid. You can only specify a scheme
	// for load balancers in a VPC.
	ErrCodeInvalidSchemeException = "InvalidScheme"

	// ErrCodeInvalidSecurityGroupException for service response error code
	// "InvalidSecurityGroup".
	//
	// One or more of the specified security groups do not exist.
	ErrCodeInvalidSecurityGroupException = "InvalidSecurityGroup"

	// ErrCodeInvalidSubnetException for service response error code
	// "InvalidSubnet".
	//
	// The specified VPC has no associated Internet gateway.
	ErrCodeInvalidSubnetException = "InvalidSubnet"

	// ErrCodeListenerNotFoundException for service response error code
	// "ListenerNotFound".
	//
	// The load balancer does not have a listener configured at the specified port.
	ErrCodeListenerNotFoundException = "ListenerNotFound"

	// ErrCodeLoadBalancerAttributeNotFoundException for service response error code
	// "LoadBalancerAttributeNotFound".
	//
	// The specified load balancer attribute does not exist.
	ErrCodeLoadBalancerAttributeNotFoundException = "LoadBalancerAttributeNotFound"

	// ErrCodePolicyNotFoundException for service response error code
	// "PolicyNotFound".
	//
	// One or more of the specified policies do not exist.
	ErrCodePolicyNotFoundException = "PolicyNotFound"

	// ErrCodePolicyTypeNotFoundException for service response error code
	// "PolicyTypeNotFound".
	//
	// One or more of the specified policy types do not exist.
	ErrCodePolicyTypeNotFoundException = "PolicyTypeNotFound"

	// ErrCodeSubnetNotFoundException for service response error code
	// "SubnetNotFound".
	//
	// One or more of the specified subnets do not exist.
	ErrCodeSubnetNotFoundException = "SubnetNotFound"

	// ErrCodeTooManyAccessPointsException for service response error code
	// "TooManyLoadBalancers".
	//
	// The quota for the number of load balancers has been reached.
	ErrCodeTooManyAccessPointsException = "TooManyLoadBalancers"

	// ErrCodeTooManyPoliciesException for service response error code
	// "TooManyPolicies".
	//
	// The quota for the number of policies for this load balancer has been reached.
	ErrCodeTooManyPoliciesException = "TooManyPolicies"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTags".
	//
	// The quota for the number of tags that can be assigned to a load balancer
	// has been reached.
	ErrCodeTooManyTagsException = "TooManyTags"

	// ErrCodeUnsupportedProtocolException for service response error code
	// "UnsupportedProtocol".
	//
	// The specified protocol or signature version is not supported.
	ErrCodeUnsupportedProtocolException = "UnsupportedProtocol"
)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
