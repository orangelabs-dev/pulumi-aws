# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetServiceQuotaResult:
    """
    A collection of values returned by getServiceQuota.
    """
    def __init__(__self__, adjustable=None, arn=None, default_value=None, global_quota=None, quota_code=None, quota_name=None, service_code=None, service_name=None, value=None, id=None):
        if adjustable and not isinstance(adjustable, bool):
            raise TypeError("Expected argument 'adjustable' to be a bool")
        __self__.adjustable = adjustable
        """
        Whether the service quota is adjustable.
        """
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Amazon Resource Name (ARN) of the service quota.
        """
        if default_value and not isinstance(default_value, float):
            raise TypeError("Expected argument 'default_value' to be a float")
        __self__.default_value = default_value
        """
        Default value of the service quota.
        """
        if global_quota and not isinstance(global_quota, bool):
            raise TypeError("Expected argument 'global_quota' to be a bool")
        __self__.global_quota = global_quota
        """
        Whether the service quota is global for the AWS account.
        """
        if quota_code and not isinstance(quota_code, str):
            raise TypeError("Expected argument 'quota_code' to be a str")
        __self__.quota_code = quota_code
        if quota_name and not isinstance(quota_name, str):
            raise TypeError("Expected argument 'quota_name' to be a str")
        __self__.quota_name = quota_name
        if service_code and not isinstance(service_code, str):
            raise TypeError("Expected argument 'service_code' to be a str")
        __self__.service_code = service_code
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        __self__.service_name = service_name
        """
        Name of the service.
        """
        if value and not isinstance(value, float):
            raise TypeError("Expected argument 'value' to be a float")
        __self__.value = value
        """
        Current value of the service quota.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_service_quota(quota_code=None,quota_name=None,service_code=None,opts=None):
    """
    Retrieve information about a Service Quota.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/servicequotas_service_quota.html.markdown.
    """
    __args__ = dict()

    __args__['quotaCode'] = quota_code
    __args__['quotaName'] = quota_name
    __args__['serviceCode'] = service_code
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:servicequotas/getServiceQuota:getServiceQuota', __args__, opts=opts).value

    return GetServiceQuotaResult(
        adjustable=__ret__.get('adjustable'),
        arn=__ret__.get('arn'),
        default_value=__ret__.get('defaultValue'),
        global_quota=__ret__.get('globalQuota'),
        quota_code=__ret__.get('quotaCode'),
        quota_name=__ret__.get('quotaName'),
        service_code=__ret__.get('serviceCode'),
        service_name=__ret__.get('serviceName'),
        value=__ret__.get('value'),
        id=__ret__.get('id'))