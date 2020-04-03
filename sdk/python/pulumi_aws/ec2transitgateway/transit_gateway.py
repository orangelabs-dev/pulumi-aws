# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class TransitGateway(pulumi.CustomResource):
    amazon_side_asn: pulumi.Output[float]
    """
    Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
    """
    arn: pulumi.Output[str]
    """
    EC2 Transit Gateway Amazon Resource Name (ARN)
    """
    association_default_route_table_id: pulumi.Output[str]
    """
    Identifier of the default association route table
    """
    auto_accept_shared_attachments: pulumi.Output[str]
    """
    Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
    """
    default_route_table_association: pulumi.Output[str]
    """
    Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
    """
    default_route_table_propagation: pulumi.Output[str]
    """
    Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
    """
    description: pulumi.Output[str]
    """
    Description of the EC2 Transit Gateway.
    """
    dns_support: pulumi.Output[str]
    """
    Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
    """
    owner_id: pulumi.Output[str]
    """
    Identifier of the AWS account that owns the EC2 Transit Gateway
    """
    propagation_default_route_table_id: pulumi.Output[str]
    """
    Identifier of the default propagation route table
    """
    tags: pulumi.Output[dict]
    """
    Key-value tags for the EC2 Transit Gateway.
    """
    vpn_ecmp_support: pulumi.Output[str]
    """
    Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
    """
    def __init__(__self__, resource_name, opts=None, amazon_side_asn=None, auto_accept_shared_attachments=None, default_route_table_association=None, default_route_table_propagation=None, description=None, dns_support=None, tags=None, vpn_ecmp_support=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an EC2 Transit Gateway.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] amazon_side_asn: Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
        :param pulumi.Input[str] auto_accept_shared_attachments: Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
        :param pulumi.Input[str] default_route_table_association: Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] default_route_table_propagation: Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] description: Description of the EC2 Transit Gateway.
        :param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[dict] tags: Key-value tags for the EC2 Transit Gateway.
        :param pulumi.Input[str] vpn_ecmp_support: Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['amazon_side_asn'] = amazon_side_asn
            __props__['auto_accept_shared_attachments'] = auto_accept_shared_attachments
            __props__['default_route_table_association'] = default_route_table_association
            __props__['default_route_table_propagation'] = default_route_table_propagation
            __props__['description'] = description
            __props__['dns_support'] = dns_support
            __props__['tags'] = tags
            __props__['vpn_ecmp_support'] = vpn_ecmp_support
            __props__['arn'] = None
            __props__['association_default_route_table_id'] = None
            __props__['owner_id'] = None
            __props__['propagation_default_route_table_id'] = None
        super(TransitGateway, __self__).__init__(
            'aws:ec2transitgateway/transitGateway:TransitGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, amazon_side_asn=None, arn=None, association_default_route_table_id=None, auto_accept_shared_attachments=None, default_route_table_association=None, default_route_table_propagation=None, description=None, dns_support=None, owner_id=None, propagation_default_route_table_id=None, tags=None, vpn_ecmp_support=None):
        """
        Get an existing TransitGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] amazon_side_asn: Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
        :param pulumi.Input[str] arn: EC2 Transit Gateway Amazon Resource Name (ARN)
        :param pulumi.Input[str] association_default_route_table_id: Identifier of the default association route table
        :param pulumi.Input[str] auto_accept_shared_attachments: Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
        :param pulumi.Input[str] default_route_table_association: Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] default_route_table_propagation: Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] description: Description of the EC2 Transit Gateway.
        :param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] owner_id: Identifier of the AWS account that owns the EC2 Transit Gateway
        :param pulumi.Input[str] propagation_default_route_table_id: Identifier of the default propagation route table
        :param pulumi.Input[dict] tags: Key-value tags for the EC2 Transit Gateway.
        :param pulumi.Input[str] vpn_ecmp_support: Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["amazon_side_asn"] = amazon_side_asn
        __props__["arn"] = arn
        __props__["association_default_route_table_id"] = association_default_route_table_id
        __props__["auto_accept_shared_attachments"] = auto_accept_shared_attachments
        __props__["default_route_table_association"] = default_route_table_association
        __props__["default_route_table_propagation"] = default_route_table_propagation
        __props__["description"] = description
        __props__["dns_support"] = dns_support
        __props__["owner_id"] = owner_id
        __props__["propagation_default_route_table_id"] = propagation_default_route_table_id
        __props__["tags"] = tags
        __props__["vpn_ecmp_support"] = vpn_ecmp_support
        return TransitGateway(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

