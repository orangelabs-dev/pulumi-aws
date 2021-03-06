// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Gamelift Alias resource.
type Alias struct {
	pulumi.CustomResourceState

	// Alias ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the alias.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the alias.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyOutput `pulumi:"routingStrategy"`
	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil || args.RoutingStrategy == nil {
		return nil, errors.New("missing required argument 'RoutingStrategy'")
	}
	if args == nil {
		args = &AliasArgs{}
	}
	var resource Alias
	err := ctx.RegisterResource("aws:gamelift/alias:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws:gamelift/alias:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
	// Alias ARN.
	Arn *string `pulumi:"arn"`
	// Description of the alias.
	Description *string `pulumi:"description"`
	// Name of the alias.
	Name *string `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy *AliasRoutingStrategy `pulumi:"routingStrategy"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

type AliasState struct {
	// Alias ARN.
	Arn pulumi.StringPtrInput
	// Description of the alias.
	Description pulumi.StringPtrInput
	// Name of the alias.
	Name pulumi.StringPtrInput
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// Description of the alias.
	Description *string `pulumi:"description"`
	// Name of the alias.
	Name *string `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategy `pulumi:"routingStrategy"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Description of the alias.
	Description pulumi.StringPtrInput
	// Name of the alias.
	Name pulumi.StringPtrInput
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}
