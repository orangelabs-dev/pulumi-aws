// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create an association between a route table and a subnet or a route table and an
// internet gateway or virtual private gateway.
type RouteTableAssociation struct {
	pulumi.CustomResourceState

	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId pulumi.StringPtrOutput `pulumi:"gatewayId"`
	// The ID of the routing table to associate with.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
}

// NewRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewRouteTableAssociation(ctx *pulumi.Context,
	name string, args *RouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*RouteTableAssociation, error) {
	if args == nil || args.RouteTableId == nil {
		return nil, errors.New("missing required argument 'RouteTableId'")
	}
	if args == nil {
		args = &RouteTableAssociationArgs{}
	}
	var resource RouteTableAssociation
	err := ctx.RegisterResource("aws:ec2/routeTableAssociation:RouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTableAssociation gets an existing RouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableAssociationState, opts ...pulumi.ResourceOption) (*RouteTableAssociation, error) {
	var resource RouteTableAssociation
	err := ctx.ReadResource("aws:ec2/routeTableAssociation:RouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTableAssociation resources.
type routeTableAssociationState struct {
	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId *string `pulumi:"gatewayId"`
	// The ID of the routing table to associate with.
	RouteTableId *string `pulumi:"routeTableId"`
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId *string `pulumi:"subnetId"`
}

type RouteTableAssociationState struct {
	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId pulumi.StringPtrInput
	// The ID of the routing table to associate with.
	RouteTableId pulumi.StringPtrInput
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId pulumi.StringPtrInput
}

func (RouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableAssociationState)(nil)).Elem()
}

type routeTableAssociationArgs struct {
	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId *string `pulumi:"gatewayId"`
	// The ID of the routing table to associate with.
	RouteTableId string `pulumi:"routeTableId"`
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId *string `pulumi:"subnetId"`
}

// The set of arguments for constructing a RouteTableAssociation resource.
type RouteTableAssociationArgs struct {
	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId pulumi.StringPtrInput
	// The ID of the routing table to associate with.
	RouteTableId pulumi.StringInput
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId pulumi.StringPtrInput
}

func (RouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableAssociationArgs)(nil)).Elem()
}
