// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package servicediscovery

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type HttpNamespace struct {
	pulumi.CustomResourceState

	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the http namespace.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewHttpNamespace registers a new resource with the given unique name, arguments, and options.
func NewHttpNamespace(ctx *pulumi.Context,
	name string, args *HttpNamespaceArgs, opts ...pulumi.ResourceOption) (*HttpNamespace, error) {
	if args == nil {
		args = &HttpNamespaceArgs{}
	}
	var resource HttpNamespace
	err := ctx.RegisterResource("aws:servicediscovery/httpNamespace:HttpNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpNamespace gets an existing HttpNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpNamespaceState, opts ...pulumi.ResourceOption) (*HttpNamespace, error) {
	var resource HttpNamespace
	err := ctx.ReadResource("aws:servicediscovery/httpNamespace:HttpNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpNamespace resources.
type httpNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn *string `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The name of the http namespace.
	Name *string `pulumi:"name"`
}

type HttpNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringPtrInput
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The name of the http namespace.
	Name pulumi.StringPtrInput
}

func (HttpNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpNamespaceState)(nil)).Elem()
}

type httpNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The name of the http namespace.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a HttpNamespace resource.
type HttpNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The name of the http namespace.
	Name pulumi.StringPtrInput
}

func (HttpNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpNamespaceArgs)(nil)).Elem()
}
