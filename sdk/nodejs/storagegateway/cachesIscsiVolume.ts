// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AWS Storage Gateway cached iSCSI volume.
 * 
 * > **NOTE:** The gateway must have cache added (e.g. via the [`aws.storagegateway.Cache`](https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html) resource) before creating volumes otherwise the Storage Gateway API will return an error.
 * 
 * > **NOTE:** The gateway must have an upload buffer added (e.g. via the [`aws.storagegateway.UploadBuffer`](https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html) resource) before the volume is operational to clients, however the Storage Gateway API will allow volume creation without error in that case and return volume status as `UPLOAD BUFFER NOT CONFIGURED`.
 * 
 * ## Example Usage
 * 
 * > **NOTE:** These examples are referencing the [`aws.storagegateway.Cache`](https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html) resource `gatewayArn` attribute to ensure this provider properly adds cache before creating the volume. If you are not using this method, you may need to declare an expicit dependency (e.g. via `dependsOn = ["aws_storagegateway_cache.example"]`) to ensure proper ordering.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/storagegateway_cached_iscsi_volume.html.markdown.
 */
export class CachesIscsiVolume extends pulumi.CustomResource {
    /**
     * Get an existing CachesIscsiVolume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CachesIscsiVolumeState, opts?: pulumi.CustomResourceOptions): CachesIscsiVolume {
        return new CachesIscsiVolume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume';

    /**
     * Returns true if the given object is an instance of CachesIscsiVolume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CachesIscsiVolume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CachesIscsiVolume.__pulumiType;
    }

    /**
     * Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Whether mutual CHAP is enabled for the iSCSI target.
     */
    public /*out*/ readonly chapEnabled!: pulumi.Output<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    public readonly gatewayArn!: pulumi.Output<string>;
    /**
     * Logical disk number.
     */
    public /*out*/ readonly lunNumber!: pulumi.Output<number>;
    /**
     * The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;
    /**
     * The port used to communicate with iSCSI targets.
     */
    public /*out*/ readonly networkInterfacePort!: pulumi.Output<number>;
    /**
     * The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
     */
    public readonly sourceVolumeArn!: pulumi.Output<string | undefined>;
    /**
     * Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
     */
    public /*out*/ readonly targetArn!: pulumi.Output<string>;
    /**
     * The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
     */
    public readonly targetName!: pulumi.Output<string>;
    /**
     * Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
     */
    public /*out*/ readonly volumeArn!: pulumi.Output<string>;
    /**
     * Volume ID, e.g. `vol-12345678`.
     */
    public /*out*/ readonly volumeId!: pulumi.Output<string>;
    /**
     * The size of the volume in bytes.
     */
    public readonly volumeSizeInBytes!: pulumi.Output<number>;

    /**
     * Create a CachesIscsiVolume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CachesIscsiVolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CachesIscsiVolumeArgs | CachesIscsiVolumeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CachesIscsiVolumeState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["chapEnabled"] = state ? state.chapEnabled : undefined;
            inputs["gatewayArn"] = state ? state.gatewayArn : undefined;
            inputs["lunNumber"] = state ? state.lunNumber : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            inputs["networkInterfacePort"] = state ? state.networkInterfacePort : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["sourceVolumeArn"] = state ? state.sourceVolumeArn : undefined;
            inputs["targetArn"] = state ? state.targetArn : undefined;
            inputs["targetName"] = state ? state.targetName : undefined;
            inputs["volumeArn"] = state ? state.volumeArn : undefined;
            inputs["volumeId"] = state ? state.volumeId : undefined;
            inputs["volumeSizeInBytes"] = state ? state.volumeSizeInBytes : undefined;
        } else {
            const args = argsOrState as CachesIscsiVolumeArgs | undefined;
            if (!args || args.gatewayArn === undefined) {
                throw new Error("Missing required property 'gatewayArn'");
            }
            if (!args || args.networkInterfaceId === undefined) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            if (!args || args.targetName === undefined) {
                throw new Error("Missing required property 'targetName'");
            }
            if (!args || args.volumeSizeInBytes === undefined) {
                throw new Error("Missing required property 'volumeSizeInBytes'");
            }
            inputs["gatewayArn"] = args ? args.gatewayArn : undefined;
            inputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
            inputs["sourceVolumeArn"] = args ? args.sourceVolumeArn : undefined;
            inputs["targetName"] = args ? args.targetName : undefined;
            inputs["volumeSizeInBytes"] = args ? args.volumeSizeInBytes : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["chapEnabled"] = undefined /*out*/;
            inputs["lunNumber"] = undefined /*out*/;
            inputs["networkInterfacePort"] = undefined /*out*/;
            inputs["targetArn"] = undefined /*out*/;
            inputs["volumeArn"] = undefined /*out*/;
            inputs["volumeId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CachesIscsiVolume.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CachesIscsiVolume resources.
 */
export interface CachesIscsiVolumeState {
    /**
     * Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Whether mutual CHAP is enabled for the iSCSI target.
     */
    readonly chapEnabled?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    readonly gatewayArn?: pulumi.Input<string>;
    /**
     * Logical disk number.
     */
    readonly lunNumber?: pulumi.Input<number>;
    /**
     * The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
     */
    readonly networkInterfaceId?: pulumi.Input<string>;
    /**
     * The port used to communicate with iSCSI targets.
     */
    readonly networkInterfacePort?: pulumi.Input<number>;
    /**
     * The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
     */
    readonly sourceVolumeArn?: pulumi.Input<string>;
    /**
     * Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
     */
    readonly targetArn?: pulumi.Input<string>;
    /**
     * The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
     */
    readonly targetName?: pulumi.Input<string>;
    /**
     * Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
     */
    readonly volumeArn?: pulumi.Input<string>;
    /**
     * Volume ID, e.g. `vol-12345678`.
     */
    readonly volumeId?: pulumi.Input<string>;
    /**
     * The size of the volume in bytes.
     */
    readonly volumeSizeInBytes?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CachesIscsiVolume resource.
 */
export interface CachesIscsiVolumeArgs {
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    readonly gatewayArn: pulumi.Input<string>;
    /**
     * The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
     */
    readonly networkInterfaceId: pulumi.Input<string>;
    /**
     * The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
     */
    readonly sourceVolumeArn?: pulumi.Input<string>;
    /**
     * The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
     */
    readonly targetName: pulumi.Input<string>;
    /**
     * The size of the volume in bytes.
     */
    readonly volumeSizeInBytes: pulumi.Input<number>;
}
