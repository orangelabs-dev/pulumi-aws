// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr
{
    /// <summary>
    /// Provides an Elastic MapReduce Cluster Instance Group configuration.
    /// See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/emr/) for more information.
    /// 
    /// &gt; **NOTE:** At this time, Instance Groups cannot be destroyed through the API nor
    /// web interface. Instance Groups are destroyed when the EMR Cluster is destroyed.
    /// this provider will resize any Instance Group to zero when destroying the resource.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/emr_instance_group.html.markdown.
    /// </summary>
    public partial class InstanceGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        /// </summary>
        [Output("autoscalingPolicy")]
        public Output<string?> AutoscalingPolicy { get; private set; } = null!;

        /// <summary>
        /// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        /// </summary>
        [Output("bidPrice")]
        public Output<string?> BidPrice { get; private set; } = null!;

        /// <summary>
        /// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
        /// </summary>
        [Output("configurationsJson")]
        public Output<string?> ConfigurationsJson { get; private set; } = null!;

        /// <summary>
        /// One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Output("ebsConfigs")]
        public Output<ImmutableArray<Outputs.InstanceGroupEbsConfigs>> EbsConfigs { get; private set; } = null!;

        /// <summary>
        /// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        /// </summary>
        [Output("ebsOptimized")]
        public Output<bool?> EbsOptimized { get; private set; } = null!;

        /// <summary>
        /// target number of instances for the instance group. defaults to 0.
        /// </summary>
        [Output("instanceCount")]
        public Output<int?> InstanceCount { get; private set; } = null!;

        /// <summary>
        /// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Human friendly name given to the instance group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("runningInstanceCount")]
        public Output<int> RunningInstanceCount { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceGroup(string name, InstanceGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:emr/instanceGroup:InstanceGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private InstanceGroup(string name, Input<string> id, InstanceGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:emr/instanceGroup:InstanceGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceGroup Get(string name, Input<string> id, InstanceGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceGroup(name, id, state, options);
        }
    }

    public sealed class InstanceGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<string>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        /// </summary>
        [Input("bidPrice")]
        public Input<string>? BidPrice { get; set; }

        /// <summary>
        /// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
        /// </summary>
        [Input("configurationsJson")]
        public Input<string>? ConfigurationsJson { get; set; }

        [Input("ebsConfigs")]
        private InputList<Inputs.InstanceGroupEbsConfigsArgs>? _ebsConfigs;

        /// <summary>
        /// One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<Inputs.InstanceGroupEbsConfigsArgs> EbsConfigs
        {
            get => _ebsConfigs ?? (_ebsConfigs = new InputList<Inputs.InstanceGroupEbsConfigsArgs>());
            set => _ebsConfigs = value;
        }

        /// <summary>
        /// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        /// <summary>
        /// target number of instances for the instance group. defaults to 0.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Human friendly name given to the instance group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public InstanceGroupArgs()
        {
        }
    }

    public sealed class InstanceGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<string>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        /// </summary>
        [Input("bidPrice")]
        public Input<string>? BidPrice { get; set; }

        /// <summary>
        /// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
        /// </summary>
        [Input("configurationsJson")]
        public Input<string>? ConfigurationsJson { get; set; }

        [Input("ebsConfigs")]
        private InputList<Inputs.InstanceGroupEbsConfigsGetArgs>? _ebsConfigs;

        /// <summary>
        /// One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<Inputs.InstanceGroupEbsConfigsGetArgs> EbsConfigs
        {
            get => _ebsConfigs ?? (_ebsConfigs = new InputList<Inputs.InstanceGroupEbsConfigsGetArgs>());
            set => _ebsConfigs = value;
        }

        /// <summary>
        /// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        /// <summary>
        /// target number of instances for the instance group. defaults to 0.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Human friendly name given to the instance group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("runningInstanceCount")]
        public Input<int>? RunningInstanceCount { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public InstanceGroupState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class InstanceGroupEbsConfigsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of I/O operations per second (IOPS) that the volume supports.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The volume size, in gibibytes (GiB). This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The volume type. Valid options are 'gp2', 'io1' and 'standard'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The number of EBS Volumes to attach per instance.
        /// </summary>
        [Input("volumesPerInstance")]
        public Input<int>? VolumesPerInstance { get; set; }

        public InstanceGroupEbsConfigsArgs()
        {
        }
    }

    public sealed class InstanceGroupEbsConfigsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of I/O operations per second (IOPS) that the volume supports.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The volume size, in gibibytes (GiB). This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The volume type. Valid options are 'gp2', 'io1' and 'standard'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The number of EBS Volumes to attach per instance.
        /// </summary>
        [Input("volumesPerInstance")]
        public Input<int>? VolumesPerInstance { get; set; }

        public InstanceGroupEbsConfigsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class InstanceGroupEbsConfigs
    {
        /// <summary>
        /// The number of I/O operations per second (IOPS) that the volume supports.
        /// </summary>
        public readonly int? Iops;
        /// <summary>
        /// The volume size, in gibibytes (GiB). This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The volume type. Valid options are 'gp2', 'io1' and 'standard'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The number of EBS Volumes to attach per instance.
        /// </summary>
        public readonly int? VolumesPerInstance;

        [OutputConstructor]
        private InstanceGroupEbsConfigs(
            int? iops,
            int size,
            string type,
            int? volumesPerInstance)
        {
            Iops = iops;
            Size = size;
            Type = type;
            VolumesPerInstance = volumesPerInstance;
        }
    }
    }
}
