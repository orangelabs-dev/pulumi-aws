// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Alb.Inputs
{

    public sealed class ListenerDefaultActionGetArgs : Pulumi.ResourceArgs
    {
        [Input("authenticateCognito")]
        public Input<Inputs.ListenerDefaultActionAuthenticateCognitoGetArgs>? AuthenticateCognito { get; set; }

        [Input("authenticateOidc")]
        public Input<Inputs.ListenerDefaultActionAuthenticateOidcGetArgs>? AuthenticateOidc { get; set; }

        /// <summary>
        /// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
        /// </summary>
        [Input("fixedResponse")]
        public Input<Inputs.ListenerDefaultActionFixedResponseGetArgs>? FixedResponse { get; set; }

        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// Information for creating a redirect action. Required if `type` is `redirect`.
        /// </summary>
        [Input("redirect")]
        public Input<Inputs.ListenerDefaultActionRedirectGetArgs>? Redirect { get; set; }

        /// <summary>
        /// The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
        /// </summary>
        [Input("targetGroupArn")]
        public Input<string>? TargetGroupArn { get; set; }

        /// <summary>
        /// The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ListenerDefaultActionGetArgs()
        {
        }
    }
}
