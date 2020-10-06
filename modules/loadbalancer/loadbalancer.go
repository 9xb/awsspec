package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
)

// ELBV2Spec contains the AWS session
type ELBV2Spec struct {
	Session *session.Session
}

var getELBV2API = func(sess *session.Session) (client elbv2iface.ELBV2API) {
	return elbv2.New(sess)
}

// New returns a new ELBV2Spec
func New(s *session.Session) ELBV2Spec {
	return ELBV2Spec{
		Session: s,
	}
}
