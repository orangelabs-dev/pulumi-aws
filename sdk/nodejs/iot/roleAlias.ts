// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an IoT role alias.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_role_alias.html.markdown.
 */
export class RoleAlias extends pulumi.CustomResource {
    /**
     * Get an existing RoleAlias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleAliasState, opts?: pulumi.CustomResourceOptions): RoleAlias {
        return new RoleAlias(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iot/roleAlias:RoleAlias';

    /**
     * Returns true if the given object is an instance of RoleAlias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleAlias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleAlias.__pulumiType;
    }

    /**
     * The name of the role alias.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * The ARN assigned by AWS to this role alias.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
     */
    public readonly credentialDuration!: pulumi.Output<number | undefined>;
    /**
     * The identity of the role to which the alias refers.
     */
    public readonly roleArn!: pulumi.Output<string>;

    /**
     * Create a RoleAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleAliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleAliasArgs | RoleAliasState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RoleAliasState | undefined;
            inputs["alias"] = state ? state.alias : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["credentialDuration"] = state ? state.credentialDuration : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
        } else {
            const args = argsOrState as RoleAliasArgs | undefined;
            if (!args || args.alias === undefined) {
                throw new Error("Missing required property 'alias'");
            }
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["alias"] = args ? args.alias : undefined;
            inputs["credentialDuration"] = args ? args.credentialDuration : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RoleAlias.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RoleAlias resources.
 */
export interface RoleAliasState {
    /**
     * The name of the role alias.
     */
    readonly alias?: pulumi.Input<string>;
    /**
     * The ARN assigned by AWS to this role alias.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
     */
    readonly credentialDuration?: pulumi.Input<number>;
    /**
     * The identity of the role to which the alias refers.
     */
    readonly roleArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RoleAlias resource.
 */
export interface RoleAliasArgs {
    /**
     * The name of the role alias.
     */
    readonly alias: pulumi.Input<string>;
    /**
     * The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
     */
    readonly credentialDuration?: pulumi.Input<number>;
    /**
     * The identity of the role to which the alias refers.
     */
    readonly roleArn: pulumi.Input<string>;
}
