// +build integration

//Package route53domains provides gucumber integration tests support.
package route53domains

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@route53domains", func() {
		gucumber.World["client"] = route53domains.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
