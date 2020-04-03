// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecretsManager
{
    public static partial class Invokes
    {
        /// <summary>
        /// Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the [`aws.secretsmanager.SecretVersion` data source](https://www.terraform.io/docs/providers/aws/d/secretsmanager_secret_version.html).
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/secretsmanager_secret.html.markdown.
        /// </summary>
        [Obsolete("Use GetSecret.InvokeAsync() instead")]
        public static Task<GetSecretResult> GetSecret(GetSecretArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretResult>("aws:secretsmanager/getSecret:getSecret", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSecret
    {
        /// <summary>
        /// Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the [`aws.secretsmanager.SecretVersion` data source](https://www.terraform.io/docs/providers/aws/d/secretsmanager_secret_version.html).
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/secretsmanager_secret.html.markdown.
        /// </summary>
        public static Task<GetSecretResult> InvokeAsync(GetSecretArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretResult>("aws:secretsmanager/getSecret:getSecret", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSecretArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the secret to retrieve.
        /// </summary>
        [Input("arn")]
        public string? Arn { get; set; }

        /// <summary>
        /// The name of the secret to retrieve.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetSecretArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSecretResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the secret.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// A description of the secret.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.
        /// </summary>
        public readonly string KmsKeyId;
        public readonly string Name;
        /// <summary>
        /// The resource-based policy document that's attached to the secret.
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// Whether rotation is enabled or not.
        /// </summary>
        public readonly bool RotationEnabled;
        /// <summary>
        /// Rotation Lambda function Amazon Resource Name (ARN) if rotation is enabled.
        /// </summary>
        public readonly string RotationLambdaArn;
        /// <summary>
        /// Rotation rules if rotation is enabled.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecretRotationRulesResult> RotationRules;
        /// <summary>
        /// Tags of the secret.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSecretResult(
            string arn,
            string description,
            string kmsKeyId,
            string name,
            string policy,
            bool rotationEnabled,
            string rotationLambdaArn,
            ImmutableArray<Outputs.GetSecretRotationRulesResult> rotationRules,
            ImmutableDictionary<string, object> tags,
            string id)
        {
            Arn = arn;
            Description = description;
            KmsKeyId = kmsKeyId;
            Name = name;
            Policy = policy;
            RotationEnabled = rotationEnabled;
            RotationLambdaArn = rotationLambdaArn;
            RotationRules = rotationRules;
            Tags = tags;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetSecretRotationRulesResult
    {
        public readonly int AutomaticallyAfterDays;

        [OutputConstructor]
        private GetSecretRotationRulesResult(int automaticallyAfterDays)
        {
            AutomaticallyAfterDays = automaticallyAfterDays;
        }
    }
    }
}
