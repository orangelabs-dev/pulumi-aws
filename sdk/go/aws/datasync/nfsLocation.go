// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an NFS Location within AWS DataSync.
//
// > **NOTE:** The DataSync Agents must be available before creating this resource.
type NfsLocation struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block containing information for connecting to the NFS File System.
	OnPremConfig NfsLocationOnPremConfigOutput `pulumi:"onPremConfig"`
	// Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
	ServerHostname pulumi.StringOutput `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapOutput    `pulumi:"tags"`
	Uri  pulumi.StringOutput `pulumi:"uri"`
}

// NewNfsLocation registers a new resource with the given unique name, arguments, and options.
func NewNfsLocation(ctx *pulumi.Context,
	name string, args *NfsLocationArgs, opts ...pulumi.ResourceOption) (*NfsLocation, error) {
	if args == nil || args.OnPremConfig == nil {
		return nil, errors.New("missing required argument 'OnPremConfig'")
	}
	if args == nil || args.ServerHostname == nil {
		return nil, errors.New("missing required argument 'ServerHostname'")
	}
	if args == nil || args.Subdirectory == nil {
		return nil, errors.New("missing required argument 'Subdirectory'")
	}
	if args == nil {
		args = &NfsLocationArgs{}
	}
	var resource NfsLocation
	err := ctx.RegisterResource("aws:datasync/nfsLocation:NfsLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNfsLocation gets an existing NfsLocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNfsLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NfsLocationState, opts ...pulumi.ResourceOption) (*NfsLocation, error) {
	var resource NfsLocation
	err := ctx.ReadResource("aws:datasync/nfsLocation:NfsLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NfsLocation resources.
type nfsLocationState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// Configuration block containing information for connecting to the NFS File System.
	OnPremConfig *NfsLocationOnPremConfig `pulumi:"onPremConfig"`
	// Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
	ServerHostname *string `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]interface{} `pulumi:"tags"`
	Uri  *string                `pulumi:"uri"`
}

type NfsLocationState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// Configuration block containing information for connecting to the NFS File System.
	OnPremConfig NfsLocationOnPremConfigPtrInput
	// Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
	ServerHostname pulumi.StringPtrInput
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapInput
	Uri  pulumi.StringPtrInput
}

func (NfsLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*nfsLocationState)(nil)).Elem()
}

type nfsLocationArgs struct {
	// Configuration block containing information for connecting to the NFS File System.
	OnPremConfig NfsLocationOnPremConfig `pulumi:"onPremConfig"`
	// Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
	ServerHostname string `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a NfsLocation resource.
type NfsLocationArgs struct {
	// Configuration block containing information for connecting to the NFS File System.
	OnPremConfig NfsLocationOnPremConfigInput
	// Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
	ServerHostname pulumi.StringInput
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapInput
}

func (NfsLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nfsLocationArgs)(nil)).Elem()
}
