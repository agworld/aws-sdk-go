// +build integration

//Package cloudhsm provides gucumber integration tests support.
package cloudhsm

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudhsm", func() {
		gucumber.World["client"] = cloudhsm.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
