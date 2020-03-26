// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elasticloadbalancing

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an application cookie stickiness policy, which allows an ELB to wed its sticky cookie's expiration to a cookie generated by your application.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/app_cookie_stickiness_policy.html.markdown.
type AppCookieStickinessPolicy struct {
	pulumi.CustomResourceState

	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName pulumi.StringOutput `pulumi:"cookieName"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntOutput `pulumi:"lbPort"`
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringOutput `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAppCookieStickinessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAppCookieStickinessPolicy(ctx *pulumi.Context,
	name string, args *AppCookieStickinessPolicyArgs, opts ...pulumi.ResourceOption) (*AppCookieStickinessPolicy, error) {
	if args == nil || args.CookieName == nil {
		return nil, errors.New("missing required argument 'CookieName'")
	}
	if args == nil || args.LbPort == nil {
		return nil, errors.New("missing required argument 'LbPort'")
	}
	if args == nil || args.LoadBalancer == nil {
		return nil, errors.New("missing required argument 'LoadBalancer'")
	}
	if args == nil {
		args = &AppCookieStickinessPolicyArgs{}
	}
	var resource AppCookieStickinessPolicy
	err := ctx.RegisterResource("aws:elasticloadbalancing/appCookieStickinessPolicy:AppCookieStickinessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppCookieStickinessPolicy gets an existing AppCookieStickinessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppCookieStickinessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppCookieStickinessPolicyState, opts ...pulumi.ResourceOption) (*AppCookieStickinessPolicy, error) {
	var resource AppCookieStickinessPolicy
	err := ctx.ReadResource("aws:elasticloadbalancing/appCookieStickinessPolicy:AppCookieStickinessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppCookieStickinessPolicy resources.
type appCookieStickinessPolicyState struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName *string `pulumi:"cookieName"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort *int `pulumi:"lbPort"`
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer *string `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name *string `pulumi:"name"`
}

type AppCookieStickinessPolicyState struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName pulumi.StringPtrInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntPtrInput
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringPtrInput
	// The name of the stickiness policy.
	Name pulumi.StringPtrInput
}

func (AppCookieStickinessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appCookieStickinessPolicyState)(nil)).Elem()
}

type appCookieStickinessPolicyArgs struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName string `pulumi:"cookieName"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort int `pulumi:"lbPort"`
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer string `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AppCookieStickinessPolicy resource.
type AppCookieStickinessPolicyArgs struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName pulumi.StringInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput
	// The name of the stickiness policy.
	Name pulumi.StringPtrInput
}

func (AppCookieStickinessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appCookieStickinessPolicyArgs)(nil)).Elem()
}
