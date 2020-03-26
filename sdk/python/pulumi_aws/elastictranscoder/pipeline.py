# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Pipeline(pulumi.CustomResource):
    arn: pulumi.Output[str]
    aws_kms_key_arn: pulumi.Output[str]
    """
    The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
    """
    content_config: pulumi.Output[dict]
    """
    The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)

      * `bucket` (`str`) - The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
      * `storage_class` (`str`) - The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.
    """
    content_config_permissions: pulumi.Output[list]
    """
    The permissions for the `content_config` object. (documented below)

      * `accesses` (`list`) - The permission that you want to give to the AWS user that you specified in `content_config_permissions.grantee`
      * `grantee` (`str`) - The AWS user or group that you want to have access to transcoded files and playlists.
      * `granteeType` (`str`) - Specify the type of value that appears in the `content_config_permissions.grantee` object. Valid values are `Canonical`, `Email` or `Group`.
    """
    input_bucket: pulumi.Output[str]
    """
    The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
    """
    name: pulumi.Output[str]
    """
    The name of the pipeline. Maximum 40 characters
    """
    notifications: pulumi.Output[dict]
    """
    The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)

      * `completed` (`str`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder has finished processing a job in this pipeline.
      * `error` (`str`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters an error condition while processing a job in this pipeline.
      * `progressing` (`str`) - The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify when Elastic Transcoder has started to process a job in this pipeline.
      * `warning` (`str`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters a warning condition while processing a job in this pipeline.
    """
    output_bucket: pulumi.Output[str]
    """
    The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
    """
    role: pulumi.Output[str]
    """
    The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
    """
    thumbnail_config: pulumi.Output[dict]
    """
    The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)

      * `bucket` (`str`) - The Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files.
      * `storage_class` (`str`) - The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the thumbnails that it stores in your Amazon S3 bucket.
    """
    thumbnail_config_permissions: pulumi.Output[list]
    """
    The permissions for the `thumbnail_config` object. (documented below)

      * `accesses` (`list`) - The permission that you want to give to the AWS user that you specified in `thumbnail_config_permissions.grantee`.
      * `grantee` (`str`) - The AWS user or group that you want to have access to thumbnail files.
      * `granteeType` (`str`) - Specify the type of value that appears in the `thumbnail_config_permissions.grantee` object.
    """
    def __init__(__self__, resource_name, opts=None, aws_kms_key_arn=None, content_config=None, content_config_permissions=None, input_bucket=None, name=None, notifications=None, output_bucket=None, role=None, thumbnail_config=None, thumbnail_config_permissions=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an Elastic Transcoder pipeline resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elastictranscoder_pipeline.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_kms_key_arn: The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        :param pulumi.Input[dict] content_config: The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        :param pulumi.Input[list] content_config_permissions: The permissions for the `content_config` object. (documented below)
        :param pulumi.Input[str] input_bucket: The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        :param pulumi.Input[str] name: The name of the pipeline. Maximum 40 characters
        :param pulumi.Input[dict] notifications: The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        :param pulumi.Input[str] output_bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        :param pulumi.Input[str] role: The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        :param pulumi.Input[dict] thumbnail_config: The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        :param pulumi.Input[list] thumbnail_config_permissions: The permissions for the `thumbnail_config` object. (documented below)

        The **content_config** object supports the following:

          * `bucket` (`pulumi.Input[str]`) - The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
          * `storage_class` (`pulumi.Input[str]`) - The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.

        The **content_config_permissions** object supports the following:

          * `accesses` (`pulumi.Input[list]`) - The permission that you want to give to the AWS user that you specified in `content_config_permissions.grantee`
          * `grantee` (`pulumi.Input[str]`) - The AWS user or group that you want to have access to transcoded files and playlists.
          * `granteeType` (`pulumi.Input[str]`) - Specify the type of value that appears in the `content_config_permissions.grantee` object. Valid values are `Canonical`, `Email` or `Group`.

        The **notifications** object supports the following:

          * `completed` (`pulumi.Input[str]`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder has finished processing a job in this pipeline.
          * `error` (`pulumi.Input[str]`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters an error condition while processing a job in this pipeline.
          * `progressing` (`pulumi.Input[str]`) - The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify when Elastic Transcoder has started to process a job in this pipeline.
          * `warning` (`pulumi.Input[str]`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters a warning condition while processing a job in this pipeline.

        The **thumbnail_config** object supports the following:

          * `bucket` (`pulumi.Input[str]`) - The Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files.
          * `storage_class` (`pulumi.Input[str]`) - The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the thumbnails that it stores in your Amazon S3 bucket.

        The **thumbnail_config_permissions** object supports the following:

          * `accesses` (`pulumi.Input[list]`) - The permission that you want to give to the AWS user that you specified in `thumbnail_config_permissions.grantee`.
          * `grantee` (`pulumi.Input[str]`) - The AWS user or group that you want to have access to thumbnail files.
          * `granteeType` (`pulumi.Input[str]`) - Specify the type of value that appears in the `thumbnail_config_permissions.grantee` object.
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

            __props__['aws_kms_key_arn'] = aws_kms_key_arn
            __props__['content_config'] = content_config
            __props__['content_config_permissions'] = content_config_permissions
            if input_bucket is None:
                raise TypeError("Missing required property 'input_bucket'")
            __props__['input_bucket'] = input_bucket
            __props__['name'] = name
            __props__['notifications'] = notifications
            __props__['output_bucket'] = output_bucket
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['thumbnail_config'] = thumbnail_config
            __props__['thumbnail_config_permissions'] = thumbnail_config_permissions
            __props__['arn'] = None
        super(Pipeline, __self__).__init__(
            'aws:elastictranscoder/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, aws_kms_key_arn=None, content_config=None, content_config_permissions=None, input_bucket=None, name=None, notifications=None, output_bucket=None, role=None, thumbnail_config=None, thumbnail_config_permissions=None):
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_kms_key_arn: The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        :param pulumi.Input[dict] content_config: The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        :param pulumi.Input[list] content_config_permissions: The permissions for the `content_config` object. (documented below)
        :param pulumi.Input[str] input_bucket: The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        :param pulumi.Input[str] name: The name of the pipeline. Maximum 40 characters
        :param pulumi.Input[dict] notifications: The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        :param pulumi.Input[str] output_bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        :param pulumi.Input[str] role: The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        :param pulumi.Input[dict] thumbnail_config: The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        :param pulumi.Input[list] thumbnail_config_permissions: The permissions for the `thumbnail_config` object. (documented below)

        The **content_config** object supports the following:

          * `bucket` (`pulumi.Input[str]`) - The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
          * `storage_class` (`pulumi.Input[str]`) - The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.

        The **content_config_permissions** object supports the following:

          * `accesses` (`pulumi.Input[list]`) - The permission that you want to give to the AWS user that you specified in `content_config_permissions.grantee`
          * `grantee` (`pulumi.Input[str]`) - The AWS user or group that you want to have access to transcoded files and playlists.
          * `granteeType` (`pulumi.Input[str]`) - Specify the type of value that appears in the `content_config_permissions.grantee` object. Valid values are `Canonical`, `Email` or `Group`.

        The **notifications** object supports the following:

          * `completed` (`pulumi.Input[str]`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder has finished processing a job in this pipeline.
          * `error` (`pulumi.Input[str]`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters an error condition while processing a job in this pipeline.
          * `progressing` (`pulumi.Input[str]`) - The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify when Elastic Transcoder has started to process a job in this pipeline.
          * `warning` (`pulumi.Input[str]`) - The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters a warning condition while processing a job in this pipeline.

        The **thumbnail_config** object supports the following:

          * `bucket` (`pulumi.Input[str]`) - The Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files.
          * `storage_class` (`pulumi.Input[str]`) - The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the thumbnails that it stores in your Amazon S3 bucket.

        The **thumbnail_config_permissions** object supports the following:

          * `accesses` (`pulumi.Input[list]`) - The permission that you want to give to the AWS user that you specified in `thumbnail_config_permissions.grantee`.
          * `grantee` (`pulumi.Input[str]`) - The AWS user or group that you want to have access to thumbnail files.
          * `granteeType` (`pulumi.Input[str]`) - Specify the type of value that appears in the `thumbnail_config_permissions.grantee` object.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["aws_kms_key_arn"] = aws_kms_key_arn
        __props__["content_config"] = content_config
        __props__["content_config_permissions"] = content_config_permissions
        __props__["input_bucket"] = input_bucket
        __props__["name"] = name
        __props__["notifications"] = notifications
        __props__["output_bucket"] = output_bucket
        __props__["role"] = role
        __props__["thumbnail_config"] = thumbnail_config
        __props__["thumbnail_config_permissions"] = thumbnail_config_permissions
        return Pipeline(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

