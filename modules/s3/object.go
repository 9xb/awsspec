package awsspec

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	// Tag represents an AWS tag
	Tag struct {
		Key   string
		Value string
	}
)

// GetKey retrieves the Tag Key
func (t Tag) GetKey() string {
	return t.Key
}

// GetValue retrieves the Tag Value
func (t Tag) GetValue() string {
	return t.Value
}

// ObjectExists returns true if the provided key exists in the specified bucket
func (s S3Spec) ObjectExists(bucket, key string) (res bool, err error) {
	svc := getS3API(s.Session)
	in := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err = svc.HeadObject(in)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				err = errors.New("bucket " + bucket + " does not exist")
				return
			case s3.ErrCodeNoSuchKey:
				return false, nil
			}
		}
		return
	}

	return true, nil
}

// ObjectHasTag verifies that an S3 object has the indicated tag
func (s S3Spec) ObjectHasTag(bucket, key string, tag TagGetter) (res bool, err error) {
	svc := getS3API(s.Session)
	in := &s3.GetObjectTaggingInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	out, err := svc.GetObjectTagging(in)
	if err != nil {
		return
	}

	for _, v := range out.TagSet {
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
