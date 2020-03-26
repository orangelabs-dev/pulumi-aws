// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a Network Interface.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/network_interface.html.markdown.
func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("aws:ec2/getNetworkInterface:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkInterface.
type LookupNetworkInterfaceArgs struct {
	// One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-network-interfaces](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-network-interfaces.html) in the AWS CLI reference.
	Filters []GetNetworkInterfaceFilter `pulumi:"filters"`
	// The identifier for the network interface.
	Id *string `pulumi:"id"`
	// Any tags assigned to the network interface.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getNetworkInterface.
type LookupNetworkInterfaceResult struct {
	// The association information for an Elastic IP address (IPv4) associated with the network interface. See supported fields below.
	Associations []GetNetworkInterfaceAssociation    `pulumi:"associations"`
	Attachments  []GetNetworkInterfaceAttachmentType `pulumi:"attachments"`
	// The Availability Zone.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Description of the network interface.
	Description string                      `pulumi:"description"`
	Filters     []GetNetworkInterfaceFilter `pulumi:"filters"`
	Id          string                      `pulumi:"id"`
	// The type of interface.
	InterfaceType string `pulumi:"interfaceType"`
	// List of IPv6 addresses to assign to the ENI.
	Ipv6Addresses []string `pulumi:"ipv6Addresses"`
	// The MAC address.
	MacAddress string `pulumi:"macAddress"`
	// The AWS account ID of the owner of the network interface.
	OwnerId string `pulumi:"ownerId"`
	// The private DNS name.
	PrivateDnsName string `pulumi:"privateDnsName"`
	// The private IPv4 address of the network interface within the subnet.
	PrivateIp string `pulumi:"privateIp"`
	// The private IPv4 addresses associated with the network interface.
	PrivateIps []string `pulumi:"privateIps"`
	// The ID of the entity that launched the instance on your behalf.
	RequesterId string `pulumi:"requesterId"`
	// The list of security groups for the network interface.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The ID of the subnet.
	SubnetId string `pulumi:"subnetId"`
	// Any tags assigned to the network interface.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC.
	VpcId string `pulumi:"vpcId"`
}
