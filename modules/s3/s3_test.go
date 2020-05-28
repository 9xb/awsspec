package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

type mockS3Client struct {
	s3iface.S3API
}

func (m mockS3Client) GetObjectTagging(*s3.GetObjectTaggingInput) (o *s3.GetObjectTaggingOutput, err error) {
	o = &s3.GetObjectTaggingOutput{
		TagSet: []*s3.Tag{
			{
				Key:   aws.String("TestKey"),
				Value: aws.String("TestValue"),
			},
		},
	}
	return
}

func TestS3ObjectHasTag(t *testing.T) {
	sess, _ := session.NewSession()
	getS3Client = func(s *session.Session) s3iface.S3API {
		return mockS3Client{}
	}

	s := New(sess)

	tag := Tag{
		Key:   "TestKey",
		Value: "Nope",
	}

	res, err := s.S3ObjectHasTag("bucket", "key", tag)
	assert.Nil(t, err)
	assert.False(t, res)

	tag = Tag{
		Key:   "TestKey",
		Value: "TestValue",
	}

	res, err = s.S3ObjectHasTag("bucket", "key", tag)
	assert.Nil(t, err)
	assert.True(t, res)
}
