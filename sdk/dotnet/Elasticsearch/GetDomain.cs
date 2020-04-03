// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticSearch
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about an Elasticsearch Domain
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elasticsearch_domain.html.markdown.
        /// </summary>
        [Obsolete("Use GetDomain.InvokeAsync() instead")]
        public static Task<GetDomainResult> GetDomain(GetDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("aws:elasticsearch/getDomain:getDomain", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetDomain
    {
        /// <summary>
        /// Use this data source to get information about an Elasticsearch Domain
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elasticsearch_domain.html.markdown.
        /// </summary>
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("aws:elasticsearch/getDomain:getDomain", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetDomainArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// The tags assigned to the domain.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetDomainArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetDomainResult
    {
        /// <summary>
        /// The policy document attached to the domain.
        /// </summary>
        public readonly string AccessPolicies;
        /// <summary>
        /// Key-value string pairs to specify advanced configuration options.
        /// </summary>
        public readonly ImmutableDictionary<string, object> AdvancedOptions;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the domain.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Cluster configuration of the domain.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainClusterConfigsResult> ClusterConfigs;
        /// <summary>
        /// Domain Amazon Cognito Authentication options for Kibana.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainCognitoOptionsResult> CognitoOptions;
        /// <summary>
        /// Status of the creation of the domain.
        /// </summary>
        public readonly bool Created;
        /// <summary>
        /// Status of the deletion of the domain.
        /// </summary>
        public readonly bool Deleted;
        /// <summary>
        /// Unique identifier for the domain.
        /// </summary>
        public readonly string DomainId;
        public readonly string DomainName;
        /// <summary>
        /// EBS Options for the instances in the domain.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainEbsOptionsResult> EbsOptions;
        /// <summary>
        /// ElasticSearch version for the domain.
        /// </summary>
        public readonly string ElasticsearchVersion;
        /// <summary>
        /// Domain encryption at rest related options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainEncryptionAtRestsResult> EncryptionAtRests;
        /// <summary>
        /// Domain-specific endpoint used to submit index, search, and data upload requests.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// Domain-specific endpoint used to access the Kibana application.
        /// </summary>
        public readonly string KibanaEndpoint;
        /// <summary>
        /// Domain log publishing related options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainLogPublishingOptionsResult> LogPublishingOptions;
        /// <summary>
        /// Domain in transit encryption related options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainNodeToNodeEncryptionsResult> NodeToNodeEncryptions;
        /// <summary>
        /// Status of a configuration change in the domain.
        /// * `snapshot_options` – Domain snapshot related options.
        /// </summary>
        public readonly string Processing;
        public readonly ImmutableArray<Outputs.GetDomainSnapshotOptionsResult> SnapshotOptions;
        /// <summary>
        /// The tags assigned to the domain.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// VPC Options for private Elasticsearch domains.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainVpcOptionsResult> VpcOptions;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDomainResult(
            string accessPolicies,
            ImmutableDictionary<string, object> advancedOptions,
            string arn,
            ImmutableArray<Outputs.GetDomainClusterConfigsResult> clusterConfigs,
            ImmutableArray<Outputs.GetDomainCognitoOptionsResult> cognitoOptions,
            bool created,
            bool deleted,
            string domainId,
            string domainName,
            ImmutableArray<Outputs.GetDomainEbsOptionsResult> ebsOptions,
            string elasticsearchVersion,
            ImmutableArray<Outputs.GetDomainEncryptionAtRestsResult> encryptionAtRests,
            string endpoint,
            string kibanaEndpoint,
            ImmutableArray<Outputs.GetDomainLogPublishingOptionsResult> logPublishingOptions,
            ImmutableArray<Outputs.GetDomainNodeToNodeEncryptionsResult> nodeToNodeEncryptions,
            string processing,
            ImmutableArray<Outputs.GetDomainSnapshotOptionsResult> snapshotOptions,
            ImmutableDictionary<string, object> tags,
            ImmutableArray<Outputs.GetDomainVpcOptionsResult> vpcOptions,
            string id)
        {
            AccessPolicies = accessPolicies;
            AdvancedOptions = advancedOptions;
            Arn = arn;
            ClusterConfigs = clusterConfigs;
            CognitoOptions = cognitoOptions;
            Created = created;
            Deleted = deleted;
            DomainId = domainId;
            DomainName = domainName;
            EbsOptions = ebsOptions;
            ElasticsearchVersion = elasticsearchVersion;
            EncryptionAtRests = encryptionAtRests;
            Endpoint = endpoint;
            KibanaEndpoint = kibanaEndpoint;
            LogPublishingOptions = logPublishingOptions;
            NodeToNodeEncryptions = nodeToNodeEncryptions;
            Processing = processing;
            SnapshotOptions = snapshotOptions;
            Tags = tags;
            VpcOptions = vpcOptions;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetDomainClusterConfigsResult
    {
        /// <summary>
        /// Number of dedicated master nodes in the cluster.
        /// </summary>
        public readonly int DedicatedMasterCount;
        /// <summary>
        /// Indicates whether dedicated master nodes are enabled for the cluster.
        /// </summary>
        public readonly bool DedicatedMasterEnabled;
        /// <summary>
        /// Instance type of the dedicated master nodes in the cluster.
        /// </summary>
        public readonly string DedicatedMasterType;
        /// <summary>
        /// Number of instances in the cluster.
        /// </summary>
        public readonly int InstanceCount;
        /// <summary>
        /// Instance type of data nodes in the cluster.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Configuration block containing zone awareness settings.
        /// </summary>
        public readonly ImmutableArray<GetDomainClusterConfigsZoneAwarenessConfigsResult> ZoneAwarenessConfigs;
        /// <summary>
        /// Indicates whether zone awareness is enabled.
        /// </summary>
        public readonly bool ZoneAwarenessEnabled;

        [OutputConstructor]
        private GetDomainClusterConfigsResult(
            int dedicatedMasterCount,
            bool dedicatedMasterEnabled,
            string dedicatedMasterType,
            int instanceCount,
            string instanceType,
            ImmutableArray<GetDomainClusterConfigsZoneAwarenessConfigsResult> zoneAwarenessConfigs,
            bool zoneAwarenessEnabled)
        {
            DedicatedMasterCount = dedicatedMasterCount;
            DedicatedMasterEnabled = dedicatedMasterEnabled;
            DedicatedMasterType = dedicatedMasterType;
            InstanceCount = instanceCount;
            InstanceType = instanceType;
            ZoneAwarenessConfigs = zoneAwarenessConfigs;
            ZoneAwarenessEnabled = zoneAwarenessEnabled;
        }
    }

    [OutputType]
    public sealed class GetDomainClusterConfigsZoneAwarenessConfigsResult
    {
        /// <summary>
        /// Number of availability zones used.
        /// </summary>
        public readonly int AvailabilityZoneCount;

        [OutputConstructor]
        private GetDomainClusterConfigsZoneAwarenessConfigsResult(int availabilityZoneCount)
        {
            AvailabilityZoneCount = availabilityZoneCount;
        }
    }

    [OutputType]
    public sealed class GetDomainCognitoOptionsResult
    {
        /// <summary>
        /// Whether node to node encryption is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The Cognito Identity pool used by the domain.
        /// </summary>
        public readonly string IdentityPoolId;
        /// <summary>
        /// The IAM Role with the AmazonESCognitoAccess policy attached.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The Cognito User pool used by the domain.
        /// </summary>
        public readonly string UserPoolId;

        [OutputConstructor]
        private GetDomainCognitoOptionsResult(
            bool enabled,
            string identityPoolId,
            string roleArn,
            string userPoolId)
        {
            Enabled = enabled;
            IdentityPoolId = identityPoolId;
            RoleArn = roleArn;
            UserPoolId = userPoolId;
        }
    }

    [OutputType]
    public sealed class GetDomainEbsOptionsResult
    {
        /// <summary>
        /// Whether EBS volumes are attached to data nodes in the domain.
        /// </summary>
        public readonly bool EbsEnabled;
        /// <summary>
        /// The baseline input/output (I/O) performance of EBS volumes
        /// attached to data nodes.
        /// </summary>
        public readonly int Iops;
        /// <summary>
        /// The size of EBS volumes attached to data nodes (in GB).
        /// </summary>
        public readonly int VolumeSize;
        /// <summary>
        /// The type of EBS volumes attached to data nodes.
        /// </summary>
        public readonly string VolumeType;

        [OutputConstructor]
        private GetDomainEbsOptionsResult(
            bool ebsEnabled,
            int iops,
            int volumeSize,
            string volumeType)
        {
            EbsEnabled = ebsEnabled;
            Iops = iops;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }

    [OutputType]
    public sealed class GetDomainEncryptionAtRestsResult
    {
        /// <summary>
        /// Whether node to node encryption is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The KMS key id used to encrypt data at rest.
        /// </summary>
        public readonly string KmsKeyId;

        [OutputConstructor]
        private GetDomainEncryptionAtRestsResult(
            bool enabled,
            string kmsKeyId)
        {
            Enabled = enabled;
            KmsKeyId = kmsKeyId;
        }
    }

    [OutputType]
    public sealed class GetDomainLogPublishingOptionsResult
    {
        /// <summary>
        /// The CloudWatch Log Group where the logs are published.
        /// </summary>
        public readonly string CloudwatchLogGroupArn;
        /// <summary>
        /// Whether node to node encryption is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The type of Elasticsearch log being published.
        /// </summary>
        public readonly string LogType;

        [OutputConstructor]
        private GetDomainLogPublishingOptionsResult(
            string cloudwatchLogGroupArn,
            bool enabled,
            string logType)
        {
            CloudwatchLogGroupArn = cloudwatchLogGroupArn;
            Enabled = enabled;
            LogType = logType;
        }
    }

    [OutputType]
    public sealed class GetDomainNodeToNodeEncryptionsResult
    {
        /// <summary>
        /// Whether node to node encryption is enabled.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private GetDomainNodeToNodeEncryptionsResult(bool enabled)
        {
            Enabled = enabled;
        }
    }

    [OutputType]
    public sealed class GetDomainSnapshotOptionsResult
    {
        /// <summary>
        /// Hour during which the service takes an automated daily
        /// snapshot of the indices in the domain.
        /// </summary>
        public readonly int AutomatedSnapshotStartHour;

        [OutputConstructor]
        private GetDomainSnapshotOptionsResult(int automatedSnapshotStartHour)
        {
            AutomatedSnapshotStartHour = automatedSnapshotStartHour;
        }
    }

    [OutputType]
    public sealed class GetDomainVpcOptionsResult
    {
        /// <summary>
        /// The availability zones used by the domain.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// The security groups used by the domain.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The subnets used by the domain.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The VPC used by the domain.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetDomainVpcOptionsResult(
            ImmutableArray<string> availabilityZones,
            ImmutableArray<string> securityGroupIds,
            ImmutableArray<string> subnetIds,
            string vpcId)
        {
            AvailabilityZones = availabilityZones;
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
    }
}
