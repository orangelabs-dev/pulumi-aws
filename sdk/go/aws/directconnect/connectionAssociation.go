// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Associates a Direct Connect Connection with a LAG.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_connection_association.html.markdown.
type ConnectionAssociation struct {
	pulumi.CustomResourceState

	// The ID of the connection.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// The ID of the LAG with which to associate the connection.
	LagId pulumi.StringOutput `pulumi:"lagId"`
}

// NewConnectionAssociation registers a new resource with the given unique name, arguments, and options.
func NewConnectionAssociation(ctx *pulumi.Context,
	name string, args *ConnectionAssociationArgs, opts ...pulumi.ResourceOption) (*ConnectionAssociation, error) {
	if args == nil || args.ConnectionId == nil {
		return nil, errors.New("missing required argument 'ConnectionId'")
	}
	if args == nil || args.LagId == nil {
		return nil, errors.New("missing required argument 'LagId'")
	}
	if args == nil {
		args = &ConnectionAssociationArgs{}
	}
	var resource ConnectionAssociation
	err := ctx.RegisterResource("aws:directconnect/connectionAssociation:ConnectionAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionAssociation gets an existing ConnectionAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionAssociationState, opts ...pulumi.ResourceOption) (*ConnectionAssociation, error) {
	var resource ConnectionAssociation
	err := ctx.ReadResource("aws:directconnect/connectionAssociation:ConnectionAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionAssociation resources.
type connectionAssociationState struct {
	// The ID of the connection.
	ConnectionId *string `pulumi:"connectionId"`
	// The ID of the LAG with which to associate the connection.
	LagId *string `pulumi:"lagId"`
}

type ConnectionAssociationState struct {
	// The ID of the connection.
	ConnectionId pulumi.StringPtrInput
	// The ID of the LAG with which to associate the connection.
	LagId pulumi.StringPtrInput
}

func (ConnectionAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionAssociationState)(nil)).Elem()
}

type connectionAssociationArgs struct {
	// The ID of the connection.
	ConnectionId string `pulumi:"connectionId"`
	// The ID of the LAG with which to associate the connection.
	LagId string `pulumi:"lagId"`
}

// The set of arguments for constructing a ConnectionAssociation resource.
type ConnectionAssociationArgs struct {
	// The ID of the connection.
	ConnectionId pulumi.StringInput
	// The ID of the LAG with which to associate the connection.
	LagId pulumi.StringInput
}

func (ConnectionAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionAssociationArgs)(nil)).Elem()
}
