package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type (
	// TagGetter is an interface that wraps the GetKey and GetValue methods
	TagGetter interface {
		GetKey() string
		GetValue() string
	}

	// S3Spec contains the AWS session
	S3Spec struct {
		Session *session.Session
	}

	// Tag represents an S3 object tag
	Tag struct {
		Key   string
		Value string
	}
)

// New creates a new AWSSpec
func New(s *session.Session) S3Spec {
	return S3Spec{
		Session: s,
	}
}

var getS3Client = func(sess *session.Session) (client s3iface.S3API) {
	return s3.New(sess)
}

// GetKey retrieves the Tag Key
func (t Tag) GetKey() string {
	return t.Key
}

// GetValue retrieves the Tag Value
func (t Tag) GetValue() string {
	return t.Value
}

// S3ObjectHasTag verifies that an S3 object has the indicated tag
func (s S3Spec) S3ObjectHasTag(bucket, key string, tag TagGetter) (res bool, err error) {
	svc := getS3Client(s.Session)
	input := &s3.GetObjectTaggingInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	output, err := svc.GetObjectTagging(input)
	if err != nil {
		return
	}

	for _, v := range output.TagSet {
		if aws.StringValue(v.Key) == tag.GetKey() {
			if aws.StringValue(v.Value) != tag.GetValue() {
				return
			}
			res = true
			return
		}
	}
	return
}
