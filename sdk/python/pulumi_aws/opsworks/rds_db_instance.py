# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RdsDbInstance(pulumi.CustomResource):
    db_password: pulumi.Output[str]
    """
    A db password
    """
    db_user: pulumi.Output[str]
    """
    A db username
    """
    rds_db_instance_arn: pulumi.Output[str]
    """
    The db instance to register for this stack. Changing this will force a new resource.
    """
    stack_id: pulumi.Output[str]
    """
    The stack to register a db instance for. Changing this will force a new resource.
    """
    def __init__(__self__, __name__, __opts__=None, db_password=None, db_user=None, rds_db_instance_arn=None, stack_id=None):
        """
        Provides an OpsWorks RDS DB Instance resource.
        
        > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] db_password: A db password
        :param pulumi.Input[str] db_user: A db username
        :param pulumi.Input[str] rds_db_instance_arn: The db instance to register for this stack. Changing this will force a new resource.
        :param pulumi.Input[str] stack_id: The stack to register a db instance for. Changing this will force a new resource.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not db_password:
            raise TypeError('Missing required property db_password')
        __props__['db_password'] = db_password

        if not db_user:
            raise TypeError('Missing required property db_user')
        __props__['db_user'] = db_user

        if not rds_db_instance_arn:
            raise TypeError('Missing required property rds_db_instance_arn')
        __props__['rds_db_instance_arn'] = rds_db_instance_arn

        if not stack_id:
            raise TypeError('Missing required property stack_id')
        __props__['stack_id'] = stack_id

        super(RdsDbInstance, __self__).__init__(
            'aws:opsworks/rdsDbInstance:RdsDbInstance',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

