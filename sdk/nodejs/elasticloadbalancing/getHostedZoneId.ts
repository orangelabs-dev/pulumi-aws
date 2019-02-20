// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
 * in a given region for the purpose of using in an AWS Route53 Alias.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const main = pulumi.output(aws.elasticloadbalancing.getHostedZoneId({}));
 * const www = new aws.route53.Record("www", {
 *     aliases: [{
 *         evaluateTargetHealth: true,
 *         name: aws_elb_main.dnsName,
 *         zoneId: main.apply(main => main.id),
 *     }],
 *     name: "example.com",
 *     type: "A",
 *     zoneId: aws_route53_zone_primary.zoneId,
 * });
 * ```
 */
export function getHostedZoneId(args?: GetHostedZoneIdArgs, opts?: pulumi.InvokeOptions): Promise<GetHostedZoneIdResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:elasticloadbalancing/getHostedZoneId:getHostedZoneId", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getHostedZoneId.
 */
export interface GetHostedZoneIdArgs {
    /**
     * Name of the region whose AWS ELB HostedZoneId is desired.
     * Defaults to the region from the AWS provider configuration.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getHostedZoneId.
 */
export interface GetHostedZoneIdResult {
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
