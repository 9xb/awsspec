package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

// LambdaSpec contains the AWS session
type LambdaSpec struct {
	Session *session.Session
}

var getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
	return lambda.New(sess)
}

// New returns a new LambdaSpec
func New(s *session.Session) LambdaSpec {
	return LambdaSpec{
		Session: s,
	}
}
