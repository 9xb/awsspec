package awsspec

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

var (
	bucket          = "testBucket"
	key             = "testObject"
	algorithm       = "testAlgorithm"
	lifecyclePrefix = "testPrefix"
)

type mockS3Client struct {
	s3iface.S3API
}

func (m mockS3Client) HeadObject(input *s3.HeadObjectInput) (o *s3.HeadObjectOutput, err error) {
	if aws.StringValue(input.Bucket) != bucket {
		err = awserr.New(s3.ErrCodeNoSuchBucket, "", err)
		return
	}

	if aws.StringValue(input.Key) == "nope" {
		err = errors.New("nope")
		return
	}

	if aws.StringValue(input.Key) != key {
		err = awserr.New(s3.ErrCodeNoSuchKey, "", err)
		return
	}

	return &s3.HeadObjectOutput{ETag: aws.String("test")}, nil
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

func (m mockS3Client) HeadBucket(input *s3.HeadBucketInput) (o *s3.HeadBucketOutput, err error) {
	if aws.StringValue(input.Bucket) == "nope" {
		return o, errors.New("nope")
	}

	if aws.StringValue(input.Bucket) != bucket {
		return o, awserr.New(s3.ErrCodeNoSuchBucket, "", err)
	}

	return
}

func (m mockS3Client) GetBucketLogging(input *s3.GetBucketLoggingInput) (o *s3.GetBucketLoggingOutput, err error) {
	if aws.StringValue(input.Bucket) == bucket {
		o = &s3.GetBucketLoggingOutput{
			LoggingEnabled: &s3.LoggingEnabled{
				TargetBucket: aws.String("test"),
			},
		}
		return
	}

	if aws.StringValue(input.Bucket) == "nope" {
		return &s3.GetBucketLoggingOutput{}, nil
	}

	return o, errors.New("nope")
}

func (m mockS3Client) GetBucketVersioning(input *s3.GetBucketVersioningInput) (o *s3.GetBucketVersioningOutput, err error) {
	if aws.StringValue(input.Bucket) == bucket {
		o = &s3.GetBucketVersioningOutput{
			Status: aws.String("Enabled"),
		}
		return
	}

	if aws.StringValue(input.Bucket) == "nope" {
		return &s3.GetBucketVersioningOutput{}, nil
	}

	return o, errors.New("nope")
}

func (m mockS3Client) GetBucketEncryption(input *s3.GetBucketEncryptionInput) (o *s3.GetBucketEncryptionOutput, err error) {
	if aws.StringValue(input.Bucket) == bucket {
		o = &s3.GetBucketEncryptionOutput{
			ServerSideEncryptionConfiguration: &s3.ServerSideEncryptionConfiguration{
				Rules: []*s3.ServerSideEncryptionRule{
					{
						ApplyServerSideEncryptionByDefault: &s3.ServerSideEncryptionByDefault{
							SSEAlgorithm: aws.String(algorithm),
						},
					},
				},
			},
		}
		return
	}

	if aws.StringValue(input.Bucket) == "nope" {
		err = awserr.NewRequestFailure(
			awserr.New("ServerSideEncryptionConfigurationNotFoundError", "The specified bucket does not have a website configuration", nil), 1, "test")
		return
	}

	return o, errors.New("nope")
}

func (m mockS3Client) GetBucketWebsite(input *s3.GetBucketWebsiteInput) (o *s3.GetBucketWebsiteOutput, err error) {
	if aws.StringValue(input.Bucket) == bucket {
		o = &s3.GetBucketWebsiteOutput{
			ErrorDocument: &s3.ErrorDocument{Key: aws.String("test")},
			IndexDocument: &s3.IndexDocument{Suffix: aws.String("test")},
		}
		return
	}

	if aws.StringValue(input.Bucket) == "nope" {
		err = awserr.NewRequestFailure(
			awserr.New("NoSuchWebsiteConfiguration", "The specified bucket does not have a website configuration", nil), 1, "test")
		return
	}

	return o, errors.New("nope")
}

func (m mockS3Client) GetBucketLifecycleConfiguration(input *s3.GetBucketLifecycleConfigurationInput) (o *s3.GetBucketLifecycleConfigurationOutput, err error) {
	if aws.StringValue(input.Bucket) == bucket {
		o = &s3.GetBucketLifecycleConfigurationOutput{
			Rules: []*s3.LifecycleRule{
				{
					Status: aws.String("Enabled"),
					Prefix: aws.String(lifecyclePrefix),
				},
			},
		}
		return
	}

	if aws.StringValue(input.Bucket) == "nope" {
		return &s3.GetBucketLifecycleConfigurationOutput{}, nil
	}

	return o, errors.New("nope")
}
