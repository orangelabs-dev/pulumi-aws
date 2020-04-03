// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS App Mesh virtual router resource.
//
// ## Breaking Changes
//
// Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92) and [here](https://github.com/awslabs/aws-app-mesh-examples/issues/94)), `appmesh.VirtualRouter` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:
//
// * Remove service `serviceNames` from the `spec` argument.
// AWS has created a `appmesh.VirtualService` resource for each of service names.
// These resource can be imported using `import`.
//
// * Add a `listener` configuration block to the `spec` argument.
//
// The state associated with existing resources will automatically be migrated.
type VirtualRouter struct {
	pulumi.CustomResourceState

	// The ARN of the virtual router.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the virtual router.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the virtual router.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The name to use for the virtual router.
	Name pulumi.StringOutput `pulumi:"name"`
	// The virtual router specification to apply.
	Spec VirtualRouterSpecOutput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewVirtualRouter registers a new resource with the given unique name, arguments, and options.
func NewVirtualRouter(ctx *pulumi.Context,
	name string, args *VirtualRouterArgs, opts ...pulumi.ResourceOption) (*VirtualRouter, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil {
		args = &VirtualRouterArgs{}
	}
	var resource VirtualRouter
	err := ctx.RegisterResource("aws:appmesh/virtualRouter:VirtualRouter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualRouter gets an existing VirtualRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRouterState, opts ...pulumi.ResourceOption) (*VirtualRouter, error) {
	var resource VirtualRouter
	err := ctx.ReadResource("aws:appmesh/virtualRouter:VirtualRouter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualRouter resources.
type virtualRouterState struct {
	// The ARN of the virtual router.
	Arn *string `pulumi:"arn"`
	// The creation date of the virtual router.
	CreatedDate *string `pulumi:"createdDate"`
	// The last update date of the virtual router.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual router.
	MeshName *string `pulumi:"meshName"`
	// The name to use for the virtual router.
	Name *string `pulumi:"name"`
	// The virtual router specification to apply.
	Spec *VirtualRouterSpec `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type VirtualRouterState struct {
	// The ARN of the virtual router.
	Arn pulumi.StringPtrInput
	// The creation date of the virtual router.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the virtual router.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringPtrInput
	// The name to use for the virtual router.
	Name pulumi.StringPtrInput
	// The virtual router specification to apply.
	Spec VirtualRouterSpecPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (VirtualRouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterState)(nil)).Elem()
}

type virtualRouterArgs struct {
	// The name of the service mesh in which to create the virtual router.
	MeshName string `pulumi:"meshName"`
	// The name to use for the virtual router.
	Name *string `pulumi:"name"`
	// The virtual router specification to apply.
	Spec VirtualRouterSpec `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualRouter resource.
type VirtualRouterArgs struct {
	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringInput
	// The name to use for the virtual router.
	Name pulumi.StringPtrInput
	// The virtual router specification to apply.
	Spec VirtualRouterSpecInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (VirtualRouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterArgs)(nil)).Elem()
}
