# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetServiceResult(object):
    """
    A collection of values returned by getService.
    """
    def __init__(__self__, arn=None, desired_count=None, launch_type=None, task_definition=None):
        if not arn:
            raise TypeError('Missing required argument arn')
        elif not isinstance(arn, basestring):
            raise TypeError('Expected argument arn to be a basestring')
        __self__.arn = arn
        """
        The ARN of the ECS Service
        """
        if not desired_count:
            raise TypeError('Missing required argument desired_count')
        elif not isinstance(desired_count, int):
            raise TypeError('Expected argument desired_count to be a int')
        __self__.desired_count = desired_count
        """
        The number of tasks for the ECS Service
        """
        if not launch_type:
            raise TypeError('Missing required argument launch_type')
        elif not isinstance(launch_type, basestring):
            raise TypeError('Expected argument launch_type to be a basestring')
        __self__.launch_type = launch_type
        """
        The launch type for the ECS Service
        """
        if not task_definition:
            raise TypeError('Missing required argument task_definition')
        elif not isinstance(task_definition, basestring):
            raise TypeError('Expected argument task_definition to be a basestring')
        __self__.task_definition = task_definition
        """
        The family for the latest ACTIVE revision
        """

def get_service(cluster_arn=None, service_name=None):
    """
    The ECS Service data source allows access to details of a specific
    Service within a AWS ECS Cluster.
    """
    __args__ = dict()

    __args__['clusterArn'] = cluster_arn
    __args__['serviceName'] = service_name
    __ret__ = pulumi.runtime.invoke('aws:ecs/getService:getService', __args__)

    return GetServiceResult(
        arn=__ret__['arn'],
        desired_count=__ret__['desiredCount'],
        launch_type=__ret__['launchType'],
        task_definition=__ret__['taskDefinition'])