// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codestar provides the client and types for making API
// requests to AWS CodeStar.
//
// This is the API reference for AWS CodeStar. This reference provides descriptions
// of the operations and data types for the AWS CodeStar API along with usage
// examples.
//
// You can use the AWS CodeStar API to work with:
//
// Projects and their resources, by calling the following:
//
//    * DeleteProject, which deletes a project.
//
//    * DescribeProject, which lists the attributes of a project.
//
//    * ListProjects, which lists all projects associated with your AWS account.
//
//    * ListResources, which lists the resources associated with a project.
//
//    * ListTagsForProject, which lists the tags associated with a project.
//
//    * TagProject, which adds tags to a project.
//
//    * UntagProject, which removes tags from a project.
//
//    * UpdateProject, which updates the attributes of a project.
//
// Teams and team members, by calling the following:
//
//    * AssociateTeamMember, which adds an IAM user to the team for a project.
//
//    * DisassociateTeamMember, which removes an IAM user from the team for
//    a project.
//
//    * ListTeamMembers, which lists all the IAM users in the team for a project,
//    including their roles and attributes.
//
//    * UpdateTeamMember, which updates a team member's attributes in a project.
//
// Users, by calling the following:
//
//    * CreateUserProfile, which creates a user profile that contains data associated
//    with the user across all projects.
//
//    * DeleteUserProfile, which deletes all user profile information across
//    all projects.
//
//    * DescribeUserProfile, which describes the profile of a user.
//
//    * ListUserProfiles, which lists all user profiles.
//
//    * UpdateUserProfile, which updates the profile for a user.
//
// See https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19 for more information on this service.
//
// See codestar package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codestar/
//
// Using the Client
//
// To AWS CodeStar with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CodeStar client CodeStar for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codestar/#New
package codestar
//Added a line for testing
//Adding another line for Git event testing part 2
