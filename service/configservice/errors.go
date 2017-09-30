// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

const (

	// ErrCodeInsufficientDeliveryPolicyException for service response error code
	// "InsufficientDeliveryPolicyException".
	//
	// Your Amazon S3 bucket policy does not permit AWS Config to write to it.
	ErrCodeInsufficientDeliveryPolicyException = "InsufficientDeliveryPolicyException"

	// ErrCodeInsufficientPermissionsException for service response error code
	// "InsufficientPermissionsException".
	//
	// Indicates one of the following errors:
	//
	//    * The rule cannot be created because the IAM role assigned to AWS Config
	//    lacks permissions to perform the config:Put* action.
	//
	//    * The AWS Lambda function cannot be invoked. Check the function ARN, and
	//    check the function's permissions.
	ErrCodeInsufficientPermissionsException = "InsufficientPermissionsException"

	// ErrCodeInvalidConfigurationRecorderNameException for service response error code
	// "InvalidConfigurationRecorderNameException".
	//
	// You have provided a configuration recorder name that is not valid.
	ErrCodeInvalidConfigurationRecorderNameException = "InvalidConfigurationRecorderNameException"

	// ErrCodeInvalidDeliveryChannelNameException for service response error code
	// "InvalidDeliveryChannelNameException".
	//
	// The specified delivery channel name is not valid.
	ErrCodeInvalidDeliveryChannelNameException = "InvalidDeliveryChannelNameException"

	// ErrCodeInvalidLimitException for service response error code
	// "InvalidLimitException".
	//
	// The specified limit is outside the allowable range.
	ErrCodeInvalidLimitException = "InvalidLimitException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The specified next token is invalid. Specify the NextToken string that was
	// returned in the previous response to get the next page of results.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// One or more of the specified parameters are invalid. Verify that your parameters
	// are valid and try again.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidRecordingGroupException for service response error code
	// "InvalidRecordingGroupException".
	//
	// AWS Config throws an exception if the recording group does not contain a
	// valid list of resource types. Invalid values could also be incorrectly formatted.
	ErrCodeInvalidRecordingGroupException = "InvalidRecordingGroupException"

	// ErrCodeInvalidResultTokenException for service response error code
	// "InvalidResultTokenException".
	//
	// The specified ResultToken is invalid.
	ErrCodeInvalidResultTokenException = "InvalidResultTokenException"

	// ErrCodeInvalidRoleException for service response error code
	// "InvalidRoleException".
	//
	// You have provided a null or empty role ARN.
	ErrCodeInvalidRoleException = "InvalidRoleException"

	// ErrCodeInvalidS3KeyPrefixException for service response error code
	// "InvalidS3KeyPrefixException".
	//
	// The specified Amazon S3 key prefix is not valid.
	ErrCodeInvalidS3KeyPrefixException = "InvalidS3KeyPrefixException"

	// ErrCodeInvalidSNSTopicARNException for service response error code
	// "InvalidSNSTopicARNException".
	//
	// The specified Amazon SNS topic does not exist.
	ErrCodeInvalidSNSTopicARNException = "InvalidSNSTopicARNException"

	// ErrCodeInvalidTimeRangeException for service response error code
	// "InvalidTimeRangeException".
	//
	// The specified time range is not valid. The earlier time is not chronologically
	// before the later time.
	ErrCodeInvalidTimeRangeException = "InvalidTimeRangeException"

	// ErrCodeLastDeliveryChannelDeleteFailedException for service response error code
	// "LastDeliveryChannelDeleteFailedException".
	//
	// You cannot delete the delivery channel you specified because the configuration
	// recorder is running.
	ErrCodeLastDeliveryChannelDeleteFailedException = "LastDeliveryChannelDeleteFailedException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// This exception is thrown if an evaluation is in progress or if you call the
	// StartConfigRulesEvaluation API more than once per minute.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMaxNumberOfConfigRulesExceededException for service response error code
	// "MaxNumberOfConfigRulesExceededException".
	//
	// Failed to add the AWS Config rule because the account already contains the
	// maximum number of 50 rules. Consider deleting any deactivated rules before
	// adding new rules.
	ErrCodeMaxNumberOfConfigRulesExceededException = "MaxNumberOfConfigRulesExceededException"

	// ErrCodeMaxNumberOfConfigurationRecordersExceededException for service response error code
	// "MaxNumberOfConfigurationRecordersExceededException".
	//
	// You have reached the limit on the number of recorders you can create.
	ErrCodeMaxNumberOfConfigurationRecordersExceededException = "MaxNumberOfConfigurationRecordersExceededException"

	// ErrCodeMaxNumberOfDeliveryChannelsExceededException for service response error code
	// "MaxNumberOfDeliveryChannelsExceededException".
	//
	// You have reached the limit on the number of delivery channels you can create.
	ErrCodeMaxNumberOfDeliveryChannelsExceededException = "MaxNumberOfDeliveryChannelsExceededException"

	// ErrCodeNoAvailableConfigurationRecorderException for service response error code
	// "NoAvailableConfigurationRecorderException".
	//
	// There are no configuration recorders available to provide the role needed
	// to describe your resources. Create a configuration recorder.
	ErrCodeNoAvailableConfigurationRecorderException = "NoAvailableConfigurationRecorderException"

	// ErrCodeNoAvailableDeliveryChannelException for service response error code
	// "NoAvailableDeliveryChannelException".
	//
	// There is no delivery channel available to record configurations.
	ErrCodeNoAvailableDeliveryChannelException = "NoAvailableDeliveryChannelException"

	// ErrCodeNoRunningConfigurationRecorderException for service response error code
	// "NoRunningConfigurationRecorderException".
	//
	// There is no configuration recorder running.
	ErrCodeNoRunningConfigurationRecorderException = "NoRunningConfigurationRecorderException"

	// ErrCodeNoSuchBucketException for service response error code
	// "NoSuchBucketException".
	//
	// The specified Amazon S3 bucket does not exist.
	ErrCodeNoSuchBucketException = "NoSuchBucketException"

	// ErrCodeNoSuchConfigRuleException for service response error code
	// "NoSuchConfigRuleException".
	//
	// One or more AWS Config rules in the request are invalid. Verify that the
	// rule names are correct and try again.
	ErrCodeNoSuchConfigRuleException = "NoSuchConfigRuleException"

	// ErrCodeNoSuchConfigurationRecorderException for service response error code
	// "NoSuchConfigurationRecorderException".
	//
	// You have specified a configuration recorder that does not exist.
	ErrCodeNoSuchConfigurationRecorderException = "NoSuchConfigurationRecorderException"

	// ErrCodeNoSuchDeliveryChannelException for service response error code
	// "NoSuchDeliveryChannelException".
	//
	// You have specified a delivery channel that does not exist.
	ErrCodeNoSuchDeliveryChannelException = "NoSuchDeliveryChannelException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The rule is currently being deleted or the rule is deleting your evaluation
	// results. Try your request again later.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotDiscoveredException for service response error code
	// "ResourceNotDiscoveredException".
	//
	// You have specified a resource that is either unknown or has not been discovered.
	ErrCodeResourceNotDiscoveredException = "ResourceNotDiscoveredException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The requested action is not valid.
	ErrCodeValidationException = "ValidationException"
)
//Added a line for testing
