// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AppSync Function.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_function.html.markdown.
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appsync/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * The ID of the associated AppSync API.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * The ARN of the Function object.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Function DataSource name.
     */
    public readonly dataSource!: pulumi.Output<string>;
    /**
     * The Function description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A unique ID representing the Function object.
     */
    public /*out*/ readonly functionId!: pulumi.Output<string>;
    /**
     * The version of the request mapping template. Currently the supported value is `2018-05-29`.
     */
    public readonly functionVersion!: pulumi.Output<string | undefined>;
    /**
     * The Function name. The function name does not have to be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
     */
    public readonly requestMappingTemplate!: pulumi.Output<string>;
    /**
     * The Function response mapping template.
     */
    public readonly responseMappingTemplate!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FunctionState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["dataSource"] = state ? state.dataSource : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["functionId"] = state ? state.functionId : undefined;
            inputs["functionVersion"] = state ? state.functionVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["requestMappingTemplate"] = state ? state.requestMappingTemplate : undefined;
            inputs["responseMappingTemplate"] = state ? state.responseMappingTemplate : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.dataSource === undefined) {
                throw new Error("Missing required property 'dataSource'");
            }
            if (!args || args.requestMappingTemplate === undefined) {
                throw new Error("Missing required property 'requestMappingTemplate'");
            }
            if (!args || args.responseMappingTemplate === undefined) {
                throw new Error("Missing required property 'responseMappingTemplate'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["dataSource"] = args ? args.dataSource : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["functionVersion"] = args ? args.functionVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["requestMappingTemplate"] = args ? args.requestMappingTemplate : undefined;
            inputs["responseMappingTemplate"] = args ? args.responseMappingTemplate : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["functionId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Function.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * The ID of the associated AppSync API.
     */
    readonly apiId?: pulumi.Input<string>;
    /**
     * The ARN of the Function object.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The Function DataSource name.
     */
    readonly dataSource?: pulumi.Input<string>;
    /**
     * The Function description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique ID representing the Function object.
     */
    readonly functionId?: pulumi.Input<string>;
    /**
     * The version of the request mapping template. Currently the supported value is `2018-05-29`.
     */
    readonly functionVersion?: pulumi.Input<string>;
    /**
     * The Function name. The function name does not have to be unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
     */
    readonly requestMappingTemplate?: pulumi.Input<string>;
    /**
     * The Function response mapping template.
     */
    readonly responseMappingTemplate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * The ID of the associated AppSync API.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * The Function DataSource name.
     */
    readonly dataSource: pulumi.Input<string>;
    /**
     * The Function description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The version of the request mapping template. Currently the supported value is `2018-05-29`.
     */
    readonly functionVersion?: pulumi.Input<string>;
    /**
     * The Function name. The function name does not have to be unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
     */
    readonly requestMappingTemplate: pulumi.Input<string>;
    /**
     * The Function response mapping template.
     */
    readonly responseMappingTemplate: pulumi.Input<string>;
}
