// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Resource Access Manager (RAM) Resource Association.
 * 
 * > *NOTE:* Certain AWS resources (e.g. EC2 Subnets) can only be shared in an AWS account that is a member of an AWS Organizations organization with organization-wide Resource Access Manager functionality enabled. See the [Resource Access Manager User Guide](https://docs.aws.amazon.com/ram/latest/userguide/what-is.html) and AWS service specific documentation for additional information.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_resource_association.html.markdown.
 */
export class ResourceAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ResourceAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceAssociationState, opts?: pulumi.CustomResourceOptions): ResourceAssociation {
        return new ResourceAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ram/resourceAssociation:ResourceAssociation';

    /**
     * Returns true if the given object is an instance of ResourceAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceAssociation.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
     */
    public readonly resourceArn!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the RAM Resource Share.
     */
    public readonly resourceShareArn!: pulumi.Output<string>;

    /**
     * Create a ResourceAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceAssociationArgs | ResourceAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourceAssociationState | undefined;
            inputs["resourceArn"] = state ? state.resourceArn : undefined;
            inputs["resourceShareArn"] = state ? state.resourceShareArn : undefined;
        } else {
            const args = argsOrState as ResourceAssociationArgs | undefined;
            if (!args || args.resourceArn === undefined) {
                throw new Error("Missing required property 'resourceArn'");
            }
            if (!args || args.resourceShareArn === undefined) {
                throw new Error("Missing required property 'resourceShareArn'");
            }
            inputs["resourceArn"] = args ? args.resourceArn : undefined;
            inputs["resourceShareArn"] = args ? args.resourceShareArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResourceAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceAssociation resources.
 */
export interface ResourceAssociationState {
    /**
     * Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
     */
    readonly resourceArn?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the RAM Resource Share.
     */
    readonly resourceShareArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceAssociation resource.
 */
export interface ResourceAssociationArgs {
    /**
     * Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
     */
    readonly resourceArn: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the RAM Resource Share.
     */
    readonly resourceShareArn: pulumi.Input<string>;
}
