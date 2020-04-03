// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ebs
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about an EBS volume for use in other
        /// resources.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_volume.html.markdown.
        /// </summary>
        [Obsolete("Use GetVolume.InvokeAsync() instead")]
        public static Task<GetVolumeResult> GetVolume(GetVolumeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("aws:ebs/getVolume:getVolume", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetVolume
    {
        /// <summary>
        /// Use this data source to get information about an EBS volume for use in other
        /// resources.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_volume.html.markdown.
        /// </summary>
        public static Task<GetVolumeResult> InvokeAsync(GetVolumeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("aws:ebs/getVolume:getVolume", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetVolumeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetVolumeFiltersArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to filter off of. There are
        /// several valid keys, for a full reference, check out
        /// [describe-volumes in the AWS CLI reference][1].
        /// </summary>
        public List<Inputs.GetVolumeFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVolumeFiltersArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// If more than one result is returned, use the most
        /// recent Volume.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

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

        public GetVolumeArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVolumeResult
    {
        /// <summary>
        /// The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The AZ where the EBS volume exists.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// Whether the disk is encrypted.
        /// </summary>
        public readonly bool Encrypted;
        public readonly ImmutableArray<Outputs.GetVolumeFiltersResult> Filters;
        /// <summary>
        /// The amount of IOPS for the disk.
        /// </summary>
        public readonly int Iops;
        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        public readonly string KmsKeyId;
        public readonly bool? MostRecent;
        /// <summary>
        /// The size of the drive in GiBs.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The snapshot_id the EBS volume is based off.
        /// </summary>
        public readonly string SnapshotId;
        /// <summary>
        /// A mapping of tags for the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The volume ID (e.g. vol-59fcb34e).
        /// </summary>
        public readonly string VolumeId;
        /// <summary>
        /// The type of EBS volume.
        /// </summary>
        public readonly string VolumeType;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVolumeResult(
            string arn,
            string availabilityZone,
            bool encrypted,
            ImmutableArray<Outputs.GetVolumeFiltersResult> filters,
            int iops,
            string kmsKeyId,
            bool? mostRecent,
            int size,
            string snapshotId,
            ImmutableDictionary<string, object> tags,
            string volumeId,
            string volumeType,
            string id)
        {
            Arn = arn;
            AvailabilityZone = availabilityZone;
            Encrypted = encrypted;
            Filters = filters;
            Iops = iops;
            KmsKeyId = kmsKeyId;
            MostRecent = mostRecent;
            Size = size;
            SnapshotId = snapshotId;
            Tags = tags;
            VolumeId = volumeId;
            VolumeType = volumeType;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetVolumeFiltersArgs : Pulumi.InvokeArgs
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

        public GetVolumeFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetVolumeFiltersResult
    {
        public readonly string Name;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetVolumeFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}
