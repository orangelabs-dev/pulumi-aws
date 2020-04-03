# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Ledger(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the QLDB Ledger
    """
    deletion_protection: pulumi.Output[bool]
    """
    The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via this provider, this value must be configured to `false` and applied first before attempting deletion.
    """
    name: pulumi.Output[str]
    """
    The friendly name for the QLDB Ledger instance.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    def __init__(__self__, resource_name, opts=None, deletion_protection=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Quantum Ledger Database (QLDB) resource

        > **NOTE:** Deletion protection is enabled by default. To successfully delete this resource via this provider, `deletion_protection = false` must be applied before attempting deletion.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deletion_protection: The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via this provider, this value must be configured to `false` and applied first before attempting deletion.
        :param pulumi.Input[str] name: The friendly name for the QLDB Ledger instance.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags
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

            __props__['deletion_protection'] = deletion_protection
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(Ledger, __self__).__init__(
            'aws:qldb/ledger:Ledger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, deletion_protection=None, name=None, tags=None):
        """
        Get an existing Ledger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the QLDB Ledger
        :param pulumi.Input[bool] deletion_protection: The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via this provider, this value must be configured to `false` and applied first before attempting deletion.
        :param pulumi.Input[str] name: The friendly name for the QLDB Ledger instance.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["deletion_protection"] = deletion_protection
        __props__["name"] = name
        __props__["tags"] = tags
        return Ledger(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

