// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `aws.wafregional.RateBasedRule` Retrieves a WAF Regional Rate Based Rule Resource Id.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = aws.wafregional.getRateBasedMod({
 *     name: "tfWAFRegionalRateBasedRule",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/wafregional_rate_based_rule.html.markdown.
 */
export function getRateBasedMod(args: GetRateBasedModArgs, opts?: pulumi.InvokeOptions): Promise<GetRateBasedModResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:wafregional/getRateBasedMod:getRateBasedMod", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRateBasedMod.
 */
export interface GetRateBasedModArgs {
    /**
     * The name of the WAF Regional rate based rule.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getRateBasedMod.
 */
export interface GetRateBasedModResult {
    readonly name: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}