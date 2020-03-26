// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Qldb
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to fetch information about a Quantum Ledger Database.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/qldb_ledger.html.markdown.
        /// </summary>
        [Obsolete("Use GetLedger.InvokeAsync() instead")]
        public static Task<GetLedgerResult> GetLedger(GetLedgerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLedgerResult>("aws:qldb/getLedger:getLedger", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetLedger
    {
        /// <summary>
        /// Use this data source to fetch information about a Quantum Ledger Database.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/qldb_ledger.html.markdown.
        /// </summary>
        public static Task<GetLedgerResult> InvokeAsync(GetLedgerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLedgerResult>("aws:qldb/getLedger:getLedger", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetLedgerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The friendly name of the ledger to match.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetLedgerArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetLedgerResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the ledger.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Deletion protection on the QLDB Ledger instance. Set to `true` by default. 
        /// </summary>
        public readonly bool DeletionProtection;
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetLedgerResult(
            string arn,
            bool deletionProtection,
            string name,
            string id)
        {
            Arn = arn;
            DeletionProtection = deletionProtection;
            Name = name;
            Id = id;
        }
    }
}
