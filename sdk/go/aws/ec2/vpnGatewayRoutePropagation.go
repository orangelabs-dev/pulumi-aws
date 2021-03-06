// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Requests automatic route propagation between a VPN gateway and a route table.
//
// > **Note:** This resource should not be used with a route table that has
// the `propagatingVgws` argument set. If that argument is set, any route
// propagation not explicitly listed in its value will be removed.
type VpnGatewayRoutePropagation struct {
	pulumi.CustomResourceState

	// The id of the `ec2.RouteTable` to propagate routes into.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The id of the `ec2.VpnGateway` to propagate routes from.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
}

// NewVpnGatewayRoutePropagation registers a new resource with the given unique name, arguments, and options.
func NewVpnGatewayRoutePropagation(ctx *pulumi.Context,
	name string, args *VpnGatewayRoutePropagationArgs, opts ...pulumi.ResourceOption) (*VpnGatewayRoutePropagation, error) {
	if args == nil || args.RouteTableId == nil {
		return nil, errors.New("missing required argument 'RouteTableId'")
	}
	if args == nil || args.VpnGatewayId == nil {
		return nil, errors.New("missing required argument 'VpnGatewayId'")
	}
	if args == nil {
		args = &VpnGatewayRoutePropagationArgs{}
	}
	var resource VpnGatewayRoutePropagation
	err := ctx.RegisterResource("aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnGatewayRoutePropagation gets an existing VpnGatewayRoutePropagation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnGatewayRoutePropagation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayRoutePropagationState, opts ...pulumi.ResourceOption) (*VpnGatewayRoutePropagation, error) {
	var resource VpnGatewayRoutePropagation
	err := ctx.ReadResource("aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnGatewayRoutePropagation resources.
type vpnGatewayRoutePropagationState struct {
	// The id of the `ec2.RouteTable` to propagate routes into.
	RouteTableId *string `pulumi:"routeTableId"`
	// The id of the `ec2.VpnGateway` to propagate routes from.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type VpnGatewayRoutePropagationState struct {
	// The id of the `ec2.RouteTable` to propagate routes into.
	RouteTableId pulumi.StringPtrInput
	// The id of the `ec2.VpnGateway` to propagate routes from.
	VpnGatewayId pulumi.StringPtrInput
}

func (VpnGatewayRoutePropagationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayRoutePropagationState)(nil)).Elem()
}

type vpnGatewayRoutePropagationArgs struct {
	// The id of the `ec2.RouteTable` to propagate routes into.
	RouteTableId string `pulumi:"routeTableId"`
	// The id of the `ec2.VpnGateway` to propagate routes from.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a VpnGatewayRoutePropagation resource.
type VpnGatewayRoutePropagationArgs struct {
	// The id of the `ec2.RouteTable` to propagate routes into.
	RouteTableId pulumi.StringInput
	// The id of the `ec2.VpnGateway` to propagate routes from.
	VpnGatewayId pulumi.StringInput
}

func (VpnGatewayRoutePropagationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayRoutePropagationArgs)(nil)).Elem()
}
