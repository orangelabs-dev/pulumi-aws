// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package inspector

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The AWS Inspector Rules Packages data source allows access to the list of AWS
// Inspector Rules Packages which can be used by AWS Inspector within the region
// configured in the provider.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/inspector_rules_packages.html.markdown.
func GetRulesPackages(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetRulesPackagesResult, error) {
	var rv GetRulesPackagesResult
	err := ctx.Invoke("aws:inspector/getRulesPackages:getRulesPackages", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getRulesPackages.
type GetRulesPackagesResult struct {
	// A list of the AWS Inspector Rules Packages arns available in the AWS region.
	Arns []string `pulumi:"arns"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
