// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package support provides the client and types for making API
// requests to AWS Support.
//
// The AWS Support API reference is intended for programmers who need detailed
// information about the AWS Support operations and data types. This service
// enables you to manage your AWS Support cases programmatically. It uses HTTP
// methods that return results in JSON format.
//
// The AWS Support service also exposes a set of Trusted Advisor (http://aws.amazon.com/premiumsupport/trustedadvisor/)
// features. You can retrieve a list of checks and their descriptions, get check
// results, specify checks to refresh, and get the refresh status of checks.
//
// The following list describes the AWS Support case management operations:
//
//    * Service names, issue categories, and available severity levels. The
//    DescribeServices and DescribeSeverityLevels operations return AWS service
//    names, service codes, service categories, and problem severity levels.
//    You use these values when you call the CreateCase operation.
//
//    * Case creation, case details, and case resolution. The CreateCase, DescribeCases,
//    DescribeAttachment, and ResolveCase operations create AWS Support cases,
//    retrieve information about cases, and resolve cases.
//
//    * Case communication. The DescribeCommunications, AddCommunicationToCase,
//    and AddAttachmentsToSet operations retrieve and add communications and
//    attachments to AWS Support cases.
//
// The following list describes the operations available from the AWS Support
// service for Trusted Advisor:
//
//    * DescribeTrustedAdvisorChecks returns the list of checks that run against
//    your AWS resources.
//
//    * Using the checkId for a specific check returned by DescribeTrustedAdvisorChecks,
//    you can call DescribeTrustedAdvisorCheckResult to obtain the results for
//    the check you specified.
//
//    * DescribeTrustedAdvisorCheckSummaries returns summarized results for
//    one or more Trusted Advisor checks.
//
//    * RefreshTrustedAdvisorCheck requests that Trusted Advisor rerun a specified
//    check.
//
//    * DescribeTrustedAdvisorCheckRefreshStatuses reports the refresh status
//    of one or more checks.
//
// For authentication of requests, AWS Support uses Signature Version 4 Signing
// Process (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
// See About the AWS Support API (http://docs.aws.amazon.com/awssupport/latest/user/Welcome.html)
// in the AWS Support User Guide for information about how to use this service
// to create and manage your support cases, and how to call Trusted Advisor
// for results of checks on your resources.
//
// See https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15 for more information on this service.
//
// See support package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/support/
//
// Using the Client
//
// To AWS Support with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Support client Support for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/support/#New
package support
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
