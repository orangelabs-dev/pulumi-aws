// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway.Outputs
{

    [OutputType]
    public sealed class MethodSettingsSettings
    {
        /// <summary>
        /// Specifies whether the cached responses are encrypted.
        /// </summary>
        public readonly bool? CacheDataEncrypted;
        /// <summary>
        /// Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
        /// </summary>
        public readonly int? CacheTtlInSeconds;
        /// <summary>
        /// Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached. 
        /// </summary>
        public readonly bool? CachingEnabled;
        /// <summary>
        /// Specifies whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
        /// </summary>
        public readonly bool? DataTraceEnabled;
        /// <summary>
        /// Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are `OFF`, `ERROR`, and `INFO`.
        /// </summary>
        public readonly string? LoggingLevel;
        /// <summary>
        /// Specifies whether Amazon CloudWatch metrics are enabled for this method.
        /// </summary>
        public readonly bool? MetricsEnabled;
        /// <summary>
        /// Specifies whether authorization is required for a cache invalidation request.
        /// </summary>
        public readonly bool? RequireAuthorizationForCacheControl;
        /// <summary>
        /// Specifies the throttling burst limit.
        /// </summary>
        public readonly int? ThrottlingBurstLimit;
        /// <summary>
        /// Specifies the throttling rate limit.
        /// </summary>
        public readonly double? ThrottlingRateLimit;
        /// <summary>
        /// Specifies how to handle unauthorized requests for cache invalidation. The available values are `FAIL_WITH_403`, `SUCCEED_WITH_RESPONSE_HEADER`, `SUCCEED_WITHOUT_RESPONSE_HEADER`.
        /// </summary>
        public readonly string? UnauthorizedCacheControlHeaderStrategy;

        [OutputConstructor]
        private MethodSettingsSettings(
            bool? cacheDataEncrypted,

            int? cacheTtlInSeconds,

            bool? cachingEnabled,

            bool? dataTraceEnabled,

            string? loggingLevel,

            bool? metricsEnabled,

            bool? requireAuthorizationForCacheControl,

            int? throttlingBurstLimit,

            double? throttlingRateLimit,

            string? unauthorizedCacheControlHeaderStrategy)
        {
            CacheDataEncrypted = cacheDataEncrypted;
            CacheTtlInSeconds = cacheTtlInSeconds;
            CachingEnabled = cachingEnabled;
            DataTraceEnabled = dataTraceEnabled;
            LoggingLevel = loggingLevel;
            MetricsEnabled = metricsEnabled;
            RequireAuthorizationForCacheControl = requireAuthorizationForCacheControl;
            ThrottlingBurstLimit = throttlingBurstLimit;
            ThrottlingRateLimit = throttlingRateLimit;
            UnauthorizedCacheControlHeaderStrategy = unauthorizedCacheControlHeaderStrategy;
        }
    }
}
