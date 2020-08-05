package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

func TestObjectExists(t *testing.T) {
	sess, _ := session.NewSession()
	getS3API = func(s *session.Session) s3iface.S3API {
		return mockS3Client{}
	}

	s := New(sess)
	res, err := s.ObjectExists(bucket, key)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = s.ObjectExists(bucket, key+"s")
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = s.ObjectExists(bucket+"s", key)
	assert.NotNil(t, err)

	_, err = s.ObjectExists(bucket, "nope")
	assert.NotNil(t, err)
}

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
