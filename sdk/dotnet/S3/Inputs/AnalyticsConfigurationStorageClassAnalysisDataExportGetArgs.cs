// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3.Inputs
{

    public sealed class AnalyticsConfigurationStorageClassAnalysisDataExportGetArgs : Pulumi.ResourceArgs
    {
        [Input("destination", required: true)]
        public Input<Inputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationGetArgs> Destination { get; set; } = null!;

        [Input("outputSchemaVersion")]
        public Input<string>? OutputSchemaVersion { get; set; }

        public AnalyticsConfigurationStorageClassAnalysisDataExportGetArgs()
        {
        }
    }
}
