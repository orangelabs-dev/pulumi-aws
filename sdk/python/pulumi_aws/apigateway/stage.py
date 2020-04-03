# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Stage(pulumi.CustomResource):
    access_log_settings: pulumi.Output[dict]
    """
    Enables access logs for the API stage. Detailed below.

      * `destination_arn` (`str`) - The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-`. Automatically removes trailing `:*` if present.
      * `format` (`str`) - The formatting and values recorded in the logs. 
        For more information on configuring the log format rules visit the AWS [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN)
    """
    cache_cluster_enabled: pulumi.Output[bool]
    """
    Specifies whether a cache cluster is enabled for the stage
    """
    cache_cluster_size: pulumi.Output[str]
    """
    The size of the cache cluster for the stage, if enabled.
    Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
    """
    client_certificate_id: pulumi.Output[str]
    """
    The identifier of a client certificate for the stage.
    """
    deployment: pulumi.Output[str]
    """
    The ID of the deployment that the stage points to
    """
    description: pulumi.Output[str]
    """
    The description of the stage
    """
    documentation_version: pulumi.Output[str]
    """
    The version of the associated API documentation
    """
    execution_arn: pulumi.Output[str]
    """
    The execution ARN to be used in [`lambda_permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `source_arn`
    when allowing API Gateway to invoke a Lambda function,
    e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
    """
    invoke_url: pulumi.Output[str]
    """
    The URL to invoke the API pointing to the stage,
    e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
    """
    rest_api: pulumi.Output[str]
    """
    The ID of the associated REST API
    """
    stage_name: pulumi.Output[str]
    """
    The name of the stage
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    variables: pulumi.Output[dict]
    """
    A map that defines the stage variables
    """
    xray_tracing_enabled: pulumi.Output[bool]
    """
    Whether active tracing with X-ray is enabled. Defaults to `false`.
    """
    def __init__(__self__, resource_name, opts=None, access_log_settings=None, cache_cluster_enabled=None, cache_cluster_size=None, client_certificate_id=None, deployment=None, description=None, documentation_version=None, rest_api=None, stage_name=None, tags=None, variables=None, xray_tracing_enabled=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway Stage.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_log_settings: Enables access logs for the API stage. Detailed below.
        :param pulumi.Input[bool] cache_cluster_enabled: Specifies whether a cache cluster is enabled for the stage
        :param pulumi.Input[str] cache_cluster_size: The size of the cache cluster for the stage, if enabled.
               Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        :param pulumi.Input[str] client_certificate_id: The identifier of a client certificate for the stage.
        :param pulumi.Input[dict] deployment: The ID of the deployment that the stage points to
        :param pulumi.Input[str] description: The description of the stage
        :param pulumi.Input[str] documentation_version: The version of the associated API documentation
        :param pulumi.Input[dict] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] stage_name: The name of the stage
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[dict] variables: A map that defines the stage variables
        :param pulumi.Input[bool] xray_tracing_enabled: Whether active tracing with X-ray is enabled. Defaults to `false`.

        The **access_log_settings** object supports the following:

          * `destination_arn` (`pulumi.Input[str]`) - The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-`. Automatically removes trailing `:*` if present.
          * `format` (`pulumi.Input[str]`) - The formatting and values recorded in the logs. 
            For more information on configuring the log format rules visit the AWS [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)
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

            __props__['access_log_settings'] = access_log_settings
            __props__['cache_cluster_enabled'] = cache_cluster_enabled
            __props__['cache_cluster_size'] = cache_cluster_size
            __props__['client_certificate_id'] = client_certificate_id
            if deployment is None:
                raise TypeError("Missing required property 'deployment'")
            __props__['deployment'] = deployment
            __props__['description'] = description
            __props__['documentation_version'] = documentation_version
            if rest_api is None:
                raise TypeError("Missing required property 'rest_api'")
            __props__['rest_api'] = rest_api
            if stage_name is None:
                raise TypeError("Missing required property 'stage_name'")
            __props__['stage_name'] = stage_name
            __props__['tags'] = tags
            __props__['variables'] = variables
            __props__['xray_tracing_enabled'] = xray_tracing_enabled
            __props__['arn'] = None
            __props__['execution_arn'] = None
            __props__['invoke_url'] = None
        super(Stage, __self__).__init__(
            'aws:apigateway/stage:Stage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_log_settings=None, arn=None, cache_cluster_enabled=None, cache_cluster_size=None, client_certificate_id=None, deployment=None, description=None, documentation_version=None, execution_arn=None, invoke_url=None, rest_api=None, stage_name=None, tags=None, variables=None, xray_tracing_enabled=None):
        """
        Get an existing Stage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_log_settings: Enables access logs for the API stage. Detailed below.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[bool] cache_cluster_enabled: Specifies whether a cache cluster is enabled for the stage
        :param pulumi.Input[str] cache_cluster_size: The size of the cache cluster for the stage, if enabled.
               Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        :param pulumi.Input[str] client_certificate_id: The identifier of a client certificate for the stage.
        :param pulumi.Input[dict] deployment: The ID of the deployment that the stage points to
        :param pulumi.Input[str] description: The description of the stage
        :param pulumi.Input[str] documentation_version: The version of the associated API documentation
        :param pulumi.Input[str] execution_arn: The execution ARN to be used in [`lambda_permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `source_arn`
               when allowing API Gateway to invoke a Lambda function,
               e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
        :param pulumi.Input[str] invoke_url: The URL to invoke the API pointing to the stage,
               e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
        :param pulumi.Input[dict] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] stage_name: The name of the stage
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[dict] variables: A map that defines the stage variables
        :param pulumi.Input[bool] xray_tracing_enabled: Whether active tracing with X-ray is enabled. Defaults to `false`.

        The **access_log_settings** object supports the following:

          * `destination_arn` (`pulumi.Input[str]`) - The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-`. Automatically removes trailing `:*` if present.
          * `format` (`pulumi.Input[str]`) - The formatting and values recorded in the logs. 
            For more information on configuring the log format rules visit the AWS [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_log_settings"] = access_log_settings
        __props__["arn"] = arn
        __props__["cache_cluster_enabled"] = cache_cluster_enabled
        __props__["cache_cluster_size"] = cache_cluster_size
        __props__["client_certificate_id"] = client_certificate_id
        __props__["deployment"] = deployment
        __props__["description"] = description
        __props__["documentation_version"] = documentation_version
        __props__["execution_arn"] = execution_arn
        __props__["invoke_url"] = invoke_url
        __props__["rest_api"] = rest_api
        __props__["stage_name"] = stage_name
        __props__["tags"] = tags
        __props__["variables"] = variables
        __props__["xray_tracing_enabled"] = xray_tracing_enabled
        return Stage(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

