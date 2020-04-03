// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Workspaces
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about a WorkSpaces Bundle.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/workspaces_bundle.html.markdown.
        /// </summary>
        [Obsolete("Use GetBundle.InvokeAsync() instead")]
        public static Task<GetBundleResult> GetBundle(GetBundleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBundleResult>("aws:workspaces/getBundle:getBundle", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetBundle
    {
        /// <summary>
        /// Use this data source to get information about a WorkSpaces Bundle.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/workspaces_bundle.html.markdown.
        /// </summary>
        public static Task<GetBundleResult> InvokeAsync(GetBundleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBundleResult>("aws:workspaces/getBundle:getBundle", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetBundleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the bundle.
        /// </summary>
        [Input("bundleId", required: true)]
        public string BundleId { get; set; } = null!;

        public GetBundleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetBundleResult
    {
        public readonly string BundleId;
        /// <summary>
        /// The compute type. See supported fields below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBundleComputeTypesResult> ComputeTypes;
        /// <summary>
        /// The description of the bundle.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The name of the compute type.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The owner of the bundle.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// The root volume. See supported fields below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBundleRootStoragesResult> RootStorages;
        /// <summary>
        /// The user storage. See supported fields below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBundleUserStoragesResult> UserStorages;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetBundleResult(
            string bundleId,
            ImmutableArray<Outputs.GetBundleComputeTypesResult> computeTypes,
            string description,
            string name,
            string owner,
            ImmutableArray<Outputs.GetBundleRootStoragesResult> rootStorages,
            ImmutableArray<Outputs.GetBundleUserStoragesResult> userStorages,
            string id)
        {
            BundleId = bundleId;
            ComputeTypes = computeTypes;
            Description = description;
            Name = name;
            Owner = owner;
            RootStorages = rootStorages;
            UserStorages = userStorages;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetBundleComputeTypesResult
    {
        /// <summary>
        /// The name of the compute type.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetBundleComputeTypesResult(string name)
        {
            Name = name;
        }
    }

    [OutputType]
    public sealed class GetBundleRootStoragesResult
    {
        /// <summary>
        /// The size of the user storage.
        /// </summary>
        public readonly string Capacity;

        [OutputConstructor]
        private GetBundleRootStoragesResult(string capacity)
        {
            Capacity = capacity;
        }
    }

    [OutputType]
    public sealed class GetBundleUserStoragesResult
    {
        /// <summary>
        /// The size of the user storage.
        /// </summary>
        public readonly string Capacity;

        [OutputConstructor]
        private GetBundleUserStoragesResult(string capacity)
        {
            Capacity = capacity;
        }
    }
    }
}
