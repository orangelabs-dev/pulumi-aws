// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a S3 bucket [inventory configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) resource.
type Inventory struct {
	pulumi.CustomResourceState

	// The name of the bucket where the inventory configuration will be stored.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Contains information about where to publish the inventory results (documented below).
	Destination InventoryDestinationOutput `pulumi:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	Filter InventoryFilterPtrOutput `pulumi:"filter"`
	// Object versions to include in the inventory list. Valid values: `All`, `Current`.
	IncludedObjectVersions pulumi.StringOutput `pulumi:"includedObjectVersions"`
	// Unique identifier of the inventory configuration for the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of optional fields that are included in the inventory results.
	// Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
	OptionalFields pulumi.StringArrayOutput `pulumi:"optionalFields"`
	// Specifies the schedule for generating inventory results (documented below).
	Schedule InventoryScheduleOutput `pulumi:"schedule"`
}

// NewInventory registers a new resource with the given unique name, arguments, and options.
func NewInventory(ctx *pulumi.Context,
	name string, args *InventoryArgs, opts ...pulumi.ResourceOption) (*Inventory, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	if args == nil || args.IncludedObjectVersions == nil {
		return nil, errors.New("missing required argument 'IncludedObjectVersions'")
	}
	if args == nil || args.Schedule == nil {
		return nil, errors.New("missing required argument 'Schedule'")
	}
	if args == nil {
		args = &InventoryArgs{}
	}
	var resource Inventory
	err := ctx.RegisterResource("aws:s3/inventory:Inventory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInventory gets an existing Inventory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInventory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InventoryState, opts ...pulumi.ResourceOption) (*Inventory, error) {
	var resource Inventory
	err := ctx.ReadResource("aws:s3/inventory:Inventory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Inventory resources.
type inventoryState struct {
	// The name of the bucket where the inventory configuration will be stored.
	Bucket *string `pulumi:"bucket"`
	// Contains information about where to publish the inventory results (documented below).
	Destination *InventoryDestination `pulumi:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	Enabled *bool `pulumi:"enabled"`
	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	Filter *InventoryFilter `pulumi:"filter"`
	// Object versions to include in the inventory list. Valid values: `All`, `Current`.
	IncludedObjectVersions *string `pulumi:"includedObjectVersions"`
	// Unique identifier of the inventory configuration for the bucket.
	Name *string `pulumi:"name"`
	// List of optional fields that are included in the inventory results.
	// Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
	OptionalFields []string `pulumi:"optionalFields"`
	// Specifies the schedule for generating inventory results (documented below).
	Schedule *InventorySchedule `pulumi:"schedule"`
}

type InventoryState struct {
	// The name of the bucket where the inventory configuration will be stored.
	Bucket pulumi.StringPtrInput
	// Contains information about where to publish the inventory results (documented below).
	Destination InventoryDestinationPtrInput
	// Specifies whether the inventory is enabled or disabled.
	Enabled pulumi.BoolPtrInput
	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	Filter InventoryFilterPtrInput
	// Object versions to include in the inventory list. Valid values: `All`, `Current`.
	IncludedObjectVersions pulumi.StringPtrInput
	// Unique identifier of the inventory configuration for the bucket.
	Name pulumi.StringPtrInput
	// List of optional fields that are included in the inventory results.
	// Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
	OptionalFields pulumi.StringArrayInput
	// Specifies the schedule for generating inventory results (documented below).
	Schedule InventorySchedulePtrInput
}

func (InventoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*inventoryState)(nil)).Elem()
}

type inventoryArgs struct {
	// The name of the bucket where the inventory configuration will be stored.
	Bucket string `pulumi:"bucket"`
	// Contains information about where to publish the inventory results (documented below).
	Destination InventoryDestination `pulumi:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	Enabled *bool `pulumi:"enabled"`
	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	Filter *InventoryFilter `pulumi:"filter"`
	// Object versions to include in the inventory list. Valid values: `All`, `Current`.
	IncludedObjectVersions string `pulumi:"includedObjectVersions"`
	// Unique identifier of the inventory configuration for the bucket.
	Name *string `pulumi:"name"`
	// List of optional fields that are included in the inventory results.
	// Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
	OptionalFields []string `pulumi:"optionalFields"`
	// Specifies the schedule for generating inventory results (documented below).
	Schedule InventorySchedule `pulumi:"schedule"`
}

// The set of arguments for constructing a Inventory resource.
type InventoryArgs struct {
	// The name of the bucket where the inventory configuration will be stored.
	Bucket pulumi.StringInput
	// Contains information about where to publish the inventory results (documented below).
	Destination InventoryDestinationInput
	// Specifies whether the inventory is enabled or disabled.
	Enabled pulumi.BoolPtrInput
	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	Filter InventoryFilterPtrInput
	// Object versions to include in the inventory list. Valid values: `All`, `Current`.
	IncludedObjectVersions pulumi.StringInput
	// Unique identifier of the inventory configuration for the bucket.
	Name pulumi.StringPtrInput
	// List of optional fields that are included in the inventory results.
	// Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
	OptionalFields pulumi.StringArrayInput
	// Specifies the schedule for generating inventory results (documented below).
	Schedule InventoryScheduleInput
}

func (InventoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inventoryArgs)(nil)).Elem()
}
