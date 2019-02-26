# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetDelegationSetResult:
    """
    A collection of values returned by getDelegationSet.
    """
    def __init__(__self__, caller_reference=None, name_servers=None):
        if caller_reference and not isinstance(caller_reference, str):
            raise TypeError('Expected argument caller_reference to be a str')
        __self__.caller_reference = caller_reference
        if name_servers and not isinstance(name_servers, list):
            raise TypeError('Expected argument name_servers to be a list')
        __self__.name_servers = name_servers

async def get_delegation_set(id=None,opts=None):
    """
    `aws_route53_delegation_set` provides details about a specific Route 53 Delegation Set.
    
    This data source allows to find a list of name servers associated with a specific delegation set.
    """
    __args__ = dict()

    __args__['id'] = id
    __ret__ = await pulumi.runtime.invoke('aws:route53/getDelegationSet:getDelegationSet', __args__, opts=opts)

    return GetDelegationSetResult(
        caller_reference=__ret__.get('callerReference'),
        name_servers=__ret__.get('nameServers'))
