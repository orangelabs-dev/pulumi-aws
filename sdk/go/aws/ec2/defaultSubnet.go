// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage a [default AWS VPC subnet](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html#default-vpc-basics)
// in the current region.
//
// The `ec2.DefaultSubnet` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead "adopts" it
// into management.
type DefaultSubnet struct {
	pulumi.CustomResourceState

	Arn                         pulumi.StringOutput `pulumi:"arn"`
	AssignIpv6AddressOnCreation pulumi.BoolOutput   `pulumi:"assignIpv6AddressOnCreation"`
	AvailabilityZone            pulumi.StringOutput `pulumi:"availabilityZone"`
	AvailabilityZoneId          pulumi.StringOutput `pulumi:"availabilityZoneId"`
	// The CIDR block for the subnet.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock              pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	Ipv6CidrBlockAssociationId pulumi.StringOutput `pulumi:"ipv6CidrBlockAssociationId"`
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address.
	MapPublicIpOnLaunch pulumi.BoolOutput `pulumi:"mapPublicIpOnLaunch"`
	// The ID of the AWS account that owns the subnet.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDefaultSubnet registers a new resource with the given unique name, arguments, and options.
func NewDefaultSubnet(ctx *pulumi.Context,
	name string, args *DefaultSubnetArgs, opts ...pulumi.ResourceOption) (*DefaultSubnet, error) {
	if args == nil || args.AvailabilityZone == nil {
		return nil, errors.New("missing required argument 'AvailabilityZone'")
	}
	if args == nil {
		args = &DefaultSubnetArgs{}
	}
	var resource DefaultSubnet
	err := ctx.RegisterResource("aws:ec2/defaultSubnet:DefaultSubnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultSubnet gets an existing DefaultSubnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultSubnetState, opts ...pulumi.ResourceOption) (*DefaultSubnet, error) {
	var resource DefaultSubnet
	err := ctx.ReadResource("aws:ec2/defaultSubnet:DefaultSubnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultSubnet resources.
type defaultSubnetState struct {
	Arn                         *string `pulumi:"arn"`
	AssignIpv6AddressOnCreation *bool   `pulumi:"assignIpv6AddressOnCreation"`
	AvailabilityZone            *string `pulumi:"availabilityZone"`
	AvailabilityZoneId          *string `pulumi:"availabilityZoneId"`
	// The CIDR block for the subnet.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock              *string `pulumi:"ipv6CidrBlock"`
	Ipv6CidrBlockAssociationId *string `pulumi:"ipv6CidrBlockAssociationId"`
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address.
	MapPublicIpOnLaunch *bool `pulumi:"mapPublicIpOnLaunch"`
	// The ID of the AWS account that owns the subnet.
	OwnerId *string `pulumi:"ownerId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type DefaultSubnetState struct {
	Arn                         pulumi.StringPtrInput
	AssignIpv6AddressOnCreation pulumi.BoolPtrInput
	AvailabilityZone            pulumi.StringPtrInput
	AvailabilityZoneId          pulumi.StringPtrInput
	// The CIDR block for the subnet.
	CidrBlock pulumi.StringPtrInput
	// The IPv6 CIDR block.
	Ipv6CidrBlock              pulumi.StringPtrInput
	Ipv6CidrBlockAssociationId pulumi.StringPtrInput
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address.
	MapPublicIpOnLaunch pulumi.BoolPtrInput
	// The ID of the AWS account that owns the subnet.
	OwnerId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (DefaultSubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultSubnetState)(nil)).Elem()
}

type defaultSubnetArgs struct {
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address.
	MapPublicIpOnLaunch *bool `pulumi:"mapPublicIpOnLaunch"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultSubnet resource.
type DefaultSubnetArgs struct {
	AvailabilityZone pulumi.StringInput
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address.
	MapPublicIpOnLaunch pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (DefaultSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultSubnetArgs)(nil)).Elem()
}
