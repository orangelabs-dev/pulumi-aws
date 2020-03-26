// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package transfer

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ARN of an AWS Transfer Server for use in other
// resources.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/transfer_server.html.markdown.
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("aws:transfer/getServer:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServer.
type LookupServerArgs struct {
	// ID for an SFTP server.
	ServerId string `pulumi:"serverId"`
}

// A collection of values returned by getServer.
type LookupServerResult struct {
	// Amazon Resource Name (ARN) of Transfer Server
	Arn string `pulumi:"arn"`
	// The endpoint of the Transfer Server (e.g. `s-12345678.server.transfer.REGION.amazonaws.com`)
	Endpoint string `pulumi:"endpoint"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
	IdentityProviderType string `pulumi:"identityProviderType"`
	// Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identityProviderType` of `API_GATEWAY`.
	InvocationRole string `pulumi:"invocationRole"`
	// Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
	LoggingRole string `pulumi:"loggingRole"`
	ServerId    string `pulumi:"serverId"`
	// URL of the service endpoint used to authenticate users with an `identityProviderType` of `API_GATEWAY`.
	Url string `pulumi:"url"`
}
