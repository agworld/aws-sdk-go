// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

const (

	// ErrCodeDuplicateResourceException for service response error code
	// "DuplicateResourceException".
	//
	// The specified resource is a duplicate.
	ErrCodeDuplicateResourceException = "DuplicateResourceException"

	// ErrCodeInvalidParametersException for service response error code
	// "InvalidParametersException".
	//
	// One or more parameters provided to the operation are invalid.
	ErrCodeInvalidParametersException = "InvalidParametersException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// An attempt was made to modify a resource that is in an invalid state. Inspect
	// the resource you are using for this operation to ensure that all resource
	// states are valid before retrying the operation.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The current limits of the service would have been exceeded by this operation.
	// Reduce the resource use or increase the service limits and retry the operation.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The operation was requested against a resource that is currently in use.
	// Free the resource from use and retry the operation.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTagOptionNotMigratedException for service response error code
	// "TagOptionNotMigratedException".
	//
	// An operation requiring TagOptions failed because the TagOptions migration
	// process has not been performed for this account. Please use the AWS console
	// to perform the migration process before retrying the operation.
	ErrCodeTagOptionNotMigratedException = "TagOptionNotMigratedException"
)
//Added a line for testing
