package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// New creates a new AWSSpec
func New(s *session.Session) S3Spec {
	return S3Spec{
		Session: s,
	}
}

var getS3API = func(sess *session.Session) (client s3iface.S3API) {
	return s3.New(sess)
}
