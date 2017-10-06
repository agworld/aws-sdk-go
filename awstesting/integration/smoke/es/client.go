// +build integration

//Package es provides gucumber integration tests support.
package es

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@es", func() {
		gucumber.World["client"] = elasticsearchservice.New(smoke.Session)
	})
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
