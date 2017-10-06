// +build integration

//Package dynamodb provides gucumber integration tests support.
package dynamodb

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@dynamodb", func() {
		gucumber.World["client"] = dynamodb.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
