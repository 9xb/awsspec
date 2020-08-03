package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

func TestObjectHasTag(t *testing.T) {
	sess, _ := session.NewSession()
	getS3API = func(s *session.Session) s3iface.S3API {
		return mockS3Client{}
	}

	s := New(sess)

	tag := Tag{
		Key:   "TestKey",
		Value: "Nope",
	}

	res, err := s.ObjectHasTag("bucket", "key", tag)
	assert.Nil(t, err)
	assert.False(t, res)

	tag = Tag{
		Key:   "TestKey",
		Value: "TestValue",
	}

	res, err = s.ObjectHasTag("bucket", "key", tag)
	assert.Nil(t, err)
	assert.True(t, res)
}
