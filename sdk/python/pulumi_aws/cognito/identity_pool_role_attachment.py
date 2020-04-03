# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class IdentityPoolRoleAttachment(pulumi.CustomResource):
    identity_pool_id: pulumi.Output[str]
    """
    An identity pool ID in the format REGION:GUID.
    """
    role_mappings: pulumi.Output[list]
    """
    A List of Role Mapping.

      * `ambiguousRoleResolution` (`str`) - Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
      * `identity_provider` (`str`) - A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
      * `mappingRules` (`list`) - The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
        * `claim` (`str`) - The claim name that must be present in the token, for example, "isAdmin" or "paid".
        * `matchType` (`str`) - The match condition that specifies how closely the claim value in the IdP token must match Value.
        * `role_arn` (`str`) - The role ARN.
        * `value` (`str`) - A brief string that the claim must match, for example, "paid" or "yes".

      * `type` (`str`) - The role mapping type.
    """
    roles: pulumi.Output[dict]
    """
    The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.

      * `authenticated` (`str`)
      * `unauthenticated` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, identity_pool_id=None, role_mappings=None, roles=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Cognito Identity Pool Roles Attachment.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_pool_id: An identity pool ID in the format REGION:GUID.
        :param pulumi.Input[list] role_mappings: A List of Role Mapping.
        :param pulumi.Input[dict] roles: The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.

        The **role_mappings** object supports the following:

          * `ambiguousRoleResolution` (`pulumi.Input[str]`) - Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
          * `identity_provider` (`pulumi.Input[str]`) - A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
          * `mappingRules` (`pulumi.Input[list]`) - The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
            * `claim` (`pulumi.Input[str]`) - The claim name that must be present in the token, for example, "isAdmin" or "paid".
            * `matchType` (`pulumi.Input[str]`) - The match condition that specifies how closely the claim value in the IdP token must match Value.
            * `role_arn` (`pulumi.Input[str]`) - The role ARN.
            * `value` (`pulumi.Input[str]`) - A brief string that the claim must match, for example, "paid" or "yes".

          * `type` (`pulumi.Input[str]`) - The role mapping type.

        The **roles** object supports the following:

          * `authenticated` (`pulumi.Input[str]`)
          * `unauthenticated` (`pulumi.Input[str]`)
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

            if identity_pool_id is None:
                raise TypeError("Missing required property 'identity_pool_id'")
            __props__['identity_pool_id'] = identity_pool_id
            __props__['role_mappings'] = role_mappings
            if roles is None:
                raise TypeError("Missing required property 'roles'")
            __props__['roles'] = roles
        super(IdentityPoolRoleAttachment, __self__).__init__(
            'aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, identity_pool_id=None, role_mappings=None, roles=None):
        """
        Get an existing IdentityPoolRoleAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_pool_id: An identity pool ID in the format REGION:GUID.
        :param pulumi.Input[list] role_mappings: A List of Role Mapping.
        :param pulumi.Input[dict] roles: The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.

        The **role_mappings** object supports the following:

          * `ambiguousRoleResolution` (`pulumi.Input[str]`) - Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
          * `identity_provider` (`pulumi.Input[str]`) - A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
          * `mappingRules` (`pulumi.Input[list]`) - The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
            * `claim` (`pulumi.Input[str]`) - The claim name that must be present in the token, for example, "isAdmin" or "paid".
            * `matchType` (`pulumi.Input[str]`) - The match condition that specifies how closely the claim value in the IdP token must match Value.
            * `role_arn` (`pulumi.Input[str]`) - The role ARN.
            * `value` (`pulumi.Input[str]`) - A brief string that the claim must match, for example, "paid" or "yes".

          * `type` (`pulumi.Input[str]`) - The role mapping type.

        The **roles** object supports the following:

          * `authenticated` (`pulumi.Input[str]`)
          * `unauthenticated` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["identity_pool_id"] = identity_pool_id
        __props__["role_mappings"] = role_mappings
        __props__["roles"] = roles
        return IdentityPoolRoleAttachment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

