package awsspec

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

// IAMSpec contains the AWS session
type (
	IAMSpec struct {
		Session *session.Session
	}

	// PolicyDocument represents an IAM policy document
	PolicyDocument struct {
		Version   string
		ID        string
		Statement []Statement
	}

	// Statement represents an IAM statement
	Statement struct {
		// TODO:
		// - Handle Principal, NotPrincipal, and Condition
		SID          string
		Principal    interface{}
		NotPrincipal interface{}
		Effect       string
		Action       *OptSlice
		NotAction    *OptSlice
		Resource     *OptSlice
		NotResource  *OptSlice
		Condition    interface{}
	}

	// OptSlice is an entity that could be either a JSON string or a slice
	// As per https://stackoverflow.com/a/38757780/543423
	OptSlice []string
)

// MarshalJSON returns o as the JSON encoding of o
func (o *OptSlice) MarshalJSON() ([]byte, error) {
	// Use normal json.Marshal for subtypes
	if len(*o) == 1 {
		return json.Marshal(([]string)(*o)[0])
	}
	return json.Marshal(*o)
}

// UnmarshalJSON sets *o to a copy of data
func (o *OptSlice) UnmarshalJSON(data []byte) error {
	// Use normal json.Unmarshal for subtypes
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		var v []string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		*o = v
		return nil
	}
	*o = []string{s}
	return nil
}

// Contains checks whether OptSlice contains the provided items slice
func (o OptSlice) Contains(items []string) (res bool) {
	if len(items) > len(o) {
		return false
	}

	for _, e := range items {
		if !find(o, e) {
			return false
		}
	}

	return true
}

var getIAMAPI = func(sess *session.Session) (client iamiface.IAMAPI) {
	return iam.New(sess)
}

// New returns a new IAMSpec
func New(s *session.Session) IAMSpec {
	return IAMSpec{
		Session: s,
	}
}

// PolicyAllows returns true if the defined actions are allowed on the provided resources.
// Please note that the check will be performed on the default policy version.
func (i IAMSpec) PolicyAllows(arn string, actions, resources []string) (res bool, err error) {
	svc := getIAMAPI(i.Session)
	ver, err := i.findDefaultPolicyVersion(arn)
	if err != nil {
		return
	}

	in := &iam.GetPolicyVersionInput{
		PolicyArn: aws.String(arn),
		VersionId: aws.String(ver),
	}
	out, err := svc.GetPolicyVersion(in)
	if err != nil {
		return
	}

	doc := PolicyDocument{}
	policy, err := url.QueryUnescape(aws.StringValue(out.PolicyVersion.Document))
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(policy), &doc)
	for _, v := range doc.Statement {
		hasActions := v.Action.Contains(actions)
		hasResources := v.Resource.Contains(resources)
		res = hasActions && hasResources
		if res {
			return
		}
	}

	return
}

// findDefaultPolicyVersion returns the version ID of default IAM policy for the specified ARN
func (i IAMSpec) findDefaultPolicyVersion(arn string) (ver string, err error) {
	svc := getIAMAPI(i.Session)
	in := &iam.ListPolicyVersionsInput{
		PolicyArn: aws.String(arn),
	}
	out, err := svc.ListPolicyVersions(in)
	if err != nil {
		return
	}

	for _, v := range out.Versions {
		if aws.BoolValue(v.IsDefaultVersion) {
			ver = aws.StringValue(v.VersionId)
			return
		}
	}

	err = errors.New("Could not find version")
	return
}

// find takes a slice and looks for an element in it.
func find(slice []string, val string) (res bool) {
	for _, item := range slice {
		if item == val {
			res = true
			return
		}
	}
	return
}
