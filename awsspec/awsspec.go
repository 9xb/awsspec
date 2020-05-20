package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWSSpec contains the AWS Session
type AWSSpec struct {
	Session *session.Session
}
