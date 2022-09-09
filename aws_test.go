package awsutils

import (
	"testing"

	"github.com/pthomison/errcheck"
)

func TestAWSGetParameter(t *testing.T) {
	param, err := AWSGetParameter("/placeholder", "us-east-2")
	errcheck.CheckTest(err, t)

	if param != "placeholderValue" {
		t.Logf("Unexpected Parameter Value: %s\n", param)
		t.Fail()
	}
}
