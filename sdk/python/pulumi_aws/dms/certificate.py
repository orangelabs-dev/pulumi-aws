# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Certificate(pulumi.CustomResource):
    certificate_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) for the certificate.
    """
    certificate_id: pulumi.Output[str]
    """
    The certificate identifier.
    """
    certificate_pem: pulumi.Output[str]
    """
    The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
    """
    certificate_wallet: pulumi.Output[str]
    """
    The contents of the Oracle Wallet certificate for use with SSL. Either `certificate_pem` or `certificate_wallet` must be set.
    """
    def __init__(__self__, __name__, __opts__=None, certificate_id=None, certificate_pem=None, certificate_wallet=None):
        """
        Provides a DMS (Data Migration Service) certificate resource. DMS certificates can be created, deleted, and imported.
        
        > **Note:** All arguments including the PEM encoded certificate will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] certificate_id: The certificate identifier.
        :param pulumi.Input[str] certificate_pem: The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
        :param pulumi.Input[str] certificate_wallet: The contents of the Oracle Wallet certificate for use with SSL. Either `certificate_pem` or `certificate_wallet` must be set.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not certificate_id:
            raise TypeError('Missing required property certificate_id')
        __props__['certificate_id'] = certificate_id

        __props__['certificate_pem'] = certificate_pem

        __props__['certificate_wallet'] = certificate_wallet

        __props__['certificate_arn'] = None

        super(Certificate, __self__).__init__(
            'aws:dms/certificate:Certificate',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

