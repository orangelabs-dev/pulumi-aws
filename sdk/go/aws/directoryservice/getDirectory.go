// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package directoryservice

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get attributes of AWS Directory Service directory (SimpleAD, Managed AD, AD Connector). It's especially useful to refer AWS Managed AD or on-premise AD in AD Connector configuration.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/directory_service_directory.html.markdown.
func LookupDirectory(ctx *pulumi.Context, args *LookupDirectoryArgs, opts ...pulumi.InvokeOption) (*LookupDirectoryResult, error) {
	var rv LookupDirectoryResult
	err := ctx.Invoke("aws:directoryservice/getDirectory:getDirectory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDirectory.
type LookupDirectoryArgs struct {
	// The ID of the directory.
	DirectoryId string `pulumi:"directoryId"`
	// A mapping of tags assigned to the directory/connector.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getDirectory.
type LookupDirectoryResult struct {
	// The access URL for the directory/connector, such as http://alias.awsapps.com.
	AccessUrl string `pulumi:"accessUrl"`
	// The alias for the directory/connector, such as `d-991708b282.awsapps.com`.
	Alias           string                       `pulumi:"alias"`
	ConnectSettings []GetDirectoryConnectSetting `pulumi:"connectSettings"`
	// A textual description for the directory/connector.
	Description string `pulumi:"description"`
	DirectoryId string `pulumi:"directoryId"`
	// A list of IP addresses of the DNS servers for the directory/connector.
	DnsIpAddresses []string `pulumi:"dnsIpAddresses"`
	// (for `MicrosoftAD`) The Microsoft AD edition (`Standard` or `Enterprise`).
	Edition string `pulumi:"edition"`
	// The directory/connector single-sign on status.
	EnableSso bool `pulumi:"enableSso"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The fully qualified name for the directory/connector.
	Name string `pulumi:"name"`
	// The ID of the security group created by the directory/connector.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The short name of the directory/connector, such as `CORP`.
	ShortName string `pulumi:"shortName"`
	// (for `SimpleAD` and `ADConnector`) The size of the directory/connector (`Small` or `Large`).
	Size string `pulumi:"size"`
	// A mapping of tags assigned to the directory/connector.
	Tags map[string]interface{} `pulumi:"tags"`
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD`).
	Type        string                   `pulumi:"type"`
	VpcSettings []GetDirectoryVpcSetting `pulumi:"vpcSettings"`
}
