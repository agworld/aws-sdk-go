package simpledb

import "github.com/aws/aws-sdk-go/aws/client"

func init() {
	initClient = func(c *client.Client) {
		// SimpleDB uses custom error unmarshaling logic
		c.Handlers.UnmarshalError.Clear()
		c.Handlers.UnmarshalError.PushBack(unmarshalError)
	}
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
