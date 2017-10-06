package cognitoidentity

import "github.com/aws/aws-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opGetOpenIdToken, opGetId, opGetCredentialsForIdentity:
			r.Handlers.Sign.Clear() // these operations are unsigned
		}
	}
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
