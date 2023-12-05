package util

import (
	"testing"

	check "gopkg.in/check.v1"
)

type S struct{}

var _ = check.Suite(&S{})

func Test(t *testing.T) {
	check.TestingT(t)
}

func (s *S) TestValidateName(c *check.C) {
	var data = []struct {
		input    string
		expected bool
	}{
		{"myappmyappmyappmpmyappmyappmyappmyappmyapp", false},
		{"myappmyappmyappmpmyappmyappmyappmyappmyap", false},
		{"myappmyappmyappmpmyappmyappmyappmyappmya", true},
		{"myApp", false},
		{"my app", false},
		{"123myapp", false},
		{"myapp", true},
		{"_theirapp", false},
		{"my-app", true},
		{"-myapp", false},
		{"my_app", false},
		{"b", true},
	}
	for _, d := range data {
		c.Check(ValidateName(d.input), check.Equals, d.expected)
	}
}