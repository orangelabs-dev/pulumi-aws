// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package msk

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Amazon Managed Streaming for Kafka configuration. More information can be found on the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration.html).
//
// > **NOTE:** The API does not support deleting MSK configurations. Removing this resource will only remove the this provider state for it.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/msk_configuration.html.markdown.
type Configuration struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions pulumi.StringArrayOutput `pulumi:"kafkaVersions"`
	// Latest revision of the configuration.
	LatestRevision pulumi.IntOutput `pulumi:"latestRevision"`
	// Name of the configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties pulumi.StringOutput `pulumi:"serverProperties"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil || args.KafkaVersions == nil {
		return nil, errors.New("missing required argument 'KafkaVersions'")
	}
	if args == nil || args.ServerProperties == nil {
		return nil, errors.New("missing required argument 'ServerProperties'")
	}
	if args == nil {
		args = &ConfigurationArgs{}
	}
	var resource Configuration
	err := ctx.RegisterResource("aws:msk/configuration:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("aws:msk/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	// Amazon Resource Name (ARN) of the configuration.
	Arn *string `pulumi:"arn"`
	// Description of the configuration.
	Description *string `pulumi:"description"`
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions []string `pulumi:"kafkaVersions"`
	// Latest revision of the configuration.
	LatestRevision *int `pulumi:"latestRevision"`
	// Name of the configuration.
	Name *string `pulumi:"name"`
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties *string `pulumi:"serverProperties"`
}

type ConfigurationState struct {
	// Amazon Resource Name (ARN) of the configuration.
	Arn pulumi.StringPtrInput
	// Description of the configuration.
	Description pulumi.StringPtrInput
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions pulumi.StringArrayInput
	// Latest revision of the configuration.
	LatestRevision pulumi.IntPtrInput
	// Name of the configuration.
	Name pulumi.StringPtrInput
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties pulumi.StringPtrInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	// Description of the configuration.
	Description *string `pulumi:"description"`
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions []string `pulumi:"kafkaVersions"`
	// Name of the configuration.
	Name *string `pulumi:"name"`
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties string `pulumi:"serverProperties"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// Description of the configuration.
	Description pulumi.StringPtrInput
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions pulumi.StringArrayInput
	// Name of the configuration.
	Name pulumi.StringPtrInput
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties pulumi.StringInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}
