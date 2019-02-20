# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventTarget(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) associated of the target.
    """
    batch_target: pulumi.Output[dict]
    """
    Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
    """
    ecs_target: pulumi.Output[dict]
    """
    Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
    """
    input: pulumi.Output[str]
    """
    Valid JSON text passed to the target.
    """
    input_path: pulumi.Output[str]
    """
    The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
    that is used for extracting part of the matched event when passing it to the target.
    """
    input_transformer: pulumi.Output[dict]
    """
    Parameters used when you are providing a custom input to a target based on certain event data.
    """
    kinesis_target: pulumi.Output[dict]
    """
    Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
    """
    role_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
    """
    rule: pulumi.Output[str]
    """
    The name of the rule you want to add targets to.
    """
    run_command_targets: pulumi.Output[list]
    """
    Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
    """
    sqs_target: pulumi.Output[dict]
    """
    Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
    """
    target_id: pulumi.Output[str]
    """
    The unique target assignment ID.  If missing, will generate a random, unique id.
    """
    def __init__(__self__, __name__, __opts__=None, arn=None, batch_target=None, ecs_target=None, input=None, input_path=None, input_transformer=None, kinesis_target=None, role_arn=None, rule=None, run_command_targets=None, sqs_target=None, target_id=None):
        """
        Provides a CloudWatch Event Target resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) associated of the target.
        :param pulumi.Input[dict] batch_target: Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[dict] ecs_target: Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] input: Valid JSON text passed to the target.
        :param pulumi.Input[str] input_path: The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
               that is used for extracting part of the matched event when passing it to the target.
        :param pulumi.Input[dict] input_transformer: Parameters used when you are providing a custom input to a target based on certain event data.
        :param pulumi.Input[dict] kinesis_target: Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
        :param pulumi.Input[str] rule: The name of the rule you want to add targets to.
        :param pulumi.Input[list] run_command_targets: Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        :param pulumi.Input[dict] sqs_target: Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] target_id: The unique target assignment ID.  If missing, will generate a random, unique id.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not arn:
            raise TypeError('Missing required property arn')
        __props__['arn'] = arn

        __props__['batch_target'] = batch_target

        __props__['ecs_target'] = ecs_target

        __props__['input'] = input

        __props__['input_path'] = input_path

        __props__['input_transformer'] = input_transformer

        __props__['kinesis_target'] = kinesis_target

        __props__['role_arn'] = role_arn

        if not rule:
            raise TypeError('Missing required property rule')
        __props__['rule'] = rule

        __props__['run_command_targets'] = run_command_targets

        __props__['sqs_target'] = sqs_target

        __props__['target_id'] = target_id

        super(EventTarget, __self__).__init__(
            'aws:cloudwatch/eventTarget:EventTarget',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

