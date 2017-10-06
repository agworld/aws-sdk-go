// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

const (

	// ErrCodeAlarmsLimitExceededException for service response error code
	// "AlarmsLimitExceededException".
	//
	// The maximum number of alarms for a deployment group (10) was exceeded.
	ErrCodeAlarmsLimitExceededException = "AlarmsLimitExceededException"

	// ErrCodeApplicationAlreadyExistsException for service response error code
	// "ApplicationAlreadyExistsException".
	//
	// An application with the specified name already exists with the applicable
	// IAM user or AWS account.
	ErrCodeApplicationAlreadyExistsException = "ApplicationAlreadyExistsException"

	// ErrCodeApplicationDoesNotExistException for service response error code
	// "ApplicationDoesNotExistException".
	//
	// The application does not exist with the applicable IAM user or AWS account.
	ErrCodeApplicationDoesNotExistException = "ApplicationDoesNotExistException"

	// ErrCodeApplicationLimitExceededException for service response error code
	// "ApplicationLimitExceededException".
	//
	// More applications were attempted to be created than are allowed.
	ErrCodeApplicationLimitExceededException = "ApplicationLimitExceededException"

	// ErrCodeApplicationNameRequiredException for service response error code
	// "ApplicationNameRequiredException".
	//
	// The minimum number of required application names was not specified.
	ErrCodeApplicationNameRequiredException = "ApplicationNameRequiredException"

	// ErrCodeBatchLimitExceededException for service response error code
	// "BatchLimitExceededException".
	//
	// The maximum number of names or IDs allowed for this request (100) was exceeded.
	ErrCodeBatchLimitExceededException = "BatchLimitExceededException"

	// ErrCodeBucketNameFilterRequiredException for service response error code
	// "BucketNameFilterRequiredException".
	//
	// A bucket name is required, but was not provided.
	ErrCodeBucketNameFilterRequiredException = "BucketNameFilterRequiredException"

	// ErrCodeDeploymentAlreadyCompletedException for service response error code
	// "DeploymentAlreadyCompletedException".
	//
	// The deployment is already complete.
	ErrCodeDeploymentAlreadyCompletedException = "DeploymentAlreadyCompletedException"

	// ErrCodeDeploymentConfigAlreadyExistsException for service response error code
	// "DeploymentConfigAlreadyExistsException".
	//
	// A deployment configuration with the specified name already exists with the
	// applicable IAM user or AWS account.
	ErrCodeDeploymentConfigAlreadyExistsException = "DeploymentConfigAlreadyExistsException"

	// ErrCodeDeploymentConfigDoesNotExistException for service response error code
	// "DeploymentConfigDoesNotExistException".
	//
	// The deployment configuration does not exist with the applicable IAM user
	// or AWS account.
	ErrCodeDeploymentConfigDoesNotExistException = "DeploymentConfigDoesNotExistException"

	// ErrCodeDeploymentConfigInUseException for service response error code
	// "DeploymentConfigInUseException".
	//
	// The deployment configuration is still in use.
	ErrCodeDeploymentConfigInUseException = "DeploymentConfigInUseException"

	// ErrCodeDeploymentConfigLimitExceededException for service response error code
	// "DeploymentConfigLimitExceededException".
	//
	// The deployment configurations limit was exceeded.
	ErrCodeDeploymentConfigLimitExceededException = "DeploymentConfigLimitExceededException"

	// ErrCodeDeploymentConfigNameRequiredException for service response error code
	// "DeploymentConfigNameRequiredException".
	//
	// The deployment configuration name was not specified.
	ErrCodeDeploymentConfigNameRequiredException = "DeploymentConfigNameRequiredException"

	// ErrCodeDeploymentDoesNotExistException for service response error code
	// "DeploymentDoesNotExistException".
	//
	// The deployment does not exist with the applicable IAM user or AWS account.
	ErrCodeDeploymentDoesNotExistException = "DeploymentDoesNotExistException"

	// ErrCodeDeploymentGroupAlreadyExistsException for service response error code
	// "DeploymentGroupAlreadyExistsException".
	//
	// A deployment group with the specified name already exists with the applicable
	// IAM user or AWS account.
	ErrCodeDeploymentGroupAlreadyExistsException = "DeploymentGroupAlreadyExistsException"

	// ErrCodeDeploymentGroupDoesNotExistException for service response error code
	// "DeploymentGroupDoesNotExistException".
	//
	// The named deployment group does not exist with the applicable IAM user or
	// AWS account.
	ErrCodeDeploymentGroupDoesNotExistException = "DeploymentGroupDoesNotExistException"

	// ErrCodeDeploymentGroupLimitExceededException for service response error code
	// "DeploymentGroupLimitExceededException".
	//
	// The deployment groups limit was exceeded.
	ErrCodeDeploymentGroupLimitExceededException = "DeploymentGroupLimitExceededException"

	// ErrCodeDeploymentGroupNameRequiredException for service response error code
	// "DeploymentGroupNameRequiredException".
	//
	// The deployment group name was not specified.
	ErrCodeDeploymentGroupNameRequiredException = "DeploymentGroupNameRequiredException"

	// ErrCodeDeploymentIdRequiredException for service response error code
	// "DeploymentIdRequiredException".
	//
	// At least one deployment ID must be specified.
	ErrCodeDeploymentIdRequiredException = "DeploymentIdRequiredException"

	// ErrCodeDeploymentIsNotInReadyStateException for service response error code
	// "DeploymentIsNotInReadyStateException".
	//
	// The deployment does not have a status of Ready and can't continue yet.
	ErrCodeDeploymentIsNotInReadyStateException = "DeploymentIsNotInReadyStateException"

	// ErrCodeDeploymentLimitExceededException for service response error code
	// "DeploymentLimitExceededException".
	//
	// The number of allowed deployments was exceeded.
	ErrCodeDeploymentLimitExceededException = "DeploymentLimitExceededException"

	// ErrCodeDeploymentNotStartedException for service response error code
	// "DeploymentNotStartedException".
	//
	// The specified deployment has not started.
	ErrCodeDeploymentNotStartedException = "DeploymentNotStartedException"

	// ErrCodeDescriptionTooLongException for service response error code
	// "DescriptionTooLongException".
	//
	// The description is too long.
	ErrCodeDescriptionTooLongException = "DescriptionTooLongException"

	// ErrCodeIamArnRequiredException for service response error code
	// "IamArnRequiredException".
	//
	// No IAM ARN was included in the request. You must use an IAM session ARN or
	// IAM user ARN in the request.
	ErrCodeIamArnRequiredException = "IamArnRequiredException"

	// ErrCodeIamSessionArnAlreadyRegisteredException for service response error code
	// "IamSessionArnAlreadyRegisteredException".
	//
	// The request included an IAM session ARN that has already been used to register
	// a different instance.
	ErrCodeIamSessionArnAlreadyRegisteredException = "IamSessionArnAlreadyRegisteredException"

	// ErrCodeIamUserArnAlreadyRegisteredException for service response error code
	// "IamUserArnAlreadyRegisteredException".
	//
	// The specified IAM user ARN is already registered with an on-premises instance.
	ErrCodeIamUserArnAlreadyRegisteredException = "IamUserArnAlreadyRegisteredException"

	// ErrCodeIamUserArnRequiredException for service response error code
	// "IamUserArnRequiredException".
	//
	// An IAM user ARN was not specified.
	ErrCodeIamUserArnRequiredException = "IamUserArnRequiredException"

	// ErrCodeInstanceDoesNotExistException for service response error code
	// "InstanceDoesNotExistException".
	//
	// The specified instance does not exist in the deployment group.
	ErrCodeInstanceDoesNotExistException = "InstanceDoesNotExistException"

	// ErrCodeInstanceIdRequiredException for service response error code
	// "InstanceIdRequiredException".
	//
	// The instance ID was not specified.
	ErrCodeInstanceIdRequiredException = "InstanceIdRequiredException"

	// ErrCodeInstanceLimitExceededException for service response error code
	// "InstanceLimitExceededException".
	//
	// The maximum number of allowed on-premises instances in a single call was
	// exceeded.
	ErrCodeInstanceLimitExceededException = "InstanceLimitExceededException"

	// ErrCodeInstanceNameAlreadyRegisteredException for service response error code
	// "InstanceNameAlreadyRegisteredException".
	//
	// The specified on-premises instance name is already registered.
	ErrCodeInstanceNameAlreadyRegisteredException = "InstanceNameAlreadyRegisteredException"

	// ErrCodeInstanceNameRequiredException for service response error code
	// "InstanceNameRequiredException".
	//
	// An on-premises instance name was not specified.
	ErrCodeInstanceNameRequiredException = "InstanceNameRequiredException"

	// ErrCodeInstanceNotRegisteredException for service response error code
	// "InstanceNotRegisteredException".
	//
	// The specified on-premises instance is not registered.
	ErrCodeInstanceNotRegisteredException = "InstanceNotRegisteredException"

	// ErrCodeInvalidAlarmConfigException for service response error code
	// "InvalidAlarmConfigException".
	//
	// The format of the alarm configuration is invalid. Possible causes include:
	//
	//    * The alarm list is null.
	//
	//    * The alarm object is null.
	//
	//    * The alarm name is empty or null or exceeds the 255 character limit.
	//
	//    * Two alarms with the same name have been specified.
	//
	//    * The alarm configuration is enabled but the alarm list is empty.
	ErrCodeInvalidAlarmConfigException = "InvalidAlarmConfigException"

	// ErrCodeInvalidApplicationNameException for service response error code
	// "InvalidApplicationNameException".
	//
	// The application name was specified in an invalid format.
	ErrCodeInvalidApplicationNameException = "InvalidApplicationNameException"

	// ErrCodeInvalidAutoRollbackConfigException for service response error code
	// "InvalidAutoRollbackConfigException".
	//
	// The automatic rollback configuration was specified in an invalid format.
	// For example, automatic rollback is enabled but an invalid triggering event
	// type or no event types were listed.
	ErrCodeInvalidAutoRollbackConfigException = "InvalidAutoRollbackConfigException"

	// ErrCodeInvalidAutoScalingGroupException for service response error code
	// "InvalidAutoScalingGroupException".
	//
	// The Auto Scaling group was specified in an invalid format or does not exist.
	ErrCodeInvalidAutoScalingGroupException = "InvalidAutoScalingGroupException"

	// ErrCodeInvalidBlueGreenDeploymentConfigurationException for service response error code
	// "InvalidBlueGreenDeploymentConfigurationException".
	//
	// The configuration for the blue/green deployment group was provided in an
	// invalid format. For information about deployment configuration format, see
	// CreateDeploymentConfig.
	ErrCodeInvalidBlueGreenDeploymentConfigurationException = "InvalidBlueGreenDeploymentConfigurationException"

	// ErrCodeInvalidBucketNameFilterException for service response error code
	// "InvalidBucketNameFilterException".
	//
	// The bucket name either doesn't exist or was specified in an invalid format.
	ErrCodeInvalidBucketNameFilterException = "InvalidBucketNameFilterException"

	// ErrCodeInvalidDeployedStateFilterException for service response error code
	// "InvalidDeployedStateFilterException".
	//
	// The deployed state filter was specified in an invalid format.
	ErrCodeInvalidDeployedStateFilterException = "InvalidDeployedStateFilterException"

	// ErrCodeInvalidDeploymentConfigNameException for service response error code
	// "InvalidDeploymentConfigNameException".
	//
	// The deployment configuration name was specified in an invalid format.
	ErrCodeInvalidDeploymentConfigNameException = "InvalidDeploymentConfigNameException"

	// ErrCodeInvalidDeploymentGroupNameException for service response error code
	// "InvalidDeploymentGroupNameException".
	//
	// The deployment group name was specified in an invalid format.
	ErrCodeInvalidDeploymentGroupNameException = "InvalidDeploymentGroupNameException"

	// ErrCodeInvalidDeploymentIdException for service response error code
	// "InvalidDeploymentIdException".
	//
	// At least one of the deployment IDs was specified in an invalid format.
	ErrCodeInvalidDeploymentIdException = "InvalidDeploymentIdException"

	// ErrCodeInvalidDeploymentInstanceTypeException for service response error code
	// "InvalidDeploymentInstanceTypeException".
	//
	// An instance type was specified for an in-place deployment. Instance types
	// are supported for blue/green deployments only.
	ErrCodeInvalidDeploymentInstanceTypeException = "InvalidDeploymentInstanceTypeException"

	// ErrCodeInvalidDeploymentStatusException for service response error code
	// "InvalidDeploymentStatusException".
	//
	// The specified deployment status doesn't exist or cannot be determined.
	ErrCodeInvalidDeploymentStatusException = "InvalidDeploymentStatusException"

	// ErrCodeInvalidDeploymentStyleException for service response error code
	// "InvalidDeploymentStyleException".
	//
	// An invalid deployment style was specified. Valid deployment types include
	// "IN_PLACE" and "BLUE_GREEN". Valid deployment options include "WITH_TRAFFIC_CONTROL"
	// and "WITHOUT_TRAFFIC_CONTROL".
	ErrCodeInvalidDeploymentStyleException = "InvalidDeploymentStyleException"

	// ErrCodeInvalidEC2TagCombinationException for service response error code
	// "InvalidEC2TagCombinationException".
	//
	// A call was submitted that specified both Ec2TagFilters and Ec2TagSet, but
	// only one of these data types can be used in a single call.
	ErrCodeInvalidEC2TagCombinationException = "InvalidEC2TagCombinationException"

	// ErrCodeInvalidEC2TagException for service response error code
	// "InvalidEC2TagException".
	//
	// The tag was specified in an invalid format.
	ErrCodeInvalidEC2TagException = "InvalidEC2TagException"

	// ErrCodeInvalidFileExistsBehaviorException for service response error code
	// "InvalidFileExistsBehaviorException".
	//
	// An invalid fileExistsBehavior option was specified to determine how AWS CodeDeploy
	// handles files or directories that already exist in a deployment target location
	// but weren't part of the previous successful deployment. Valid values include
	// "DISALLOW", "OVERWRITE", and "RETAIN".
	ErrCodeInvalidFileExistsBehaviorException = "InvalidFileExistsBehaviorException"

	// ErrCodeInvalidIamSessionArnException for service response error code
	// "InvalidIamSessionArnException".
	//
	// The IAM session ARN was specified in an invalid format.
	ErrCodeInvalidIamSessionArnException = "InvalidIamSessionArnException"

	// ErrCodeInvalidIamUserArnException for service response error code
	// "InvalidIamUserArnException".
	//
	// The IAM user ARN was specified in an invalid format.
	ErrCodeInvalidIamUserArnException = "InvalidIamUserArnException"

	// ErrCodeInvalidInstanceNameException for service response error code
	// "InvalidInstanceNameException".
	//
	// The specified on-premises instance name was specified in an invalid format.
	ErrCodeInvalidInstanceNameException = "InvalidInstanceNameException"

	// ErrCodeInvalidInstanceStatusException for service response error code
	// "InvalidInstanceStatusException".
	//
	// The specified instance status does not exist.
	ErrCodeInvalidInstanceStatusException = "InvalidInstanceStatusException"

	// ErrCodeInvalidInstanceTypeException for service response error code
	// "InvalidInstanceTypeException".
	//
	// An invalid instance type was specified for instances in a blue/green deployment.
	// Valid values include "Blue" for an original environment and "Green" for a
	// replacement environment.
	ErrCodeInvalidInstanceTypeException = "InvalidInstanceTypeException"

	// ErrCodeInvalidKeyPrefixFilterException for service response error code
	// "InvalidKeyPrefixFilterException".
	//
	// The specified key prefix filter was specified in an invalid format.
	ErrCodeInvalidKeyPrefixFilterException = "InvalidKeyPrefixFilterException"

	// ErrCodeInvalidLoadBalancerInfoException for service response error code
	// "InvalidLoadBalancerInfoException".
	//
	// An invalid load balancer name, or no load balancer name, was specified.
	ErrCodeInvalidLoadBalancerInfoException = "InvalidLoadBalancerInfoException"

	// ErrCodeInvalidMinimumHealthyHostValueException for service response error code
	// "InvalidMinimumHealthyHostValueException".
	//
	// The minimum healthy instance value was specified in an invalid format.
	ErrCodeInvalidMinimumHealthyHostValueException = "InvalidMinimumHealthyHostValueException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The next token was specified in an invalid format.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidOnPremisesTagCombinationException for service response error code
	// "InvalidOnPremisesTagCombinationException".
	//
	// A call was submitted that specified both OnPremisesTagFilters and OnPremisesTagSet,
	// but only one of these data types can be used in a single call.
	ErrCodeInvalidOnPremisesTagCombinationException = "InvalidOnPremisesTagCombinationException"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// An invalid operation was detected.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeInvalidRegistrationStatusException for service response error code
	// "InvalidRegistrationStatusException".
	//
	// The registration status was specified in an invalid format.
	ErrCodeInvalidRegistrationStatusException = "InvalidRegistrationStatusException"

	// ErrCodeInvalidRevisionException for service response error code
	// "InvalidRevisionException".
	//
	// The revision was specified in an invalid format.
	ErrCodeInvalidRevisionException = "InvalidRevisionException"

	// ErrCodeInvalidRoleException for service response error code
	// "InvalidRoleException".
	//
	// The service role ARN was specified in an invalid format. Or, if an Auto Scaling
	// group was specified, the specified service role does not grant the appropriate
	// permissions to Auto Scaling.
	ErrCodeInvalidRoleException = "InvalidRoleException"

	// ErrCodeInvalidSortByException for service response error code
	// "InvalidSortByException".
	//
	// The column name to sort by is either not present or was specified in an invalid
	// format.
	ErrCodeInvalidSortByException = "InvalidSortByException"

	// ErrCodeInvalidSortOrderException for service response error code
	// "InvalidSortOrderException".
	//
	// The sort order was specified in an invalid format.
	ErrCodeInvalidSortOrderException = "InvalidSortOrderException"

	// ErrCodeInvalidTagException for service response error code
	// "InvalidTagException".
	//
	// The specified tag was specified in an invalid format.
	ErrCodeInvalidTagException = "InvalidTagException"

	// ErrCodeInvalidTagFilterException for service response error code
	// "InvalidTagFilterException".
	//
	// The specified tag filter was specified in an invalid format.
	ErrCodeInvalidTagFilterException = "InvalidTagFilterException"

	// ErrCodeInvalidTargetInstancesException for service response error code
	// "InvalidTargetInstancesException".
	//
	// The target instance configuration is invalid. Possible causes include:
	//
	//    * Configuration data for target instances was entered for an in-place
	//    deployment.
	//
	//    * The limit of 10 tags for a tag type was exceeded.
	//
	//    * The combined length of the tag names exceeded the limit.
	//
	//    * A specified tag is not currently applied to any instances.
	ErrCodeInvalidTargetInstancesException = "InvalidTargetInstancesException"

	// ErrCodeInvalidTimeRangeException for service response error code
	// "InvalidTimeRangeException".
	//
	// The specified time range was specified in an invalid format.
	ErrCodeInvalidTimeRangeException = "InvalidTimeRangeException"

	// ErrCodeInvalidTriggerConfigException for service response error code
	// "InvalidTriggerConfigException".
	//
	// The trigger was specified in an invalid format.
	ErrCodeInvalidTriggerConfigException = "InvalidTriggerConfigException"

	// ErrCodeLifecycleHookLimitExceededException for service response error code
	// "LifecycleHookLimitExceededException".
	//
	// The limit for lifecycle hooks was exceeded.
	ErrCodeLifecycleHookLimitExceededException = "LifecycleHookLimitExceededException"

	// ErrCodeMultipleIamArnsProvidedException for service response error code
	// "MultipleIamArnsProvidedException".
	//
	// Both an IAM user ARN and an IAM session ARN were included in the request.
	// Use only one ARN type.
	ErrCodeMultipleIamArnsProvidedException = "MultipleIamArnsProvidedException"

	// ErrCodeResourceValidationException for service response error code
	// "ResourceValidationException".
	//
	// The specified resource could not be validated.
	ErrCodeResourceValidationException = "ResourceValidationException"

	// ErrCodeRevisionDoesNotExistException for service response error code
	// "RevisionDoesNotExistException".
	//
	// The named revision does not exist with the applicable IAM user or AWS account.
	ErrCodeRevisionDoesNotExistException = "RevisionDoesNotExistException"

	// ErrCodeRevisionRequiredException for service response error code
	// "RevisionRequiredException".
	//
	// The revision ID was not specified.
	ErrCodeRevisionRequiredException = "RevisionRequiredException"

	// ErrCodeRoleRequiredException for service response error code
	// "RoleRequiredException".
	//
	// The role ID was not specified.
	ErrCodeRoleRequiredException = "RoleRequiredException"

	// ErrCodeTagLimitExceededException for service response error code
	// "TagLimitExceededException".
	//
	// The maximum allowed number of tags was exceeded.
	ErrCodeTagLimitExceededException = "TagLimitExceededException"

	// ErrCodeTagRequiredException for service response error code
	// "TagRequiredException".
	//
	// A tag was not specified.
	ErrCodeTagRequiredException = "TagRequiredException"

	// ErrCodeTagSetListLimitExceededException for service response error code
	// "TagSetListLimitExceededException".
	//
	// The number of tag groups included in the tag set list exceeded the maximum
	// allowed limit of 3.
	ErrCodeTagSetListLimitExceededException = "TagSetListLimitExceededException"

	// ErrCodeTriggerTargetsLimitExceededException for service response error code
	// "TriggerTargetsLimitExceededException".
	//
	// The maximum allowed number of triggers was exceeded.
	ErrCodeTriggerTargetsLimitExceededException = "TriggerTargetsLimitExceededException"

	// ErrCodeUnsupportedActionForDeploymentTypeException for service response error code
	// "UnsupportedActionForDeploymentTypeException".
	//
	// A call was submitted that is not supported for the specified deployment type.
	ErrCodeUnsupportedActionForDeploymentTypeException = "UnsupportedActionForDeploymentTypeException"
)
//Added a line for testing
//Adding another line for Git event testing part 2
