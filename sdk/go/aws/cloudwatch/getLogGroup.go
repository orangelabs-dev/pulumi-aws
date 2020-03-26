// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cloudwatch

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about an AWS Cloudwatch Log Group
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudwatch_log_group.html.markdown.
func LookupLogGroup(ctx *pulumi.Context, args *LookupLogGroupArgs, opts ...pulumi.InvokeOption) (*LookupLogGroupResult, error) {
	var rv LookupLogGroupResult
	err := ctx.Invoke("aws:cloudwatch/getLogGroup:getLogGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogGroup.
type LookupLogGroupArgs struct {
	// The name of the Cloudwatch log group
	Name string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getLogGroup.
type LookupLogGroupResult struct {
	// The ARN of the Cloudwatch log group
	Arn string `pulumi:"arn"`
	// The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.
	CreationTime int `pulumi:"creationTime"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ARN of the KMS Key to use when encrypting log data.
	KmsKeyId string `pulumi:"kmsKeyId"`
	Name     string `pulumi:"name"`
	// The number of days log events retained in the specified log group.
	RetentionInDays int `pulumi:"retentionInDays"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}
