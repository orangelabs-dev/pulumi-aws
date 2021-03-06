// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Inputs
{

    public sealed class ServiceServiceRegistriesGetArgs : Pulumi.ResourceArgs
    {
        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        [Input("containerPort")]
        public Input<int>? ContainerPort { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("registryArn", required: true)]
        public Input<string> RegistryArn { get; set; } = null!;

        public ServiceServiceRegistriesGetArgs()
        {
        }
    }
}
