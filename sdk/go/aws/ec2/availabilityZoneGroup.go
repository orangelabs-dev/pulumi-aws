// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EC2 Availability Zone Group, such as updating its opt-in status.
//
// > **NOTE:** This is an advanced resource. The provider will automatically assume management of the EC2 Availability Zone Group without import and perform no actions on removal from configuration.
type AvailabilityZoneGroup struct {
	pulumi.CustomResourceState

	// Name of the Availability Zone Group.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
	OptInStatus pulumi.StringOutput `pulumi:"optInStatus"`
}

// NewAvailabilityZoneGroup registers a new resource with the given unique name, arguments, and options.
func NewAvailabilityZoneGroup(ctx *pulumi.Context,
	name string, args *AvailabilityZoneGroupArgs, opts ...pulumi.ResourceOption) (*AvailabilityZoneGroup, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.OptInStatus == nil {
		return nil, errors.New("missing required argument 'OptInStatus'")
	}
	if args == nil {
		args = &AvailabilityZoneGroupArgs{}
	}
	var resource AvailabilityZoneGroup
	err := ctx.RegisterResource("aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvailabilityZoneGroup gets an existing AvailabilityZoneGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvailabilityZoneGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilityZoneGroupState, opts ...pulumi.ResourceOption) (*AvailabilityZoneGroup, error) {
	var resource AvailabilityZoneGroup
	err := ctx.ReadResource("aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvailabilityZoneGroup resources.
type availabilityZoneGroupState struct {
	// Name of the Availability Zone Group.
	GroupName *string `pulumi:"groupName"`
	// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
	OptInStatus *string `pulumi:"optInStatus"`
}

type AvailabilityZoneGroupState struct {
	// Name of the Availability Zone Group.
	GroupName pulumi.StringPtrInput
	// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
	OptInStatus pulumi.StringPtrInput
}

func (AvailabilityZoneGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilityZoneGroupState)(nil)).Elem()
}

type availabilityZoneGroupArgs struct {
	// Name of the Availability Zone Group.
	GroupName string `pulumi:"groupName"`
	// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
	OptInStatus string `pulumi:"optInStatus"`
}

// The set of arguments for constructing a AvailabilityZoneGroup resource.
type AvailabilityZoneGroupArgs struct {
	// Name of the Availability Zone Group.
	GroupName pulumi.StringInput
	// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
	OptInStatus pulumi.StringInput
}

func (AvailabilityZoneGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilityZoneGroupArgs)(nil)).Elem()
}
