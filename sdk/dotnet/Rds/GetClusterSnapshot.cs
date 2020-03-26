// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about a DB Cluster Snapshot for use when provisioning DB clusters.
        /// 
        /// &gt; **NOTE:** This data source does not apply to snapshots created on DB Instances. 
        /// See the [`aws.rds.Snapshot` data source](https://www.terraform.io/docs/providers/aws/d/db_snapshot.html) for DB Instance snapshots.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/db_cluster_snapshot.html.markdown.
        /// </summary>
        [Obsolete("Use GetClusterSnapshot.InvokeAsync() instead")]
        public static Task<GetClusterSnapshotResult> GetClusterSnapshot(GetClusterSnapshotArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterSnapshotResult>("aws:rds/getClusterSnapshot:getClusterSnapshot", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetClusterSnapshot
    {
        /// <summary>
        /// Use this data source to get information about a DB Cluster Snapshot for use when provisioning DB clusters.
        /// 
        /// &gt; **NOTE:** This data source does not apply to snapshots created on DB Instances. 
        /// See the [`aws.rds.Snapshot` data source](https://www.terraform.io/docs/providers/aws/d/db_snapshot.html) for DB Instance snapshots.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/db_cluster_snapshot.html.markdown.
        /// </summary>
        public static Task<GetClusterSnapshotResult> InvokeAsync(GetClusterSnapshotArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterSnapshotResult>("aws:rds/getClusterSnapshot:getClusterSnapshot", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterSnapshotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns the list of snapshots created by the specific db_cluster
        /// </summary>
        [Input("dbClusterIdentifier")]
        public string? DbClusterIdentifier { get; set; }

        /// <summary>
        /// Returns information on a specific snapshot_id.
        /// </summary>
        [Input("dbClusterSnapshotIdentifier")]
        public string? DbClusterSnapshotIdentifier { get; set; }

        /// <summary>
        /// Set this value to true to include manual DB Cluster Snapshots that are public and can be
        /// copied or restored by any AWS account, otherwise set this value to false. The default is `false`.
        /// </summary>
        [Input("includePublic")]
        public bool? IncludePublic { get; set; }

        /// <summary>
        /// Set this value to true to include shared manual DB Cluster Snapshots from other
        /// AWS accounts that this AWS account has been given permission to copy or restore, otherwise set this value to false.
        /// The default is `false`.
        /// </summary>
        [Input("includeShared")]
        public bool? IncludeShared { get; set; }

        /// <summary>
        /// If more than one result is returned, use the most recent Snapshot.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// The type of snapshots to be returned. If you don't specify a SnapshotType
        /// value, then both automated and manual DB cluster snapshots are returned. Shared and public DB Cluster Snapshots are not
        /// included in the returned results by default. Possible values are, `automated`, `manual`, `shared` and `public`.
        /// </summary>
        [Input("snapshotType")]
        public string? SnapshotType { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags for the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetClusterSnapshotArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterSnapshotResult
    {
        /// <summary>
        /// Specifies the allocated storage size in gigabytes (GB).
        /// </summary>
        public readonly int AllocatedStorage;
        /// <summary>
        /// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// Specifies the DB cluster identifier of the DB cluster that this DB cluster snapshot was created from.
        /// </summary>
        public readonly string? DbClusterIdentifier;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
        /// </summary>
        public readonly string DbClusterSnapshotArn;
        public readonly string? DbClusterSnapshotIdentifier;
        /// <summary>
        /// Specifies the name of the database engine.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// Version of the database engine for this DB cluster snapshot.
        /// </summary>
        public readonly string EngineVersion;
        public readonly bool? IncludePublic;
        public readonly bool? IncludeShared;
        /// <summary>
        /// If storage_encrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// License model information for the restored DB cluster.
        /// </summary>
        public readonly string LicenseModel;
        public readonly bool? MostRecent;
        /// <summary>
        /// Port that the DB cluster was listening on at the time of the snapshot.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Time when the snapshot was taken, in Universal Coordinated Time (UTC).
        /// </summary>
        public readonly string SnapshotCreateTime;
        public readonly string? SnapshotType;
        public readonly string SourceDbClusterSnapshotArn;
        /// <summary>
        /// The status of this DB Cluster Snapshot.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Specifies whether the DB cluster snapshot is encrypted.
        /// </summary>
        public readonly bool StorageEncrypted;
        /// <summary>
        /// A mapping of tags for the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The VPC ID associated with the DB cluster snapshot.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterSnapshotResult(
            int allocatedStorage,
            ImmutableArray<string> availabilityZones,
            string? dbClusterIdentifier,
            string dbClusterSnapshotArn,
            string? dbClusterSnapshotIdentifier,
            string engine,
            string engineVersion,
            bool? includePublic,
            bool? includeShared,
            string kmsKeyId,
            string licenseModel,
            bool? mostRecent,
            int port,
            string snapshotCreateTime,
            string? snapshotType,
            string sourceDbClusterSnapshotArn,
            string status,
            bool storageEncrypted,
            ImmutableDictionary<string, object> tags,
            string vpcId,
            string id)
        {
            AllocatedStorage = allocatedStorage;
            AvailabilityZones = availabilityZones;
            DbClusterIdentifier = dbClusterIdentifier;
            DbClusterSnapshotArn = dbClusterSnapshotArn;
            DbClusterSnapshotIdentifier = dbClusterSnapshotIdentifier;
            Engine = engine;
            EngineVersion = engineVersion;
            IncludePublic = includePublic;
            IncludeShared = includeShared;
            KmsKeyId = kmsKeyId;
            LicenseModel = licenseModel;
            MostRecent = mostRecent;
            Port = port;
            SnapshotCreateTime = snapshotCreateTime;
            SnapshotType = snapshotType;
            SourceDbClusterSnapshotArn = sourceDbClusterSnapshotArn;
            Status = status;
            StorageEncrypted = storageEncrypted;
            Tags = tags;
            VpcId = vpcId;
            Id = id;
        }
    }
}
