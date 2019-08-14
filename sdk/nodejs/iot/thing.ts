// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages an AWS IoT Thing.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_thing.html.markdown.
 */
export class Thing extends pulumi.CustomResource {
    /**
     * Get an existing Thing resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ThingState, opts?: pulumi.CustomResourceOptions): Thing {
        return new Thing(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iot/thing:Thing';

    /**
     * Returns true if the given object is an instance of Thing.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Thing {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Thing.__pulumiType;
    }

    /**
     * The ARN of the thing.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Map of attributes of the thing.
     */
    public readonly attributes!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The default client ID.
     */
    public /*out*/ readonly defaultClientId!: pulumi.Output<string>;
    /**
     * The name of the thing.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The thing type name.
     */
    public readonly thingTypeName!: pulumi.Output<string | undefined>;
    /**
     * The current version of the thing record in the registry.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a Thing resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ThingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ThingArgs | ThingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ThingState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["attributes"] = state ? state.attributes : undefined;
            inputs["defaultClientId"] = state ? state.defaultClientId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["thingTypeName"] = state ? state.thingTypeName : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ThingArgs | undefined;
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["thingTypeName"] = args ? args.thingTypeName : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["defaultClientId"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Thing.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Thing resources.
 */
export interface ThingState {
    /**
     * The ARN of the thing.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Map of attributes of the thing.
     */
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * The default client ID.
     */
    readonly defaultClientId?: pulumi.Input<string>;
    /**
     * The name of the thing.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The thing type name.
     */
    readonly thingTypeName?: pulumi.Input<string>;
    /**
     * The current version of the thing record in the registry.
     */
    readonly version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Thing resource.
 */
export interface ThingArgs {
    /**
     * Map of attributes of the thing.
     */
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the thing.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The thing type name.
     */
    readonly thingTypeName?: pulumi.Input<string>;
}
