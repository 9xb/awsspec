package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

func setupTestEnv() S3Spec {
	sess, _ := session.NewSession()
	getS3API = func(s *session.Session) s3iface.S3API {
		return mockS3Client{}
	}

	return New(sess)
}

func TestBucketExists(t *testing.T) {
	s := setupTestEnv()
	res, err := s.BucketExists(bucket)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = s.BucketExists(bucket + "s")
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = s.BucketExists("nope")
	assert.NotNil(t, err)
}

func TestBucketHasLoggingEnabled(t *testing.T) {
	s := setupTestEnv()
	res, err := s.BucketHasLoggingEnabled(bucket)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = s.BucketHasLoggingEnabled("nope")
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = s.BucketHasLoggingEnabled(bucket + "s")
	assert.NotNil(t, err)
}

func TestBucketHasVersioningEnabled(t *testing.T) {
	s := setupTestEnv()
	res, err := s.BucketHasVersioningEnabled(bucket)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = s.BucketHasVersioningEnabled("nope")
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = s.BucketHasVersioningEnabled(bucket + "s")
	assert.NotNil(t, err)
}

func TestBucketHasServerSideEncryption(t *testing.T) {
	s := setupTestEnv()
	res, err := s.BucketHasServerSideEncryption(bucket, algorithm)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = s.BucketHasServerSideEncryption(bucket, algorithm+"s")
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = s.BucketHasServerSideEncryption("nope", algorithm)
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = s.BucketHasServerSideEncryption(bucket+"s", algorithm)
	assert.NotNil(t, err)
}

func TestBucketHasWebsiteEnabled(t *testing.T) {
	s := setupTestEnv()
	res, err := s.BucketHasWebsiteEnabled(bucket)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = s.BucketHasWebsiteEnabled("nope")
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = s.BucketHasWebsiteEnabled(bucket + "s")
	assert.NotNil(t, err)
}

func TestBucketHasLifecycleRule(t *testing.T) {
	s := setupTestEnv()
	rule := s3.LifecycleRule{
		Status: aws.String("Enabled"),
		Prefix: aws.String(lifecyclePrefix),
	}

	res, err := s.BucketHasLifecycleRule(bucket, rule)
	assert.Nil(t, err)
	assert.True(t, res)

	rule.Status = aws.String("Disabled")
	res, err = s.BucketHasLifecycleRule(bucket, rule)
	assert.Nil(t, err)
	assert.False(t, res)

	rule = s3.LifecycleRule{
		Status: aws.String("Enabled"),
		Prefix: aws.String(lifecyclePrefix + "s"),
	}
	res, err = s.BucketHasLifecycleRule(bucket, rule)
	assert.Nil(t, err)
	assert.False(t, res)

	rule = s3.LifecycleRule{
		Status: aws.String("Enabled"),
		Prefix: aws.String(lifecyclePrefix),
	}
	res, err = s.BucketHasLifecycleRule("nope", rule)
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = s.BucketHasLifecycleRule(bucket+"s", rule)
	assert.NotNil(t, err)
}
