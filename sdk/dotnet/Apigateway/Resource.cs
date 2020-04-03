// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides an API Gateway Resource.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_resource.html.markdown.
    /// </summary>
    public partial class Resource : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the parent API resource
        /// </summary>
        [Output("parentId")]
        public Output<string> ParentId { get; private set; } = null!;

        /// <summary>
        /// The complete path for this API resource, including all parent paths.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// The last path segment of this API resource.
        /// </summary>
        [Output("pathPart")]
        public Output<string> PathPart { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Output("restApi")]
        public Output<string> RestApi { get; private set; } = null!;


        /// <summary>
        /// Create a Resource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resource(string name, ResourceArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/resource:Resource", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Resource(string name, Input<string> id, ResourceState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/resource:Resource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Resource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resource Get(string name, Input<string> id, ResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new Resource(name, id, state, options);
        }
    }

    public sealed class ResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the parent API resource
        /// </summary>
        [Input("parentId", required: true)]
        public Input<string> ParentId { get; set; } = null!;

        /// <summary>
        /// The last path segment of this API resource.
        /// </summary>
        [Input("pathPart", required: true)]
        public Input<string> PathPart { get; set; } = null!;

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Input("restApi", required: true)]
        public Input<string> RestApi { get; set; } = null!;

        public ResourceArgs()
        {
        }
    }

    public sealed class ResourceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the parent API resource
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// The complete path for this API resource, including all parent paths.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The last path segment of this API resource.
        /// </summary>
        [Input("pathPart")]
        public Input<string>? PathPart { get; set; }

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Input("restApi")]
        public Input<string>? RestApi { get; set; }

        public ResourceState()
        {
        }
    }
}
