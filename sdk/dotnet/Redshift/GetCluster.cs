// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedShift
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides details about a specific redshift cluster.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/redshift_cluster.html.markdown.
        /// </summary>
        [Obsolete("Use GetCluster.InvokeAsync() instead")]
        public static Task<GetClusterResult> GetCluster(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:redshift/getCluster:getCluster", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetCluster
    {
        /// <summary>
        /// Provides details about a specific redshift cluster.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/redshift_cluster.html.markdown.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:redshift/getCluster:getCluster", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster identifier
        /// </summary>
        [Input("clusterIdentifier", required: true)]
        public string ClusterIdentifier { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// The tags associated to the cluster
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetClusterArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// Whether major version upgrades can be applied during maintenance period
        /// </summary>
        public readonly bool AllowVersionUpgrade;
        /// <summary>
        /// The backup retention period
        /// </summary>
        public readonly int AutomatedSnapshotRetentionPeriod;
        /// <summary>
        /// The availability zone of the cluster
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// The name of the S3 bucket where the log files are to be stored
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// The cluster identifier
        /// </summary>
        public readonly string ClusterIdentifier;
        /// <summary>
        /// The name of the parameter group to be associated with this cluster
        /// </summary>
        public readonly string ClusterParameterGroupName;
        /// <summary>
        /// The public key for the cluster
        /// </summary>
        public readonly string ClusterPublicKey;
        /// <summary>
        /// The cluster revision number
        /// </summary>
        public readonly string ClusterRevisionNumber;
        /// <summary>
        /// The security groups associated with the cluster
        /// </summary>
        public readonly ImmutableArray<string> ClusterSecurityGroups;
        /// <summary>
        /// The name of a cluster subnet group to be associated with this cluster
        /// </summary>
        public readonly string ClusterSubnetGroupName;
        /// <summary>
        /// The cluster type
        /// </summary>
        public readonly string ClusterType;
        public readonly string ClusterVersion;
        /// <summary>
        /// The name of the default database in the cluster
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The Elastic IP of the cluster
        /// </summary>
        public readonly string ElasticIp;
        /// <summary>
        /// Whether cluster logging is enabled
        /// </summary>
        public readonly bool EnableLogging;
        /// <summary>
        /// Whether the cluster data is encrypted
        /// </summary>
        public readonly bool Encrypted;
        /// <summary>
        /// The cluster endpoint
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// Whether enhanced VPC routing is enabled
        /// </summary>
        public readonly bool EnhancedVpcRouting;
        /// <summary>
        /// The IAM roles associated to the cluster
        /// </summary>
        public readonly ImmutableArray<string> IamRoles;
        /// <summary>
        /// The KMS encryption key associated to the cluster
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// Username for the master DB user
        /// </summary>
        public readonly string MasterUsername;
        /// <summary>
        /// The cluster node type
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// The number of nodes in the cluster
        /// </summary>
        public readonly int NumberOfNodes;
        /// <summary>
        /// The port the cluster responds on
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The maintenance window
        /// </summary>
        public readonly string PreferredMaintenanceWindow;
        /// <summary>
        /// Whether the cluster is publicly accessible
        /// </summary>
        public readonly bool PubliclyAccessible;
        /// <summary>
        /// The folder inside the S3 bucket where the log files are stored
        /// </summary>
        public readonly string S3KeyPrefix;
        /// <summary>
        /// The tags associated to the cluster
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// The VPC Id associated with the cluster
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The VPC security group Ids associated with the cluster
        /// </summary>
        public readonly ImmutableArray<string> VpcSecurityGroupIds;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterResult(
            bool allowVersionUpgrade,
            int automatedSnapshotRetentionPeriod,
            string availabilityZone,
            string bucketName,
            string clusterIdentifier,
            string clusterParameterGroupName,
            string clusterPublicKey,
            string clusterRevisionNumber,
            ImmutableArray<string> clusterSecurityGroups,
            string clusterSubnetGroupName,
            string clusterType,
            string clusterVersion,
            string databaseName,
            string elasticIp,
            bool enableLogging,
            bool encrypted,
            string endpoint,
            bool enhancedVpcRouting,
            ImmutableArray<string> iamRoles,
            string kmsKeyId,
            string masterUsername,
            string nodeType,
            int numberOfNodes,
            int port,
            string preferredMaintenanceWindow,
            bool publiclyAccessible,
            string s3KeyPrefix,
            ImmutableDictionary<string, object>? tags,
            string vpcId,
            ImmutableArray<string> vpcSecurityGroupIds,
            string id)
        {
            AllowVersionUpgrade = allowVersionUpgrade;
            AutomatedSnapshotRetentionPeriod = automatedSnapshotRetentionPeriod;
            AvailabilityZone = availabilityZone;
            BucketName = bucketName;
            ClusterIdentifier = clusterIdentifier;
            ClusterParameterGroupName = clusterParameterGroupName;
            ClusterPublicKey = clusterPublicKey;
            ClusterRevisionNumber = clusterRevisionNumber;
            ClusterSecurityGroups = clusterSecurityGroups;
            ClusterSubnetGroupName = clusterSubnetGroupName;
            ClusterType = clusterType;
            ClusterVersion = clusterVersion;
            DatabaseName = databaseName;
            ElasticIp = elasticIp;
            EnableLogging = enableLogging;
            Encrypted = encrypted;
            Endpoint = endpoint;
            EnhancedVpcRouting = enhancedVpcRouting;
            IamRoles = iamRoles;
            KmsKeyId = kmsKeyId;
            MasterUsername = masterUsername;
            NodeType = nodeType;
            NumberOfNodes = numberOfNodes;
            Port = port;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            PubliclyAccessible = publiclyAccessible;
            S3KeyPrefix = s3KeyPrefix;
            Tags = tags;
            VpcId = vpcId;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
            Id = id;
        }
    }
}
