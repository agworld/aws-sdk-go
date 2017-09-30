// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

const (

	// ErrCodeDomainLimitExceeded for service response error code
	// "DomainLimitExceeded".
	//
	// The number of domains has exceeded the allowed threshold for the account.
	ErrCodeDomainLimitExceeded = "DomainLimitExceeded"

	// ErrCodeDuplicateRequest for service response error code
	// "DuplicateRequest".
	//
	// The request is already in progress for the domain.
	ErrCodeDuplicateRequest = "DuplicateRequest"

	// ErrCodeInvalidInput for service response error code
	// "InvalidInput".
	//
	// The requested item is not acceptable. For example, for an OperationId it
	// may refer to the ID of an operation that is already completed. For a domain
	// name, it may not be a valid domain name or belong to the requester account.
	ErrCodeInvalidInput = "InvalidInput"

	// ErrCodeOperationLimitExceeded for service response error code
	// "OperationLimitExceeded".
	//
	// The number of operations or jobs running exceeded the allowed threshold for
	// the account.
	ErrCodeOperationLimitExceeded = "OperationLimitExceeded"

	// ErrCodeTLDRulesViolation for service response error code
	// "TLDRulesViolation".
	//
	// The top-level domain does not support this operation.
	ErrCodeTLDRulesViolation = "TLDRulesViolation"

	// ErrCodeUnsupportedTLD for service response error code
	// "UnsupportedTLD".
	//
	// Amazon Route 53 does not support this top-level domain.
	ErrCodeUnsupportedTLD = "UnsupportedTLD"
)
//Added a line for testing
