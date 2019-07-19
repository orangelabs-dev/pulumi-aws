// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicequotas

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an individual Service Quota.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/servicequotas_service_quota.html.markdown.
type ServiceQuota struct {
	s *pulumi.ResourceState
}

// NewServiceQuota registers a new resource with the given unique name, arguments, and options.
func NewServiceQuota(ctx *pulumi.Context,
	name string, args *ServiceQuotaArgs, opts ...pulumi.ResourceOpt) (*ServiceQuota, error) {
	if args == nil || args.QuotaCode == nil {
		return nil, errors.New("missing required argument 'QuotaCode'")
	}
	if args == nil || args.ServiceCode == nil {
		return nil, errors.New("missing required argument 'ServiceCode'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["quotaCode"] = nil
		inputs["serviceCode"] = nil
		inputs["value"] = nil
	} else {
		inputs["quotaCode"] = args.QuotaCode
		inputs["serviceCode"] = args.ServiceCode
		inputs["value"] = args.Value
	}
	inputs["adjustable"] = nil
	inputs["arn"] = nil
	inputs["defaultValue"] = nil
	inputs["quotaName"] = nil
	inputs["requestId"] = nil
	inputs["requestStatus"] = nil
	inputs["serviceName"] = nil
	s, err := ctx.RegisterResource("aws:servicequotas/serviceQuota:ServiceQuota", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServiceQuota{s: s}, nil
}

// GetServiceQuota gets an existing ServiceQuota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceQuota(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServiceQuotaState, opts ...pulumi.ResourceOpt) (*ServiceQuota, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adjustable"] = state.Adjustable
		inputs["arn"] = state.Arn
		inputs["defaultValue"] = state.DefaultValue
		inputs["quotaCode"] = state.QuotaCode
		inputs["quotaName"] = state.QuotaName
		inputs["requestId"] = state.RequestId
		inputs["requestStatus"] = state.RequestStatus
		inputs["serviceCode"] = state.ServiceCode
		inputs["serviceName"] = state.ServiceName
		inputs["value"] = state.Value
	}
	s, err := ctx.ReadResource("aws:servicequotas/serviceQuota:ServiceQuota", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServiceQuota{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ServiceQuota) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ServiceQuota) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Whether the service quota can be increased.
func (r *ServiceQuota) Adjustable() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["adjustable"])
}

// Amazon Resource Name (ARN) of the service quota.
func (r *ServiceQuota) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Default value of the service quota.
func (r *ServiceQuota) DefaultValue() *pulumi.Float64Output {
	return (*pulumi.Float64Output)(r.s.State["defaultValue"])
}

// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
func (r *ServiceQuota) QuotaCode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["quotaCode"])
}

// Name of the quota.
func (r *ServiceQuota) QuotaName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["quotaName"])
}

func (r *ServiceQuota) RequestId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["requestId"])
}

func (r *ServiceQuota) RequestStatus() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["requestStatus"])
}

// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
func (r *ServiceQuota) ServiceCode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceCode"])
}

// Name of the service.
func (r *ServiceQuota) ServiceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceName"])
}

// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
func (r *ServiceQuota) Value() *pulumi.Float64Output {
	return (*pulumi.Float64Output)(r.s.State["value"])
}

// Input properties used for looking up and filtering ServiceQuota resources.
type ServiceQuotaState struct {
	// Whether the service quota can be increased.
	Adjustable interface{}
	// Amazon Resource Name (ARN) of the service quota.
	Arn interface{}
	// Default value of the service quota.
	DefaultValue interface{}
	// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode interface{}
	// Name of the quota.
	QuotaName interface{}
	RequestId interface{}
	RequestStatus interface{}
	// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode interface{}
	// Name of the service.
	ServiceName interface{}
	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value interface{}
}

// The set of arguments for constructing a ServiceQuota resource.
type ServiceQuotaArgs struct {
	// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode interface{}
	// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode interface{}
	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value interface{}
}
