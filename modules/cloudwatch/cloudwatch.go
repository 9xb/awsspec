package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

// CWSpec contains the AWS session
type CWSpec struct {
	Session *session.Session
}

var getCWAPI = func(sess *session.Session) (client cloudwatchiface.CloudWatchAPI) {
	return cloudwatch.New(sess)
}

// New returns a new CWSpec
func New(s *session.Session) CWSpec {
	return CWSpec{
		Session: s,
	}
}
