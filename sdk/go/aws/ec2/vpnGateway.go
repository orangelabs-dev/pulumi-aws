// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a VPC VPN Gateway.
type VpnGateway struct {
	pulumi.CustomResourceState

	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn pulumi.StringOutput `pulumi:"amazonSideAsn"`
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VPC ID to create in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewVpnGateway(ctx *pulumi.Context,
	name string, args *VpnGatewayArgs, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	if args == nil {
		args = &VpnGatewayArgs{}
	}
	var resource VpnGateway
	err := ctx.RegisterResource("aws:ec2/vpnGateway:VpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnGateway gets an existing VpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayState, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	var resource VpnGateway
	err := ctx.ReadResource("aws:ec2/vpnGateway:VpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnGateway resources.
type vpnGatewayState struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn *string `pulumi:"amazonSideAsn"`
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID to create in.
	VpcId *string `pulumi:"vpcId"`
}

type VpnGatewayState struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn pulumi.StringPtrInput
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The VPC ID to create in.
	VpcId pulumi.StringPtrInput
}

func (VpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayState)(nil)).Elem()
}

type vpnGatewayArgs struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn *string `pulumi:"amazonSideAsn"`
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID to create in.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpnGateway resource.
type VpnGatewayArgs struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn pulumi.StringPtrInput
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The VPC ID to create in.
	VpcId pulumi.StringPtrInput
}

func (VpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayArgs)(nil)).Elem()
}
