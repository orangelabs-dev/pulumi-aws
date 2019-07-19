# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DomainName(pulumi.CustomResource):
    certificate_arn: pulumi.Output[str]
    """
    The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
    """
    certificate_body: pulumi.Output[str]
    """
    The certificate issued for the domain name
    being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
    `regional_certificate_name`.
    """
    certificate_chain: pulumi.Output[str]
    """
    The certificate for the CA that issued the
    certificate, along with any intermediate CA certificates required to
    create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
    `regional_certificate_arn`, and `regional_certificate_name`.
    """
    certificate_name: pulumi.Output[str]
    """
    The unique name to use when registering this
    certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
    `regional_certificate_name`. Required if `certificate_arn` is not set.
    """
    certificate_private_key: pulumi.Output[str]
    """
    The private key associated with the
    domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
    """
    certificate_upload_date: pulumi.Output[str]
    """
    The upload date associated with the domain certificate.
    """
    cloudfront_domain_name: pulumi.Output[str]
    """
    The hostname created by Cloudfront to represent
    the distribution that implements this domain name mapping.
    """
    cloudfront_zone_id: pulumi.Output[str]
    """
    For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
    that can be used to create a Route53 alias record for the distribution.
    """
    domain_name: pulumi.Output[str]
    """
    The fully-qualified domain name to register
    """
    endpoint_configuration: pulumi.Output[dict]
    """
    Configuration block defining API endpoint information including type. Defined below.
    """
    regional_certificate_arn: pulumi.Output[str]
    """
    The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
    """
    regional_certificate_name: pulumi.Output[str]
    """
    The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
    `certificate_private_key`.
    """
    regional_domain_name: pulumi.Output[str]
    """
    The hostname for the custom domain's regional endpoint.
    """
    regional_zone_id: pulumi.Output[str]
    """
    The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
    """
    security_policy: pulumi.Output[str]
    """
    The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
    """
    def __init__(__self__, resource_name, opts=None, certificate_arn=None, certificate_body=None, certificate_chain=None, certificate_name=None, certificate_private_key=None, domain_name=None, endpoint_configuration=None, regional_certificate_arn=None, regional_certificate_name=None, security_policy=None, __name__=None, __opts__=None):
        """
        Registers a custom domain name for use with AWS API Gateway. Additional information about this functionality
        can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
        
        This resource just establishes ownership of and the TLS settings for
        a particular domain name. An API can be attached to a particular path
        under the registered domain name using
        the `aws_api_gateway_base_path_mapping` resource.
        
        API Gateway domains can be defined as either 'edge-optimized' or 'regional'.  In an edge-optimized configuration,
        API Gateway internally creates and manages a CloudFront distribution to route requests on the given hostname. In
        addition to this resource it's necessary to create a DNS record corresponding to the given domain name which is an alias
        (either Route53 alias or traditional CNAME) to the Cloudfront domain name exported in the `cloudfront_domain_name`
        attribute.
        
        In a regional configuration, API Gateway does not create a CloudFront distribution to route requests to the API, though
        a distribution can be created if needed. In either case, it is necessary to create a DNS record corresponding to the
        given domain name which is an alias (either Route53 alias or traditional CNAME) to the regional domain name exported in
        the `regional_domain_name` attribute.
        
        > **Note:** API Gateway requires the use of AWS Certificate Manager (ACM) certificates instead of Identity and Access Management (IAM) certificates in regions that support ACM. Regions that support ACM can be found in the [Regions and Endpoints Documentation](https://docs.aws.amazon.com/general/latest/gr/rande.html#acm_region). To import an existing private key and certificate into ACM or request an ACM certificate, see the [`aws_acm_certificate` resource](https://www.terraform.io/docs/providers/aws/r/acm_certificate.html).
        
        > **Note:** All arguments including the private key will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] certificate_body: The certificate issued for the domain name
               being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
               `regional_certificate_name`.
        :param pulumi.Input[str] certificate_chain: The certificate for the CA that issued the
               certificate, along with any intermediate CA certificates required to
               create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
               `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] certificate_name: The unique name to use when registering this
               certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
               `regional_certificate_name`. Required if `certificate_arn` is not set.
        :param pulumi.Input[str] certificate_private_key: The private key associated with the
               domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] domain_name: The fully-qualified domain name to register
        :param pulumi.Input[dict] endpoint_configuration: Configuration block defining API endpoint information including type. Defined below.
        :param pulumi.Input[str] regional_certificate_arn: The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        :param pulumi.Input[str] regional_certificate_name: The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
               `certificate_private_key`.
        :param pulumi.Input[str] security_policy: The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_domain_name.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['certificate_arn'] = certificate_arn

        __props__['certificate_body'] = certificate_body

        __props__['certificate_chain'] = certificate_chain

        __props__['certificate_name'] = certificate_name

        __props__['certificate_private_key'] = certificate_private_key

        if domain_name is None:
            raise TypeError("Missing required property 'domain_name'")
        __props__['domain_name'] = domain_name

        __props__['endpoint_configuration'] = endpoint_configuration

        __props__['regional_certificate_arn'] = regional_certificate_arn

        __props__['regional_certificate_name'] = regional_certificate_name

        __props__['security_policy'] = security_policy

        __props__['certificate_upload_date'] = None
        __props__['cloudfront_domain_name'] = None
        __props__['cloudfront_zone_id'] = None
        __props__['regional_domain_name'] = None
        __props__['regional_zone_id'] = None

        super(DomainName, __self__).__init__(
            'aws:apigateway/domainName:DomainName',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

