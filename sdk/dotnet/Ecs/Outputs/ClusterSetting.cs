// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Outputs
{

    [OutputType]
    public sealed class ClusterSetting
    {
        /// <summary>
        /// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
        /// </summary>
        public readonly string Name;
        public readonly string Value;

        [OutputConstructor]
        private ClusterSetting(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
