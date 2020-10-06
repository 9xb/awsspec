package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
)

// Route53Spec contains the AWS session
type Route53Spec struct {
	Session *session.Session
}

var getRoute53API = func(sess *session.Session) (client route53iface.Route53API) {
	return route53.New(sess)
}

// New returns a new Route53Spec
func New(s *session.Session) Route53Spec {
	return Route53Spec{
		Session: s,
	}
}
