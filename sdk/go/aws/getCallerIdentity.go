// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the access to the effective Account ID, User ID, and ARN in
// which this provider is authorized.
func GetCallerIdentity(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCallerIdentityResult, error) {
	var rv GetCallerIdentityResult
	err := ctx.Invoke("aws:index/getCallerIdentity:getCallerIdentity", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCallerIdentity.
type GetCallerIdentityResult struct {
	// The AWS Account ID number of the account that owns or contains the calling entity.
	AccountId string `pulumi:"accountId"`
	// The AWS ARN associated with the calling entity.
	Arn string `pulumi:"arn"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The unique identifier of the calling entity.
	UserId string `pulumi:"userId"`
}
