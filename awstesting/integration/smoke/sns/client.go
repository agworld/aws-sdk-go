// +build integration

//Package sns provides gucumber integration tests support.
package sns

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sns", func() {
		gucumber.World["client"] = sns.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
