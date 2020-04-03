// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache
{
    /// <summary>
    /// Provides an ElastiCache Subnet Group resource.
    /// 
    /// &gt; **NOTE:** ElastiCache Subnet Groups are only for use when working with an
    /// ElastiCache cluster **inside** of a VPC. If you are on EC2 Classic, see the
    /// ElastiCache Security Group resource.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elasticache_subnet_group.html.markdown.
    /// </summary>
    public partial class SubnetGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Description for the cache subnet group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Name for the cache subnet group. Elasticache converts this name to lowercase.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of VPC Subnet IDs for the cache subnet group
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;


        /// <summary>
        /// Create a SubnetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubnetGroup(string name, SubnetGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticache/subnetGroup:SubnetGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SubnetGroup(string name, Input<string> id, SubnetGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticache/subnetGroup:SubnetGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SubnetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubnetGroup Get(string name, Input<string> id, SubnetGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SubnetGroup(name, id, state, options);
        }
    }

    public sealed class SubnetGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the cache subnet group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name for the cache subnet group. Elasticache converts this name to lowercase.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// List of VPC Subnet IDs for the cache subnet group
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public SubnetGroupArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class SubnetGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the cache subnet group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name for the cache subnet group. Elasticache converts this name to lowercase.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// List of VPC Subnet IDs for the cache subnet group
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public SubnetGroupState()
        {
            Description = "Managed by Pulumi";
        }
    }
}
