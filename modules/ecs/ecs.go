package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

// ECSSpec contains the AWS session
type ECSSpec struct {
	Session *session.Session
}

var getECSAPI = func(sess *session.Session) (client ecsiface.ECSAPI) {
	return ecs.New(sess)
}

// New returns a new ECSSpec
func New(s *session.Session) ECSSpec {
	return ECSSpec{
		Session: s,
	}
}
