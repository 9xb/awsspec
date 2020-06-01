package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/stretchr/testify/assert"
)

func TestUserHasPolicy(t *testing.T) {
	sess, _ := session.NewSession()
	getIAMAPI = func(s *session.Session) iamiface.IAMAPI {
		return mockIAMAPI{}
	}
	user := "testuser"

	i := New(sess)

	res, err := i.UserHasPolicy(user, policyPrefix+userPolicies[0])
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = i.UserHasPolicy(user, policyPrefix+"Nope")
	assert.Nil(t, err)
	assert.False(t, res)
}
