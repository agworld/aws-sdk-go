// +build integration

//Package kinesis provides gucumber integration tests support.
package kinesis

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@kinesis", func() {
		gucumber.World["client"] = kinesis.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
