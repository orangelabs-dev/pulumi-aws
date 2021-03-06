// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an AWS Config Rule.
 * 
 * > **Note:** Config Rule requires an existing [Configuration Recorder](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `dependsOn` is recommended (as shown below) to avoid race conditions.
 * 
 * ## Example Usage
 * 
 * ### AWS Managed Rules
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const role = new aws.iam.Role("r", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "config.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const foo = new aws.cfg.Recorder("foo", {
 *     roleArn: role.arn,
 * });
 * const rule = new aws.cfg.Rule("r", {
 *     source: {
 *         owner: "AWS",
 *         sourceIdentifier: "S3_BUCKET_VERSIONING_ENABLED",
 *     },
 * }, {dependsOn: [foo]});
 * const rolePolicy = new aws.iam.RolePolicy("p", {
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *   	{
 *   		"Action": "config:Put*",
 *   		"Effect": "Allow",
 *   		"Resource": "*"
 * 
 *   	}
 *   ]
 * }
 * `,
 *     role: role.id,
 * });
 * ```
 * 
 * ### Custom Rules
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleRecorder = new aws.cfg.Recorder("example", {});
 * const exampleFunction = new aws.lambda.Function("example", {});
 * const examplePermission = new aws.lambda.Permission("example", {
 *     action: "lambda:InvokeFunction",
 *     function: exampleFunction.arn,
 *     principal: "config.amazonaws.com",
 * });
 * const exampleRule = new aws.cfg.Rule("example", {
 *     source: {
 *         owner: "CUSTOM_LAMBDA",
 *         sourceIdentifier: exampleFunction.arn,
 *     },
 * }, {dependsOn: [exampleRecorder, examplePermission]});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_config_rule.html.markdown.
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cfg/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * The ARN of the config rule
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of the rule
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A string in JSON format that is passed to the AWS Config rule Lambda function.
     */
    public readonly inputParameters!: pulumi.Output<string | undefined>;
    /**
     * The frequency that you want AWS Config to run evaluations for a rule that
     * is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
     */
    public readonly maximumExecutionFrequency!: pulumi.Output<string | undefined>;
    /**
     * The name of the rule
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the config rule
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * Scope defines which resources can trigger an evaluation for the rule as documented below.
     */
    public readonly scope!: pulumi.Output<outputs.cfg.RuleScope | undefined>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    public readonly source!: pulumi.Output<outputs.cfg.RuleSource>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["inputParameters"] = state ? state.inputParameters : undefined;
            inputs["maximumExecutionFrequency"] = state ? state.maximumExecutionFrequency : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ruleId"] = state ? state.ruleId : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["inputParameters"] = args ? args.inputParameters : undefined;
            inputs["maximumExecutionFrequency"] = args ? args.maximumExecutionFrequency : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ruleId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Rule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * The ARN of the config rule
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Description of the rule
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A string in JSON format that is passed to the AWS Config rule Lambda function.
     */
    readonly inputParameters?: pulumi.Input<string>;
    /**
     * The frequency that you want AWS Config to run evaluations for a rule that
     * is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
     */
    readonly maximumExecutionFrequency?: pulumi.Input<string>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the config rule
     */
    readonly ruleId?: pulumi.Input<string>;
    /**
     * Scope defines which resources can trigger an evaluation for the rule as documented below.
     */
    readonly scope?: pulumi.Input<inputs.cfg.RuleScope>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    readonly source?: pulumi.Input<inputs.cfg.RuleSource>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * Description of the rule
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A string in JSON format that is passed to the AWS Config rule Lambda function.
     */
    readonly inputParameters?: pulumi.Input<string>;
    /**
     * The frequency that you want AWS Config to run evaluations for a rule that
     * is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
     */
    readonly maximumExecutionFrequency?: pulumi.Input<string>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Scope defines which resources can trigger an evaluation for the rule as documented below.
     */
    readonly scope?: pulumi.Input<inputs.cfg.RuleScope>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    readonly source: pulumi.Input<inputs.cfg.RuleSource>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
