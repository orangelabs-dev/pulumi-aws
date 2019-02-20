# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VpcEndpointRouteTableAssociation(pulumi.CustomResource):
    route_table_id: pulumi.Output[str]
    """
    Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
    """
    vpc_endpoint_id: pulumi.Output[str]
    """
    Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
    """
    def __init__(__self__, __name__, __opts__=None, route_table_id=None, vpc_endpoint_id=None):
        """
        Manages a VPC Endpoint Route Table Association
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] route_table_id: Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
        :param pulumi.Input[str] vpc_endpoint_id: Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not route_table_id:
            raise TypeError('Missing required property route_table_id')
        __props__['route_table_id'] = route_table_id

        if not vpc_endpoint_id:
            raise TypeError('Missing required property vpc_endpoint_id')
        __props__['vpc_endpoint_id'] = vpc_endpoint_id

        super(VpcEndpointRouteTableAssociation, __self__).__init__(
            'aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

