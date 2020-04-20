// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Outputs
{

    [OutputType]
    public sealed class TablePointInTimeRecovery
    {
        /// <summary>
        /// Indicates whether ttl is enabled (true) or disabled (false).
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private TablePointInTimeRecovery(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
