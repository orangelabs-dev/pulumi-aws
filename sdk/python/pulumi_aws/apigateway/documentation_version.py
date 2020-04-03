# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DocumentationVersion(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the API documentation version.
    """
    rest_api_id: pulumi.Output[str]
    """
    The ID of the associated Rest API
    """
    version: pulumi.Output[str]
    """
    The version identifier of the API documentation snapshot.
    """
    def __init__(__self__, resource_name, opts=None, description=None, rest_api_id=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage an API Gateway Documentation Version.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the API documentation version.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API
        :param pulumi.Input[str] version: The version identifier of the API documentation snapshot.
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

            __props__['description'] = description
            if rest_api_id is None:
                raise TypeError("Missing required property 'rest_api_id'")
            __props__['rest_api_id'] = rest_api_id
            if version is None:
                raise TypeError("Missing required property 'version'")
            __props__['version'] = version
        super(DocumentationVersion, __self__).__init__(
            'aws:apigateway/documentationVersion:DocumentationVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, rest_api_id=None, version=None):
        """
        Get an existing DocumentationVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the API documentation version.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API
        :param pulumi.Input[str] version: The version identifier of the API documentation snapshot.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["rest_api_id"] = rest_api_id
        __props__["version"] = version
        return DocumentationVersion(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

