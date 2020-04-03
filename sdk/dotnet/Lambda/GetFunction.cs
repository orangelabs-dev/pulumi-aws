// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information about a Lambda Function.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_function.html.markdown.
        /// </summary>
        [Obsolete("Use GetFunction.InvokeAsync() instead")]
        public static Task<GetFunctionResult> GetFunction(GetFunctionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("aws:lambda/getFunction:getFunction", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetFunction
    {
        /// <summary>
        /// Provides information about a Lambda Function.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_function.html.markdown.
        /// </summary>
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("aws:lambda/getFunction:getFunction", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetFunctionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the lambda function.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        /// <summary>
        /// Alias name or version number of the lambda function. e.g. `$LATEST`, `my-alias`, or `1`
        /// </summary>
        [Input("qualifier")]
        public string? Qualifier { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetFunctionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetFunctionResult
    {
        /// <summary>
        /// Unqualified (no `:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `qualified_arn`.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Configure the function's *dead letter queue*.
        /// </summary>
        public readonly Outputs.GetFunctionDeadLetterConfigResult DeadLetterConfig;
        /// <summary>
        /// Description of what your Lambda Function does.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Lambda environment's configuration settings.
        /// </summary>
        public readonly Outputs.GetFunctionEnvironmentResult Environment;
        public readonly string FunctionName;
        /// <summary>
        /// The function entrypoint in your code.
        /// </summary>
        public readonly string Handler;
        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway.
        /// </summary>
        public readonly string InvokeArn;
        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        public readonly string KmsKeyArn;
        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// A list of Lambda Layer ARNs attached to your Lambda Function.
        /// </summary>
        public readonly ImmutableArray<string> Layers;
        /// <summary>
        /// Amount of memory in MB your Lambda Function can use at runtime.
        /// </summary>
        public readonly int MemorySize;
        /// <summary>
        /// Qualified (`:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `arn`.
        /// </summary>
        public readonly string QualifiedArn;
        public readonly string? Qualifier;
        /// <summary>
        /// The amount of reserved concurrent executions for this lambda function or `-1` if unreserved.
        /// </summary>
        public readonly int ReservedConcurrentExecutions;
        /// <summary>
        /// IAM role attached to the Lambda Function.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// The runtime environment for the Lambda function..
        /// </summary>
        public readonly string Runtime;
        /// <summary>
        /// Base64-encoded representation of raw SHA-256 sum of the zip file.
        /// </summary>
        public readonly string SourceCodeHash;
        /// <summary>
        /// The size in bytes of the function .zip file.
        /// </summary>
        public readonly int SourceCodeSize;
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The function execution time at which Lambda should terminate the function.
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// Tracing settings of the function.
        /// </summary>
        public readonly Outputs.GetFunctionTracingConfigResult TracingConfig;
        /// <summary>
        /// The version of the Lambda function.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// VPC configuration associated with your Lambda function.
        /// </summary>
        public readonly Outputs.GetFunctionVpcConfigResult VpcConfig;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFunctionResult(
            string arn,
            Outputs.GetFunctionDeadLetterConfigResult deadLetterConfig,
            string description,
            Outputs.GetFunctionEnvironmentResult environment,
            string functionName,
            string handler,
            string invokeArn,
            string kmsKeyArn,
            string lastModified,
            ImmutableArray<string> layers,
            int memorySize,
            string qualifiedArn,
            string? qualifier,
            int reservedConcurrentExecutions,
            string role,
            string runtime,
            string sourceCodeHash,
            int sourceCodeSize,
            ImmutableDictionary<string, object> tags,
            int timeout,
            Outputs.GetFunctionTracingConfigResult tracingConfig,
            string version,
            Outputs.GetFunctionVpcConfigResult vpcConfig,
            string id)
        {
            Arn = arn;
            DeadLetterConfig = deadLetterConfig;
            Description = description;
            Environment = environment;
            FunctionName = functionName;
            Handler = handler;
            InvokeArn = invokeArn;
            KmsKeyArn = kmsKeyArn;
            LastModified = lastModified;
            Layers = layers;
            MemorySize = memorySize;
            QualifiedArn = qualifiedArn;
            Qualifier = qualifier;
            ReservedConcurrentExecutions = reservedConcurrentExecutions;
            Role = role;
            Runtime = runtime;
            SourceCodeHash = sourceCodeHash;
            SourceCodeSize = sourceCodeSize;
            Tags = tags;
            Timeout = timeout;
            TracingConfig = tracingConfig;
            Version = version;
            VpcConfig = vpcConfig;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetFunctionDeadLetterConfigResult
    {
        public readonly string TargetArn;

        [OutputConstructor]
        private GetFunctionDeadLetterConfigResult(string targetArn)
        {
            TargetArn = targetArn;
        }
    }

    [OutputType]
    public sealed class GetFunctionEnvironmentResult
    {
        public readonly ImmutableDictionary<string, string> Variables;

        [OutputConstructor]
        private GetFunctionEnvironmentResult(ImmutableDictionary<string, string> variables)
        {
            Variables = variables;
        }
    }

    [OutputType]
    public sealed class GetFunctionTracingConfigResult
    {
        public readonly string Mode;

        [OutputConstructor]
        private GetFunctionTracingConfigResult(string mode)
        {
            Mode = mode;
        }
    }

    [OutputType]
    public sealed class GetFunctionVpcConfigResult
    {
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly ImmutableArray<string> SubnetIds;
        public readonly string VpcId;

        [OutputConstructor]
        private GetFunctionVpcConfigResult(
            ImmutableArray<string> securityGroupIds,
            ImmutableArray<string> subnetIds,
            string vpcId)
        {
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
    }
}
