// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
// By using this data source, you can reference SQS queues without having to hardcode
// the ARNs as input.
func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	var rv LookupQueueResult
	err := ctx.Invoke("aws:sqs/getQueue:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueue.
type LookupQueueArgs struct {
	// The name of the queue to match.
	Name string `pulumi:"name"`
	// A mapping of tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getQueue.
type LookupQueueResult struct {
	// The Amazon Resource Name (ARN) of the queue.
	Arn string `pulumi:"arn"`
	// id is the provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// A mapping of tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The URL of the queue.
	Url string `pulumi:"url"`
}
