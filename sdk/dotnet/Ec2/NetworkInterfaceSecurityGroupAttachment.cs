// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// This resource attaches a security group to an Elastic Network Interface (ENI).
    /// It can be used to attach a security group to any existing ENI, be it a
    /// secondary ENI or one attached as the primary interface on an instance.
    /// 
    /// &gt; **NOTE on instances, interfaces, and security groups:** This provider currently
    /// provides the capability to assign security groups via the [`aws.ec2.Instance`][1]
    /// and the [`aws.ec2.NetworkInterface`][2] resources. Using this resource in
    /// conjunction with security groups provided in-line in those resources will cause
    /// conflicts, and will lead to spurious diffs and undefined behavior - please use
    /// one or the other.
    /// 
    /// [1]: https://www.terraform.io/docs/providers/aws/d/instance.html
    /// [2]: https://www.terraform.io/docs/providers/aws/r/network_interface.html
    /// 
    /// 
    /// ## Output Reference
    /// 
    /// There are no outputs for this resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/network_interface_sg_attachment.html.markdown.
    /// </summary>
    public partial class NetworkInterfaceSecurityGroupAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the network interface to attach to.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterfaceSecurityGroupAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfaceSecurityGroupAttachment(string name, NetworkInterfaceSecurityGroupAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfaceSecurityGroupAttachment(string name, Input<string> id, NetworkInterfaceSecurityGroupAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterfaceSecurityGroupAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfaceSecurityGroupAttachment Get(string name, Input<string> id, NetworkInterfaceSecurityGroupAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterfaceSecurityGroupAttachment(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceSecurityGroupAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the network interface to attach to.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        public NetworkInterfaceSecurityGroupAttachmentArgs()
        {
        }
    }

    public sealed class NetworkInterfaceSecurityGroupAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the network interface to attach to.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public NetworkInterfaceSecurityGroupAttachmentState()
        {
        }
    }
}
