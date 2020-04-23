// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpsWorks.Outputs
{

    [OutputType]
    public sealed class StackCustomCookbooksSource
    {
        /// <summary>
        /// Password to use when authenticating to the source. The provider cannot perform drift detection of this configuration.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// For sources that are version-aware, the revision to use.
        /// </summary>
        public readonly string? Revision;
        /// <summary>
        /// SSH key to use when authenticating to the source. The provider cannot perform drift detection of this configuration.
        /// </summary>
        public readonly string? SshKey;
        /// <summary>
        /// The type of source to use. For example, "archive".
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The URL where the cookbooks resource can be found.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Username to use when authenticating to the source.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private StackCustomCookbooksSource(
            string? password,

            string? revision,

            string? sshKey,

            string type,

            string url,

            string? username)
        {
            Password = password;
            Revision = revision;
            SshKey = sshKey;
            Type = type;
            Url = url;
            Username = username;
        }
    }
}