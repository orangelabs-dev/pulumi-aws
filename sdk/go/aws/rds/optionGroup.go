// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an RDS DB option group resource. Documentation of the available options for various RDS engines can be found at:
//
// * [MariaDB Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Options.html)
// * [Microsoft SQL Server Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.html)
// * [MySQL Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.Options.html)
// * [Oracle Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.Options.html)
type OptionGroup struct {
	pulumi.CustomResourceState

	// The ARN of the db option group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies the name of the engine that this option group should be associated with.
	EngineName pulumi.StringOutput `pulumi:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion pulumi.StringOutput `pulumi:"majorEngineVersion"`
	// The Name of the setting.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription pulumi.StringOutput `pulumi:"optionGroupDescription"`
	// A list of Options to apply.
	Options OptionGroupOptionArrayOutput `pulumi:"options"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewOptionGroup registers a new resource with the given unique name, arguments, and options.
func NewOptionGroup(ctx *pulumi.Context,
	name string, args *OptionGroupArgs, opts ...pulumi.ResourceOption) (*OptionGroup, error) {
	if args == nil || args.EngineName == nil {
		return nil, errors.New("missing required argument 'EngineName'")
	}
	if args == nil || args.MajorEngineVersion == nil {
		return nil, errors.New("missing required argument 'MajorEngineVersion'")
	}
	if args == nil {
		args = &OptionGroupArgs{}
	}
	if args.OptionGroupDescription == nil {
		args.OptionGroupDescription = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource OptionGroup
	err := ctx.RegisterResource("aws:rds/optionGroup:OptionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOptionGroup gets an existing OptionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOptionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OptionGroupState, opts ...pulumi.ResourceOption) (*OptionGroup, error) {
	var resource OptionGroup
	err := ctx.ReadResource("aws:rds/optionGroup:OptionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OptionGroup resources.
type optionGroupState struct {
	// The ARN of the db option group.
	Arn *string `pulumi:"arn"`
	// Specifies the name of the engine that this option group should be associated with.
	EngineName *string `pulumi:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion *string `pulumi:"majorEngineVersion"`
	// The Name of the setting.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix *string `pulumi:"namePrefix"`
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription *string `pulumi:"optionGroupDescription"`
	// A list of Options to apply.
	Options []OptionGroupOption `pulumi:"options"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type OptionGroupState struct {
	// The ARN of the db option group.
	Arn pulumi.StringPtrInput
	// Specifies the name of the engine that this option group should be associated with.
	EngineName pulumi.StringPtrInput
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion pulumi.StringPtrInput
	// The Name of the setting.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix pulumi.StringPtrInput
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription pulumi.StringPtrInput
	// A list of Options to apply.
	Options OptionGroupOptionArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (OptionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*optionGroupState)(nil)).Elem()
}

type optionGroupArgs struct {
	// Specifies the name of the engine that this option group should be associated with.
	EngineName string `pulumi:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion string `pulumi:"majorEngineVersion"`
	// The Name of the setting.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix *string `pulumi:"namePrefix"`
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription *string `pulumi:"optionGroupDescription"`
	// A list of Options to apply.
	Options []OptionGroupOption `pulumi:"options"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a OptionGroup resource.
type OptionGroupArgs struct {
	// Specifies the name of the engine that this option group should be associated with.
	EngineName pulumi.StringInput
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion pulumi.StringInput
	// The Name of the setting.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix pulumi.StringPtrInput
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription pulumi.StringPtrInput
	// A list of Options to apply.
	Options OptionGroupOptionArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (OptionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*optionGroupArgs)(nil)).Elem()
}
