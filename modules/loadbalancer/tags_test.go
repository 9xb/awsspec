package awsspec

import (
	"testing"

	s3Spec "github.com/9xb/awsspec/modules/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/stretchr/testify/assert"
)

func TestResourceHasTag(t *testing.T) {
	sess, _ := session.NewSession()
	getELBV2API = func(sess *session.Session) (client elbv2iface.ELBV2API) {
		return mockELBV2API{}
	}

	e := New(sess)
	tag := s3Spec.Tag{
		Key:   aws.StringValue(tags[0].Key),
		Value: aws.StringValue(tags[0].Value),
	}

	res, err := e.ResourceHasTag(resourceARN, tag)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = e.ResourceHasTag("nope", tag)
	assert.Nil(t, err)
	assert.False(t, res)
}
