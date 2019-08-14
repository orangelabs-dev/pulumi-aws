// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a AWS Transfer User SSH Key resource.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/transfer_ssh_key.html.markdown.
 */
export class SshKey extends pulumi.CustomResource {
    /**
     * Get an existing SshKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SshKeyState, opts?: pulumi.CustomResourceOptions): SshKey {
        return new SshKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:transfer/sshKey:SshKey';

    /**
     * Returns true if the given object is an instance of SshKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SshKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SshKey.__pulumiType;
    }

    /**
     * The public key portion of an SSH key pair.
     */
    public readonly body!: pulumi.Output<string>;
    /**
     * The Server ID of the Transfer Server (e.g. `s-12345678`)
     */
    public readonly serverId!: pulumi.Output<string>;
    /**
     * The name of the user account that is assigned to one or more servers.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a SshKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SshKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SshKeyArgs | SshKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SshKeyState | undefined;
            inputs["body"] = state ? state.body : undefined;
            inputs["serverId"] = state ? state.serverId : undefined;
            inputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as SshKeyArgs | undefined;
            if (!args || args.body === undefined) {
                throw new Error("Missing required property 'body'");
            }
            if (!args || args.serverId === undefined) {
                throw new Error("Missing required property 'serverId'");
            }
            if (!args || args.userName === undefined) {
                throw new Error("Missing required property 'userName'");
            }
            inputs["body"] = args ? args.body : undefined;
            inputs["serverId"] = args ? args.serverId : undefined;
            inputs["userName"] = args ? args.userName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SshKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SshKey resources.
 */
export interface SshKeyState {
    /**
     * The public key portion of an SSH key pair.
     */
    readonly body?: pulumi.Input<string>;
    /**
     * The Server ID of the Transfer Server (e.g. `s-12345678`)
     */
    readonly serverId?: pulumi.Input<string>;
    /**
     * The name of the user account that is assigned to one or more servers.
     */
    readonly userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SshKey resource.
 */
export interface SshKeyArgs {
    /**
     * The public key portion of an SSH key pair.
     */
    readonly body: pulumi.Input<string>;
    /**
     * The Server ID of the Transfer Server (e.g. `s-12345678`)
     */
    readonly serverId: pulumi.Input<string>;
    /**
     * The name of the user account that is assigned to one or more servers.
     */
    readonly userName: pulumi.Input<string>;
}
