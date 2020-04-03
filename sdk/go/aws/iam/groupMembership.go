// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **WARNING:** Multiple iam.GroupMembership resources with the same group name will produce inconsistent behavior!
//
// Provides a top level resource to manage IAM Group membership for IAM Users. For
// more information on managing IAM Groups or IAM Users, see [IAM Groups][1] or
// [IAM Users][2]
//
// > **Note:** `iam.GroupMembership` will conflict with itself if used more than once with the same group. To non-exclusively manage the users in a group, see the
// [`iam.UserGroupMembership` resource][3].
type GroupMembership struct {
	pulumi.CustomResourceState

	// The IAM Group name to attach the list of `users` to
	Group pulumi.StringOutput `pulumi:"group"`
	// The name to identify the Group Membership
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of IAM User names to associate with the Group
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewGroupMembership(ctx *pulumi.Context,
	name string, args *GroupMembershipArgs, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	if args == nil || args.Group == nil {
		return nil, errors.New("missing required argument 'Group'")
	}
	if args == nil || args.Users == nil {
		return nil, errors.New("missing required argument 'Users'")
	}
	if args == nil {
		args = &GroupMembershipArgs{}
	}
	var resource GroupMembership
	err := ctx.RegisterResource("aws:iam/groupMembership:GroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembership gets an existing GroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipState, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	var resource GroupMembership
	err := ctx.ReadResource("aws:iam/groupMembership:GroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembership resources.
type groupMembershipState struct {
	// The IAM Group name to attach the list of `users` to
	Group *string `pulumi:"group"`
	// The name to identify the Group Membership
	Name *string `pulumi:"name"`
	// A list of IAM User names to associate with the Group
	Users []string `pulumi:"users"`
}

type GroupMembershipState struct {
	// The IAM Group name to attach the list of `users` to
	Group pulumi.StringPtrInput
	// The name to identify the Group Membership
	Name pulumi.StringPtrInput
	// A list of IAM User names to associate with the Group
	Users pulumi.StringArrayInput
}

func (GroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipState)(nil)).Elem()
}

type groupMembershipArgs struct {
	// The IAM Group name to attach the list of `users` to
	Group string `pulumi:"group"`
	// The name to identify the Group Membership
	Name *string `pulumi:"name"`
	// A list of IAM User names to associate with the Group
	Users []string `pulumi:"users"`
}

// The set of arguments for constructing a GroupMembership resource.
type GroupMembershipArgs struct {
	// The IAM Group name to attach the list of `users` to
	Group pulumi.StringInput
	// The name to identify the Group Membership
	Name pulumi.StringPtrInput
	// A list of IAM User names to associate with the Group
	Users pulumi.StringArrayInput
}

func (GroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipArgs)(nil)).Elem()
}
