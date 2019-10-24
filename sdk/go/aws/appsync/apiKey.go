// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AppSync API Key.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_api_key.html.markdown.
type ApiKey struct {
	s *pulumi.ResourceState
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOpt) (*ApiKey, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	inputs := make(map[string]interface{})
	inputs["description"] = "Managed by Pulumi"
	if args == nil {
		inputs["apiId"] = nil
		inputs["expires"] = nil
	} else {
		inputs["apiId"] = args.ApiId
		inputs["description"] = args.Description
		inputs["expires"] = args.Expires
	}
	inputs["key"] = nil
	s, err := ctx.RegisterResource("aws:appsync/apiKey:ApiKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiKey{s: s}, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApiKeyState, opts ...pulumi.ResourceOpt) (*ApiKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiId"] = state.ApiId
		inputs["description"] = state.Description
		inputs["expires"] = state.Expires
		inputs["key"] = state.Key
	}
	s, err := ctx.ReadResource("aws:appsync/apiKey:ApiKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ApiKey) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ApiKey) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the associated AppSync API
func (r *ApiKey) ApiId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiId"])
}

// The API key description. Defaults to "Managed by Pulumi".
func (r *ApiKey) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
func (r *ApiKey) Expires() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["expires"])
}

// The API key
func (r *ApiKey) Key() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["key"])
}

// Input properties used for looking up and filtering ApiKey resources.
type ApiKeyState struct {
	// The ID of the associated AppSync API
	ApiId interface{}
	// The API key description. Defaults to "Managed by Pulumi".
	Description interface{}
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires interface{}
	// The API key
	Key interface{}
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// The ID of the associated AppSync API
	ApiId interface{}
	// The API key description. Defaults to "Managed by Pulumi".
	Description interface{}
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires interface{}
}
