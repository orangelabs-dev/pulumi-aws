# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class TargetGroupAttachment(pulumi.CustomResource):
    availability_zone: pulumi.Output[str]
    """
    The Availability Zone where the IP address of the target is to be registered.
    """
    port: pulumi.Output[int]
    """
    The port on which targets receive traffic.
    """
    target_group_arn: pulumi.Output[str]
    """
    The ARN of the target group with which to register targets
    """
    target_id: pulumi.Output[str]
    """
    The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
    """
    def __init__(__self__, __name__, __opts__=None, availability_zone=None, port=None, target_group_arn=None, target_id=None):
        """
        Provides the ability to register instances and containers with an Application Load Balancer (ALB) or Network Load Balancer (NLB) target group. For attaching resources with Elastic Load Balancer (ELB), see the [`aws_elb_attachment` resource](https://www.terraform.io/docs/providers/aws/r/elb_attachment.html).
        
        > **Note:** `aws_alb_target_group_attachment` is known as `aws_lb_target_group_attachment`. The functionality is identical.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] availability_zone: The Availability Zone where the IP address of the target is to be registered.
        :param pulumi.Input[int] port: The port on which targets receive traffic.
        :param pulumi.Input[str] target_group_arn: The ARN of the target group with which to register targets
        :param pulumi.Input[str] target_id: The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['availability_zone'] = availability_zone

        __props__['port'] = port

        if not target_group_arn:
            raise TypeError('Missing required property target_group_arn')
        __props__['target_group_arn'] = target_group_arn

        if not target_id:
            raise TypeError('Missing required property target_id')
        __props__['target_id'] = target_id

        super(TargetGroupAttachment, __self__).__init__(
            'aws:applicationloadbalancing/targetGroupAttachment:TargetGroupAttachment',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

