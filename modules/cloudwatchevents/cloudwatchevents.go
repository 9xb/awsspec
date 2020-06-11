package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
)

// CWEventsSpec contains the AWS session
type CWEventsSpec struct {
	Session *session.Session
}

var getCWEventsAPI = func(sess *session.Session) (client cloudwatcheventsiface.CloudWatchEventsAPI) {
	return cloudwatchevents.New(sess)
}

// New returns a new CWEventsSpec
func New(s *session.Session) CWEventsSpec {
	return CWEventsSpec{
		Session: s,
	}
}
