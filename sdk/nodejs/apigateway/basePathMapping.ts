// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./restApi";

/**
 * Connects a custom domain name registered via `aws.apigateway.DomainName`
 * with a deployed API so that its methods can be called via the
 * custom domain name.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_base_path_mapping.html.markdown.
 */
export class BasePathMapping extends pulumi.CustomResource {
    /**
     * Get an existing BasePathMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BasePathMappingState, opts?: pulumi.CustomResourceOptions): BasePathMapping {
        return new BasePathMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/basePathMapping:BasePathMapping';

    /**
     * Returns true if the given object is an instance of BasePathMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BasePathMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BasePathMapping.__pulumiType;
    }

    /**
     * The id of the API to connect.
     */
    public readonly restApi!: pulumi.Output<RestApi>;
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    public readonly basePath!: pulumi.Output<string | undefined>;
    /**
     * The already-registered domain name to connect the API to.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    public readonly stageName!: pulumi.Output<string | undefined>;

    /**
     * Create a BasePathMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasePathMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BasePathMappingArgs | BasePathMappingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BasePathMappingState | undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["basePath"] = state ? state.basePath : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["stageName"] = state ? state.stageName : undefined;
        } else {
            const args = argsOrState as BasePathMappingArgs | undefined;
            if (!args || args.restApi === undefined) {
                throw new Error("Missing required property 'restApi'");
            }
            if (!args || args.domainName === undefined) {
                throw new Error("Missing required property 'domainName'");
            }
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["basePath"] = args ? args.basePath : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["stageName"] = args ? args.stageName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BasePathMapping.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BasePathMapping resources.
 */
export interface BasePathMappingState {
    /**
     * The id of the API to connect.
     */
    readonly restApi?: pulumi.Input<RestApi>;
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    readonly basePath?: pulumi.Input<string>;
    /**
     * The already-registered domain name to connect the API to.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    readonly stageName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BasePathMapping resource.
 */
export interface BasePathMappingArgs {
    /**
     * The id of the API to connect.
     */
    readonly restApi: pulumi.Input<RestApi>;
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    readonly basePath?: pulumi.Input<string>;
    /**
     * The already-registered domain name to connect the API to.
     */
    readonly domainName: pulumi.Input<string>;
    /**
     * The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    readonly stageName?: pulumi.Input<string>;
}
