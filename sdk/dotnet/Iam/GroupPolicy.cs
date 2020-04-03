// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    /// <summary>
    /// Provides an IAM policy attached to a group.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_group_policy.html.markdown.
    /// </summary>
    public partial class GroupPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The IAM group to attach to the policy.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The name of the policy. If omitted, this provider will
        /// assign a random, unique name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a GroupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupPolicy(string name, GroupPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/groupPolicy:GroupPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GroupPolicy(string name, Input<string> id, GroupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/groupPolicy:GroupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupPolicy Get(string name, Input<string> id, GroupPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupPolicy(name, id, state, options);
        }
    }

    public sealed class GroupPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IAM group to attach to the policy.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The name of the policy. If omitted, this provider will
        /// assign a random, unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public GroupPolicyArgs()
        {
        }
    }

    public sealed class GroupPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IAM group to attach to the policy.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name of the policy. If omitted, this provider will
        /// assign a random, unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public GroupPolicyState()
        {
        }
    }
}
