# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VpcEndpoint(pulumi.CustomResource):
    auto_accept: pulumi.Output[bool]
    """
    Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
    """
    cidr_blocks: pulumi.Output[list]
    """
    The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
    """
    dns_entries: pulumi.Output[list]
    """
    The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
    """
    network_interface_ids: pulumi.Output[list]
    """
    One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
    """
    policy: pulumi.Output[str]
    """
    A policy to attach to the endpoint that controls access to the service. Applicable for endpoints of type `Gateway`. Defaults to full access. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
    """
    prefix_list_id: pulumi.Output[str]
    """
    The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
    """
    private_dns_enabled: pulumi.Output[bool]
    """
    Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
    Defaults to `false`.
    """
    route_table_ids: pulumi.Output[list]
    """
    One or more route table IDs. Applicable for endpoints of type `Gateway`.
    """
    security_group_ids: pulumi.Output[list]
    """
    The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
    """
    service_name: pulumi.Output[str]
    """
    The service name, in the form `com.amazonaws.region.service` for AWS services.
    """
    state: pulumi.Output[str]
    """
    The state of the VPC endpoint.
    """
    subnet_ids: pulumi.Output[list]
    """
    The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
    """
    vpc_endpoint_type: pulumi.Output[str]
    """
    The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
    """
    vpc_id: pulumi.Output[str]
    """
    The ID of the VPC in which the endpoint will be used.
    """
    def __init__(__self__, __name__, __opts__=None, auto_accept=None, policy=None, private_dns_enabled=None, route_table_ids=None, security_group_ids=None, service_name=None, subnet_ids=None, vpc_endpoint_type=None, vpc_id=None):
        """
        Provides a VPC Endpoint resource.
        
        > **NOTE on VPC Endpoints and VPC Endpoint Associations:** Terraform provides both standalone VPC Endpoint Associations for
        Route Tables - (an association between a VPC endpoint and a single `route_table_id`) and
        Subnets - (an association between a VPC endpoint and a single `subnet_id`) and
        a VPC Endpoint resource with `route_table_ids` and `subnet_ids` attributes.
        Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
        Doing so will cause a conflict of associations and will overwrite the association.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[bool] auto_accept: Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
        :param pulumi.Input[str] policy: A policy to attach to the endpoint that controls access to the service. Applicable for endpoints of type `Gateway`. Defaults to full access. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
        :param pulumi.Input[bool] private_dns_enabled: Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
               Defaults to `false`.
        :param pulumi.Input[list] route_table_ids: One or more route table IDs. Applicable for endpoints of type `Gateway`.
        :param pulumi.Input[list] security_group_ids: The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
        :param pulumi.Input[str] service_name: The service name, in the form `com.amazonaws.region.service` for AWS services.
        :param pulumi.Input[list] subnet_ids: The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
        :param pulumi.Input[str] vpc_endpoint_type: The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
        :param pulumi.Input[str] vpc_id: The ID of the VPC in which the endpoint will be used.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['auto_accept'] = auto_accept

        __props__['policy'] = policy

        __props__['private_dns_enabled'] = private_dns_enabled

        __props__['route_table_ids'] = route_table_ids

        __props__['security_group_ids'] = security_group_ids

        if not service_name:
            raise TypeError('Missing required property service_name')
        __props__['service_name'] = service_name

        __props__['subnet_ids'] = subnet_ids

        __props__['vpc_endpoint_type'] = vpc_endpoint_type

        if not vpc_id:
            raise TypeError('Missing required property vpc_id')
        __props__['vpc_id'] = vpc_id

        __props__['cidr_blocks'] = None
        __props__['dns_entries'] = None
        __props__['network_interface_ids'] = None
        __props__['prefix_list_id'] = None
        __props__['state'] = None

        super(VpcEndpoint, __self__).__init__(
            'aws:ec2/vpcEndpoint:VpcEndpoint',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

