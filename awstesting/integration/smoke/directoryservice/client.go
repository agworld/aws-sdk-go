// +build integration

//Package directoryservice provides gucumber integration tests support.
package directoryservice

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@directoryservice", func() {
		gucumber.World["client"] = directoryservice.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
