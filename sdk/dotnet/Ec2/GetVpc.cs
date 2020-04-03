// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.ec2.Vpc` provides details about a specific VPC.
        /// 
        /// This resource can prove useful when a module accepts a vpc id as
        /// an input variable and needs to, for example, determine the CIDR block of that
        /// VPC.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc.html.markdown.
        /// </summary>
        [Obsolete("Use GetVpc.InvokeAsync() instead")]
        public static Task<GetVpcResult> GetVpc(GetVpcArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcResult>("aws:ec2/getVpc:getVpc", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetVpc
    {
        /// <summary>
        /// `aws.ec2.Vpc` provides details about a specific VPC.
        /// 
        /// This resource can prove useful when a module accepts a vpc id as
        /// an input variable and needs to, for example, determine the CIDR block of that
        /// VPC.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc.html.markdown.
        /// </summary>
        public static Task<GetVpcResult> InvokeAsync(GetVpcArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcResult>("aws:ec2/getVpc:getVpc", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetVpcArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cidr block of the desired VPC.
        /// </summary>
        [Input("cidrBlock")]
        public string? CidrBlock { get; set; }

        /// <summary>
        /// Boolean constraint on whether the desired VPC is
        /// the default VPC for the region.
        /// </summary>
        [Input("default")]
        public bool? Default { get; set; }

        /// <summary>
        /// The DHCP options id of the desired VPC.
        /// </summary>
        [Input("dhcpOptionsId")]
        public string? DhcpOptionsId { get; set; }

        [Input("filters")]
        private List<Inputs.GetVpcFiltersArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetVpcFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcFiltersArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The id of the specific VPC to retrieve.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The current state of the desired VPC.
        /// Can be either `"pending"` or `"available"`.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags, each pair of which must exactly match
        /// a pair on the desired VPC.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetVpcArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVpcResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of VPC
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The CIDR block for the association.
        /// </summary>
        public readonly string CidrBlock;
        public readonly ImmutableArray<Outputs.GetVpcCidrBlockAssociationsResult> CidrBlockAssociations;
        public readonly bool Default;
        public readonly string DhcpOptionsId;
        /// <summary>
        /// Whether or not the VPC has DNS hostname support
        /// </summary>
        public readonly bool EnableDnsHostnames;
        /// <summary>
        /// Whether or not the VPC has DNS support
        /// </summary>
        public readonly bool EnableDnsSupport;
        public readonly ImmutableArray<Outputs.GetVpcFiltersResult> Filters;
        public readonly string Id;
        /// <summary>
        /// The allowed tenancy of instances launched into the
        /// selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
        /// </summary>
        public readonly string InstanceTenancy;
        /// <summary>
        /// The association ID for the IPv6 CIDR block.
        /// </summary>
        public readonly string Ipv6AssociationId;
        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        public readonly string Ipv6CidrBlock;
        /// <summary>
        /// The ID of the main route table associated with this VPC.
        /// </summary>
        public readonly string MainRouteTableId;
        /// <summary>
        /// The ID of the AWS account that owns the VPC.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// The State of the association.
        /// </summary>
        public readonly string State;
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private GetVpcResult(
            string arn,
            string cidrBlock,
            ImmutableArray<Outputs.GetVpcCidrBlockAssociationsResult> cidrBlockAssociations,
            bool @default,
            string dhcpOptionsId,
            bool enableDnsHostnames,
            bool enableDnsSupport,
            ImmutableArray<Outputs.GetVpcFiltersResult> filters,
            string id,
            string instanceTenancy,
            string ipv6AssociationId,
            string ipv6CidrBlock,
            string mainRouteTableId,
            string ownerId,
            string state,
            ImmutableDictionary<string, object> tags)
        {
            Arn = arn;
            CidrBlock = cidrBlock;
            CidrBlockAssociations = cidrBlockAssociations;
            Default = @default;
            DhcpOptionsId = dhcpOptionsId;
            EnableDnsHostnames = enableDnsHostnames;
            EnableDnsSupport = enableDnsSupport;
            Filters = filters;
            Id = id;
            InstanceTenancy = instanceTenancy;
            Ipv6AssociationId = ipv6AssociationId;
            Ipv6CidrBlock = ipv6CidrBlock;
            MainRouteTableId = mainRouteTableId;
            OwnerId = ownerId;
            State = state;
            Tags = tags;
        }
    }

    namespace Inputs
    {

    public sealed class GetVpcFiltersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the field to filter by, as defined by
        /// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html).
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// Set of values that are accepted for the given field.
        /// A VPC will be selected if any one of the given values matches.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetVpcFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetVpcCidrBlockAssociationsResult
    {
        /// <summary>
        /// The association ID for the the IPv4 CIDR block.
        /// </summary>
        public readonly string AssociationId;
        /// <summary>
        /// The cidr block of the desired VPC.
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// The current state of the desired VPC.
        /// Can be either `"pending"` or `"available"`.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetVpcCidrBlockAssociationsResult(
            string associationId,
            string cidrBlock,
            string state)
        {
            AssociationId = associationId;
            CidrBlock = cidrBlock;
            State = state;
        }
    }

    [OutputType]
    public sealed class GetVpcFiltersResult
    {
        /// <summary>
        /// The name of the field to filter by, as defined by
        /// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Set of values that are accepted for the given field.
        /// A VPC will be selected if any one of the given values matches.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetVpcFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}
