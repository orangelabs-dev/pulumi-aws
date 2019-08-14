// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieve information about a Secrets Manager secret version, including its secret value. To retrieve secret metadata, see the [`aws.secretsmanager.Secret` data source](https://www.terraform.io/docs/providers/aws/d/secretsmanager_secret.html).
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/secretsmanager_secret_version.html.markdown.
 */
export function getSecretVersion(args: GetSecretVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretVersionResult> & GetSecretVersionResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetSecretVersionResult> = pulumi.runtime.invoke("aws:secretsmanager/getSecretVersion:getSecretVersion", {
        "secretId": args.secretId,
        "versionId": args.versionId,
        "versionStage": args.versionStage,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getSecretVersion.
 */
export interface GetSecretVersionArgs {
    /**
     * Specifies the secret containing the version that you want to retrieve. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret.
     */
    readonly secretId: string;
    /**
     * Specifies the unique identifier of the version of the secret that you want to retrieve. Overrides `versionStage`.
     */
    readonly versionId?: string;
    /**
     * Specifies the secret version that you want to retrieve by the staging label attached to the version. Defaults to `AWSCURRENT`.
     */
    readonly versionStage?: string;
}

/**
 * A collection of values returned by getSecretVersion.
 */
export interface GetSecretVersionResult {
    /**
     * The ARN of the secret.
     */
    readonly arn: string;
    /**
     * The decrypted part of the protected secret information that was originally provided as a binary. Base64 encoded.
     */
    readonly secretBinary: string;
    readonly secretId: string;
    /**
     * The decrypted part of the protected secret information that was originally provided as a string.
     */
    readonly secretString: string;
    /**
     * The unique identifier of this version of the secret.
     */
    readonly versionId: string;
    readonly versionStage?: string;
    readonly versionStages: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
