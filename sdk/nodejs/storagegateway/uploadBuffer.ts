// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an AWS Storage Gateway upload buffer.
 * 
 * > **NOTE:** The Storage Gateway API provides no method to remove an upload buffer disk. Destroying this resource does not perform any Storage Gateway actions.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.storagegateway.UploadBuffer("example", {
 *     diskId: aws_storagegateway_local_disk_example.id,
 *     gatewayArn: aws_storagegateway_gateway_example.arn,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/storagegateway_upload_buffer.html.markdown.
 */
export class UploadBuffer extends pulumi.CustomResource {
    /**
     * Get an existing UploadBuffer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UploadBufferState, opts?: pulumi.CustomResourceOptions): UploadBuffer {
        return new UploadBuffer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:storagegateway/uploadBuffer:UploadBuffer';

    /**
     * Returns true if the given object is an instance of UploadBuffer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UploadBuffer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UploadBuffer.__pulumiType;
    }

    /**
     * Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    public readonly gatewayArn!: pulumi.Output<string>;

    /**
     * Create a UploadBuffer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UploadBufferArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UploadBufferArgs | UploadBufferState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UploadBufferState | undefined;
            inputs["diskId"] = state ? state.diskId : undefined;
            inputs["gatewayArn"] = state ? state.gatewayArn : undefined;
        } else {
            const args = argsOrState as UploadBufferArgs | undefined;
            if (!args || args.diskId === undefined) {
                throw new Error("Missing required property 'diskId'");
            }
            if (!args || args.gatewayArn === undefined) {
                throw new Error("Missing required property 'gatewayArn'");
            }
            inputs["diskId"] = args ? args.diskId : undefined;
            inputs["gatewayArn"] = args ? args.gatewayArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UploadBuffer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UploadBuffer resources.
 */
export interface UploadBufferState {
    /**
     * Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
     */
    readonly diskId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    readonly gatewayArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UploadBuffer resource.
 */
export interface UploadBufferArgs {
    /**
     * Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
     */
    readonly diskId: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    readonly gatewayArn: pulumi.Input<string>;
}
