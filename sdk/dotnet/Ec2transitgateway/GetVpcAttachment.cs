// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get information on an EC2 Transit Gateway VPC Attachment.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_vpc_attachment.html.markdown.
        /// </summary>
        [Obsolete("Use GetVpcAttachment.InvokeAsync() instead")]
        public static Task<GetVpcAttachmentResult> GetVpcAttachment(GetVpcAttachmentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcAttachmentResult>("aws:ec2transitgateway/getVpcAttachment:getVpcAttachment", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetVpcAttachment
    {
        /// <summary>
        /// Get information on an EC2 Transit Gateway VPC Attachment.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_vpc_attachment.html.markdown.
        /// </summary>
        public static Task<GetVpcAttachmentResult> InvokeAsync(GetVpcAttachmentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcAttachmentResult>("aws:ec2transitgateway/getVpcAttachment:getVpcAttachment", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetVpcAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetVpcAttachmentFiltersArgs>? _filters;

        /// <summary>
        /// One or more configuration blocks containing name-values filters. Detailed below.
        /// </summary>
        public List<Inputs.GetVpcAttachmentFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcAttachmentFiltersArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Identifier of the EC2 Transit Gateway VPC Attachment.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway VPC Attachment
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetVpcAttachmentArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVpcAttachmentResult
    {
        /// <summary>
        /// Whether DNS support is enabled.
        /// </summary>
        public readonly string DnsSupport;
        public readonly ImmutableArray<Outputs.GetVpcAttachmentFiltersResult> Filters;
        /// <summary>
        /// EC2 Transit Gateway VPC Attachment identifier
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Whether IPv6 support is enabled.
        /// </summary>
        public readonly string Ipv6Support;
        /// <summary>
        /// Identifiers of EC2 Subnets.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway VPC Attachment
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// EC2 Transit Gateway identifier
        /// </summary>
        public readonly string TransitGatewayId;
        /// <summary>
        /// Identifier of EC2 VPC.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// Identifier of the AWS account that owns the EC2 VPC.
        /// </summary>
        public readonly string VpcOwnerId;

        [OutputConstructor]
        private GetVpcAttachmentResult(
            string dnsSupport,
            ImmutableArray<Outputs.GetVpcAttachmentFiltersResult> filters,
            string? id,
            string ipv6Support,
            ImmutableArray<string> subnetIds,
            ImmutableDictionary<string, object> tags,
            string transitGatewayId,
            string vpcId,
            string vpcOwnerId)
        {
            DnsSupport = dnsSupport;
            Filters = filters;
            Id = id;
            Ipv6Support = ipv6Support;
            SubnetIds = subnetIds;
            Tags = tags;
            TransitGatewayId = transitGatewayId;
            VpcId = vpcId;
            VpcOwnerId = vpcOwnerId;
        }
    }

    namespace Inputs
    {

    public sealed class GetVpcAttachmentFiltersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the filter.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// List of one or more values for the filter.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetVpcAttachmentFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetVpcAttachmentFiltersResult
    {
        /// <summary>
        /// Name of the filter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of one or more values for the filter.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetVpcAttachmentFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}
