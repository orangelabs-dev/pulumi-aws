// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glacier

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type VaultNotification struct {
	// You can configure a vault to publish a notification for `ArchiveRetrievalCompleted` and `InventoryRetrievalCompleted` events.
	Events []string `pulumi:"events"`
	// The SNS Topic ARN.
	SnsTopic string `pulumi:"snsTopic"`
}

type VaultNotificationInput interface {
	pulumi.Input

	ToVaultNotificationOutput() VaultNotificationOutput
	ToVaultNotificationOutputWithContext(context.Context) VaultNotificationOutput
}

type VaultNotificationArgs struct {
	// You can configure a vault to publish a notification for `ArchiveRetrievalCompleted` and `InventoryRetrievalCompleted` events.
	Events pulumi.StringArrayInput `pulumi:"events"`
	// The SNS Topic ARN.
	SnsTopic pulumi.StringInput `pulumi:"snsTopic"`
}

func (VaultNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultNotification)(nil)).Elem()
}

func (i VaultNotificationArgs) ToVaultNotificationOutput() VaultNotificationOutput {
	return i.ToVaultNotificationOutputWithContext(context.Background())
}

func (i VaultNotificationArgs) ToVaultNotificationOutputWithContext(ctx context.Context) VaultNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultNotificationOutput)
}

type VaultNotificationArrayInput interface {
	pulumi.Input

	ToVaultNotificationArrayOutput() VaultNotificationArrayOutput
	ToVaultNotificationArrayOutputWithContext(context.Context) VaultNotificationArrayOutput
}

type VaultNotificationArray []VaultNotificationInput

func (VaultNotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultNotification)(nil)).Elem()
}

func (i VaultNotificationArray) ToVaultNotificationArrayOutput() VaultNotificationArrayOutput {
	return i.ToVaultNotificationArrayOutputWithContext(context.Background())
}

func (i VaultNotificationArray) ToVaultNotificationArrayOutputWithContext(ctx context.Context) VaultNotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultNotificationArrayOutput)
}

type VaultNotificationOutput struct{ *pulumi.OutputState }

func (VaultNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultNotification)(nil)).Elem()
}

func (o VaultNotificationOutput) ToVaultNotificationOutput() VaultNotificationOutput {
	return o
}

func (o VaultNotificationOutput) ToVaultNotificationOutputWithContext(ctx context.Context) VaultNotificationOutput {
	return o
}

// You can configure a vault to publish a notification for `ArchiveRetrievalCompleted` and `InventoryRetrievalCompleted` events.
func (o VaultNotificationOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VaultNotification) []string { return v.Events }).(pulumi.StringArrayOutput)
}

// The SNS Topic ARN.
func (o VaultNotificationOutput) SnsTopic() pulumi.StringOutput {
	return o.ApplyT(func(v VaultNotification) string { return v.SnsTopic }).(pulumi.StringOutput)
}

type VaultNotificationArrayOutput struct{ *pulumi.OutputState }

func (VaultNotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultNotification)(nil)).Elem()
}

func (o VaultNotificationArrayOutput) ToVaultNotificationArrayOutput() VaultNotificationArrayOutput {
	return o
}

func (o VaultNotificationArrayOutput) ToVaultNotificationArrayOutputWithContext(ctx context.Context) VaultNotificationArrayOutput {
	return o
}

func (o VaultNotificationArrayOutput) Index(i pulumi.IntInput) VaultNotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultNotification {
		return vs[0].([]VaultNotification)[vs[1].(int)]
	}).(VaultNotificationOutput)
}

func init() {
	pulumi.RegisterOutputType(VaultNotificationOutput{})
	pulumi.RegisterOutputType(VaultNotificationArrayOutput{})
}
