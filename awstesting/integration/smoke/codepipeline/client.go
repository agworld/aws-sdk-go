// +build integration

//Package codepipeline provides gucumber integration tests support.
package codepipeline

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@codepipeline", func() {
		gucumber.World["client"] = codepipeline.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
