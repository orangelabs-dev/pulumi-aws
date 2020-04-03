// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicequotas

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Retrieve information about a Service Quota.
func LookupServiceQuota(ctx *pulumi.Context, args *LookupServiceQuotaArgs, opts ...pulumi.InvokeOption) (*LookupServiceQuotaResult, error) {
	var rv LookupServiceQuotaResult
	err := ctx.Invoke("aws:servicequotas/getServiceQuota:getServiceQuota", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceQuota.
type LookupServiceQuotaArgs struct {
	// Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode *string `pulumi:"quotaCode"`
	// Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaName *string `pulumi:"quotaName"`
	// Service code for the quota. Available values can be found with the [`servicequotas.getService` data source](https://www.terraform.io/docs/providers/aws/d/servicequotas_service.html) or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode string `pulumi:"serviceCode"`
}

// A collection of values returned by getServiceQuota.
type LookupServiceQuotaResult struct {
	// Whether the service quota is adjustable.
	Adjustable bool `pulumi:"adjustable"`
	// Amazon Resource Name (ARN) of the service quota.
	Arn string `pulumi:"arn"`
	// Default value of the service quota.
	DefaultValue float64 `pulumi:"defaultValue"`
	// Whether the service quota is global for the AWS account.
	GlobalQuota bool `pulumi:"globalQuota"`
	// id is the provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	QuotaCode   string `pulumi:"quotaCode"`
	QuotaName   string `pulumi:"quotaName"`
	ServiceCode string `pulumi:"serviceCode"`
	// Name of the service.
	ServiceName string `pulumi:"serviceName"`
	// Current value of the service quota.
	Value float64 `pulumi:"value"`
}
