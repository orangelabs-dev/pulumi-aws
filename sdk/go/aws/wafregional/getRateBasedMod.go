// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `wafregional.RateBasedRule` Retrieves a WAF Regional Rate Based Rule Resource Id.
func GetRateBasedMod(ctx *pulumi.Context, args *GetRateBasedModArgs, opts ...pulumi.InvokeOption) (*GetRateBasedModResult, error) {
	var rv GetRateBasedModResult
	err := ctx.Invoke("aws:wafregional/getRateBasedMod:getRateBasedMod", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRateBasedMod.
type GetRateBasedModArgs struct {
	// The name of the WAF Regional rate based rule.
	Name string `pulumi:"name"`
}

// A collection of values returned by getRateBasedMod.
type GetRateBasedModResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
