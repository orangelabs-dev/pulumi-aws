// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The Canonical User ID data source allows access to the [canonical user ID](http://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html)
// for the effective account in which this provider is working.
func GetCanonicalUserId(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCanonicalUserIdResult, error) {
	var rv GetCanonicalUserIdResult
	err := ctx.Invoke("aws:index/getCanonicalUserId:getCanonicalUserId", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCanonicalUserId.
type GetCanonicalUserIdResult struct {
	// The human-friendly name linked to the canonical user ID. The bucket owner's display name. **NOTE:** [This value](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTServiceGET.html) is only included in the response in the US East (N. Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo), EU (Ireland), and South America (São Paulo) regions.
	DisplayName string `pulumi:"displayName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
