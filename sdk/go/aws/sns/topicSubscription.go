// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

//   Provides a resource for subscribing to SNS topics. Requires that an SNS topic exist for the subscription to attach to.
// This resource allows you to automatically place messages sent to SNS topics in SQS queues, send them as HTTP(S) POST requests
// to a given endpoint, send SMS messages, or notify devices / applications. The most likely use case will
// probably be SQS queues.
//
// > **NOTE:** If the SNS topic and SQS queue are in different AWS regions, it is important for the "sns.TopicSubscription" to use an AWS provider that is in the same region of the SNS topic. If the "sns.TopicSubscription" is using a provider with a different region than the SNS topic, the subscription will fail to create.
//
// > **NOTE:** Setup of cross-account subscriptions from SNS topics to SQS queues requires the provider to have access to BOTH accounts.
//
// > **NOTE:** If SNS topic and SQS queue are in different AWS accounts but the same region it is important for the "sns.TopicSubscription" to use the AWS provider of the account with the SQS queue. If "sns.TopicSubscription" is using a Provider with a different account than the SQS queue, the provider creates the subscriptions but does not keep state and tries to re-create the subscription at every apply.
//
// > **NOTE:** If SNS topic and SQS queue are in different AWS accounts and different AWS regions it is important to recognize that the subscription needs to be initiated from the account with the SQS queue but in the region of the SNS topic.
type TopicSubscription struct {
	pulumi.CustomResourceState

	// The ARN of the subscription stored as a more user-friendly property
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
	ConfirmationTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"confirmationTimeoutInMinutes"`
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringPtrOutput `pulumi:"deliveryPolicy"`
	// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
	EndpointAutoConfirms pulumi.BoolPtrOutput `pulumi:"endpointAutoConfirms"`
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringPtrOutput `pulumi:"filterPolicy"`
	// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
	RawMessageDelivery pulumi.BoolPtrOutput `pulumi:"rawMessageDelivery"`
	// The ARN of the SNS topic to subscribe to
	Topic pulumi.StringOutput `pulumi:"topic"`
}

// NewTopicSubscription registers a new resource with the given unique name, arguments, and options.
func NewTopicSubscription(ctx *pulumi.Context,
	name string, args *TopicSubscriptionArgs, opts ...pulumi.ResourceOption) (*TopicSubscription, error) {
	if args == nil || args.Endpoint == nil {
		return nil, errors.New("missing required argument 'Endpoint'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.Topic == nil {
		return nil, errors.New("missing required argument 'Topic'")
	}
	if args == nil {
		args = &TopicSubscriptionArgs{}
	}
	var resource TopicSubscription
	err := ctx.RegisterResource("aws:sns/topicSubscription:TopicSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicSubscription gets an existing TopicSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicSubscriptionState, opts ...pulumi.ResourceOption) (*TopicSubscription, error) {
	var resource TopicSubscription
	err := ctx.ReadResource("aws:sns/topicSubscription:TopicSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicSubscription resources.
type topicSubscriptionState struct {
	// The ARN of the subscription stored as a more user-friendly property
	Arn *string `pulumi:"arn"`
	// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
	ConfirmationTimeoutInMinutes *int `pulumi:"confirmationTimeoutInMinutes"`
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy *string `pulumi:"deliveryPolicy"`
	// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
	Endpoint *string `pulumi:"endpoint"`
	// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
	EndpointAutoConfirms *bool `pulumi:"endpointAutoConfirms"`
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy *string `pulumi:"filterPolicy"`
	// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
	Protocol *string `pulumi:"protocol"`
	// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
	RawMessageDelivery *bool `pulumi:"rawMessageDelivery"`
	// The ARN of the SNS topic to subscribe to
	Topic *string `pulumi:"topic"`
}

type TopicSubscriptionState struct {
	// The ARN of the subscription stored as a more user-friendly property
	Arn pulumi.StringPtrInput
	// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
	ConfirmationTimeoutInMinutes pulumi.IntPtrInput
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringPtrInput
	// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
	Endpoint pulumi.StringPtrInput
	// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
	EndpointAutoConfirms pulumi.BoolPtrInput
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringPtrInput
	// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
	Protocol pulumi.StringPtrInput
	// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
	RawMessageDelivery pulumi.BoolPtrInput
	// The ARN of the SNS topic to subscribe to
	Topic pulumi.StringPtrInput
}

func (TopicSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicSubscriptionState)(nil)).Elem()
}

type topicSubscriptionArgs struct {
	// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
	ConfirmationTimeoutInMinutes *int `pulumi:"confirmationTimeoutInMinutes"`
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy *string `pulumi:"deliveryPolicy"`
	// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
	Endpoint string `pulumi:"endpoint"`
	// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
	EndpointAutoConfirms *bool `pulumi:"endpointAutoConfirms"`
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy *string `pulumi:"filterPolicy"`
	// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
	Protocol string `pulumi:"protocol"`
	// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
	RawMessageDelivery *bool `pulumi:"rawMessageDelivery"`
	// The ARN of the SNS topic to subscribe to
	Topic interface{} `pulumi:"topic"`
}

// The set of arguments for constructing a TopicSubscription resource.
type TopicSubscriptionArgs struct {
	// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
	ConfirmationTimeoutInMinutes pulumi.IntPtrInput
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringPtrInput
	// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
	Endpoint pulumi.StringInput
	// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
	EndpointAutoConfirms pulumi.BoolPtrInput
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringPtrInput
	// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
	Protocol pulumi.StringInput
	// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
	RawMessageDelivery pulumi.BoolPtrInput
	// The ARN of the SNS topic to subscribe to
	Topic pulumi.Input
}

func (TopicSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicSubscriptionArgs)(nil)).Elem()
}
