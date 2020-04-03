// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regional Geo Match Set Resource
type GeoMatchSet struct {
	pulumi.CustomResourceState

	// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayOutput `pulumi:"geoMatchConstraints"`
	// The name or description of the Geo Match Set.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewGeoMatchSet registers a new resource with the given unique name, arguments, and options.
func NewGeoMatchSet(ctx *pulumi.Context,
	name string, args *GeoMatchSetArgs, opts ...pulumi.ResourceOption) (*GeoMatchSet, error) {
	if args == nil {
		args = &GeoMatchSetArgs{}
	}
	var resource GeoMatchSet
	err := ctx.RegisterResource("aws:wafregional/geoMatchSet:GeoMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeoMatchSet gets an existing GeoMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeoMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeoMatchSetState, opts ...pulumi.ResourceOption) (*GeoMatchSet, error) {
	var resource GeoMatchSet
	err := ctx.ReadResource("aws:wafregional/geoMatchSet:GeoMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeoMatchSet resources.
type geoMatchSetState struct {
	// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints []GeoMatchSetGeoMatchConstraint `pulumi:"geoMatchConstraints"`
	// The name or description of the Geo Match Set.
	Name *string `pulumi:"name"`
}

type GeoMatchSetState struct {
	// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayInput
	// The name or description of the Geo Match Set.
	Name pulumi.StringPtrInput
}

func (GeoMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*geoMatchSetState)(nil)).Elem()
}

type geoMatchSetArgs struct {
	// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints []GeoMatchSetGeoMatchConstraint `pulumi:"geoMatchConstraints"`
	// The name or description of the Geo Match Set.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a GeoMatchSet resource.
type GeoMatchSetArgs struct {
	// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayInput
	// The name or description of the Geo Match Set.
	Name pulumi.StringPtrInput
}

func (GeoMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geoMatchSetArgs)(nil)).Elem()
}
