// +build 1.6,codegen

package api_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/private/model/api"
)

func TestShapeTagJoin(t *testing.T) {
	s := api.ShapeTags{
		{Key: "location", Val: "query"},
		{Key: "locationName", Val: "abc"},
		{Key: "type", Val: "string"},
	}

	expected := `location:"query" locationName:"abc" type:"string"`

	o := s.Join(" ")
	o2 := s.String()
	if expected != o {
		t.Errorf("Expected %s, but received %s", expected, o)
	}
	if expected != o2 {
		t.Errorf("Expected %s, but received %s", expected, o2)
	}
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
