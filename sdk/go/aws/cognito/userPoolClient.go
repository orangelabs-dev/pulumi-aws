// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Cognito User Pool Client resource.
type UserPoolClient struct {
	pulumi.CustomResourceState

	// List of allowed OAuth flows (code, implicit, client_credentials).
	AllowedOauthFlows pulumi.StringArrayOutput `pulumi:"allowedOauthFlows"`
	// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
	AllowedOauthFlowsUserPoolClient pulumi.BoolPtrOutput `pulumi:"allowedOauthFlowsUserPoolClient"`
	// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
	AllowedOauthScopes pulumi.StringArrayOutput `pulumi:"allowedOauthScopes"`
	// The Amazon Pinpoint analytics configuration for collecting metrics for this user pool.
	AnalyticsConfiguration UserPoolClientAnalyticsConfigurationPtrOutput `pulumi:"analyticsConfiguration"`
	// List of allowed callback URLs for the identity providers.
	CallbackUrls pulumi.StringArrayOutput `pulumi:"callbackUrls"`
	// The client secret of the user pool client.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The default redirect URI. Must be in the list of callback URLs.
	DefaultRedirectUri pulumi.StringPtrOutput `pulumi:"defaultRedirectUri"`
	// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY,  USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
	ExplicitAuthFlows pulumi.StringArrayOutput `pulumi:"explicitAuthFlows"`
	// Should an application secret be generated.
	GenerateSecret pulumi.BoolPtrOutput `pulumi:"generateSecret"`
	// List of allowed logout URLs for the identity providers.
	LogoutUrls pulumi.StringArrayOutput `pulumi:"logoutUrls"`
	// The name of the application client.
	Name pulumi.StringOutput `pulumi:"name"`
	// Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors pulumi.StringOutput `pulumi:"preventUserExistenceErrors"`
	// List of user pool attributes the application client can read from.
	ReadAttributes pulumi.StringArrayOutput `pulumi:"readAttributes"`
	// The time limit in days refresh tokens are valid for.
	RefreshTokenValidity pulumi.IntPtrOutput `pulumi:"refreshTokenValidity"`
	// List of provider names for the identity providers that are supported on this client.
	SupportedIdentityProviders pulumi.StringArrayOutput `pulumi:"supportedIdentityProviders"`
	// The user pool the client belongs to.
	UserPoolId pulumi.StringOutput `pulumi:"userPoolId"`
	// List of user pool attributes the application client can write to.
	WriteAttributes pulumi.StringArrayOutput `pulumi:"writeAttributes"`
}

// NewUserPoolClient registers a new resource with the given unique name, arguments, and options.
func NewUserPoolClient(ctx *pulumi.Context,
	name string, args *UserPoolClientArgs, opts ...pulumi.ResourceOption) (*UserPoolClient, error) {
	if args == nil || args.UserPoolId == nil {
		return nil, errors.New("missing required argument 'UserPoolId'")
	}
	if args == nil {
		args = &UserPoolClientArgs{}
	}
	var resource UserPoolClient
	err := ctx.RegisterResource("aws:cognito/userPoolClient:UserPoolClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPoolClient gets an existing UserPoolClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPoolClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPoolClientState, opts ...pulumi.ResourceOption) (*UserPoolClient, error) {
	var resource UserPoolClient
	err := ctx.ReadResource("aws:cognito/userPoolClient:UserPoolClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPoolClient resources.
type userPoolClientState struct {
	// List of allowed OAuth flows (code, implicit, client_credentials).
	AllowedOauthFlows []string `pulumi:"allowedOauthFlows"`
	// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
	AllowedOauthFlowsUserPoolClient *bool `pulumi:"allowedOauthFlowsUserPoolClient"`
	// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
	AllowedOauthScopes []string `pulumi:"allowedOauthScopes"`
	// The Amazon Pinpoint analytics configuration for collecting metrics for this user pool.
	AnalyticsConfiguration *UserPoolClientAnalyticsConfiguration `pulumi:"analyticsConfiguration"`
	// List of allowed callback URLs for the identity providers.
	CallbackUrls []string `pulumi:"callbackUrls"`
	// The client secret of the user pool client.
	ClientSecret *string `pulumi:"clientSecret"`
	// The default redirect URI. Must be in the list of callback URLs.
	DefaultRedirectUri *string `pulumi:"defaultRedirectUri"`
	// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY,  USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
	ExplicitAuthFlows []string `pulumi:"explicitAuthFlows"`
	// Should an application secret be generated.
	GenerateSecret *bool `pulumi:"generateSecret"`
	// List of allowed logout URLs for the identity providers.
	LogoutUrls []string `pulumi:"logoutUrls"`
	// The name of the application client.
	Name *string `pulumi:"name"`
	// Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors *string `pulumi:"preventUserExistenceErrors"`
	// List of user pool attributes the application client can read from.
	ReadAttributes []string `pulumi:"readAttributes"`
	// The time limit in days refresh tokens are valid for.
	RefreshTokenValidity *int `pulumi:"refreshTokenValidity"`
	// List of provider names for the identity providers that are supported on this client.
	SupportedIdentityProviders []string `pulumi:"supportedIdentityProviders"`
	// The user pool the client belongs to.
	UserPoolId *string `pulumi:"userPoolId"`
	// List of user pool attributes the application client can write to.
	WriteAttributes []string `pulumi:"writeAttributes"`
}

type UserPoolClientState struct {
	// List of allowed OAuth flows (code, implicit, client_credentials).
	AllowedOauthFlows pulumi.StringArrayInput
	// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
	AllowedOauthFlowsUserPoolClient pulumi.BoolPtrInput
	// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
	AllowedOauthScopes pulumi.StringArrayInput
	// The Amazon Pinpoint analytics configuration for collecting metrics for this user pool.
	AnalyticsConfiguration UserPoolClientAnalyticsConfigurationPtrInput
	// List of allowed callback URLs for the identity providers.
	CallbackUrls pulumi.StringArrayInput
	// The client secret of the user pool client.
	ClientSecret pulumi.StringPtrInput
	// The default redirect URI. Must be in the list of callback URLs.
	DefaultRedirectUri pulumi.StringPtrInput
	// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY,  USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
	ExplicitAuthFlows pulumi.StringArrayInput
	// Should an application secret be generated.
	GenerateSecret pulumi.BoolPtrInput
	// List of allowed logout URLs for the identity providers.
	LogoutUrls pulumi.StringArrayInput
	// The name of the application client.
	Name pulumi.StringPtrInput
	// Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors pulumi.StringPtrInput
	// List of user pool attributes the application client can read from.
	ReadAttributes pulumi.StringArrayInput
	// The time limit in days refresh tokens are valid for.
	RefreshTokenValidity pulumi.IntPtrInput
	// List of provider names for the identity providers that are supported on this client.
	SupportedIdentityProviders pulumi.StringArrayInput
	// The user pool the client belongs to.
	UserPoolId pulumi.StringPtrInput
	// List of user pool attributes the application client can write to.
	WriteAttributes pulumi.StringArrayInput
}

func (UserPoolClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolClientState)(nil)).Elem()
}

type userPoolClientArgs struct {
	// List of allowed OAuth flows (code, implicit, client_credentials).
	AllowedOauthFlows []string `pulumi:"allowedOauthFlows"`
	// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
	AllowedOauthFlowsUserPoolClient *bool `pulumi:"allowedOauthFlowsUserPoolClient"`
	// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
	AllowedOauthScopes []string `pulumi:"allowedOauthScopes"`
	// The Amazon Pinpoint analytics configuration for collecting metrics for this user pool.
	AnalyticsConfiguration *UserPoolClientAnalyticsConfiguration `pulumi:"analyticsConfiguration"`
	// List of allowed callback URLs for the identity providers.
	CallbackUrls []string `pulumi:"callbackUrls"`
	// The default redirect URI. Must be in the list of callback URLs.
	DefaultRedirectUri *string `pulumi:"defaultRedirectUri"`
	// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY,  USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
	ExplicitAuthFlows []string `pulumi:"explicitAuthFlows"`
	// Should an application secret be generated.
	GenerateSecret *bool `pulumi:"generateSecret"`
	// List of allowed logout URLs for the identity providers.
	LogoutUrls []string `pulumi:"logoutUrls"`
	// The name of the application client.
	Name *string `pulumi:"name"`
	// Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors *string `pulumi:"preventUserExistenceErrors"`
	// List of user pool attributes the application client can read from.
	ReadAttributes []string `pulumi:"readAttributes"`
	// The time limit in days refresh tokens are valid for.
	RefreshTokenValidity *int `pulumi:"refreshTokenValidity"`
	// List of provider names for the identity providers that are supported on this client.
	SupportedIdentityProviders []string `pulumi:"supportedIdentityProviders"`
	// The user pool the client belongs to.
	UserPoolId string `pulumi:"userPoolId"`
	// List of user pool attributes the application client can write to.
	WriteAttributes []string `pulumi:"writeAttributes"`
}

// The set of arguments for constructing a UserPoolClient resource.
type UserPoolClientArgs struct {
	// List of allowed OAuth flows (code, implicit, client_credentials).
	AllowedOauthFlows pulumi.StringArrayInput
	// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
	AllowedOauthFlowsUserPoolClient pulumi.BoolPtrInput
	// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
	AllowedOauthScopes pulumi.StringArrayInput
	// The Amazon Pinpoint analytics configuration for collecting metrics for this user pool.
	AnalyticsConfiguration UserPoolClientAnalyticsConfigurationPtrInput
	// List of allowed callback URLs for the identity providers.
	CallbackUrls pulumi.StringArrayInput
	// The default redirect URI. Must be in the list of callback URLs.
	DefaultRedirectUri pulumi.StringPtrInput
	// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY,  USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
	ExplicitAuthFlows pulumi.StringArrayInput
	// Should an application secret be generated.
	GenerateSecret pulumi.BoolPtrInput
	// List of allowed logout URLs for the identity providers.
	LogoutUrls pulumi.StringArrayInput
	// The name of the application client.
	Name pulumi.StringPtrInput
	// Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors pulumi.StringPtrInput
	// List of user pool attributes the application client can read from.
	ReadAttributes pulumi.StringArrayInput
	// The time limit in days refresh tokens are valid for.
	RefreshTokenValidity pulumi.IntPtrInput
	// List of provider names for the identity providers that are supported on this client.
	SupportedIdentityProviders pulumi.StringArrayInput
	// The user pool the client belongs to.
	UserPoolId pulumi.StringInput
	// List of user pool attributes the application client can write to.
	WriteAttributes pulumi.StringArrayInput
}

func (UserPoolClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolClientArgs)(nil)).Elem()
}
