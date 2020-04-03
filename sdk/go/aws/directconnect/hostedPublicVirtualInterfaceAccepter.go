// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the accepter's side of a Direct Connect hosted public virtual interface.
// This resource accepts ownership of a public virtual interface created by another AWS account.
type HostedPublicVirtualInterfaceAccepter struct {
	pulumi.CustomResourceState

	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringOutput `pulumi:"virtualInterfaceId"`
}

// NewHostedPublicVirtualInterfaceAccepter registers a new resource with the given unique name, arguments, and options.
func NewHostedPublicVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, args *HostedPublicVirtualInterfaceAccepterArgs, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterfaceAccepter, error) {
	if args == nil || args.VirtualInterfaceId == nil {
		return nil, errors.New("missing required argument 'VirtualInterfaceId'")
	}
	if args == nil {
		args = &HostedPublicVirtualInterfaceAccepterArgs{}
	}
	var resource HostedPublicVirtualInterfaceAccepter
	err := ctx.RegisterResource("aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedPublicVirtualInterfaceAccepter gets an existing HostedPublicVirtualInterfaceAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPublicVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedPublicVirtualInterfaceAccepterState, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterfaceAccepter, error) {
	var resource HostedPublicVirtualInterfaceAccepter
	err := ctx.ReadResource("aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedPublicVirtualInterfaceAccepter resources.
type hostedPublicVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId *string `pulumi:"virtualInterfaceId"`
}

type HostedPublicVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringPtrInput
}

func (HostedPublicVirtualInterfaceAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPublicVirtualInterfaceAccepterState)(nil)).Elem()
}

type hostedPublicVirtualInterfaceAccepterArgs struct {
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId string `pulumi:"virtualInterfaceId"`
}

// The set of arguments for constructing a HostedPublicVirtualInterfaceAccepter resource.
type HostedPublicVirtualInterfaceAccepterArgs struct {
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringInput
}

func (HostedPublicVirtualInterfaceAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPublicVirtualInterfaceAccepterArgs)(nil)).Elem()
}
