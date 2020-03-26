// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The IAM Account Alias data source allows access to the account alias
// for the effective account in which this provider is working.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_account_alias.html.markdown.
func LookupAccountAlias(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupAccountAliasResult, error) {
	var rv LookupAccountAliasResult
	err := ctx.Invoke("aws:iam/getAccountAlias:getAccountAlias", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAccountAlias.
type LookupAccountAliasResult struct {
	// The alias associated with the AWS account.
	AccountAlias string `pulumi:"accountAlias"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
