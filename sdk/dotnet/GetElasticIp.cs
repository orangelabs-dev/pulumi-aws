// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.ec2.Eip` provides details about a specific Elastic IP.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/eip.html.markdown.
        /// </summary>
        public static Task<GetElasticIpResult> GetElasticIp(GetElasticIpArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetElasticIpResult>("aws:index/getElasticIp:getElasticIp", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetElasticIpArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetElasticIpFiltersArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html).
        /// </summary>
        public List<Inputs.GetElasticIpFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetElasticIpFiltersArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The allocation id of the specific VPC EIP to retrieve. If a classic EIP is required, do NOT set `id`, only set `public_ip`
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The public IP of the specific EIP to retrieve.
        /// </summary>
        [Input("publicIp")]
        public string? PublicIp { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags, each pair of which must exactly match a pair on the desired Elastic IP
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetElasticIpArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetElasticIpResult
    {
        /// <summary>
        /// The ID representing the association of the address with an instance in a VPC.
        /// </summary>
        public readonly string AssociationId;
        /// <summary>
        /// Indicates whether the address is for use in EC2-Classic (standard) or in a VPC (vpc).
        /// </summary>
        public readonly string Domain;
        public readonly ImmutableArray<Outputs.GetElasticIpFiltersResult> Filters;
        /// <summary>
        /// If VPC Elastic IP, the allocation identifier. If EC2-Classic Elastic IP, the public IP address.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the instance that the address is associated with (if any).
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The ID of the network interface.
        /// </summary>
        public readonly string NetworkInterfaceId;
        /// <summary>
        /// The ID of the AWS account that owns the network interface.
        /// </summary>
        public readonly string NetworkInterfaceOwnerId;
        /// <summary>
        /// The Private DNS associated with the Elastic IP address.
        /// </summary>
        public readonly string PrivateDns;
        /// <summary>
        /// The private IP address associated with the Elastic IP address.
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// Public DNS associated with the Elastic IP address.
        /// </summary>
        public readonly string PublicDns;
        /// <summary>
        /// Public IP address of Elastic IP.
        /// </summary>
        public readonly string PublicIp;
        /// <summary>
        /// The ID of an address pool.
        /// </summary>
        public readonly string PublicIpv4Pool;
        /// <summary>
        /// Key-value map of tags associated with Elastic IP.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private GetElasticIpResult(
            string associationId,
            string domain,
            ImmutableArray<Outputs.GetElasticIpFiltersResult> filters,
            string id,
            string instanceId,
            string networkInterfaceId,
            string networkInterfaceOwnerId,
            string privateDns,
            string privateIp,
            string publicDns,
            string publicIp,
            string publicIpv4Pool,
            ImmutableDictionary<string, object> tags)
        {
            AssociationId = associationId;
            Domain = domain;
            Filters = filters;
            Id = id;
            InstanceId = instanceId;
            NetworkInterfaceId = networkInterfaceId;
            NetworkInterfaceOwnerId = networkInterfaceOwnerId;
            PrivateDns = privateDns;
            PrivateIp = privateIp;
            PublicDns = publicDns;
            PublicIp = publicIp;
            PublicIpv4Pool = publicIpv4Pool;
            Tags = tags;
        }
    }

    namespace Inputs
    {

    public sealed class GetElasticIpFiltersArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetElasticIpFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetElasticIpFiltersResult
    {
        public readonly string Name;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetElasticIpFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}