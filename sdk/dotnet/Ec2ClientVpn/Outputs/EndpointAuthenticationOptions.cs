// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2ClientVpn.Outputs
{

    [OutputType]
    public sealed class EndpointAuthenticationOptions
    {
        /// <summary>
        /// The ID of the Active Directory to be used for authentication if type is `directory-service-authentication`.
        /// </summary>
        public readonly string? ActiveDirectoryId;
        /// <summary>
        /// The ARN of the client certificate. The certificate must be signed by a certificate authority (CA) and it must be provisioned in AWS Certificate Manager (ACM). Only necessary when type is set to `certificate-authentication`.
        /// </summary>
        public readonly string? RootCertificateChainArn;
        /// <summary>
        /// The type of client authentication to be used. Specify `certificate-authentication` to use certificate-based authentication, or `directory-service-authentication` to use Active Directory authentication.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private EndpointAuthenticationOptions(
            string? activeDirectoryId,

            string? rootCertificateChainArn,

            string type)
        {
            ActiveDirectoryId = activeDirectoryId;
            RootCertificateChainArn = rootCertificateChainArn;
            Type = type;
        }
    }
}
