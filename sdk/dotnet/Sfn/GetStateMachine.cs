// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sfn
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ARN of a State Machine in AWS Step
        /// Function (SFN). By using this data source, you can reference a
        /// state machine without having to hard code the ARNs as input.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/sfn_state_machine.html.markdown.
        /// </summary>
        [Obsolete("Use GetStateMachine.InvokeAsync() instead")]
        public static Task<GetStateMachineResult> GetStateMachine(GetStateMachineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStateMachineResult>("aws:sfn/getStateMachine:getStateMachine", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetStateMachine
    {
        /// <summary>
        /// Use this data source to get the ARN of a State Machine in AWS Step
        /// Function (SFN). By using this data source, you can reference a
        /// state machine without having to hard code the ARNs as input.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/sfn_state_machine.html.markdown.
        /// </summary>
        public static Task<GetStateMachineResult> InvokeAsync(GetStateMachineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStateMachineResult>("aws:sfn/getStateMachine:getStateMachine", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetStateMachineArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The friendly name of the state machine to match.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetStateMachineArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetStateMachineResult
    {
        /// <summary>
        /// Set to the arn of the state function.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The date the state machine was created.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// Set to the state machine definition.
        /// </summary>
        public readonly string Definition;
        public readonly string Name;
        /// <summary>
        /// Set to the role_arn used by the state function.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// Set to the current status of the state machine.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetStateMachineResult(
            string arn,
            string creationDate,
            string definition,
            string name,
            string roleArn,
            string status,
            string id)
        {
            Arn = arn;
            CreationDate = creationDate;
            Definition = definition;
            Name = name;
            RoleArn = roleArn;
            Status = status;
            Id = id;
        }
    }
}
