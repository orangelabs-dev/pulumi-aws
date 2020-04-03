// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic Container Registry Repository.
type Repository struct {
	pulumi.CustomResourceState

	// Full ARN of the repository.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration RepositoryImageScanningConfigurationPtrOutput `pulumi:"imageScanningConfiguration"`
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability pulumi.StringPtrOutput `pulumi:"imageTagMutability"`
	// Name of the repository.
	Name pulumi.StringOutput `pulumi:"name"`
	// The registry ID where the repository was created.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
	RepositoryUrl pulumi.StringOutput `pulumi:"repositoryUrl"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		args = &RepositoryArgs{}
	}
	var resource Repository
	err := ctx.RegisterResource("aws:ecr/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("aws:ecr/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Full ARN of the repository.
	Arn *string `pulumi:"arn"`
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration *RepositoryImageScanningConfiguration `pulumi:"imageScanningConfiguration"`
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability *string `pulumi:"imageTagMutability"`
	// Name of the repository.
	Name *string `pulumi:"name"`
	// The registry ID where the repository was created.
	RegistryId *string `pulumi:"registryId"`
	// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type RepositoryState struct {
	// Full ARN of the repository.
	Arn pulumi.StringPtrInput
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration RepositoryImageScanningConfigurationPtrInput
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability pulumi.StringPtrInput
	// Name of the repository.
	Name pulumi.StringPtrInput
	// The registry ID where the repository was created.
	RegistryId pulumi.StringPtrInput
	// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
	RepositoryUrl pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration *RepositoryImageScanningConfiguration `pulumi:"imageScanningConfiguration"`
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability *string `pulumi:"imageTagMutability"`
	// Name of the repository.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration RepositoryImageScanningConfigurationPtrInput
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability pulumi.StringPtrInput
	// Name of the repository.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}
