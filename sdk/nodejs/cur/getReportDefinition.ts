// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an AWS Cost and Usage Report Definition.
 * 
 * > *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.
 * 
 * > *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cur_report_definition.html.markdown.
 */
export function getReportDefinition(args: GetReportDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetReportDefinitionResult> & GetReportDefinitionResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetReportDefinitionResult> = pulumi.runtime.invoke("aws:cur/getReportDefinition:getReportDefinition", {
        "reportName": args.reportName,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getReportDefinition.
 */
export interface GetReportDefinitionArgs {
    /**
     * The name of the report definition to match.
     */
    readonly reportName: string;
}

/**
 * A collection of values returned by getReportDefinition.
 */
export interface GetReportDefinitionResult {
    /**
     * A list of additional artifacts.
     */
    readonly additionalArtifacts: string[];
    /**
     * A list of schema elements.
     */
    readonly additionalSchemaElements: string[];
    /**
     * Preferred format for report.
     */
    readonly compression: string;
    /**
     * Preferred compression format for report.
     */
    readonly format: string;
    readonly reportName: string;
    /**
     * Name of customer S3 bucket.
     */
    readonly s3Bucket: string;
    /**
     * Preferred report path prefix.
     */
    readonly s3Prefix: string;
    /**
     * Region of customer S3 bucket.
     */
    readonly s3Region: string;
    /**
     * The frequency on which report data are measured and displayed.
     */
    readonly timeUnit: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
