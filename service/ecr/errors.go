// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

const (

	// ErrCodeEmptyUploadException for service response error code
	// "EmptyUploadException".
	//
	// The specified layer upload does not contain any layer parts.
	ErrCodeEmptyUploadException = "EmptyUploadException"

	// ErrCodeImageAlreadyExistsException for service response error code
	// "ImageAlreadyExistsException".
	//
	// The specified image has already been pushed, and there are no changes to
	// the manifest or image tag since the last push.
	ErrCodeImageAlreadyExistsException = "ImageAlreadyExistsException"

	// ErrCodeImageNotFoundException for service response error code
	// "ImageNotFoundException".
	//
	// The image requested does not exist in the specified repository.
	ErrCodeImageNotFoundException = "ImageNotFoundException"

	// ErrCodeInvalidLayerException for service response error code
	// "InvalidLayerException".
	//
	// The layer digest calculation performed by Amazon ECR upon receipt of the
	// image layer does not match the digest specified.
	ErrCodeInvalidLayerException = "InvalidLayerException"

	// ErrCodeInvalidLayerPartException for service response error code
	// "InvalidLayerPartException".
	//
	// The layer part size is not valid, or the first byte specified is not consecutive
	// to the last byte of a previous layer part upload.
	ErrCodeInvalidLayerPartException = "InvalidLayerPartException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// The specified parameter is invalid. Review the available parameters for the
	// API request.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeLayerAlreadyExistsException for service response error code
	// "LayerAlreadyExistsException".
	//
	// The image layer already exists in the associated repository.
	ErrCodeLayerAlreadyExistsException = "LayerAlreadyExistsException"

	// ErrCodeLayerInaccessibleException for service response error code
	// "LayerInaccessibleException".
	//
	// The specified layer is not available because it is not associated with an
	// image. Unassociated image layers may be cleaned up at any time.
	ErrCodeLayerInaccessibleException = "LayerInaccessibleException"

	// ErrCodeLayerPartTooSmallException for service response error code
	// "LayerPartTooSmallException".
	//
	// Layer parts must be at least 5 MiB in size.
	ErrCodeLayerPartTooSmallException = "LayerPartTooSmallException"

	// ErrCodeLayersNotFoundException for service response error code
	// "LayersNotFoundException".
	//
	// The specified layers could not be found, or the specified layer is not valid
	// for this repository.
	ErrCodeLayersNotFoundException = "LayersNotFoundException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The operation did not succeed because it would have exceeded a service limit
	// for your account. For more information, see Amazon ECR Default Service Limits
	// (http://docs.aws.amazon.com/AmazonECR/latest/userguide/service_limits.html)
	// in the Amazon EC2 Container Registry User Guide.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeRepositoryAlreadyExistsException for service response error code
	// "RepositoryAlreadyExistsException".
	//
	// The specified repository already exists in the specified registry.
	ErrCodeRepositoryAlreadyExistsException = "RepositoryAlreadyExistsException"

	// ErrCodeRepositoryNotEmptyException for service response error code
	// "RepositoryNotEmptyException".
	//
	// The specified repository contains images. To delete a repository that contains
	// images, you must force the deletion with the force parameter.
	ErrCodeRepositoryNotEmptyException = "RepositoryNotEmptyException"

	// ErrCodeRepositoryNotFoundException for service response error code
	// "RepositoryNotFoundException".
	//
	// The specified repository could not be found. Check the spelling of the specified
	// repository and ensure that you are performing operations on the correct registry.
	ErrCodeRepositoryNotFoundException = "RepositoryNotFoundException"

	// ErrCodeRepositoryPolicyNotFoundException for service response error code
	// "RepositoryPolicyNotFoundException".
	//
	// The specified repository and registry combination does not have an associated
	// repository policy.
	ErrCodeRepositoryPolicyNotFoundException = "RepositoryPolicyNotFoundException"

	// ErrCodeServerException for service response error code
	// "ServerException".
	//
	// These errors are usually caused by a server-side issue.
	ErrCodeServerException = "ServerException"

	// ErrCodeUploadNotFoundException for service response error code
	// "UploadNotFoundException".
	//
	// The upload could not be found, or the specified upload id is not valid for
	// this repository.
	ErrCodeUploadNotFoundException = "UploadNotFoundException"
)
//Added a line for testing
//Adding another line for Git event testing part 2
