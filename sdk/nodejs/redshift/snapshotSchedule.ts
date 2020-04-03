// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const defaultSnapshotSchedule = new aws.redshift.SnapshotSchedule("default", {
 *     definitions: ["rate(12 hours)"],
 *     identifier: "tf-redshift-snapshot-schedule",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/redshift_snapshot_schedule.html.markdown.
 */
export class SnapshotSchedule extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotSchedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotScheduleState, opts?: pulumi.CustomResourceOptions): SnapshotSchedule {
        return new SnapshotSchedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:redshift/snapshotSchedule:SnapshotSchedule';

    /**
     * Returns true if the given object is an instance of SnapshotSchedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnapshotSchedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnapshotSchedule.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
     */
    public readonly definitions!: pulumi.Output<string[]>;
    /**
     * The description of the snapshot schedule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * Creates a unique
     * identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    public readonly identifierPrefix!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a SnapshotSchedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotScheduleArgs | SnapshotScheduleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SnapshotScheduleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["definitions"] = state ? state.definitions : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["identifier"] = state ? state.identifier : undefined;
            inputs["identifierPrefix"] = state ? state.identifierPrefix : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SnapshotScheduleArgs | undefined;
            if (!args || args.definitions === undefined) {
                throw new Error("Missing required property 'definitions'");
            }
            inputs["definitions"] = args ? args.definitions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["identifierPrefix"] = args ? args.identifierPrefix : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SnapshotSchedule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnapshotSchedule resources.
 */
export interface SnapshotScheduleState {
    readonly arn?: pulumi.Input<string>;
    /**
     * The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
     */
    readonly definitions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the snapshot schedule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
     */
    readonly identifier?: pulumi.Input<string>;
    /**
     * Creates a unique
     * identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    readonly identifierPrefix?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a SnapshotSchedule resource.
 */
export interface SnapshotScheduleArgs {
    /**
     * The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
     */
    readonly definitions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the snapshot schedule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
     */
    readonly identifier?: pulumi.Input<string>;
    /**
     * Creates a unique
     * identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    readonly identifierPrefix?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
