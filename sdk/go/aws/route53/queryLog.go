// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Route53 query logging configuration resource.
//
// > **NOTE:** There are restrictions on the configuration of query logging. Notably,
// the CloudWatch log group must be in the `us-east-1` region,
// a permissive CloudWatch log resource policy must be in place, and
// the Route53 hosted zone must be public.
// See [Configuring Logging for DNS Queries](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html?console_help=true#query-logs-configuring) for additional details.
type QueryLog struct {
	pulumi.CustomResourceState

	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn pulumi.StringOutput `pulumi:"cloudwatchLogGroupArn"`
	// Route53 hosted zone ID to enable query logs.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewQueryLog registers a new resource with the given unique name, arguments, and options.
func NewQueryLog(ctx *pulumi.Context,
	name string, args *QueryLogArgs, opts ...pulumi.ResourceOption) (*QueryLog, error) {
	if args == nil || args.CloudwatchLogGroupArn == nil {
		return nil, errors.New("missing required argument 'CloudwatchLogGroupArn'")
	}
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	if args == nil {
		args = &QueryLogArgs{}
	}
	var resource QueryLog
	err := ctx.RegisterResource("aws:route53/queryLog:QueryLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueryLog gets an existing QueryLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueryLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryLogState, opts ...pulumi.ResourceOption) (*QueryLog, error) {
	var resource QueryLog
	err := ctx.ReadResource("aws:route53/queryLog:QueryLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueryLog resources.
type queryLogState struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn *string `pulumi:"cloudwatchLogGroupArn"`
	// Route53 hosted zone ID to enable query logs.
	ZoneId *string `pulumi:"zoneId"`
}

type QueryLogState struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn pulumi.StringPtrInput
	// Route53 hosted zone ID to enable query logs.
	ZoneId pulumi.StringPtrInput
}

func (QueryLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryLogState)(nil)).Elem()
}

type queryLogArgs struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn string `pulumi:"cloudwatchLogGroupArn"`
	// Route53 hosted zone ID to enable query logs.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a QueryLog resource.
type QueryLogArgs struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn pulumi.StringInput
	// Route53 hosted zone ID to enable query logs.
	ZoneId pulumi.StringInput
}

func (QueryLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryLogArgs)(nil)).Elem()
}
