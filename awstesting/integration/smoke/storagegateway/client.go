// +build integration

//Package storagegateway provides gucumber integration tests support.
package storagegateway

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@storagegateway", func() {
		gucumber.World["client"] = storagegateway.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
