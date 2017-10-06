package sqs

import "github.com/aws/aws-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
//Added a line for testing
//Adding another line for Git event testing part 2
