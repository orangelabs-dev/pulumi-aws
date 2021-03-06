// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the ARN of a KMS key alias.
// By using this data source, you can reference key alias
// without having to hard code the ARN as input.
func LookupAlias(ctx *pulumi.Context, args *LookupAliasArgs, opts ...pulumi.InvokeOption) (*LookupAliasResult, error) {
	var rv LookupAliasResult
	err := ctx.Invoke("aws:kms/getAlias:getAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlias.
type LookupAliasArgs struct {
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name string `pulumi:"name"`
}

// A collection of values returned by getAlias.
type LookupAliasResult struct {
	// The Amazon Resource Name(ARN) of the key alias.
	Arn string `pulumi:"arn"`
	// id is the provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// ARN pointed to by the alias.
	TargetKeyArn string `pulumi:"targetKeyArn"`
	// Key identifier pointed to by the alias.
	TargetKeyId string `pulumi:"targetKeyId"`
}
