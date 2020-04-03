// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppSync
{
    /// <summary>
    /// Provides an AppSync GraphQL API.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_graphql_api.html.markdown.
    /// </summary>
    public partial class GraphQLApi : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more additional authentication providers for the GraphqlApi. Defined below.
        /// </summary>
        [Output("additionalAuthenticationProviders")]
        public Output<ImmutableArray<Outputs.GraphQLApiAdditionalAuthenticationProviders>> AdditionalAuthenticationProviders { get; private set; } = null!;

        /// <summary>
        /// The ARN
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        /// </summary>
        [Output("authenticationType")]
        public Output<string> AuthenticationType { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing logging configuration. Defined below.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.GraphQLApiLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// A user-supplied name for the GraphqlApi.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        [Output("openidConnectConfig")]
        public Output<Outputs.GraphQLApiOpenidConnectConfig?> OpenidConnectConfig { get; private set; } = null!;

        /// <summary>
        /// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        /// </summary>
        [Output("schema")]
        public Output<string?> Schema { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        /// </summary>
        [Output("uris")]
        public Output<ImmutableDictionary<string, string>> Uris { get; private set; } = null!;

        /// <summary>
        /// The Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        [Output("userPoolConfig")]
        public Output<Outputs.GraphQLApiUserPoolConfig?> UserPoolConfig { get; private set; } = null!;

        /// <summary>
        /// Whether tracing with X-ray is enabled. Defaults to false.
        /// </summary>
        [Output("xrayEnabled")]
        public Output<bool?> XrayEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a GraphQLApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GraphQLApi(string name, GraphQLApiArgs args, CustomResourceOptions? options = null)
            : base("aws:appsync/graphQLApi:GraphQLApi", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GraphQLApi(string name, Input<string> id, GraphQLApiState? state = null, CustomResourceOptions? options = null)
            : base("aws:appsync/graphQLApi:GraphQLApi", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GraphQLApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GraphQLApi Get(string name, Input<string> id, GraphQLApiState? state = null, CustomResourceOptions? options = null)
        {
            return new GraphQLApi(name, id, state, options);
        }
    }

    public sealed class GraphQLApiArgs : Pulumi.ResourceArgs
    {
        [Input("additionalAuthenticationProviders")]
        private InputList<Inputs.GraphQLApiAdditionalAuthenticationProvidersArgs>? _additionalAuthenticationProviders;

        /// <summary>
        /// One or more additional authentication providers for the GraphqlApi. Defined below.
        /// </summary>
        public InputList<Inputs.GraphQLApiAdditionalAuthenticationProvidersArgs> AdditionalAuthenticationProviders
        {
            get => _additionalAuthenticationProviders ?? (_additionalAuthenticationProviders = new InputList<Inputs.GraphQLApiAdditionalAuthenticationProvidersArgs>());
            set => _additionalAuthenticationProviders = value;
        }

        /// <summary>
        /// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// Nested argument containing logging configuration. Defined below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.GraphQLApiLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// A user-supplied name for the GraphqlApi.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        [Input("openidConnectConfig")]
        public Input<Inputs.GraphQLApiOpenidConnectConfigArgs>? OpenidConnectConfig { get; set; }

        /// <summary>
        /// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        [Input("userPoolConfig")]
        public Input<Inputs.GraphQLApiUserPoolConfigArgs>? UserPoolConfig { get; set; }

        /// <summary>
        /// Whether tracing with X-ray is enabled. Defaults to false.
        /// </summary>
        [Input("xrayEnabled")]
        public Input<bool>? XrayEnabled { get; set; }

        public GraphQLApiArgs()
        {
        }
    }

    public sealed class GraphQLApiState : Pulumi.ResourceArgs
    {
        [Input("additionalAuthenticationProviders")]
        private InputList<Inputs.GraphQLApiAdditionalAuthenticationProvidersGetArgs>? _additionalAuthenticationProviders;

        /// <summary>
        /// One or more additional authentication providers for the GraphqlApi. Defined below.
        /// </summary>
        public InputList<Inputs.GraphQLApiAdditionalAuthenticationProvidersGetArgs> AdditionalAuthenticationProviders
        {
            get => _additionalAuthenticationProviders ?? (_additionalAuthenticationProviders = new InputList<Inputs.GraphQLApiAdditionalAuthenticationProvidersGetArgs>());
            set => _additionalAuthenticationProviders = value;
        }

        /// <summary>
        /// The ARN
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        /// </summary>
        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        /// <summary>
        /// Nested argument containing logging configuration. Defined below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.GraphQLApiLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// A user-supplied name for the GraphqlApi.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        [Input("openidConnectConfig")]
        public Input<Inputs.GraphQLApiOpenidConnectConfigGetArgs>? OpenidConnectConfig { get; set; }

        /// <summary>
        /// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("uris")]
        private InputMap<string>? _uris;

        /// <summary>
        /// Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        /// </summary>
        public InputMap<string> Uris
        {
            get => _uris ?? (_uris = new InputMap<string>());
            set => _uris = value;
        }

        /// <summary>
        /// The Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        [Input("userPoolConfig")]
        public Input<Inputs.GraphQLApiUserPoolConfigGetArgs>? UserPoolConfig { get; set; }

        /// <summary>
        /// Whether tracing with X-ray is enabled. Defaults to false.
        /// </summary>
        [Input("xrayEnabled")]
        public Input<bool>? XrayEnabled { get; set; }

        public GraphQLApiState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class GraphQLApiAdditionalAuthenticationProvidersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        [Input("openidConnectConfig")]
        public Input<GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfigArgs>? OpenidConnectConfig { get; set; }

        /// <summary>
        /// The Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        [Input("userPoolConfig")]
        public Input<GraphQLApiAdditionalAuthenticationProvidersUserPoolConfigArgs>? UserPoolConfig { get; set; }

        public GraphQLApiAdditionalAuthenticationProvidersArgs()
        {
        }
    }

    public sealed class GraphQLApiAdditionalAuthenticationProvidersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        [Input("openidConnectConfig")]
        public Input<GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfigGetArgs>? OpenidConnectConfig { get; set; }

        /// <summary>
        /// The Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        [Input("userPoolConfig")]
        public Input<GraphQLApiAdditionalAuthenticationProvidersUserPoolConfigGetArgs>? UserPoolConfig { get; set; }

        public GraphQLApiAdditionalAuthenticationProvidersGetArgs()
        {
        }
    }

    public sealed class GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of milliseconds a token is valid after being authenticated.
        /// </summary>
        [Input("authTtl")]
        public Input<int>? AuthTtl { get; set; }

        /// <summary>
        /// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Number of milliseconds a token is valid after being issued to a user.
        /// </summary>
        [Input("iatTtl")]
        public Input<int>? IatTtl { get; set; }

        /// <summary>
        /// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        public GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfigArgs()
        {
        }
    }

    public sealed class GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of milliseconds a token is valid after being authenticated.
        /// </summary>
        [Input("authTtl")]
        public Input<int>? AuthTtl { get; set; }

        /// <summary>
        /// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Number of milliseconds a token is valid after being issued to a user.
        /// </summary>
        [Input("iatTtl")]
        public Input<int>? IatTtl { get; set; }

        /// <summary>
        /// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        public GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfigGetArgs()
        {
        }
    }

    public sealed class GraphQLApiAdditionalAuthenticationProvidersUserPoolConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
        /// </summary>
        [Input("appIdClientRegex")]
        public Input<string>? AppIdClientRegex { get; set; }

        /// <summary>
        /// The AWS region in which the user pool was created.
        /// </summary>
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        /// <summary>
        /// The user pool ID.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public GraphQLApiAdditionalAuthenticationProvidersUserPoolConfigArgs()
        {
        }
    }

    public sealed class GraphQLApiAdditionalAuthenticationProvidersUserPoolConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
        /// </summary>
        [Input("appIdClientRegex")]
        public Input<string>? AppIdClientRegex { get; set; }

        /// <summary>
        /// The AWS region in which the user pool was created.
        /// </summary>
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        /// <summary>
        /// The user pool ID.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public GraphQLApiAdditionalAuthenticationProvidersUserPoolConfigGetArgs()
        {
        }
    }

    public sealed class GraphQLApiLogConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
        /// </summary>
        [Input("cloudwatchLogsRoleArn", required: true)]
        public Input<string> CloudwatchLogsRoleArn { get; set; } = null!;

        /// <summary>
        /// Field logging level. Valid values: `ALL`, `ERROR`, `NONE`.
        /// </summary>
        [Input("fieldLogLevel", required: true)]
        public Input<string> FieldLogLevel { get; set; } = null!;

        public GraphQLApiLogConfigArgs()
        {
        }
    }

    public sealed class GraphQLApiLogConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
        /// </summary>
        [Input("cloudwatchLogsRoleArn", required: true)]
        public Input<string> CloudwatchLogsRoleArn { get; set; } = null!;

        /// <summary>
        /// Field logging level. Valid values: `ALL`, `ERROR`, `NONE`.
        /// </summary>
        [Input("fieldLogLevel", required: true)]
        public Input<string> FieldLogLevel { get; set; } = null!;

        public GraphQLApiLogConfigGetArgs()
        {
        }
    }

    public sealed class GraphQLApiOpenidConnectConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of milliseconds a token is valid after being authenticated.
        /// </summary>
        [Input("authTtl")]
        public Input<int>? AuthTtl { get; set; }

        /// <summary>
        /// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Number of milliseconds a token is valid after being issued to a user.
        /// </summary>
        [Input("iatTtl")]
        public Input<int>? IatTtl { get; set; }

        /// <summary>
        /// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        public GraphQLApiOpenidConnectConfigArgs()
        {
        }
    }

    public sealed class GraphQLApiOpenidConnectConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of milliseconds a token is valid after being authenticated.
        /// </summary>
        [Input("authTtl")]
        public Input<int>? AuthTtl { get; set; }

        /// <summary>
        /// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Number of milliseconds a token is valid after being issued to a user.
        /// </summary>
        [Input("iatTtl")]
        public Input<int>? IatTtl { get; set; }

        /// <summary>
        /// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        public GraphQLApiOpenidConnectConfigGetArgs()
        {
        }
    }

    public sealed class GraphQLApiUserPoolConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
        /// </summary>
        [Input("appIdClientRegex")]
        public Input<string>? AppIdClientRegex { get; set; }

        /// <summary>
        /// The AWS region in which the user pool was created.
        /// </summary>
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        /// <summary>
        /// The action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: `ALLOW` and `DENY`
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        /// <summary>
        /// The user pool ID.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public GraphQLApiUserPoolConfigArgs()
        {
        }
    }

    public sealed class GraphQLApiUserPoolConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
        /// </summary>
        [Input("appIdClientRegex")]
        public Input<string>? AppIdClientRegex { get; set; }

        /// <summary>
        /// The AWS region in which the user pool was created.
        /// </summary>
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        /// <summary>
        /// The action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: `ALLOW` and `DENY`
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        /// <summary>
        /// The user pool ID.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public GraphQLApiUserPoolConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GraphQLApiAdditionalAuthenticationProviders
    {
        /// <summary>
        /// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        /// </summary>
        public readonly string AuthenticationType;
        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        public readonly GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfig? OpenidConnectConfig;
        /// <summary>
        /// The Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        public readonly GraphQLApiAdditionalAuthenticationProvidersUserPoolConfig? UserPoolConfig;

        [OutputConstructor]
        private GraphQLApiAdditionalAuthenticationProviders(
            string authenticationType,
            GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfig? openidConnectConfig,
            GraphQLApiAdditionalAuthenticationProvidersUserPoolConfig? userPoolConfig)
        {
            AuthenticationType = authenticationType;
            OpenidConnectConfig = openidConnectConfig;
            UserPoolConfig = userPoolConfig;
        }
    }

    [OutputType]
    public sealed class GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfig
    {
        /// <summary>
        /// Number of milliseconds a token is valid after being authenticated.
        /// </summary>
        public readonly int? AuthTtl;
        /// <summary>
        /// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// Number of milliseconds a token is valid after being issued to a user.
        /// </summary>
        public readonly int? IatTtl;
        /// <summary>
        /// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
        /// </summary>
        public readonly string Issuer;

        [OutputConstructor]
        private GraphQLApiAdditionalAuthenticationProvidersOpenidConnectConfig(
            int? authTtl,
            string? clientId,
            int? iatTtl,
            string issuer)
        {
            AuthTtl = authTtl;
            ClientId = clientId;
            IatTtl = iatTtl;
            Issuer = issuer;
        }
    }

    [OutputType]
    public sealed class GraphQLApiAdditionalAuthenticationProvidersUserPoolConfig
    {
        /// <summary>
        /// A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
        /// </summary>
        public readonly string? AppIdClientRegex;
        /// <summary>
        /// The AWS region in which the user pool was created.
        /// </summary>
        public readonly string AwsRegion;
        /// <summary>
        /// The user pool ID.
        /// </summary>
        public readonly string UserPoolId;

        [OutputConstructor]
        private GraphQLApiAdditionalAuthenticationProvidersUserPoolConfig(
            string? appIdClientRegex,
            string awsRegion,
            string userPoolId)
        {
            AppIdClientRegex = appIdClientRegex;
            AwsRegion = awsRegion;
            UserPoolId = userPoolId;
        }
    }

    [OutputType]
    public sealed class GraphQLApiLogConfig
    {
        /// <summary>
        /// Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
        /// </summary>
        public readonly string CloudwatchLogsRoleArn;
        /// <summary>
        /// Field logging level. Valid values: `ALL`, `ERROR`, `NONE`.
        /// </summary>
        public readonly string FieldLogLevel;

        [OutputConstructor]
        private GraphQLApiLogConfig(
            string cloudwatchLogsRoleArn,
            string fieldLogLevel)
        {
            CloudwatchLogsRoleArn = cloudwatchLogsRoleArn;
            FieldLogLevel = fieldLogLevel;
        }
    }

    [OutputType]
    public sealed class GraphQLApiOpenidConnectConfig
    {
        /// <summary>
        /// Number of milliseconds a token is valid after being authenticated.
        /// </summary>
        public readonly int? AuthTtl;
        /// <summary>
        /// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// Number of milliseconds a token is valid after being issued to a user.
        /// </summary>
        public readonly int? IatTtl;
        /// <summary>
        /// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
        /// </summary>
        public readonly string Issuer;

        [OutputConstructor]
        private GraphQLApiOpenidConnectConfig(
            int? authTtl,
            string? clientId,
            int? iatTtl,
            string issuer)
        {
            AuthTtl = authTtl;
            ClientId = clientId;
            IatTtl = iatTtl;
            Issuer = issuer;
        }
    }

    [OutputType]
    public sealed class GraphQLApiUserPoolConfig
    {
        /// <summary>
        /// A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
        /// </summary>
        public readonly string? AppIdClientRegex;
        /// <summary>
        /// The AWS region in which the user pool was created.
        /// </summary>
        public readonly string AwsRegion;
        /// <summary>
        /// The action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: `ALLOW` and `DENY`
        /// </summary>
        public readonly string DefaultAction;
        /// <summary>
        /// The user pool ID.
        /// </summary>
        public readonly string UserPoolId;

        [OutputConstructor]
        private GraphQLApiUserPoolConfig(
            string? appIdClientRegex,
            string awsRegion,
            string defaultAction,
            string userPoolId)
        {
            AppIdClientRegex = appIdClientRegex;
            AwsRegion = awsRegion;
            DefaultAction = defaultAction;
            UserPoolId = userPoolId;
        }
    }
    }
}
