// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SES receipt filter resource
type ReceiptFilter struct {
	pulumi.CustomResourceState

	// The IP address or address range to filter, in CIDR notation
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// The name of the filter
	Name pulumi.StringOutput `pulumi:"name"`
	// Block or Allow
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewReceiptFilter registers a new resource with the given unique name, arguments, and options.
func NewReceiptFilter(ctx *pulumi.Context,
	name string, args *ReceiptFilterArgs, opts ...pulumi.ResourceOption) (*ReceiptFilter, error) {
	if args == nil || args.Cidr == nil {
		return nil, errors.New("missing required argument 'Cidr'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil {
		args = &ReceiptFilterArgs{}
	}
	var resource ReceiptFilter
	err := ctx.RegisterResource("aws:ses/receiptFilter:ReceiptFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReceiptFilter gets an existing ReceiptFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReceiptFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReceiptFilterState, opts ...pulumi.ResourceOption) (*ReceiptFilter, error) {
	var resource ReceiptFilter
	err := ctx.ReadResource("aws:ses/receiptFilter:ReceiptFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReceiptFilter resources.
type receiptFilterState struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr *string `pulumi:"cidr"`
	// The name of the filter
	Name *string `pulumi:"name"`
	// Block or Allow
	Policy *string `pulumi:"policy"`
}

type ReceiptFilterState struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr pulumi.StringPtrInput
	// The name of the filter
	Name pulumi.StringPtrInput
	// Block or Allow
	Policy pulumi.StringPtrInput
}

func (ReceiptFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptFilterState)(nil)).Elem()
}

type receiptFilterArgs struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr string `pulumi:"cidr"`
	// The name of the filter
	Name *string `pulumi:"name"`
	// Block or Allow
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a ReceiptFilter resource.
type ReceiptFilterArgs struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr pulumi.StringInput
	// The name of the filter
	Name pulumi.StringPtrInput
	// Block or Allow
	Policy pulumi.StringInput
}

func (ReceiptFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptFilterArgs)(nil)).Elem()
}
