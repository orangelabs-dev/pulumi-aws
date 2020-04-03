// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.route53.ResolverRule` provides details about a specific Route53 Resolver rule.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route53_resolver_rule.html.markdown.
        /// </summary>
        [Obsolete("Use GetResolverRule.InvokeAsync() instead")]
        public static Task<GetResolverRuleResult> GetResolverRule(GetResolverRuleArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverRuleResult>("aws:route53/getResolverRule:getResolverRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetResolverRule
    {
        /// <summary>
        /// `aws.route53.ResolverRule` provides details about a specific Route53 Resolver rule.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route53_resolver_rule.html.markdown.
        /// </summary>
        public static Task<GetResolverRuleResult> InvokeAsync(GetResolverRuleArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverRuleResult>("aws:route53/getResolverRule:getResolverRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetResolverRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain name the desired resolver rule forwards DNS queries for. Conflicts with `resolver_rule_id`.
        /// </summary>
        [Input("domainName")]
        public string? DomainName { get; set; }

        /// <summary>
        /// The friendly name of the desired resolver rule. Conflicts with `resolver_rule_id`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the outbound resolver endpoint of the desired resolver rule. Conflicts with `resolver_rule_id`.
        /// </summary>
        [Input("resolverEndpointId")]
        public string? ResolverEndpointId { get; set; }

        /// <summary>
        /// The ID of the desired resolver rule. Conflicts with `domain_name`, `name`, `resolver_endpoint_id` and `rule_type`.
        /// </summary>
        [Input("resolverRuleId")]
        public string? ResolverRuleId { get; set; }

        /// <summary>
        /// The rule type of the desired resolver rule. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`. Conflicts with `resolver_rule_id`.
        /// </summary>
        [Input("ruleType")]
        public string? RuleType { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the resolver rule.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetResolverRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetResolverRuleResult
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) for the resolver rule.
        /// </summary>
        public readonly string Arn;
        public readonly string DomainName;
        public readonly string Name;
        /// <summary>
        /// When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
        /// </summary>
        public readonly string OwnerId;
        public readonly string ResolverEndpointId;
        public readonly string ResolverRuleId;
        public readonly string RuleType;
        /// <summary>
        /// Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
        /// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        /// </summary>
        public readonly string ShareStatus;
        /// <summary>
        /// A mapping of tags assigned to the resolver rule.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetResolverRuleResult(
            string arn,
            string domainName,
            string name,
            string ownerId,
            string resolverEndpointId,
            string resolverRuleId,
            string ruleType,
            string shareStatus,
            ImmutableDictionary<string, object> tags,
            string id)
        {
            Arn = arn;
            DomainName = domainName;
            Name = name;
            OwnerId = ownerId;
            ResolverEndpointId = resolverEndpointId;
            ResolverRuleId = resolverRuleId;
            RuleType = ruleType;
            ShareStatus = shareStatus;
            Tags = tags;
            Id = id;
        }
    }
}
