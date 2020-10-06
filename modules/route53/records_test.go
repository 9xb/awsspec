package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/stretchr/testify/assert"
)

func TestRecordExists(t *testing.T) {
	sess, _ := session.NewSession()
	getRoute53API = func(sess *session.Session) (client route53iface.Route53API) {
		return mockRoute53API{}
	}

	r := New(sess)

	res, err := r.RecordExists(domainName, hostedZone, recordType)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = r.RecordExists("nope", hostedZone, domainName)
	assert.Nil(t, err)
	assert.False(t, res)

}
