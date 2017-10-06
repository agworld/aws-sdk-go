// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

const (

	// ErrCodeCreationLimitExceededException for service response error code
	// "CreationLimitExceededException".
	//
	// The exception is thrown when customer tries to create a record (e.g. budget),
	// but the number this record already exceeds the limitation.
	ErrCodeCreationLimitExceededException = "CreationLimitExceededException"

	// ErrCodeDuplicateRecordException for service response error code
	// "DuplicateRecordException".
	//
	// The exception is thrown when customer tries to create a record (e.g. budget)
	// that already exists.
	ErrCodeDuplicateRecordException = "DuplicateRecordException"

	// ErrCodeExpiredNextTokenException for service response error code
	// "ExpiredNextTokenException".
	//
	// This exception is thrown if the paging token is expired - past its TTL
	ErrCodeExpiredNextTokenException = "ExpiredNextTokenException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// This exception is thrown on an unknown internal failure.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// This exception is thrown if paging token signature didn't match the token,
	// or the paging token isn't for this request
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// This exception is thrown if any request is given an invalid parameter. E.g.,
	// if a required Date field is null.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// This exception is thrown if a requested entity is not found. E.g., if a budget
	// id doesn't exist for an account ID.
	ErrCodeNotFoundException = "NotFoundException"
)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
