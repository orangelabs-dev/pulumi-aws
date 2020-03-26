// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
 * By using this data source, you can reference SQS queues without having to hardcode
 * the ARNs as input.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = aws.sqs.getQueue({
 *     name: "queue",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/sqs_queue.html.markdown.
 */
export function getQueue(args: GetQueueArgs, opts?: pulumi.InvokeOptions): Promise<GetQueueResult> & GetQueueResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetQueueResult> = pulumi.runtime.invoke("aws:sqs/getQueue:getQueue", {
        "name": args.name,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getQueue.
 */
export interface GetQueueArgs {
    /**
     * The name of the queue to match.
     */
    readonly name: string;
    /**
     * A mapping of tags for the resource.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getQueue.
 */
export interface GetQueueResult {
    /**
     * The Amazon Resource Name (ARN) of the queue.
     */
    readonly arn: string;
    readonly name: string;
    /**
     * A mapping of tags for the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * The URL of the queue.
     */
    readonly url: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
