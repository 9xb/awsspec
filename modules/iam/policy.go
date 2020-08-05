package awsspec

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

type (
	ConditionOperator interface {
		GetOperator() string
		GetVariable() string
		GetValue() interface{}
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
		Condition    map[ConditionType]map[ConditionVariable]OptSlice `json:",omitempty"`
	}

	// OptSlice is an entity that could be either a JSON string or a slice
	// As per https://stackoverflow.com/a/38757780/543423
	OptSlice []string

	// ConditionType represents all the possible comparison types for the
	// Condition of a Policy Statement
	// Inspired by github.com/gwkunze/goiam/policy
	ConditionType string

	// ConditionVariable represent the available variables used in Conditions
	// Inspired by github.com/gwkunze/goiam/policy
	ConditionVariable string
)

const (
	ConditionStringEquals              ConditionType = "StringEquals"
	ConditionStringNotEquals           ConditionType = "StringNotEquals"
	ConditionStringEqualsIgnoreCase    ConditionType = "StringEqualsIgnoreCase"
	ConditionStringNotEqualsIgnoreCase ConditionType = "StringNotEqualsIgnoreCase"
	ConditionStringLike                ConditionType = "StringLike"
	ConditionStringNotLike             ConditionType = "StringNotLike"
	ConditionNumericEquals             ConditionType = "NumericEquals"
	ConditionNumericNotEquals          ConditionType = "NumericNotEquals"
	ConditionNumericLessThan           ConditionType = "NumericLessThan"
	ConditionNumericLessThanEquals     ConditionType = "NumericLessThanEquals"
	ConditionNumericGreaterThan        ConditionType = "NumericGreaterThan"
	ConditionNumericGreaterThanEquals  ConditionType = "NumericGreaterThanEquals"
	ConditionDateEquals                ConditionType = "DateEquals"
	ConditionDateNotEquals             ConditionType = "DateNotEquals"
	ConditionDateLessThan              ConditionType = "DateLessThan"
	ConditionDateLessThanEquals        ConditionType = "DateLessThanEquals"
	ConditionDateGreaterThan           ConditionType = "DateGreaterThan"
	ConditionDateGreaterThanEquals     ConditionType = "DateGreaterThanEquals"
	ConditionBool                      ConditionType = "Bool"
	ConditionIpAddress                 ConditionType = "IpAddress"
	ConditionNotIpAddress              ConditionType = "NotIpAddress"
	ConditionArnEquals                 ConditionType = "ArnEquals"
	ConditionArnNotEquals              ConditionType = "ArnNotEquals"
	ConditionArnLike                   ConditionType = "ArnLike"
	ConditionArnNotLike                ConditionType = "ArnNotLike"
	ConditionNull                      ConditionType = "Null"
)

const (
	VarCurrentTime        ConditionVariable = "AWS:CurrentTime"
	VarEpochTime          ConditionVariable = "AWS:EpochTime"
	VarMultiFactorAuthAge ConditionVariable = "AWS:MultiFactorAuthAge"
	VarPrincipalType      ConditionVariable = "AWS:principaltype"
	VarSecureTransport    ConditionVariable = "AWS:SecureTransport"
	VarSourceArn          ConditionVariable = "AWS:SourceArn"
	VarSourceIp           ConditionVariable = "AWS:SourceIp"
	VarUserAgent          ConditionVariable = "AWS:UserAgent"
	VarUsedId             ConditionVariable = "AWS:userid"
	VarUsername           ConditionVariable = "AWS:username"
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
