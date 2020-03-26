// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Information about single EC2 Instance Type Offering.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_instance_type_offering.html.markdown.
func GetInstanceTypeOffering(ctx *pulumi.Context, args *GetInstanceTypeOfferingArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypeOfferingResult, error) {
	var rv GetInstanceTypeOfferingResult
	err := ctx.Invoke("aws:ec2/getInstanceTypeOffering:getInstanceTypeOffering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypeOffering.
type GetInstanceTypeOfferingArgs struct {
	// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypeOfferings.html) for supported filters. Detailed below.
	Filters []GetInstanceTypeOfferingFilter `pulumi:"filters"`
	// Location type. Defaults to `region`. Valid values: `availability-zone`, `availability-zone-id`, and `region`.
	LocationType *string `pulumi:"locationType"`
	// Ordered list of preferred EC2 Instance Types. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
	PreferredInstanceTypes []string `pulumi:"preferredInstanceTypes"`
}

// A collection of values returned by getInstanceTypeOffering.
type GetInstanceTypeOfferingResult struct {
	Filters []GetInstanceTypeOfferingFilter `pulumi:"filters"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// EC2 Instance Type.
	InstanceType           string   `pulumi:"instanceType"`
	LocationType           *string  `pulumi:"locationType"`
	PreferredInstanceTypes []string `pulumi:"preferredInstanceTypes"`
}
