# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SmsChannel(pulumi.CustomResource):
    application_id: pulumi.Output[str]
    """
    The application ID.
    """
    enabled: pulumi.Output[bool]
    """
    Whether the channel is enabled or disabled. Defaults to `true`.
    """
    promotional_messages_per_second: pulumi.Output[float]
    """
    Promotional messages per second that can be sent.
    """
    sender_id: pulumi.Output[str]
    """
    Sender identifier of your messages.
    """
    short_code: pulumi.Output[str]
    """
    The Short Code registered with the phone provider.
    """
    transactional_messages_per_second: pulumi.Output[float]
    """
    Transactional messages per second that can be sent.
    """
    def __init__(__self__, resource_name, opts=None, application_id=None, enabled=None, sender_id=None, short_code=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Pinpoint SMS Channel resource.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] sender_id: Sender identifier of your messages.
        :param pulumi.Input[str] short_code: The Short Code registered with the phone provider.
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

            if application_id is None:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            __props__['enabled'] = enabled
            __props__['sender_id'] = sender_id
            __props__['short_code'] = short_code
            __props__['promotional_messages_per_second'] = None
            __props__['transactional_messages_per_second'] = None
        super(SmsChannel, __self__).__init__(
            'aws:pinpoint/smsChannel:SmsChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, application_id=None, enabled=None, promotional_messages_per_second=None, sender_id=None, short_code=None, transactional_messages_per_second=None):
        """
        Get an existing SmsChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[float] promotional_messages_per_second: Promotional messages per second that can be sent.
        :param pulumi.Input[str] sender_id: Sender identifier of your messages.
        :param pulumi.Input[str] short_code: The Short Code registered with the phone provider.
        :param pulumi.Input[float] transactional_messages_per_second: Transactional messages per second that can be sent.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_id"] = application_id
        __props__["enabled"] = enabled
        __props__["promotional_messages_per_second"] = promotional_messages_per_second
        __props__["sender_id"] = sender_id
        __props__["short_code"] = short_code
        __props__["transactional_messages_per_second"] = transactional_messages_per_second
        return SmsChannel(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

