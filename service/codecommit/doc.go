// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codecommit provides the client and types for making API
// requests to AWS CodeCommit.
//
// This is the AWS CodeCommit API Reference. This reference provides descriptions
// of the operations and data types for AWS CodeCommit API along with usage
// examples.
//
// You can use the AWS CodeCommit API to work with the following objects:
//
// Repositories, by calling the following:
//
//    * BatchGetRepositories, which returns information about one or more repositories
//    associated with your AWS account
//
//    * CreateRepository, which creates an AWS CodeCommit repository
//
//    * DeleteRepository, which deletes an AWS CodeCommit repository
//
//    * GetRepository, which returns information about a specified repository
//
//    * ListRepositories, which lists all AWS CodeCommit repositories associated
//    with your AWS account
//
//    * UpdateRepositoryDescription, which sets or updates the description of
//    the repository
//
//    * UpdateRepositoryName, which changes the name of the repository. If you
//    change the name of a repository, no other users of that repository will
//    be able to access it until you send them the new HTTPS or SSH URL to use.
//
// Branches, by calling the following:
//
//    * CreateBranch, which creates a new branch in a specified repository
//
//    * GetBranch, which returns information about a specified branch
//
//    * ListBranches, which lists all branches for a specified repository
//
//    * UpdateDefaultBranch, which changes the default branch for a repository
//
// Information about committed code in a repository, by calling the following:
//
//    * GetBlob, which returns the base-64 encoded content of an individual
//    Git blob object within a repository
//
//    * GetCommit, which returns information about a commit, including commit
//    messages and author and committer information
//
//    * GetDifferences, which returns information about the differences in a
//    valid commit specifier (such as a branch, tag, HEAD, commit ID or other
//    fully qualified reference)
//
// Triggers, by calling the following:
//
//    * GetRepositoryTriggers, which returns information about triggers configured
//    for a repository
//
//    * PutRepositoryTriggers, which replaces all triggers for a repository
//    and can be used to create or delete triggers
//
//    * TestRepositoryTriggers, which tests the functionality of a repository
//    trigger by sending data to the trigger target
//
// For information about how to use AWS CodeCommit, see the AWS CodeCommit User
// Guide (http://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13 for more information on this service.
//
// See codecommit package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codecommit/
//
// Using the Client
//
// To AWS CodeCommit with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CodeCommit client CodeCommit for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codecommit/#New
package codecommit
//Added a line for testing
