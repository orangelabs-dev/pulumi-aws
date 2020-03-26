// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get a list of AMI IDs matching the specified criteria.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ami_ids.html.markdown.
func GetAmiIds(ctx *pulumi.Context, args *GetAmiIdsArgs, opts ...pulumi.InvokeOption) (*GetAmiIdsResult, error) {
	var rv GetAmiIdsResult
	err := ctx.Invoke("aws:index/getAmiIds:getAmiIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAmiIds.
type GetAmiIdsArgs struct {
	// Limit search to users with *explicit* launch
	// permission on  the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers []string `pulumi:"executableUsers"`
	// One or more name/value pairs to filter off of. There
	// are several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters []GetAmiIdsFilter `pulumi:"filters"`
	// A regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API.
	// This filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. It is recommended to combine this with other
	// options to narrow down the list AWS returns.
	NameRegex *string `pulumi:"nameRegex"`
	// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
	Owners []string `pulumi:"owners"`
	// Used to sort AMIs by creation time.
	SortAscending *bool `pulumi:"sortAscending"`
}

// A collection of values returned by getAmiIds.
type GetAmiIdsResult struct {
	ExecutableUsers []string          `pulumi:"executableUsers"`
	Filters         []GetAmiIdsFilter `pulumi:"filters"`
	// id is the provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	Ids           []string `pulumi:"ids"`
	NameRegex     *string  `pulumi:"nameRegex"`
	Owners        []string `pulumi:"owners"`
	SortAscending *bool    `pulumi:"sortAscending"`
}
