// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sns
{
    /// <summary>
    ///   Provides a resource for subscribing to SNS topics. Requires that an SNS topic exist for the subscription to attach to.
    /// This resource allows you to automatically place messages sent to SNS topics in SQS queues, send them as HTTP(S) POST requests
    /// to a given endpoint, send SMS messages, or notify devices / applications. The most likely use case will
    /// probably be SQS queues.
    /// 
    /// &gt; **NOTE:** If the SNS topic and SQS queue are in different AWS regions, it is important for the "aws.sns.TopicSubscription" to use an AWS provider that is in the same region of the SNS topic. If the "aws.sns.TopicSubscription" is using a provider with a different region than the SNS topic, the subscription will fail to create.
    /// 
    /// &gt; **NOTE:** Setup of cross-account subscriptions from SNS topics to SQS queues requires the provider to have access to BOTH accounts.
    /// 
    /// &gt; **NOTE:** If SNS topic and SQS queue are in different AWS accounts but the same region it is important for the "aws.sns.TopicSubscription" to use the AWS provider of the account with the SQS queue. If "aws.sns.TopicSubscription" is using a Provider with a different account than the SQS queue, the provider creates the subscriptions but does not keep state and tries to re-create the subscription at every apply.
    /// 
    /// &gt; **NOTE:** If SNS topic and SQS queue are in different AWS accounts and different AWS regions it is important to recognize that the subscription needs to be initiated from the account with the SQS queue but in the region of the SNS topic.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sns_topic_subscription.html.markdown.
    /// </summary>
    public partial class TopicSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the subscription stored as a more user-friendly property
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
        /// </summary>
        [Output("confirmationTimeoutInMinutes")]
        public Output<int?> ConfirmationTimeoutInMinutes { get; private set; } = null!;

        /// <summary>
        /// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
        /// </summary>
        [Output("deliveryPolicy")]
        public Output<string?> DeliveryPolicy { get; private set; } = null!;

        /// <summary>
        /// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
        /// </summary>
        [Output("endpointAutoConfirms")]
        public Output<bool?> EndpointAutoConfirms { get; private set; } = null!;

        /// <summary>
        /// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
        /// </summary>
        [Output("filterPolicy")]
        public Output<string?> FilterPolicy { get; private set; } = null!;

        /// <summary>
        /// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
        /// </summary>
        [Output("rawMessageDelivery")]
        public Output<bool?> RawMessageDelivery { get; private set; } = null!;

        /// <summary>
        /// The ARN of the SNS topic to subscribe to
        /// </summary>
        [Output("topic")]
        public Output<string> Topic { get; private set; } = null!;


        /// <summary>
        /// Create a TopicSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TopicSubscription(string name, TopicSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("aws:sns/topicSubscription:TopicSubscription", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private TopicSubscription(string name, Input<string> id, TopicSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("aws:sns/topicSubscription:TopicSubscription", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TopicSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TopicSubscription Get(string name, Input<string> id, TopicSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new TopicSubscription(name, id, state, options);
        }
    }

    public sealed class TopicSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
        /// </summary>
        [Input("confirmationTimeoutInMinutes")]
        public Input<int>? ConfirmationTimeoutInMinutes { get; set; }

        /// <summary>
        /// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
        /// </summary>
        [Input("deliveryPolicy")]
        public Input<string>? DeliveryPolicy { get; set; }

        /// <summary>
        /// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
        /// </summary>
        [Input("endpointAutoConfirms")]
        public Input<bool>? EndpointAutoConfirms { get; set; }

        /// <summary>
        /// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
        /// </summary>
        [Input("filterPolicy")]
        public Input<string>? FilterPolicy { get; set; }

        /// <summary>
        /// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
        /// </summary>
        [Input("rawMessageDelivery")]
        public Input<bool>? RawMessageDelivery { get; set; }

        /// <summary>
        /// The ARN of the SNS topic to subscribe to
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public TopicSubscriptionArgs()
        {
        }
    }

    public sealed class TopicSubscriptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the subscription stored as a more user-friendly property
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
        /// </summary>
        [Input("confirmationTimeoutInMinutes")]
        public Input<int>? ConfirmationTimeoutInMinutes { get; set; }

        /// <summary>
        /// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
        /// </summary>
        [Input("deliveryPolicy")]
        public Input<string>? DeliveryPolicy { get; set; }

        /// <summary>
        /// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
        /// </summary>
        [Input("endpointAutoConfirms")]
        public Input<bool>? EndpointAutoConfirms { get; set; }

        /// <summary>
        /// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
        /// </summary>
        [Input("filterPolicy")]
        public Input<string>? FilterPolicy { get; set; }

        /// <summary>
        /// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
        /// </summary>
        [Input("rawMessageDelivery")]
        public Input<bool>? RawMessageDelivery { get; set; }

        /// <summary>
        /// The ARN of the SNS topic to subscribe to
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public TopicSubscriptionState()
        {
        }
    }
}
