// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

const (

	// ErrCodeInternalError for service response error code
	// "InternalError".
	//
	// An internal error has occured.
	ErrCodeInternalError = "InternalError"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// A parameter specified in the request is not valid, is unsupported, or cannot
	// be used.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeMissingRequiredParameterException for service response error code
	// "MissingRequiredParameterException".
	//
	// The request is missing a required parameter. Ensure that you have supplied
	// all the required parameters for the request.
	ErrCodeMissingRequiredParameterException = "MissingRequiredParameterException"

	// ErrCodeNoConnectorsAvailableException for service response error code
	// "NoConnectorsAvailableException".
	//
	// No connectors are available to handle this request. Please associate connector(s)
	// and verify any existing connectors are healthy and can respond to requests.
	ErrCodeNoConnectorsAvailableException = "NoConnectorsAvailableException"

	// ErrCodeOperationNotPermittedException for service response error code
	// "OperationNotPermittedException".
	//
	// The specified operation is not allowed. This error can occur for a number
	// of reasons; for example, you might be trying to start a Replication Run before
	// seed Replication Run.
	ErrCodeOperationNotPermittedException = "OperationNotPermittedException"

	// ErrCodeReplicationJobAlreadyExistsException for service response error code
	// "ReplicationJobAlreadyExistsException".
	//
	// An active Replication Job already exists for the specified server.
	ErrCodeReplicationJobAlreadyExistsException = "ReplicationJobAlreadyExistsException"

	// ErrCodeReplicationJobNotFoundException for service response error code
	// "ReplicationJobNotFoundException".
	//
	// The specified Replication Job cannot be found.
	ErrCodeReplicationJobNotFoundException = "ReplicationJobNotFoundException"

	// ErrCodeReplicationRunLimitExceededException for service response error code
	// "ReplicationRunLimitExceededException".
	//
	// This user has exceeded the maximum allowed Replication Run limit.
	ErrCodeReplicationRunLimitExceededException = "ReplicationRunLimitExceededException"

	// ErrCodeServerCannotBeReplicatedException for service response error code
	// "ServerCannotBeReplicatedException".
	//
	// The provided server cannot be replicated.
	ErrCodeServerCannotBeReplicatedException = "ServerCannotBeReplicatedException"

	// ErrCodeUnauthorizedOperationException for service response error code
	// "UnauthorizedOperationException".
	//
	// This user does not have permissions to perform this operation.
	ErrCodeUnauthorizedOperationException = "UnauthorizedOperationException"
)
//Added a line for testing
