# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Gateway(pulumi.CustomResource):
    amazon_side_asn: pulumi.Output[str]
    """
    The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
    """
    name: pulumi.Output[str]
    """
    The name of the connection.
    """
    owner_account_id: pulumi.Output[str]
    """
    AWS Account ID of the gateway.
    """
    def __init__(__self__, resource_name, opts=None, amazon_side_asn=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Direct Connect Gateway.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] amazon_side_asn: The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
        :param pulumi.Input[str] name: The name of the connection.
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

            if amazon_side_asn is None:
                raise TypeError("Missing required property 'amazon_side_asn'")
            __props__['amazon_side_asn'] = amazon_side_asn
            __props__['name'] = name
            __props__['owner_account_id'] = None
        super(Gateway, __self__).__init__(
            'aws:directconnect/gateway:Gateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, amazon_side_asn=None, name=None, owner_account_id=None):
        """
        Get an existing Gateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] amazon_side_asn: The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[str] owner_account_id: AWS Account ID of the gateway.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["amazon_side_asn"] = amazon_side_asn
        __props__["name"] = name
        __props__["owner_account_id"] = owner_account_id
        return Gateway(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

