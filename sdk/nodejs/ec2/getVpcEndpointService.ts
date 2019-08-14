// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The VPC Endpoint Service data source details about a specific service that
 * can be specified when creating a VPC endpoint within the region configured in the provider.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc_endpoint_service.html.markdown.
 */
export function getVpcEndpointService(args?: GetVpcEndpointServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcEndpointServiceResult> & GetVpcEndpointServiceResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetVpcEndpointServiceResult> = pulumi.runtime.invoke("aws:ec2/getVpcEndpointService:getVpcEndpointService", {
        "service": args.service,
        "serviceName": args.serviceName,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getVpcEndpointService.
 */
export interface GetVpcEndpointServiceArgs {
    /**
     * The common name of an AWS service (e.g. `s3`).
     */
    readonly service?: string;
    /**
     * The service name that can be specified when creating a VPC endpoint.
     */
    readonly serviceName?: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getVpcEndpointService.
 */
export interface GetVpcEndpointServiceResult {
    /**
     * Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
     */
    readonly acceptanceRequired: boolean;
    /**
     * The Availability Zones in which the service is available.
     */
    readonly availabilityZones: string[];
    /**
     * The DNS names for the service.
     */
    readonly baseEndpointDnsNames: string[];
    /**
     * Whether or not the service manages its VPC endpoints - `true` or `false`.
     */
    readonly managesVpcEndpoints: boolean;
    /**
     * The AWS account ID of the service owner or `amazon`.
     */
    readonly owner: string;
    /**
     * The private DNS name for the service.
     */
    readonly privateDnsName: string;
    readonly service?: string;
    /**
     * The ID of the endpoint service.
     */
    readonly serviceId: string;
    readonly serviceName: string;
    /**
     * The service type, `Gateway` or `Interface`.
     */
    readonly serviceType: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * Whether or not the service supports endpoint policies - `true` or `false`.
     */
    readonly vpcEndpointPolicySupported: boolean;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
