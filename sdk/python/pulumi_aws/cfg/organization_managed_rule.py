# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class OrganizationManagedRule(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the rule
    """
    description: pulumi.Output[str]
    """
    Description of the rule
    """
    excluded_accounts: pulumi.Output[list]
    """
    List of AWS account identifiers to exclude from the rule
    """
    input_parameters: pulumi.Output[str]
    """
    A string in JSON format that is passed to the AWS Config Rule Lambda Function
    """
    maximum_execution_frequency: pulumi.Output[str]
    """
    The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
    """
    name: pulumi.Output[str]
    """
    The name of the rule
    """
    resource_id_scope: pulumi.Output[str]
    """
    Identifier of the AWS resource to evaluate
    """
    resource_types_scopes: pulumi.Output[list]
    """
    List of types of AWS resources to evaluate
    """
    rule_identifier: pulumi.Output[str]
    """
    Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
    """
    tag_key_scope: pulumi.Output[str]
    """
    Tag key of AWS resources to evaluate
    """
    tag_value_scope: pulumi.Output[str]
    """
    Tag value of AWS resources to evaluate
    """
    def __init__(__self__, resource_name, opts=None, description=None, excluded_accounts=None, input_parameters=None, maximum_execution_frequency=None, name=None, resource_id_scope=None, resource_types_scopes=None, rule_identifier=None, tag_key_scope=None, tag_value_scope=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Config Organization Managed Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Custom Rules (those invoking a custom Lambda Function), see the [`cfg.OrganizationCustomRule` resource](https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html).

        > **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.

        > **NOTE:** Every Organization account except those configured in the `excluded_accounts` argument must have a Configuration Recorder with proper IAM permissions before the rule will successfully create or update. See also the [`cfg.Recorder` resource](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html).



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the rule
        :param pulumi.Input[list] excluded_accounts: List of AWS account identifiers to exclude from the rule
        :param pulumi.Input[str] input_parameters: A string in JSON format that is passed to the AWS Config Rule Lambda Function
        :param pulumi.Input[str] maximum_execution_frequency: The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[str] resource_id_scope: Identifier of the AWS resource to evaluate
        :param pulumi.Input[list] resource_types_scopes: List of types of AWS resources to evaluate
        :param pulumi.Input[str] rule_identifier: Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        :param pulumi.Input[str] tag_key_scope: Tag key of AWS resources to evaluate
        :param pulumi.Input[str] tag_value_scope: Tag value of AWS resources to evaluate
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
            __props__['excluded_accounts'] = excluded_accounts
            __props__['input_parameters'] = input_parameters
            __props__['maximum_execution_frequency'] = maximum_execution_frequency
            __props__['name'] = name
            __props__['resource_id_scope'] = resource_id_scope
            __props__['resource_types_scopes'] = resource_types_scopes
            if rule_identifier is None:
                raise TypeError("Missing required property 'rule_identifier'")
            __props__['rule_identifier'] = rule_identifier
            __props__['tag_key_scope'] = tag_key_scope
            __props__['tag_value_scope'] = tag_value_scope
            __props__['arn'] = None
        super(OrganizationManagedRule, __self__).__init__(
            'aws:cfg/organizationManagedRule:OrganizationManagedRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, description=None, excluded_accounts=None, input_parameters=None, maximum_execution_frequency=None, name=None, resource_id_scope=None, resource_types_scopes=None, rule_identifier=None, tag_key_scope=None, tag_value_scope=None):
        """
        Get an existing OrganizationManagedRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the rule
        :param pulumi.Input[str] description: Description of the rule
        :param pulumi.Input[list] excluded_accounts: List of AWS account identifiers to exclude from the rule
        :param pulumi.Input[str] input_parameters: A string in JSON format that is passed to the AWS Config Rule Lambda Function
        :param pulumi.Input[str] maximum_execution_frequency: The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[str] resource_id_scope: Identifier of the AWS resource to evaluate
        :param pulumi.Input[list] resource_types_scopes: List of types of AWS resources to evaluate
        :param pulumi.Input[str] rule_identifier: Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        :param pulumi.Input[str] tag_key_scope: Tag key of AWS resources to evaluate
        :param pulumi.Input[str] tag_value_scope: Tag value of AWS resources to evaluate
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["excluded_accounts"] = excluded_accounts
        __props__["input_parameters"] = input_parameters
        __props__["maximum_execution_frequency"] = maximum_execution_frequency
        __props__["name"] = name
        __props__["resource_id_scope"] = resource_id_scope
        __props__["resource_types_scopes"] = resource_types_scopes
        __props__["rule_identifier"] = rule_identifier
        __props__["tag_key_scope"] = tag_key_scope
        __props__["tag_value_scope"] = tag_value_scope
        return OrganizationManagedRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

