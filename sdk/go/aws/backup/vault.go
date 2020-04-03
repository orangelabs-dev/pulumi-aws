// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS Backup vault resource.
type Vault struct {
	pulumi.CustomResourceState

	// The ARN of the vault.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`
	// Name of the backup vault to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints pulumi.IntOutput `pulumi:"recoveryPoints"`
	// Metadata that you can assign to help organize the resources that you create.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		args = &VaultArgs{}
	}
	var resource Vault
	err := ctx.RegisterResource("aws:backup/vault:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("aws:backup/vault:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
	// The ARN of the vault.
	Arn *string `pulumi:"arn"`
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// Name of the backup vault to create.
	Name *string `pulumi:"name"`
	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints *int `pulumi:"recoveryPoints"`
	// Metadata that you can assign to help organize the resources that you create.
	Tags map[string]interface{} `pulumi:"tags"`
}

type VaultState struct {
	// The ARN of the vault.
	Arn pulumi.StringPtrInput
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringPtrInput
	// Name of the backup vault to create.
	Name pulumi.StringPtrInput
	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints pulumi.IntPtrInput
	// Metadata that you can assign to help organize the resources that you create.
	Tags pulumi.MapInput
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// Name of the backup vault to create.
	Name *string `pulumi:"name"`
	// Metadata that you can assign to help organize the resources that you create.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringPtrInput
	// Name of the backup vault to create.
	Name pulumi.StringPtrInput
	// Metadata that you can assign to help organize the resources that you create.
	Tags pulumi.MapInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}
