// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the id and rootResourceId of a REST API in
 * API Gateway. To fetch the REST API you must provide a name to match against. 
 * As there is no unique name constraint on REST APIs this data source will 
 * error if there is more than one match.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_rest_api.html.markdown.
 */
export function getRestApi(args: GetRestApiArgs, opts?: pulumi.InvokeOptions): Promise<GetRestApiResult> & GetRestApiResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetRestApiResult> = pulumi.runtime.invoke("aws:apigateway/getRestApi:getRestApi", {
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getRestApi.
 */
export interface GetRestApiArgs {
    /**
     * The name of the REST API to look up. If no REST API is found with this name, an error will be returned. 
     * If multiple REST APIs are found with this name, an error will be returned.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getRestApi.
 */
export interface GetRestApiResult {
    readonly name: string;
    /**
     * Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.
     */
    readonly rootResourceId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
