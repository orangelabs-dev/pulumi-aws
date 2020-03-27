# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class TrafficMirrorFilterRule(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A description of the traffic mirror filter rule.
    """
    destination_cidr_block: pulumi.Output[str]
    """
    The destination CIDR block to assign to the Traffic Mirror rule.
    """
    destination_port_range: pulumi.Output[dict]
    """
    The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below

      * `from_port` (`float`) - Starting port of the range
      * `to_port` (`float`) - Ending port of the range
    """
    protocol: pulumi.Output[float]
    """
    The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
    """
    rule_action: pulumi.Output[str]
    """
    The action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
    """
    rule_number: pulumi.Output[float]
    """
    The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
    """
    source_cidr_block: pulumi.Output[str]
    """
    The source CIDR block to assign to the Traffic Mirror rule.
    """
    source_port_range: pulumi.Output[dict]
    """
    The source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below

      * `from_port` (`float`) - Starting port of the range
      * `to_port` (`float`) - Ending port of the range
    """
    traffic_direction: pulumi.Output[str]
    """
    The direction of traffic to be captured. Valid values are `ingress` and `egress`
    """
    traffic_mirror_filter_id: pulumi.Output[str]
    """
    ID of the traffic mirror filter to which this rule should be added
    """
    def __init__(__self__, resource_name, opts=None, description=None, destination_cidr_block=None, destination_port_range=None, protocol=None, rule_action=None, rule_number=None, source_cidr_block=None, source_port_range=None, traffic_direction=None, traffic_mirror_filter_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an Traffic mirror filter rule.  
        Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_traffic_mirror_filter_rule.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the traffic mirror filter rule.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block to assign to the Traffic Mirror rule.
        :param pulumi.Input[dict] destination_port_range: The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
        :param pulumi.Input[float] protocol: The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
        :param pulumi.Input[str] rule_action: The action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
        :param pulumi.Input[float] rule_number: The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
        :param pulumi.Input[str] source_cidr_block: The source CIDR block to assign to the Traffic Mirror rule.
        :param pulumi.Input[dict] source_port_range: The source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
        :param pulumi.Input[str] traffic_direction: The direction of traffic to be captured. Valid values are `ingress` and `egress`
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to which this rule should be added

        The **destination_port_range** object supports the following:

          * `from_port` (`pulumi.Input[float]`) - Starting port of the range
          * `to_port` (`pulumi.Input[float]`) - Ending port of the range

        The **source_port_range** object supports the following:

          * `from_port` (`pulumi.Input[float]`) - Starting port of the range
          * `to_port` (`pulumi.Input[float]`) - Ending port of the range
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
            if destination_cidr_block is None:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__['destination_cidr_block'] = destination_cidr_block
            __props__['destination_port_range'] = destination_port_range
            __props__['protocol'] = protocol
            if rule_action is None:
                raise TypeError("Missing required property 'rule_action'")
            __props__['rule_action'] = rule_action
            if rule_number is None:
                raise TypeError("Missing required property 'rule_number'")
            __props__['rule_number'] = rule_number
            if source_cidr_block is None:
                raise TypeError("Missing required property 'source_cidr_block'")
            __props__['source_cidr_block'] = source_cidr_block
            __props__['source_port_range'] = source_port_range
            if traffic_direction is None:
                raise TypeError("Missing required property 'traffic_direction'")
            __props__['traffic_direction'] = traffic_direction
            if traffic_mirror_filter_id is None:
                raise TypeError("Missing required property 'traffic_mirror_filter_id'")
            __props__['traffic_mirror_filter_id'] = traffic_mirror_filter_id
        super(TrafficMirrorFilterRule, __self__).__init__(
            'aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, destination_cidr_block=None, destination_port_range=None, protocol=None, rule_action=None, rule_number=None, source_cidr_block=None, source_port_range=None, traffic_direction=None, traffic_mirror_filter_id=None):
        """
        Get an existing TrafficMirrorFilterRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the traffic mirror filter rule.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block to assign to the Traffic Mirror rule.
        :param pulumi.Input[dict] destination_port_range: The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
        :param pulumi.Input[float] protocol: The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
        :param pulumi.Input[str] rule_action: The action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
        :param pulumi.Input[float] rule_number: The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
        :param pulumi.Input[str] source_cidr_block: The source CIDR block to assign to the Traffic Mirror rule.
        :param pulumi.Input[dict] source_port_range: The source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
        :param pulumi.Input[str] traffic_direction: The direction of traffic to be captured. Valid values are `ingress` and `egress`
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to which this rule should be added

        The **destination_port_range** object supports the following:

          * `from_port` (`pulumi.Input[float]`) - Starting port of the range
          * `to_port` (`pulumi.Input[float]`) - Ending port of the range

        The **source_port_range** object supports the following:

          * `from_port` (`pulumi.Input[float]`) - Starting port of the range
          * `to_port` (`pulumi.Input[float]`) - Ending port of the range
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["destination_cidr_block"] = destination_cidr_block
        __props__["destination_port_range"] = destination_port_range
        __props__["protocol"] = protocol
        __props__["rule_action"] = rule_action
        __props__["rule_number"] = rule_number
        __props__["source_cidr_block"] = source_cidr_block
        __props__["source_port_range"] = source_port_range
        __props__["traffic_direction"] = traffic_direction
        __props__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
        return TrafficMirrorFilterRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
