// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApplicationLoadBalancing
{
    /// <summary>
    /// Provides a Load Balancer Listener Rule resource.
    /// 
    /// &gt; **Note:** `aws.alb.ListenerRule` is known as `aws.lb.ListenerRule`. The functionality is identical.
    /// </summary>
    public partial class ListenerRule : Pulumi.CustomResource
    {
        /// <summary>
        /// An Action block. Action blocks are documented below.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.ListenerRuleAction>> Actions { get; private set; } = null!;

        /// <summary>
        /// The ARN of the rule (matches `id`)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.ListenerRuleCondition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// The ARN of the listener to which to attach the rule.
        /// </summary>
        [Output("listenerArn")]
        public Output<string> ListenerArn { get; private set; } = null!;

        /// <summary>
        /// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;


        /// <summary>
        /// Create a ListenerRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ListenerRule(string name, ListenerRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:applicationloadbalancing/listenerRule:ListenerRule", name, args ?? new ListenerRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ListenerRule(string name, Input<string> id, ListenerRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:applicationloadbalancing/listenerRule:ListenerRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ListenerRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ListenerRule Get(string name, Input<string> id, ListenerRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ListenerRule(name, id, state, options);
        }
    }

    public sealed class ListenerRuleArgs : Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.ListenerRuleActionArgs>? _actions;

        /// <summary>
        /// An Action block. Action blocks are documented below.
        /// </summary>
        public InputList<Inputs.ListenerRuleActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.ListenerRuleActionArgs>());
            set => _actions = value;
        }

        [Input("conditions", required: true)]
        private InputList<Inputs.ListenerRuleConditionArgs>? _conditions;

        /// <summary>
        /// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
        /// </summary>
        public InputList<Inputs.ListenerRuleConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ListenerRuleConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The ARN of the listener to which to attach the rule.
        /// </summary>
        [Input("listenerArn", required: true)]
        public Input<string> ListenerArn { get; set; } = null!;

        /// <summary>
        /// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        public ListenerRuleArgs()
        {
        }
    }

    public sealed class ListenerRuleState : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.ListenerRuleActionGetArgs>? _actions;

        /// <summary>
        /// An Action block. Action blocks are documented below.
        /// </summary>
        public InputList<Inputs.ListenerRuleActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.ListenerRuleActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// The ARN of the rule (matches `id`)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("conditions")]
        private InputList<Inputs.ListenerRuleConditionGetArgs>? _conditions;

        /// <summary>
        /// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
        /// </summary>
        public InputList<Inputs.ListenerRuleConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ListenerRuleConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The ARN of the listener to which to attach the rule.
        /// </summary>
        [Input("listenerArn")]
        public Input<string>? ListenerArn { get; set; }

        /// <summary>
        /// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        public ListenerRuleState()
        {
        }
    }
}
