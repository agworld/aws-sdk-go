// +build integration

//Package lambda provides gucumber integration tests support.
package lambda

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@lambda", func() {
		gucumber.World["client"] = lambda.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
