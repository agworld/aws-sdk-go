// +build integration

//Package acm provides gucumber integration tests support.
package acm

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@acm", func() {
		gucumber.World["client"] = acm.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
