// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Organizations
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get information about the organization that the user's account belongs to
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/organizations_organization.html.markdown.
        /// </summary>
        [Obsolete("Use GetOrganization.InvokeAsync() instead")]
        public static Task<GetOrganizationResult> GetOrganization(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationResult>("aws:organizations/getOrganization:getOrganization", InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetOrganization
    {
        /// <summary>
        /// Get information about the organization that the user's account belongs to
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/organizations_organization.html.markdown.
        /// </summary>
        public static Task<GetOrganizationResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationResult>("aws:organizations/getOrganization:getOrganization", InvokeArgs.Empty, options.WithVersion());
    }

    [OutputType]
    public sealed class GetOrganizationResult
    {
        /// <summary>
        /// List of organization accounts including the master account. For a list excluding the master account, see the `non_master_accounts` attribute. All elements have these attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationAccountsResult> Accounts;
        /// <summary>
        /// ARN of the root
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// A list of AWS service principal names that have integration enabled with your organization. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
        /// </summary>
        public readonly ImmutableArray<string> AwsServiceAccessPrincipals;
        /// <summary>
        /// A list of Organizations policy types that are enabled in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g. `SERVICE_CONTROL_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
        /// </summary>
        public readonly ImmutableArray<string> EnabledPolicyTypes;
        /// <summary>
        /// The FeatureSet of the organization.
        /// </summary>
        public readonly string FeatureSet;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the account that is designated as the master account for the organization.
        /// </summary>
        public readonly string MasterAccountArn;
        /// <summary>
        /// The email address that is associated with the AWS account that is designated as the master account for the organization.
        /// </summary>
        public readonly string MasterAccountEmail;
        /// <summary>
        /// The unique identifier (ID) of the master account of an organization.
        /// </summary>
        public readonly string MasterAccountId;
        /// <summary>
        /// List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationNonMasterAccountsResult> NonMasterAccounts;
        /// <summary>
        /// List of organization roots. All elements have these attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationRootsResult> Roots;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetOrganizationResult(
            ImmutableArray<Outputs.GetOrganizationAccountsResult> accounts,
            string arn,
            ImmutableArray<string> awsServiceAccessPrincipals,
            ImmutableArray<string> enabledPolicyTypes,
            string featureSet,
            string masterAccountArn,
            string masterAccountEmail,
            string masterAccountId,
            ImmutableArray<Outputs.GetOrganizationNonMasterAccountsResult> nonMasterAccounts,
            ImmutableArray<Outputs.GetOrganizationRootsResult> roots,
            string id)
        {
            Accounts = accounts;
            Arn = arn;
            AwsServiceAccessPrincipals = awsServiceAccessPrincipals;
            EnabledPolicyTypes = enabledPolicyTypes;
            FeatureSet = featureSet;
            MasterAccountArn = masterAccountArn;
            MasterAccountEmail = masterAccountEmail;
            MasterAccountId = masterAccountId;
            NonMasterAccounts = nonMasterAccounts;
            Roots = roots;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetOrganizationAccountsResult
    {
        /// <summary>
        /// ARN of the root
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Email of the account
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Identifier of the root
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the policy type
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of the policy type as it relates to the associated root
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetOrganizationAccountsResult(
            string arn,
            string email,
            string id,
            string name,
            string status)
        {
            Arn = arn;
            Email = email;
            Id = id;
            Name = name;
            Status = status;
        }
    }

    [OutputType]
    public sealed class GetOrganizationNonMasterAccountsResult
    {
        /// <summary>
        /// ARN of the root
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Email of the account
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Identifier of the root
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the policy type
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of the policy type as it relates to the associated root
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetOrganizationNonMasterAccountsResult(
            string arn,
            string email,
            string id,
            string name,
            string status)
        {
            Arn = arn;
            Email = email;
            Id = id;
            Name = name;
            Status = status;
        }
    }

    [OutputType]
    public sealed class GetOrganizationRootsPolicyTypesResult
    {
        /// <summary>
        /// The status of the policy type as it relates to the associated root
        /// </summary>
        public readonly string Status;
        public readonly string Type;

        [OutputConstructor]
        private GetOrganizationRootsPolicyTypesResult(
            string status,
            string type)
        {
            Status = status;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetOrganizationRootsResult
    {
        /// <summary>
        /// ARN of the root
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Identifier of the root
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the policy type
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of policy types enabled for this root. All elements have these attributes:
        /// </summary>
        public readonly ImmutableArray<GetOrganizationRootsPolicyTypesResult> PolicyTypes;

        [OutputConstructor]
        private GetOrganizationRootsResult(
            string arn,
            string id,
            string name,
            ImmutableArray<GetOrganizationRootsPolicyTypesResult> policyTypes)
        {
            Arn = arn;
            Id = id;
            Name = name;
            PolicyTypes = policyTypes;
        }
    }
    }
}
