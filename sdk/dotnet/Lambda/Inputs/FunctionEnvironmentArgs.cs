// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda.Inputs
{

    public sealed class FunctionEnvironmentArgs : Pulumi.ResourceArgs
    {
        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// A map that defines environment variables for the Lambda function.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public FunctionEnvironmentArgs()
        {
        }
    }
}
