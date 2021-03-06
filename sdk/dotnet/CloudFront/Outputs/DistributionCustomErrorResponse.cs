// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Outputs
{

    [OutputType]
    public sealed class DistributionCustomErrorResponse
    {
        /// <summary>
        /// The minimum amount of time you want
        /// HTTP error codes to stay in CloudFront caches before CloudFront queries your
        /// origin to see whether the object has been updated.
        /// </summary>
        public readonly int? ErrorCachingMinTtl;
        /// <summary>
        /// The 4xx or 5xx HTTP status code that you want to
        /// customize.
        /// </summary>
        public readonly int ErrorCode;
        /// <summary>
        /// The HTTP status code that you want CloudFront
        /// to return with the custom error page to the viewer.
        /// </summary>
        public readonly int? ResponseCode;
        /// <summary>
        /// The path of the custom error page (for
        /// example, `/custom_404.html`).
        /// </summary>
        public readonly string? ResponsePagePath;

        [OutputConstructor]
        private DistributionCustomErrorResponse(
            int? errorCachingMinTtl,

            int errorCode,

            int? responseCode,

            string? responsePagePath)
        {
            ErrorCachingMinTtl = errorCachingMinTtl;
            ErrorCode = errorCode;
            ResponseCode = responseCode;
            ResponsePagePath = responsePagePath;
        }
    }
}
