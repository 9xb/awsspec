package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

// IAMSpec contains the AWS session
type IAMSpec struct {
	Session *session.Session
}

var getIAMAPI = func(sess *session.Session) (client iamiface.IAMAPI) {
	return iam.New(sess)
}

// New returns a new IAMSpec
func New(s *session.Session) IAMSpec {
	return IAMSpec{
		Session: s,
	}
}
