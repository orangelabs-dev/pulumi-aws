// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceDiscovery
{
    /// <summary>
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/service_discovery_http_namespace.html.markdown.
    /// </summary>
    public partial class HttpNamespace : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN that Amazon Route 53 assigns to the namespace when you create it.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description that you specify for the namespace when you create it.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the http namespace.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a HttpNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HttpNamespace(string name, HttpNamespaceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:servicediscovery/httpNamespace:HttpNamespace", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private HttpNamespace(string name, Input<string> id, HttpNamespaceState? state = null, CustomResourceOptions? options = null)
            : base("aws:servicediscovery/httpNamespace:HttpNamespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HttpNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HttpNamespace Get(string name, Input<string> id, HttpNamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new HttpNamespace(name, id, state, options);
        }
    }

    public sealed class HttpNamespaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description that you specify for the namespace when you create it.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the http namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public HttpNamespaceArgs()
        {
        }
    }

    public sealed class HttpNamespaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN that Amazon Route 53 assigns to the namespace when you create it.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description that you specify for the namespace when you create it.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the http namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public HttpNamespaceState()
        {
        }
    }
}
