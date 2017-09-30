// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

const (

	// ErrCodeActionNotFoundException for service response error code
	// "ActionNotFoundException".
	//
	// The specified action cannot be found.
	ErrCodeActionNotFoundException = "ActionNotFoundException"

	// ErrCodeActionTypeNotFoundException for service response error code
	// "ActionTypeNotFoundException".
	//
	// The specified action type cannot be found.
	ErrCodeActionTypeNotFoundException = "ActionTypeNotFoundException"

	// ErrCodeApprovalAlreadyCompletedException for service response error code
	// "ApprovalAlreadyCompletedException".
	//
	// The approval action has already been approved or rejected.
	ErrCodeApprovalAlreadyCompletedException = "ApprovalAlreadyCompletedException"

	// ErrCodeInvalidActionDeclarationException for service response error code
	// "InvalidActionDeclarationException".
	//
	// The specified action declaration was specified in an invalid format.
	ErrCodeInvalidActionDeclarationException = "InvalidActionDeclarationException"

	// ErrCodeInvalidApprovalTokenException for service response error code
	// "InvalidApprovalTokenException".
	//
	// The approval request already received a response or has expired.
	ErrCodeInvalidApprovalTokenException = "InvalidApprovalTokenException"

	// ErrCodeInvalidBlockerDeclarationException for service response error code
	// "InvalidBlockerDeclarationException".
	//
	// Reserved for future use.
	ErrCodeInvalidBlockerDeclarationException = "InvalidBlockerDeclarationException"

	// ErrCodeInvalidClientTokenException for service response error code
	// "InvalidClientTokenException".
	//
	// The client token was specified in an invalid format
	ErrCodeInvalidClientTokenException = "InvalidClientTokenException"

	// ErrCodeInvalidJobException for service response error code
	// "InvalidJobException".
	//
	// The specified job was specified in an invalid format or cannot be found.
	ErrCodeInvalidJobException = "InvalidJobException"

	// ErrCodeInvalidJobStateException for service response error code
	// "InvalidJobStateException".
	//
	// The specified job state was specified in an invalid format.
	ErrCodeInvalidJobStateException = "InvalidJobStateException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The next token was specified in an invalid format. Make sure that the next
	// token you provided is the token returned by a previous call.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidNonceException for service response error code
	// "InvalidNonceException".
	//
	// The specified nonce was specified in an invalid format.
	ErrCodeInvalidNonceException = "InvalidNonceException"

	// ErrCodeInvalidStageDeclarationException for service response error code
	// "InvalidStageDeclarationException".
	//
	// The specified stage declaration was specified in an invalid format.
	ErrCodeInvalidStageDeclarationException = "InvalidStageDeclarationException"

	// ErrCodeInvalidStructureException for service response error code
	// "InvalidStructureException".
	//
	// The specified structure was specified in an invalid format.
	ErrCodeInvalidStructureException = "InvalidStructureException"

	// ErrCodeJobNotFoundException for service response error code
	// "JobNotFoundException".
	//
	// The specified job was specified in an invalid format or cannot be found.
	ErrCodeJobNotFoundException = "JobNotFoundException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The number of pipelines associated with the AWS account has exceeded the
	// limit allowed for the account.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotLatestPipelineExecutionException for service response error code
	// "NotLatestPipelineExecutionException".
	//
	// The stage has failed in a later run of the pipeline and the pipelineExecutionId
	// associated with the request is out of date.
	ErrCodeNotLatestPipelineExecutionException = "NotLatestPipelineExecutionException"

	// ErrCodePipelineExecutionNotFoundException for service response error code
	// "PipelineExecutionNotFoundException".
	//
	// The pipeline execution was specified in an invalid format or cannot be found,
	// or an execution ID does not belong to the specified pipeline.
	ErrCodePipelineExecutionNotFoundException = "PipelineExecutionNotFoundException"

	// ErrCodePipelineNameInUseException for service response error code
	// "PipelineNameInUseException".
	//
	// The specified pipeline name is already in use.
	ErrCodePipelineNameInUseException = "PipelineNameInUseException"

	// ErrCodePipelineNotFoundException for service response error code
	// "PipelineNotFoundException".
	//
	// The specified pipeline was specified in an invalid format or cannot be found.
	ErrCodePipelineNotFoundException = "PipelineNotFoundException"

	// ErrCodePipelineVersionNotFoundException for service response error code
	// "PipelineVersionNotFoundException".
	//
	// The specified pipeline version was specified in an invalid format or cannot
	// be found.
	ErrCodePipelineVersionNotFoundException = "PipelineVersionNotFoundException"

	// ErrCodeStageNotFoundException for service response error code
	// "StageNotFoundException".
	//
	// The specified stage was specified in an invalid format or cannot be found.
	ErrCodeStageNotFoundException = "StageNotFoundException"

	// ErrCodeStageNotRetryableException for service response error code
	// "StageNotRetryableException".
	//
	// The specified stage can't be retried because the pipeline structure or stage
	// state changed after the stage was not completed; the stage contains no failed
	// actions; one or more actions are still in progress; or another retry attempt
	// is already in progress.
	ErrCodeStageNotRetryableException = "StageNotRetryableException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The validation was specified in an invalid format.
	ErrCodeValidationException = "ValidationException"
)
//Added a line for testing
