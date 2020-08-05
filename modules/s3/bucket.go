package awsspec

import (
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

// BucketExists returns true if the specified S3 bucket exists
func (s S3Spec) BucketExists(bucket string) (res bool, err error) {
	svc := getS3API(s.Session)
	in := &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	}

	_, err = svc.HeadBucket(in)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchBucket {
				return false, nil
			}
		}
		return
	}

	return true, nil
}

// BucketHasLoggingEnabled returns true if the specified S3 bucket has logging enabled. Throws an error if the bucket does not exist.
func (s S3Spec) BucketHasLoggingEnabled(bucket string) (res bool, err error) {
	svc := getS3API(s.Session)
	in := &s3.GetBucketLoggingInput{
		Bucket: aws.String(bucket),
	}

	out, err := svc.GetBucketLogging(in)
	if err != nil {
		return
	}

	if out.LoggingEnabled != nil {
		return true, nil
	}

	return
}

// BucketHasVersioningEnabled returns true if the specified S3 bucket has versioning enabled. Throws an error if the bucket does not exist.
func (s S3Spec) BucketHasVersioningEnabled(bucket string) (res bool, err error) {
	svc := getS3API(s.Session)
	in := &s3.GetBucketVersioningInput{
		Bucket: aws.String(bucket),
	}

	out, err := svc.GetBucketVersioning(in)
	if err != nil {
		return
	}

	if aws.StringValue(out.Status) != "Enabled" {
		return
	}

	return true, nil
}

// BucketHasServerSideEncryption returns true if the specified S3 bucket has the provided Server Side Encryption method. Throws an error if the bucket does not exist.
func (s S3Spec) BucketHasServerSideEncryption(bucket, algorithm string) (res bool, err error) {
	svc := getS3API(s.Session)
	in := &s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucket),
	}
	out, err := svc.GetBucketEncryption(in)
	if err != nil {
		if reqerr, ok := err.(awserr.RequestFailure); ok {
			if reqerr.Code() == "ServerSideEncryptionConfigurationNotFoundError" {
				return false, nil
			}
			return
		}
		return
	}

	if out.ServerSideEncryptionConfiguration == nil {
		return
	}

	rules := out.ServerSideEncryptionConfiguration.Rules
	if len(rules) == 0 {
		return
	}

	for _, v := range rules {
		if aws.StringValue(v.ApplyServerSideEncryptionByDefault.SSEAlgorithm) == algorithm {
			return true, nil
		}
	}

	return
}

// BucketHasWebsiteEnabled returns true if the specified S3 bucket is setup as a website. Throws an error if the bucket does not exist.
func (s S3Spec) BucketHasWebsiteEnabled(bucket string) (res bool, err error) {
	svc := getS3API(s.Session)
	in := &s3.GetBucketWebsiteInput{
		Bucket: aws.String(bucket),
	}
	out, err := svc.GetBucketWebsite(in)
	if err != nil {
		if reqerr, ok := err.(awserr.RequestFailure); ok {
			if reqerr.Code() == "NoSuchWebsiteConfiguration" {
				return false, nil
			}
			return
		}
		return
	}

	if out.IndexDocument == nil {
		return
	}

	return true, nil
}

// BucketHasLifecycleRule returns true if the specified S3 bucket has the provided lifecycle rule. Throws an error if the bucket does not exist.
func (s S3Spec) BucketHasLifecycleRule(bucket string, rule s3.LifecycleRule) (res bool, err error) {
	svc := getS3API(s.Session)
	in := &s3.GetBucketLifecycleConfigurationInput{
		Bucket: aws.String(bucket),
	}
	out, err := svc.GetBucketLifecycleConfiguration(in)
	if err != nil {
		return
	}

	rules := out.Rules
	if len(rules) < 1 {
		return
	}

	for _, v := range rules {
		if reflect.DeepEqual(v, &rule) {
			return true, nil
		}
	}

	return
}
