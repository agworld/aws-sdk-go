// +build integration

//Package sts provides gucumber integration tests support.
package sts

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sts", func() {
		gucumber.World["client"] = sts.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
