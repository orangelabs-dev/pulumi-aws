// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm.Outputs
{

    [OutputType]
    public sealed class MaintenanceWindowTaskLoggingInfo
    {
        public readonly string S3BucketName;
        public readonly string? S3BucketPrefix;
        public readonly string S3Region;

        [OutputConstructor]
        private MaintenanceWindowTaskLoggingInfo(
            string s3BucketName,

            string? s3BucketPrefix,

            string s3Region)
        {
            S3BucketName = s3BucketName;
            S3BucketPrefix = s3BucketPrefix;
            S3Region = s3Region;
        }
    }
}
