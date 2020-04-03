# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class CachesIscsiVolume(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
    """
    chap_enabled: pulumi.Output[bool]
    """
    Whether mutual CHAP is enabled for the iSCSI target.
    """
    gateway_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the gateway.
    """
    lun_number: pulumi.Output[float]
    """
    Logical disk number.
    """
    network_interface_id: pulumi.Output[str]
    """
    The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
    """
    network_interface_port: pulumi.Output[float]
    """
    The port used to communicate with iSCSI targets.
    """
    snapshot_id: pulumi.Output[str]
    """
    The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
    """
    source_volume_arn: pulumi.Output[str]
    """
    The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volume_size_in_bytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    target_arn: pulumi.Output[str]
    """
    Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
    """
    target_name: pulumi.Output[str]
    """
    The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
    """
    volume_arn: pulumi.Output[str]
    """
    Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
    """
    volume_id: pulumi.Output[str]
    """
    Volume ID, e.g. `vol-12345678`.
    """
    volume_size_in_bytes: pulumi.Output[float]
    """
    The size of the volume in bytes.
    """
    def __init__(__self__, resource_name, opts=None, gateway_arn=None, network_interface_id=None, snapshot_id=None, source_volume_arn=None, tags=None, target_name=None, volume_size_in_bytes=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an AWS Storage Gateway cached iSCSI volume.

        > **NOTE:** The gateway must have cache added (e.g. via the [`storagegateway.Cache`](https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html) resource) before creating volumes otherwise the Storage Gateway API will return an error.

        > **NOTE:** The gateway must have an upload buffer added (e.g. via the [`storagegateway.UploadBuffer`](https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html) resource) before the volume is operational to clients, however the Storage Gateway API will allow volume creation without error in that case and return volume status as `UPLOAD BUFFER NOT CONFIGURED`.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        :param pulumi.Input[str] network_interface_id: The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
        :param pulumi.Input[str] snapshot_id: The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
        :param pulumi.Input[str] source_volume_arn: The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volume_size_in_bytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] target_name: The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
        :param pulumi.Input[float] volume_size_in_bytes: The size of the volume in bytes.
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

            if gateway_arn is None:
                raise TypeError("Missing required property 'gateway_arn'")
            __props__['gateway_arn'] = gateway_arn
            if network_interface_id is None:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__['network_interface_id'] = network_interface_id
            __props__['snapshot_id'] = snapshot_id
            __props__['source_volume_arn'] = source_volume_arn
            __props__['tags'] = tags
            if target_name is None:
                raise TypeError("Missing required property 'target_name'")
            __props__['target_name'] = target_name
            if volume_size_in_bytes is None:
                raise TypeError("Missing required property 'volume_size_in_bytes'")
            __props__['volume_size_in_bytes'] = volume_size_in_bytes
            __props__['arn'] = None
            __props__['chap_enabled'] = None
            __props__['lun_number'] = None
            __props__['network_interface_port'] = None
            __props__['target_arn'] = None
            __props__['volume_arn'] = None
            __props__['volume_id'] = None
        super(CachesIscsiVolume, __self__).__init__(
            'aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, chap_enabled=None, gateway_arn=None, lun_number=None, network_interface_id=None, network_interface_port=None, snapshot_id=None, source_volume_arn=None, tags=None, target_arn=None, target_name=None, volume_arn=None, volume_id=None, volume_size_in_bytes=None):
        """
        Get an existing CachesIscsiVolume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
        :param pulumi.Input[bool] chap_enabled: Whether mutual CHAP is enabled for the iSCSI target.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        :param pulumi.Input[float] lun_number: Logical disk number.
        :param pulumi.Input[str] network_interface_id: The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
        :param pulumi.Input[float] network_interface_port: The port used to communicate with iSCSI targets.
        :param pulumi.Input[str] snapshot_id: The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
        :param pulumi.Input[str] source_volume_arn: The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volume_size_in_bytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] target_arn: Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
        :param pulumi.Input[str] target_name: The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
        :param pulumi.Input[str] volume_arn: Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
        :param pulumi.Input[str] volume_id: Volume ID, e.g. `vol-12345678`.
        :param pulumi.Input[float] volume_size_in_bytes: The size of the volume in bytes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["chap_enabled"] = chap_enabled
        __props__["gateway_arn"] = gateway_arn
        __props__["lun_number"] = lun_number
        __props__["network_interface_id"] = network_interface_id
        __props__["network_interface_port"] = network_interface_port
        __props__["snapshot_id"] = snapshot_id
        __props__["source_volume_arn"] = source_volume_arn
        __props__["tags"] = tags
        __props__["target_arn"] = target_arn
        __props__["target_name"] = target_name
        __props__["volume_arn"] = volume_arn
        __props__["volume_id"] = volume_id
        __props__["volume_size_in_bytes"] = volume_size_in_bytes
        return CachesIscsiVolume(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

