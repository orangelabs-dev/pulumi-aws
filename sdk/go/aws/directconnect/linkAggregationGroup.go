// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Direct Connect LAG. Connections can be added to the LAG via the [`directconnect.Connection`](https://www.terraform.io/docs/providers/aws/r/dx_connection.html) and [`directconnect.ConnectionAssociation`](https://www.terraform.io/docs/providers/aws/r/dx_connection_association.html) resources.
//
// > *NOTE:* When creating a LAG, Direct Connect requires creating a Connection. This provider will remove this unmanaged connection during resource creation.
type LinkAggregationGroup struct {
	pulumi.CustomResourceState

	// The ARN of the LAG.
	// * `jumboFrameCapable` -Indicates whether jumbo frames (9001 MTU) are supported.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The bandwidth of the individual physical connections bundled by the LAG. Available values: 1Gbps, 10Gbps. Case sensitive.
	ConnectionsBandwidth pulumi.StringOutput `pulumi:"connectionsBandwidth"`
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy pulumi.StringOutput `pulumi:"hasLogicalRedundancy"`
	JumboFrameCapable    pulumi.BoolOutput   `pulumi:"jumboFrameCapable"`
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the LAG.
	Name pulumi.StringOutput `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewLinkAggregationGroup registers a new resource with the given unique name, arguments, and options.
func NewLinkAggregationGroup(ctx *pulumi.Context,
	name string, args *LinkAggregationGroupArgs, opts ...pulumi.ResourceOption) (*LinkAggregationGroup, error) {
	if args == nil || args.ConnectionsBandwidth == nil {
		return nil, errors.New("missing required argument 'ConnectionsBandwidth'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil {
		args = &LinkAggregationGroupArgs{}
	}
	var resource LinkAggregationGroup
	err := ctx.RegisterResource("aws:directconnect/linkAggregationGroup:LinkAggregationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkAggregationGroup gets an existing LinkAggregationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkAggregationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkAggregationGroupState, opts ...pulumi.ResourceOption) (*LinkAggregationGroup, error) {
	var resource LinkAggregationGroup
	err := ctx.ReadResource("aws:directconnect/linkAggregationGroup:LinkAggregationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkAggregationGroup resources.
type linkAggregationGroupState struct {
	// The ARN of the LAG.
	// * `jumboFrameCapable` -Indicates whether jumbo frames (9001 MTU) are supported.
	Arn *string `pulumi:"arn"`
	// The bandwidth of the individual physical connections bundled by the LAG. Available values: 1Gbps, 10Gbps. Case sensitive.
	ConnectionsBandwidth *string `pulumi:"connectionsBandwidth"`
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy *string `pulumi:"hasLogicalRedundancy"`
	JumboFrameCapable    *bool   `pulumi:"jumboFrameCapable"`
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location *string `pulumi:"location"`
	// The name of the LAG.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type LinkAggregationGroupState struct {
	// The ARN of the LAG.
	// * `jumboFrameCapable` -Indicates whether jumbo frames (9001 MTU) are supported.
	Arn pulumi.StringPtrInput
	// The bandwidth of the individual physical connections bundled by the LAG. Available values: 1Gbps, 10Gbps. Case sensitive.
	ConnectionsBandwidth pulumi.StringPtrInput
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy pulumi.StringPtrInput
	JumboFrameCapable    pulumi.BoolPtrInput
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringPtrInput
	// The name of the LAG.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (LinkAggregationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkAggregationGroupState)(nil)).Elem()
}

type linkAggregationGroupArgs struct {
	// The bandwidth of the individual physical connections bundled by the LAG. Available values: 1Gbps, 10Gbps. Case sensitive.
	ConnectionsBandwidth string `pulumi:"connectionsBandwidth"`
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location string `pulumi:"location"`
	// The name of the LAG.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a LinkAggregationGroup resource.
type LinkAggregationGroupArgs struct {
	// The bandwidth of the individual physical connections bundled by the LAG. Available values: 1Gbps, 10Gbps. Case sensitive.
	ConnectionsBandwidth pulumi.StringInput
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringInput
	// The name of the LAG.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (LinkAggregationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkAggregationGroupArgs)(nil)).Elem()
}
