// Package endpoints contains the models for endpoints that should be used
// to generate endpoint definition files for the SDK.
package endpoints

//go:generate go run -tags codegen ../../private/model/cli/gen-endpoints/main.go -model ./endpoints.json -out ../../aws/endpoints/defaults.go
//go:generate gofmt -s -w ../../aws/endpoints
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
