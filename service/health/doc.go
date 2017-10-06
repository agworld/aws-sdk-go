// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package health provides the client and types for making API
// requests to AWS Health APIs and Notifications.
//
// The AWS Health API provides programmatic access to the AWS Health information
// that is presented in the AWS Personal Health Dashboard (https://phd.aws.amazon.com/phd/home#/).
// You can get information about events that affect your AWS resources:
//
//    * DescribeEvents: Summary information about events.
//
//    * DescribeEventDetails: Detailed information about one or more events.
//
//    * DescribeAffectedEntities: Information about AWS resources that are affected
//    by one or more events.
//
// In addition, these operations provide information about event types and summary
// counts of events or affected entities:
//
//    * DescribeEventTypes: Information about the kinds of events that AWS Health
//    tracks.
//
//    * DescribeEventAggregates: A count of the number of events that meet specified
//    criteria.
//
//    * DescribeEntityAggregates: A count of the number of affected entities
//    that meet specified criteria.
//
// The Health API requires a Business or Enterprise support plan from AWS Support
// (http://aws.amazon.com/premiumsupport/). Calling the Health API from an account
// that does not have a Business or Enterprise support plan causes a SubscriptionRequiredException.
//
// For authentication of requests, AWS Health uses the Signature Version 4 Signing
// Process (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
// See the AWS Health User Guide (http://docs.aws.amazon.com/health/latest/ug/what-is-aws-health.html)
// for information about how to use the API.
//
// Service Endpoint
//
// The HTTP endpoint for the AWS Health API is:
//
//    * https://health.us-east-1.amazonaws.com
//
// See https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04 for more information on this service.
//
// See health package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/health/
//
// Using the Client
//
// To AWS Health APIs and Notifications with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Health APIs and Notifications client Health for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/health/#New
package health
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
