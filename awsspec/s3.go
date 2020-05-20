package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type (
	TagGetter interface {
		GetKey() string
		GetValue() string
	}

	Tag struct {
		Key   string
		Value string
	}
)

var getS3Client = func(sess *session.Session) (client s3iface.S3API) {
	return s3.New(sess)
}

func (t Tag) GetKey() string {
	return t.Key
}

func (t Tag) GetValue() string {
	return t.Value
}

// S3ObjectHasTag verifies that an S3 object has the indicated tag
func (a AWSSpec) S3ObjectHasTag(bucket, key string, tag TagGetter) (res bool, err error) {
	svc := getS3Client(a.Session)
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
