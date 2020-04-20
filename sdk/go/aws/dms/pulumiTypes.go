// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type EndpointElasticsearchSettings struct {
	EndpointUri             string `pulumi:"endpointUri"`
	ErrorRetryDuration      *int   `pulumi:"errorRetryDuration"`
	FullLoadErrorPercentage *int   `pulumi:"fullLoadErrorPercentage"`
	ServiceAccessRoleArn    string `pulumi:"serviceAccessRoleArn"`
}

// EndpointElasticsearchSettingsInput is an input type that accepts EndpointElasticsearchSettingsArgs and EndpointElasticsearchSettingsOutput values.
// You can construct a concrete instance of `EndpointElasticsearchSettingsInput` via:
//
// 		 EndpointElasticsearchSettingsArgs{...}
//
type EndpointElasticsearchSettingsInput interface {
	pulumi.Input

	ToEndpointElasticsearchSettingsOutput() EndpointElasticsearchSettingsOutput
	ToEndpointElasticsearchSettingsOutputWithContext(context.Context) EndpointElasticsearchSettingsOutput
}

type EndpointElasticsearchSettingsArgs struct {
	EndpointUri             pulumi.StringInput `pulumi:"endpointUri"`
	ErrorRetryDuration      pulumi.IntPtrInput `pulumi:"errorRetryDuration"`
	FullLoadErrorPercentage pulumi.IntPtrInput `pulumi:"fullLoadErrorPercentage"`
	ServiceAccessRoleArn    pulumi.StringInput `pulumi:"serviceAccessRoleArn"`
}

func (EndpointElasticsearchSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointElasticsearchSettings)(nil)).Elem()
}

func (i EndpointElasticsearchSettingsArgs) ToEndpointElasticsearchSettingsOutput() EndpointElasticsearchSettingsOutput {
	return i.ToEndpointElasticsearchSettingsOutputWithContext(context.Background())
}

func (i EndpointElasticsearchSettingsArgs) ToEndpointElasticsearchSettingsOutputWithContext(ctx context.Context) EndpointElasticsearchSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointElasticsearchSettingsOutput)
}

func (i EndpointElasticsearchSettingsArgs) ToEndpointElasticsearchSettingsPtrOutput() EndpointElasticsearchSettingsPtrOutput {
	return i.ToEndpointElasticsearchSettingsPtrOutputWithContext(context.Background())
}

func (i EndpointElasticsearchSettingsArgs) ToEndpointElasticsearchSettingsPtrOutputWithContext(ctx context.Context) EndpointElasticsearchSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointElasticsearchSettingsOutput).ToEndpointElasticsearchSettingsPtrOutputWithContext(ctx)
}

// EndpointElasticsearchSettingsPtrInput is an input type that accepts EndpointElasticsearchSettingsArgs, EndpointElasticsearchSettingsPtr and EndpointElasticsearchSettingsPtrOutput values.
// You can construct a concrete instance of `EndpointElasticsearchSettingsPtrInput` via:
//
// 		 EndpointElasticsearchSettingsArgs{...}
//
//  or:
//
// 		 nil
//
type EndpointElasticsearchSettingsPtrInput interface {
	pulumi.Input

	ToEndpointElasticsearchSettingsPtrOutput() EndpointElasticsearchSettingsPtrOutput
	ToEndpointElasticsearchSettingsPtrOutputWithContext(context.Context) EndpointElasticsearchSettingsPtrOutput
}

type endpointElasticsearchSettingsPtrType EndpointElasticsearchSettingsArgs

func EndpointElasticsearchSettingsPtr(v *EndpointElasticsearchSettingsArgs) EndpointElasticsearchSettingsPtrInput {
	return (*endpointElasticsearchSettingsPtrType)(v)
}

func (*endpointElasticsearchSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointElasticsearchSettings)(nil)).Elem()
}

func (i *endpointElasticsearchSettingsPtrType) ToEndpointElasticsearchSettingsPtrOutput() EndpointElasticsearchSettingsPtrOutput {
	return i.ToEndpointElasticsearchSettingsPtrOutputWithContext(context.Background())
}

func (i *endpointElasticsearchSettingsPtrType) ToEndpointElasticsearchSettingsPtrOutputWithContext(ctx context.Context) EndpointElasticsearchSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointElasticsearchSettingsPtrOutput)
}

type EndpointElasticsearchSettingsOutput struct{ *pulumi.OutputState }

func (EndpointElasticsearchSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointElasticsearchSettings)(nil)).Elem()
}

func (o EndpointElasticsearchSettingsOutput) ToEndpointElasticsearchSettingsOutput() EndpointElasticsearchSettingsOutput {
	return o
}

func (o EndpointElasticsearchSettingsOutput) ToEndpointElasticsearchSettingsOutputWithContext(ctx context.Context) EndpointElasticsearchSettingsOutput {
	return o
}

func (o EndpointElasticsearchSettingsOutput) ToEndpointElasticsearchSettingsPtrOutput() EndpointElasticsearchSettingsPtrOutput {
	return o.ToEndpointElasticsearchSettingsPtrOutputWithContext(context.Background())
}

func (o EndpointElasticsearchSettingsOutput) ToEndpointElasticsearchSettingsPtrOutputWithContext(ctx context.Context) EndpointElasticsearchSettingsPtrOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) *EndpointElasticsearchSettings {
		return &v
	}).(EndpointElasticsearchSettingsPtrOutput)
}
func (o EndpointElasticsearchSettingsOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) string { return v.EndpointUri }).(pulumi.StringOutput)
}

func (o EndpointElasticsearchSettingsOutput) ErrorRetryDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) *int { return v.ErrorRetryDuration }).(pulumi.IntPtrOutput)
}

func (o EndpointElasticsearchSettingsOutput) FullLoadErrorPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) *int { return v.FullLoadErrorPercentage }).(pulumi.IntPtrOutput)
}

func (o EndpointElasticsearchSettingsOutput) ServiceAccessRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) string { return v.ServiceAccessRoleArn }).(pulumi.StringOutput)
}

type EndpointElasticsearchSettingsPtrOutput struct{ *pulumi.OutputState }

func (EndpointElasticsearchSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointElasticsearchSettings)(nil)).Elem()
}

func (o EndpointElasticsearchSettingsPtrOutput) ToEndpointElasticsearchSettingsPtrOutput() EndpointElasticsearchSettingsPtrOutput {
	return o
}

func (o EndpointElasticsearchSettingsPtrOutput) ToEndpointElasticsearchSettingsPtrOutputWithContext(ctx context.Context) EndpointElasticsearchSettingsPtrOutput {
	return o
}

func (o EndpointElasticsearchSettingsPtrOutput) Elem() EndpointElasticsearchSettingsOutput {
	return o.ApplyT(func(v *EndpointElasticsearchSettings) EndpointElasticsearchSettings { return *v }).(EndpointElasticsearchSettingsOutput)
}

func (o EndpointElasticsearchSettingsPtrOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) string { return v.EndpointUri }).(pulumi.StringOutput)
}

func (o EndpointElasticsearchSettingsPtrOutput) ErrorRetryDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) *int { return v.ErrorRetryDuration }).(pulumi.IntPtrOutput)
}

func (o EndpointElasticsearchSettingsPtrOutput) FullLoadErrorPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) *int { return v.FullLoadErrorPercentage }).(pulumi.IntPtrOutput)
}

func (o EndpointElasticsearchSettingsPtrOutput) ServiceAccessRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointElasticsearchSettings) string { return v.ServiceAccessRoleArn }).(pulumi.StringOutput)
}

type EndpointKinesisSettings struct {
	MessageFormat        *string `pulumi:"messageFormat"`
	ServiceAccessRoleArn *string `pulumi:"serviceAccessRoleArn"`
	StreamArn            *string `pulumi:"streamArn"`
}

// EndpointKinesisSettingsInput is an input type that accepts EndpointKinesisSettingsArgs and EndpointKinesisSettingsOutput values.
// You can construct a concrete instance of `EndpointKinesisSettingsInput` via:
//
// 		 EndpointKinesisSettingsArgs{...}
//
type EndpointKinesisSettingsInput interface {
	pulumi.Input

	ToEndpointKinesisSettingsOutput() EndpointKinesisSettingsOutput
	ToEndpointKinesisSettingsOutputWithContext(context.Context) EndpointKinesisSettingsOutput
}

type EndpointKinesisSettingsArgs struct {
	MessageFormat        pulumi.StringPtrInput `pulumi:"messageFormat"`
	ServiceAccessRoleArn pulumi.StringPtrInput `pulumi:"serviceAccessRoleArn"`
	StreamArn            pulumi.StringPtrInput `pulumi:"streamArn"`
}

func (EndpointKinesisSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointKinesisSettings)(nil)).Elem()
}

func (i EndpointKinesisSettingsArgs) ToEndpointKinesisSettingsOutput() EndpointKinesisSettingsOutput {
	return i.ToEndpointKinesisSettingsOutputWithContext(context.Background())
}

func (i EndpointKinesisSettingsArgs) ToEndpointKinesisSettingsOutputWithContext(ctx context.Context) EndpointKinesisSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointKinesisSettingsOutput)
}

func (i EndpointKinesisSettingsArgs) ToEndpointKinesisSettingsPtrOutput() EndpointKinesisSettingsPtrOutput {
	return i.ToEndpointKinesisSettingsPtrOutputWithContext(context.Background())
}

func (i EndpointKinesisSettingsArgs) ToEndpointKinesisSettingsPtrOutputWithContext(ctx context.Context) EndpointKinesisSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointKinesisSettingsOutput).ToEndpointKinesisSettingsPtrOutputWithContext(ctx)
}

// EndpointKinesisSettingsPtrInput is an input type that accepts EndpointKinesisSettingsArgs, EndpointKinesisSettingsPtr and EndpointKinesisSettingsPtrOutput values.
// You can construct a concrete instance of `EndpointKinesisSettingsPtrInput` via:
//
// 		 EndpointKinesisSettingsArgs{...}
//
//  or:
//
// 		 nil
//
type EndpointKinesisSettingsPtrInput interface {
	pulumi.Input

	ToEndpointKinesisSettingsPtrOutput() EndpointKinesisSettingsPtrOutput
	ToEndpointKinesisSettingsPtrOutputWithContext(context.Context) EndpointKinesisSettingsPtrOutput
}

type endpointKinesisSettingsPtrType EndpointKinesisSettingsArgs

func EndpointKinesisSettingsPtr(v *EndpointKinesisSettingsArgs) EndpointKinesisSettingsPtrInput {
	return (*endpointKinesisSettingsPtrType)(v)
}

func (*endpointKinesisSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointKinesisSettings)(nil)).Elem()
}

func (i *endpointKinesisSettingsPtrType) ToEndpointKinesisSettingsPtrOutput() EndpointKinesisSettingsPtrOutput {
	return i.ToEndpointKinesisSettingsPtrOutputWithContext(context.Background())
}

func (i *endpointKinesisSettingsPtrType) ToEndpointKinesisSettingsPtrOutputWithContext(ctx context.Context) EndpointKinesisSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointKinesisSettingsPtrOutput)
}

type EndpointKinesisSettingsOutput struct{ *pulumi.OutputState }

func (EndpointKinesisSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointKinesisSettings)(nil)).Elem()
}

func (o EndpointKinesisSettingsOutput) ToEndpointKinesisSettingsOutput() EndpointKinesisSettingsOutput {
	return o
}

func (o EndpointKinesisSettingsOutput) ToEndpointKinesisSettingsOutputWithContext(ctx context.Context) EndpointKinesisSettingsOutput {
	return o
}

func (o EndpointKinesisSettingsOutput) ToEndpointKinesisSettingsPtrOutput() EndpointKinesisSettingsPtrOutput {
	return o.ToEndpointKinesisSettingsPtrOutputWithContext(context.Background())
}

func (o EndpointKinesisSettingsOutput) ToEndpointKinesisSettingsPtrOutputWithContext(ctx context.Context) EndpointKinesisSettingsPtrOutput {
	return o.ApplyT(func(v EndpointKinesisSettings) *EndpointKinesisSettings {
		return &v
	}).(EndpointKinesisSettingsPtrOutput)
}
func (o EndpointKinesisSettingsOutput) MessageFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointKinesisSettings) *string { return v.MessageFormat }).(pulumi.StringPtrOutput)
}

func (o EndpointKinesisSettingsOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointKinesisSettings) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

func (o EndpointKinesisSettingsOutput) StreamArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointKinesisSettings) *string { return v.StreamArn }).(pulumi.StringPtrOutput)
}

type EndpointKinesisSettingsPtrOutput struct{ *pulumi.OutputState }

func (EndpointKinesisSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointKinesisSettings)(nil)).Elem()
}

func (o EndpointKinesisSettingsPtrOutput) ToEndpointKinesisSettingsPtrOutput() EndpointKinesisSettingsPtrOutput {
	return o
}

func (o EndpointKinesisSettingsPtrOutput) ToEndpointKinesisSettingsPtrOutputWithContext(ctx context.Context) EndpointKinesisSettingsPtrOutput {
	return o
}

func (o EndpointKinesisSettingsPtrOutput) Elem() EndpointKinesisSettingsOutput {
	return o.ApplyT(func(v *EndpointKinesisSettings) EndpointKinesisSettings { return *v }).(EndpointKinesisSettingsOutput)
}

func (o EndpointKinesisSettingsPtrOutput) MessageFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointKinesisSettings) *string { return v.MessageFormat }).(pulumi.StringPtrOutput)
}

func (o EndpointKinesisSettingsPtrOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointKinesisSettings) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

func (o EndpointKinesisSettingsPtrOutput) StreamArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointKinesisSettings) *string { return v.StreamArn }).(pulumi.StringPtrOutput)
}

type EndpointMongodbSettings struct {
	AuthMechanism     *string `pulumi:"authMechanism"`
	AuthSource        *string `pulumi:"authSource"`
	AuthType          *string `pulumi:"authType"`
	DocsToInvestigate *string `pulumi:"docsToInvestigate"`
	ExtractDocId      *string `pulumi:"extractDocId"`
	NestingLevel      *string `pulumi:"nestingLevel"`
}

// EndpointMongodbSettingsInput is an input type that accepts EndpointMongodbSettingsArgs and EndpointMongodbSettingsOutput values.
// You can construct a concrete instance of `EndpointMongodbSettingsInput` via:
//
// 		 EndpointMongodbSettingsArgs{...}
//
type EndpointMongodbSettingsInput interface {
	pulumi.Input

	ToEndpointMongodbSettingsOutput() EndpointMongodbSettingsOutput
	ToEndpointMongodbSettingsOutputWithContext(context.Context) EndpointMongodbSettingsOutput
}

type EndpointMongodbSettingsArgs struct {
	AuthMechanism     pulumi.StringPtrInput `pulumi:"authMechanism"`
	AuthSource        pulumi.StringPtrInput `pulumi:"authSource"`
	AuthType          pulumi.StringPtrInput `pulumi:"authType"`
	DocsToInvestigate pulumi.StringPtrInput `pulumi:"docsToInvestigate"`
	ExtractDocId      pulumi.StringPtrInput `pulumi:"extractDocId"`
	NestingLevel      pulumi.StringPtrInput `pulumi:"nestingLevel"`
}

func (EndpointMongodbSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMongodbSettings)(nil)).Elem()
}

func (i EndpointMongodbSettingsArgs) ToEndpointMongodbSettingsOutput() EndpointMongodbSettingsOutput {
	return i.ToEndpointMongodbSettingsOutputWithContext(context.Background())
}

func (i EndpointMongodbSettingsArgs) ToEndpointMongodbSettingsOutputWithContext(ctx context.Context) EndpointMongodbSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsOutput)
}

func (i EndpointMongodbSettingsArgs) ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput {
	return i.ToEndpointMongodbSettingsPtrOutputWithContext(context.Background())
}

func (i EndpointMongodbSettingsArgs) ToEndpointMongodbSettingsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsOutput).ToEndpointMongodbSettingsPtrOutputWithContext(ctx)
}

// EndpointMongodbSettingsPtrInput is an input type that accepts EndpointMongodbSettingsArgs, EndpointMongodbSettingsPtr and EndpointMongodbSettingsPtrOutput values.
// You can construct a concrete instance of `EndpointMongodbSettingsPtrInput` via:
//
// 		 EndpointMongodbSettingsArgs{...}
//
//  or:
//
// 		 nil
//
type EndpointMongodbSettingsPtrInput interface {
	pulumi.Input

	ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput
	ToEndpointMongodbSettingsPtrOutputWithContext(context.Context) EndpointMongodbSettingsPtrOutput
}

type endpointMongodbSettingsPtrType EndpointMongodbSettingsArgs

func EndpointMongodbSettingsPtr(v *EndpointMongodbSettingsArgs) EndpointMongodbSettingsPtrInput {
	return (*endpointMongodbSettingsPtrType)(v)
}

func (*endpointMongodbSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointMongodbSettings)(nil)).Elem()
}

func (i *endpointMongodbSettingsPtrType) ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput {
	return i.ToEndpointMongodbSettingsPtrOutputWithContext(context.Background())
}

func (i *endpointMongodbSettingsPtrType) ToEndpointMongodbSettingsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsPtrOutput)
}

type EndpointMongodbSettingsOutput struct{ *pulumi.OutputState }

func (EndpointMongodbSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMongodbSettings)(nil)).Elem()
}

func (o EndpointMongodbSettingsOutput) ToEndpointMongodbSettingsOutput() EndpointMongodbSettingsOutput {
	return o
}

func (o EndpointMongodbSettingsOutput) ToEndpointMongodbSettingsOutputWithContext(ctx context.Context) EndpointMongodbSettingsOutput {
	return o
}

func (o EndpointMongodbSettingsOutput) ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput {
	return o.ToEndpointMongodbSettingsPtrOutputWithContext(context.Background())
}

func (o EndpointMongodbSettingsOutput) ToEndpointMongodbSettingsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *EndpointMongodbSettings {
		return &v
	}).(EndpointMongodbSettingsPtrOutput)
}
func (o EndpointMongodbSettingsOutput) AuthMechanism() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthMechanism }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) AuthSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthSource }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) DocsToInvestigate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.DocsToInvestigate }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) ExtractDocId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.ExtractDocId }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) NestingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.NestingLevel }).(pulumi.StringPtrOutput)
}

type EndpointMongodbSettingsPtrOutput struct{ *pulumi.OutputState }

func (EndpointMongodbSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointMongodbSettings)(nil)).Elem()
}

func (o EndpointMongodbSettingsPtrOutput) ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput {
	return o
}

func (o EndpointMongodbSettingsPtrOutput) ToEndpointMongodbSettingsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsPtrOutput {
	return o
}

func (o EndpointMongodbSettingsPtrOutput) Elem() EndpointMongodbSettingsOutput {
	return o.ApplyT(func(v *EndpointMongodbSettings) EndpointMongodbSettings { return *v }).(EndpointMongodbSettingsOutput)
}

func (o EndpointMongodbSettingsPtrOutput) AuthMechanism() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthMechanism }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) AuthSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthSource }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) DocsToInvestigate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.DocsToInvestigate }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) ExtractDocId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.ExtractDocId }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) NestingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.NestingLevel }).(pulumi.StringPtrOutput)
}

type EndpointS3Settings struct {
	BucketFolder            *string `pulumi:"bucketFolder"`
	BucketName              *string `pulumi:"bucketName"`
	CompressionType         *string `pulumi:"compressionType"`
	CsvDelimiter            *string `pulumi:"csvDelimiter"`
	CsvRowDelimiter         *string `pulumi:"csvRowDelimiter"`
	ExternalTableDefinition *string `pulumi:"externalTableDefinition"`
	ServiceAccessRoleArn    *string `pulumi:"serviceAccessRoleArn"`
}

// EndpointS3SettingsInput is an input type that accepts EndpointS3SettingsArgs and EndpointS3SettingsOutput values.
// You can construct a concrete instance of `EndpointS3SettingsInput` via:
//
// 		 EndpointS3SettingsArgs{...}
//
type EndpointS3SettingsInput interface {
	pulumi.Input

	ToEndpointS3SettingsOutput() EndpointS3SettingsOutput
	ToEndpointS3SettingsOutputWithContext(context.Context) EndpointS3SettingsOutput
}

type EndpointS3SettingsArgs struct {
	BucketFolder            pulumi.StringPtrInput `pulumi:"bucketFolder"`
	BucketName              pulumi.StringPtrInput `pulumi:"bucketName"`
	CompressionType         pulumi.StringPtrInput `pulumi:"compressionType"`
	CsvDelimiter            pulumi.StringPtrInput `pulumi:"csvDelimiter"`
	CsvRowDelimiter         pulumi.StringPtrInput `pulumi:"csvRowDelimiter"`
	ExternalTableDefinition pulumi.StringPtrInput `pulumi:"externalTableDefinition"`
	ServiceAccessRoleArn    pulumi.StringPtrInput `pulumi:"serviceAccessRoleArn"`
}

func (EndpointS3SettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointS3Settings)(nil)).Elem()
}

func (i EndpointS3SettingsArgs) ToEndpointS3SettingsOutput() EndpointS3SettingsOutput {
	return i.ToEndpointS3SettingsOutputWithContext(context.Background())
}

func (i EndpointS3SettingsArgs) ToEndpointS3SettingsOutputWithContext(ctx context.Context) EndpointS3SettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsOutput)
}

func (i EndpointS3SettingsArgs) ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput {
	return i.ToEndpointS3SettingsPtrOutputWithContext(context.Background())
}

func (i EndpointS3SettingsArgs) ToEndpointS3SettingsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsOutput).ToEndpointS3SettingsPtrOutputWithContext(ctx)
}

// EndpointS3SettingsPtrInput is an input type that accepts EndpointS3SettingsArgs, EndpointS3SettingsPtr and EndpointS3SettingsPtrOutput values.
// You can construct a concrete instance of `EndpointS3SettingsPtrInput` via:
//
// 		 EndpointS3SettingsArgs{...}
//
//  or:
//
// 		 nil
//
type EndpointS3SettingsPtrInput interface {
	pulumi.Input

	ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput
	ToEndpointS3SettingsPtrOutputWithContext(context.Context) EndpointS3SettingsPtrOutput
}

type endpointS3SettingsPtrType EndpointS3SettingsArgs

func EndpointS3SettingsPtr(v *EndpointS3SettingsArgs) EndpointS3SettingsPtrInput {
	return (*endpointS3SettingsPtrType)(v)
}

func (*endpointS3SettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointS3Settings)(nil)).Elem()
}

func (i *endpointS3SettingsPtrType) ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput {
	return i.ToEndpointS3SettingsPtrOutputWithContext(context.Background())
}

func (i *endpointS3SettingsPtrType) ToEndpointS3SettingsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsPtrOutput)
}

type EndpointS3SettingsOutput struct{ *pulumi.OutputState }

func (EndpointS3SettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointS3Settings)(nil)).Elem()
}

func (o EndpointS3SettingsOutput) ToEndpointS3SettingsOutput() EndpointS3SettingsOutput {
	return o
}

func (o EndpointS3SettingsOutput) ToEndpointS3SettingsOutputWithContext(ctx context.Context) EndpointS3SettingsOutput {
	return o
}

func (o EndpointS3SettingsOutput) ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput {
	return o.ToEndpointS3SettingsPtrOutputWithContext(context.Background())
}

func (o EndpointS3SettingsOutput) ToEndpointS3SettingsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *EndpointS3Settings {
		return &v
	}).(EndpointS3SettingsPtrOutput)
}
func (o EndpointS3SettingsOutput) BucketFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.BucketFolder }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) CompressionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CompressionType }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) CsvDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CsvDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) CsvRowDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CsvRowDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) ExternalTableDefinition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.ExternalTableDefinition }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

type EndpointS3SettingsPtrOutput struct{ *pulumi.OutputState }

func (EndpointS3SettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointS3Settings)(nil)).Elem()
}

func (o EndpointS3SettingsPtrOutput) ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput {
	return o
}

func (o EndpointS3SettingsPtrOutput) ToEndpointS3SettingsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsPtrOutput {
	return o
}

func (o EndpointS3SettingsPtrOutput) Elem() EndpointS3SettingsOutput {
	return o.ApplyT(func(v *EndpointS3Settings) EndpointS3Settings { return *v }).(EndpointS3SettingsOutput)
}

func (o EndpointS3SettingsPtrOutput) BucketFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.BucketFolder }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) CompressionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CompressionType }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) CsvDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CsvDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) CsvRowDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CsvRowDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) ExternalTableDefinition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.ExternalTableDefinition }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointElasticsearchSettingsOutput{})
	pulumi.RegisterOutputType(EndpointElasticsearchSettingsPtrOutput{})
	pulumi.RegisterOutputType(EndpointKinesisSettingsOutput{})
	pulumi.RegisterOutputType(EndpointKinesisSettingsPtrOutput{})
	pulumi.RegisterOutputType(EndpointMongodbSettingsOutput{})
	pulumi.RegisterOutputType(EndpointMongodbSettingsPtrOutput{})
	pulumi.RegisterOutputType(EndpointS3SettingsOutput{})
	pulumi.RegisterOutputType(EndpointS3SettingsPtrOutput{})
}
