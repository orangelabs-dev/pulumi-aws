// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package sagemaker

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a lifecycle configuration for SageMaker Notebook Instances.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_notebook_instance_lifecycle_configuration.html.markdown.
type NotebookInstanceLifecycleConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate pulumi.StringPtrOutput `pulumi:"onCreate"`
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart pulumi.StringPtrOutput `pulumi:"onStart"`
}

// NewNotebookInstanceLifecycleConfiguration registers a new resource with the given unique name, arguments, and options.
func NewNotebookInstanceLifecycleConfiguration(ctx *pulumi.Context,
	name string, args *NotebookInstanceLifecycleConfigurationArgs, opts ...pulumi.ResourceOption) (*NotebookInstanceLifecycleConfiguration, error) {
	if args == nil {
		args = &NotebookInstanceLifecycleConfigurationArgs{}
	}
	var resource NotebookInstanceLifecycleConfiguration
	err := ctx.RegisterResource("aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookInstanceLifecycleConfiguration gets an existing NotebookInstanceLifecycleConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookInstanceLifecycleConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookInstanceLifecycleConfigurationState, opts ...pulumi.ResourceOption) (*NotebookInstanceLifecycleConfiguration, error) {
	var resource NotebookInstanceLifecycleConfiguration
	err := ctx.ReadResource("aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookInstanceLifecycleConfiguration resources.
type notebookInstanceLifecycleConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
	Arn *string `pulumi:"arn"`
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate *string `pulumi:"onCreate"`
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart *string `pulumi:"onStart"`
}

type NotebookInstanceLifecycleConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
	Arn pulumi.StringPtrInput
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate pulumi.StringPtrInput
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart pulumi.StringPtrInput
}

func (NotebookInstanceLifecycleConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceLifecycleConfigurationState)(nil)).Elem()
}

type notebookInstanceLifecycleConfigurationArgs struct {
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate *string `pulumi:"onCreate"`
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart *string `pulumi:"onStart"`
}

// The set of arguments for constructing a NotebookInstanceLifecycleConfiguration resource.
type NotebookInstanceLifecycleConfigurationArgs struct {
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate pulumi.StringPtrInput
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart pulumi.StringPtrInput
}

func (NotebookInstanceLifecycleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceLifecycleConfigurationArgs)(nil)).Elem()
}
