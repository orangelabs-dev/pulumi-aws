# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class App(pulumi.CustomResource):
    application_id: pulumi.Output[str]
    """
    The Application ID of the Pinpoint App.
    """
    campaign_hook: pulumi.Output[dict]
    """
    The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
    """
    limits: pulumi.Output[dict]
    """
    The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
    """
    name: pulumi.Output[str]
    """
    The application name. By default generated by Terraform
    """
    name_prefix: pulumi.Output[str]
    """
    The name of the Pinpoint application. Conflicts with `name`
    """
    quiet_time: pulumi.Output[dict]
    """
    The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
    """
    def __init__(__self__, __name__, __opts__=None, campaign_hook=None, limits=None, name=None, name_prefix=None, quiet_time=None):
        """
        Provides a Pinpoint App resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] campaign_hook: The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
        :param pulumi.Input[dict] limits: The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
        :param pulumi.Input[str] name: The application name. By default generated by Terraform
        :param pulumi.Input[str] name_prefix: The name of the Pinpoint application. Conflicts with `name`
        :param pulumi.Input[dict] quiet_time: The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['campaign_hook'] = campaign_hook

        __props__['limits'] = limits

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        __props__['quiet_time'] = quiet_time

        __props__['application_id'] = None

        super(App, __self__).__init__(
            'aws:pinpoint/app:App',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

