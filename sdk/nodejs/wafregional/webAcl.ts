// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ipset = new aws.wafregional.IpSet("ipset", {
 *     ipSetDescriptors: [{
 *         type: "IPV4",
 *         value: "192.0.7.0/24",
 *     }],
 *     name: "tfIPSet",
 * });
 * const wafrule = new aws.wafregional.Rule("wafrule", {
 *     metricName: "tfWAFRule",
 *     name: "tfWAFRule",
 *     predicates: [{
 *         dataId: ipset.id,
 *         negated: false,
 *         type: "IPMatch",
 *     }],
 * });
 * const wafacl = new aws.wafregional.WebAcl("wafacl", {
 *     defaultAction: {
 *         type: "ALLOW",
 *     },
 *     metricName: "tfWebACL",
 *     name: "tfWebACL",
 *     rules: [{
 *         action: {
 *             type: "BLOCK",
 *         },
 *         priority: 1,
 *         ruleId: wafrule.id,
 *         type: "REGULAR",
 *     }],
 * });
 * ```
 * 
 * ## Nested Fields
 * 
 * ### `rule`
 * 
 * See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_ActivatedRule.html) for all details and supported values.
 * 
 * #### Arguments
 * 
 * * `action` - (Required) The action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`.
 * * `override_action` - (Required) Override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`.
 * * `priority` - (Required) Specifies the order in which the rules in a WebACL are evaluated.
 *   Rules with a lower value are evaluated before rules with a higher value.
 * * `rule_id` - (Required) ID of the associated WAF (Regional) rule (e.g. [`aws_wafregional_rule`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)). WAF (Global) rules cannot be used.
 * * `type` - (Optional) The rule type, either `REGULAR`, as defined by [Rule](http://docs.aws.amazon.com/waf/latest/APIReference/API_Rule.html), `RATE_BASED`, as defined by [RateBasedRule](http://docs.aws.amazon.com/waf/latest/APIReference/API_RateBasedRule.html), or `GROUP`, as defined by [RuleGroup](https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html). The default is REGULAR. If you add a RATE_BASED rule, you need to set `type` as `RATE_BASED`. If you add a GROUP rule, you need to set `type` as `GROUP`.
 * 
 * ### `default_action` / `action`
 * 
 * #### Arguments
 * 
 * * `type` - (Required) Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule.
 *   e.g. `ALLOW`, `BLOCK` or `COUNT`
 */
export class WebAcl extends pulumi.CustomResource {
    /**
     * Get an existing WebAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAclState, opts?: pulumi.CustomResourceOptions): WebAcl {
        return new WebAcl(name, <any>state, { ...opts, id: id });
    }

    /**
     * The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
     */
    public readonly defaultAction: pulumi.Output<{ type: string }>;
    /**
     * The name or description for the Amazon CloudWatch metric of this web ACL.
     */
    public readonly metricName: pulumi.Output<string>;
    /**
     * The name or description of the web ACL.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The rules to associate with the web ACL and the settings for each rule.
     */
    public readonly rules: pulumi.Output<{ action?: { type: string }, overrideAction?: { type: string }, priority: number, ruleId: string, type?: string }[] | undefined>;

    /**
     * Create a WebAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAclArgs | WebAclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: WebAclState = argsOrState as WebAclState | undefined;
            inputs["defaultAction"] = state ? state.defaultAction : undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as WebAclArgs | undefined;
            if (!args || args.defaultAction === undefined) {
                throw new Error("Missing required property 'defaultAction'");
            }
            if (!args || args.metricName === undefined) {
                throw new Error("Missing required property 'metricName'");
            }
            inputs["defaultAction"] = args ? args.defaultAction : undefined;
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rules"] = args ? args.rules : undefined;
        }
        super("aws:wafregional/webAcl:WebAcl", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebAcl resources.
 */
export interface WebAclState {
    /**
     * The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
     */
    readonly defaultAction?: pulumi.Input<{ type: pulumi.Input<string> }>;
    /**
     * The name or description for the Amazon CloudWatch metric of this web ACL.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * The name or description of the web ACL.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The rules to associate with the web ACL and the settings for each rule.
     */
    readonly rules?: pulumi.Input<pulumi.Input<{ action?: pulumi.Input<{ type: pulumi.Input<string> }>, overrideAction?: pulumi.Input<{ type: pulumi.Input<string> }>, priority: pulumi.Input<number>, ruleId: pulumi.Input<string>, type?: pulumi.Input<string> }>[]>;
}

/**
 * The set of arguments for constructing a WebAcl resource.
 */
export interface WebAclArgs {
    /**
     * The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
     */
    readonly defaultAction: pulumi.Input<{ type: pulumi.Input<string> }>;
    /**
     * The name or description for the Amazon CloudWatch metric of this web ACL.
     */
    readonly metricName: pulumi.Input<string>;
    /**
     * The name or description of the web ACL.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The rules to associate with the web ACL and the settings for each rule.
     */
    readonly rules?: pulumi.Input<pulumi.Input<{ action?: pulumi.Input<{ type: pulumi.Input<string> }>, overrideAction?: pulumi.Input<{ type: pulumi.Input<string> }>, priority: pulumi.Input<number>, ruleId: pulumi.Input<string>, type?: pulumi.Input<string> }>[]>;
}
