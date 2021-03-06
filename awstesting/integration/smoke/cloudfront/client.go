// +build integration

//Package cloudfront provides gucumber integration tests support.
package cloudfront

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudfront", func() {
		gucumber.World["client"] = cloudfront.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
