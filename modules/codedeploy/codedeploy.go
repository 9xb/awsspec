package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
)

// CodeDeploySpec contains the AWS session
type CodeDeploySpec struct {
	Session *session.Session
}

var getCodeDeployAPI = func(sess *session.Session) (client codedeployiface.CodeDeployAPI) {
	return codedeploy.New(sess)
}

// New returns a new IAMSpec
func New(s *session.Session) CodeDeploySpec {
	return CodeDeploySpec{
		Session: s,
	}
}
