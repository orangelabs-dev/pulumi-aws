// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a conditional forwarder for managed Microsoft AD in AWS Directory Service.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/directory_service_conditional_forwarder.html.markdown.
 */
export class ConditionalForwader extends pulumi.CustomResource {
    /**
     * Get an existing ConditionalForwader resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConditionalForwaderState, opts?: pulumi.CustomResourceOptions): ConditionalForwader {
        return new ConditionalForwader(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directoryservice/conditionalForwader:ConditionalForwader';

    /**
     * Returns true if the given object is an instance of ConditionalForwader.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConditionalForwader {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConditionalForwader.__pulumiType;
    }

    /**
     * The id of directory.
     */
    public readonly directoryId!: pulumi.Output<string>;
    /**
     * A list of forwarder IP addresses.
     */
    public readonly dnsIps!: pulumi.Output<string[]>;
    /**
     * The fully qualified domain name of the remote domain for which forwarders will be used.
     */
    public readonly remoteDomainName!: pulumi.Output<string>;

    /**
     * Create a ConditionalForwader resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConditionalForwaderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConditionalForwaderArgs | ConditionalForwaderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConditionalForwaderState | undefined;
            inputs["directoryId"] = state ? state.directoryId : undefined;
            inputs["dnsIps"] = state ? state.dnsIps : undefined;
            inputs["remoteDomainName"] = state ? state.remoteDomainName : undefined;
        } else {
            const args = argsOrState as ConditionalForwaderArgs | undefined;
            if (!args || args.directoryId === undefined) {
                throw new Error("Missing required property 'directoryId'");
            }
            if (!args || args.dnsIps === undefined) {
                throw new Error("Missing required property 'dnsIps'");
            }
            if (!args || args.remoteDomainName === undefined) {
                throw new Error("Missing required property 'remoteDomainName'");
            }
            inputs["directoryId"] = args ? args.directoryId : undefined;
            inputs["dnsIps"] = args ? args.dnsIps : undefined;
            inputs["remoteDomainName"] = args ? args.remoteDomainName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ConditionalForwader.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConditionalForwader resources.
 */
export interface ConditionalForwaderState {
    /**
     * The id of directory.
     */
    readonly directoryId?: pulumi.Input<string>;
    /**
     * A list of forwarder IP addresses.
     */
    readonly dnsIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fully qualified domain name of the remote domain for which forwarders will be used.
     */
    readonly remoteDomainName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConditionalForwader resource.
 */
export interface ConditionalForwaderArgs {
    /**
     * The id of directory.
     */
    readonly directoryId: pulumi.Input<string>;
    /**
     * A list of forwarder IP addresses.
     */
    readonly dnsIps: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fully qualified domain name of the remote domain for which forwarders will be used.
     */
    readonly remoteDomainName: pulumi.Input<string>;
}
