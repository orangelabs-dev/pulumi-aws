// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an AWS Config Configuration Aggregator
type ConfigurationAggregator struct {
	pulumi.CustomResourceState

	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource ConfigurationAggregatorAccountAggregationSourcePtrOutput `pulumi:"accountAggregationSource"`
	// The ARN of the aggregator
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the configuration aggregator.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrOutput `pulumi:"organizationAggregationSource"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewConfigurationAggregator registers a new resource with the given unique name, arguments, and options.
func NewConfigurationAggregator(ctx *pulumi.Context,
	name string, args *ConfigurationAggregatorArgs, opts ...pulumi.ResourceOption) (*ConfigurationAggregator, error) {
	if args == nil {
		args = &ConfigurationAggregatorArgs{}
	}
	var resource ConfigurationAggregator
	err := ctx.RegisterResource("aws:cfg/configurationAggregator:ConfigurationAggregator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationAggregator gets an existing ConfigurationAggregator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationAggregator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAggregatorState, opts ...pulumi.ResourceOption) (*ConfigurationAggregator, error) {
	var resource ConfigurationAggregator
	err := ctx.ReadResource("aws:cfg/configurationAggregator:ConfigurationAggregator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationAggregator resources.
type configurationAggregatorState struct {
	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource *ConfigurationAggregatorAccountAggregationSource `pulumi:"accountAggregationSource"`
	// The ARN of the aggregator
	Arn *string `pulumi:"arn"`
	// The name of the configuration aggregator.
	Name *string `pulumi:"name"`
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource *ConfigurationAggregatorOrganizationAggregationSource `pulumi:"organizationAggregationSource"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ConfigurationAggregatorState struct {
	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource ConfigurationAggregatorAccountAggregationSourcePtrInput
	// The ARN of the aggregator
	Arn pulumi.StringPtrInput
	// The name of the configuration aggregator.
	Name pulumi.StringPtrInput
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ConfigurationAggregatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAggregatorState)(nil)).Elem()
}

type configurationAggregatorArgs struct {
	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource *ConfigurationAggregatorAccountAggregationSource `pulumi:"accountAggregationSource"`
	// The name of the configuration aggregator.
	Name *string `pulumi:"name"`
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource *ConfigurationAggregatorOrganizationAggregationSource `pulumi:"organizationAggregationSource"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationAggregator resource.
type ConfigurationAggregatorArgs struct {
	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource ConfigurationAggregatorAccountAggregationSourcePtrInput
	// The name of the configuration aggregator.
	Name pulumi.StringPtrInput
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ConfigurationAggregatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAggregatorArgs)(nil)).Elem()
}
