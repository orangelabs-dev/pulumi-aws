// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a VPC Endpoint Route Table Association
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_endpoint_route_table_association.html.markdown.
type VpcEndpointRouteTableAssociation struct {
	pulumi.CustomResourceState

	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId pulumi.StringOutput `pulumi:"vpcEndpointId"`
}

// NewVpcEndpointRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointRouteTableAssociation(ctx *pulumi.Context,
	name string, args *VpcEndpointRouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*VpcEndpointRouteTableAssociation, error) {
	if args == nil || args.RouteTableId == nil {
		return nil, errors.New("missing required argument 'RouteTableId'")
	}
	if args == nil || args.VpcEndpointId == nil {
		return nil, errors.New("missing required argument 'VpcEndpointId'")
	}
	if args == nil {
		args = &VpcEndpointRouteTableAssociationArgs{}
	}
	var resource VpcEndpointRouteTableAssociation
	err := ctx.RegisterResource("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointRouteTableAssociation gets an existing VpcEndpointRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointRouteTableAssociationState, opts ...pulumi.ResourceOption) (*VpcEndpointRouteTableAssociation, error) {
	var resource VpcEndpointRouteTableAssociation
	err := ctx.ReadResource("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointRouteTableAssociation resources.
type vpcEndpointRouteTableAssociationState struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId *string `pulumi:"routeTableId"`
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
}

type VpcEndpointRouteTableAssociationState struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId pulumi.StringPtrInput
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId pulumi.StringPtrInput
}

func (VpcEndpointRouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointRouteTableAssociationState)(nil)).Elem()
}

type vpcEndpointRouteTableAssociationArgs struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId string `pulumi:"routeTableId"`
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId string `pulumi:"vpcEndpointId"`
}

// The set of arguments for constructing a VpcEndpointRouteTableAssociation resource.
type VpcEndpointRouteTableAssociationArgs struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId pulumi.StringInput
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId pulumi.StringInput
}

func (VpcEndpointRouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointRouteTableAssociationArgs)(nil)).Elem()
}
