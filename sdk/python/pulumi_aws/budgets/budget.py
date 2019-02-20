# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Budget(pulumi.CustomResource):
    account_id: pulumi.Output[str]
    """
    The ID of the target account for budget. Will use current user's account_id by default if omitted.
    """
    budget_type: pulumi.Output[str]
    """
    Whether this budget tracks monetary cost or usage.
    """
    cost_filters: pulumi.Output[dict]
    """
    Map of CostFilters key/value pairs to apply to the budget.
    """
    cost_types: pulumi.Output[dict]
    """
    Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions..
    """
    limit_amount: pulumi.Output[str]
    """
    The amount of cost or usage being measured for a budget.
    """
    limit_unit: pulumi.Output[str]
    """
    The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
    """
    name: pulumi.Output[str]
    """
    The name of a budget. Unique within accounts.
    """
    name_prefix: pulumi.Output[str]
    """
    The prefix of the name of a budget. Unique within accounts.
    """
    time_period_end: pulumi.Output[str]
    """
    The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
    """
    time_period_start: pulumi.Output[str]
    """
    The start of the time period covered by the budget. The start date must come before the end date. Format: `2017-01-01_12:00`.
    """
    time_unit: pulumi.Output[str]
    """
    The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`.
    """
    def __init__(__self__, __name__, __opts__=None, account_id=None, budget_type=None, cost_filters=None, cost_types=None, limit_amount=None, limit_unit=None, name=None, name_prefix=None, time_period_end=None, time_period_start=None, time_unit=None):
        """
        Provides a budgets budget resource. Budgets use the cost visualisation provided by Cost Explorer to show you the status of your budgets, to provide forecasts of your estimated costs, and to track your AWS usage, including your free tier usage.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the target account for budget. Will use current user's account_id by default if omitted.
        :param pulumi.Input[str] budget_type: Whether this budget tracks monetary cost or usage.
        :param pulumi.Input[dict] cost_filters: Map of CostFilters key/value pairs to apply to the budget.
        :param pulumi.Input[dict] cost_types: Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions..
        :param pulumi.Input[str] limit_amount: The amount of cost or usage being measured for a budget.
        :param pulumi.Input[str] limit_unit: The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
        :param pulumi.Input[str] name: The name of a budget. Unique within accounts.
        :param pulumi.Input[str] name_prefix: The prefix of the name of a budget. Unique within accounts.
        :param pulumi.Input[str] time_period_end: The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
        :param pulumi.Input[str] time_period_start: The start of the time period covered by the budget. The start date must come before the end date. Format: `2017-01-01_12:00`.
        :param pulumi.Input[str] time_unit: The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['account_id'] = account_id

        if not budget_type:
            raise TypeError('Missing required property budget_type')
        __props__['budget_type'] = budget_type

        __props__['cost_filters'] = cost_filters

        __props__['cost_types'] = cost_types

        if not limit_amount:
            raise TypeError('Missing required property limit_amount')
        __props__['limit_amount'] = limit_amount

        if not limit_unit:
            raise TypeError('Missing required property limit_unit')
        __props__['limit_unit'] = limit_unit

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        __props__['time_period_end'] = time_period_end

        if not time_period_start:
            raise TypeError('Missing required property time_period_start')
        __props__['time_period_start'] = time_period_start

        if not time_unit:
            raise TypeError('Missing required property time_unit')
        __props__['time_unit'] = time_unit

        super(Budget, __self__).__init__(
            'aws:budgets/budget:Budget',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

