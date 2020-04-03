// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource for managing the main routing table of a VPC.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const mainRouteTableAssociation = new aws.ec2.MainRouteTableAssociation("a", {
 *     routeTableId: aws_route_table_bar.id,
 *     vpcId: aws_vpc_foo.id,
 * });
 * ```
 * 
 * ## Notes
 * 
 * On VPC creation, the AWS API always creates an initial Main Route Table. This
 * resource records the ID of that Route Table under `originalRouteTableId`.
 * The "Delete" action for a `mainRouteTableAssociation` consists of resetting
 * this original table as the Main Route Table for the VPC. You'll see this
 * additional Route Table in the AWS console; it must remain intact in order for
 * the `mainRouteTableAssociation` delete to work properly.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/main_route_table_association.html.markdown.
 */
export class MainRouteTableAssociation extends pulumi.CustomResource {
    /**
     * Get an existing MainRouteTableAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MainRouteTableAssociationState, opts?: pulumi.CustomResourceOptions): MainRouteTableAssociation {
        return new MainRouteTableAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/mainRouteTableAssociation:MainRouteTableAssociation';

    /**
     * Returns true if the given object is an instance of MainRouteTableAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MainRouteTableAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MainRouteTableAssociation.__pulumiType;
    }

    /**
     * Used internally, see __Notes__ below
     */
    public /*out*/ readonly originalRouteTableId!: pulumi.Output<string>;
    /**
     * The ID of the Route Table to set as the new
     * main route table for the target VPC
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * The ID of the VPC whose main route table should be set
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a MainRouteTableAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MainRouteTableAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MainRouteTableAssociationArgs | MainRouteTableAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MainRouteTableAssociationState | undefined;
            inputs["originalRouteTableId"] = state ? state.originalRouteTableId : undefined;
            inputs["routeTableId"] = state ? state.routeTableId : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as MainRouteTableAssociationArgs | undefined;
            if (!args || args.routeTableId === undefined) {
                throw new Error("Missing required property 'routeTableId'");
            }
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["routeTableId"] = args ? args.routeTableId : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["originalRouteTableId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MainRouteTableAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MainRouteTableAssociation resources.
 */
export interface MainRouteTableAssociationState {
    /**
     * Used internally, see __Notes__ below
     */
    readonly originalRouteTableId?: pulumi.Input<string>;
    /**
     * The ID of the Route Table to set as the new
     * main route table for the target VPC
     */
    readonly routeTableId?: pulumi.Input<string>;
    /**
     * The ID of the VPC whose main route table should be set
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MainRouteTableAssociation resource.
 */
export interface MainRouteTableAssociationArgs {
    /**
     * The ID of the Route Table to set as the new
     * main route table for the target VPC
     */
    readonly routeTableId: pulumi.Input<string>;
    /**
     * The ID of the VPC whose main route table should be set
     */
    readonly vpcId: pulumi.Input<string>;
}
