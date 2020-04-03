// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Route53 Hosted Zone VPC association. VPC associations can only be made on private zones.
//
// > **NOTE:** Unless explicit association ordering is required (e.g. a separate cross-account association authorization), usage of this resource is not recommended. Use the `vpc` configuration blocks available within the [`route53.Zone` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone.html) instead.
//
// > **NOTE:** This provider provides both this standalone Zone VPC Association resource and exclusive VPC associations defined in-line in the [`route53.Zone` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone.html) via `vpc` configuration blocks. At this time, you cannot use those in-line VPC associations in conjunction with this resource and the same zone ID otherwise it will cause a perpetual difference in plan output. You can optionally use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) in the `route53.Zone` resource to manage additional associations via this resource.
type ZoneAssociation struct {
	pulumi.CustomResourceState

	// The VPC to associate with the private hosted zone.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringOutput `pulumi:"vpcRegion"`
	// The private hosted zone to associate.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewZoneAssociation registers a new resource with the given unique name, arguments, and options.
func NewZoneAssociation(ctx *pulumi.Context,
	name string, args *ZoneAssociationArgs, opts ...pulumi.ResourceOption) (*ZoneAssociation, error) {
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	if args == nil {
		args = &ZoneAssociationArgs{}
	}
	var resource ZoneAssociation
	err := ctx.RegisterResource("aws:route53/zoneAssociation:ZoneAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneAssociation gets an existing ZoneAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneAssociationState, opts ...pulumi.ResourceOption) (*ZoneAssociation, error) {
	var resource ZoneAssociation
	err := ctx.ReadResource("aws:route53/zoneAssociation:ZoneAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneAssociation resources.
type zoneAssociationState struct {
	// The VPC to associate with the private hosted zone.
	VpcId *string `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion *string `pulumi:"vpcRegion"`
	// The private hosted zone to associate.
	ZoneId *string `pulumi:"zoneId"`
}

type ZoneAssociationState struct {
	// The VPC to associate with the private hosted zone.
	VpcId pulumi.StringPtrInput
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringPtrInput
	// The private hosted zone to associate.
	ZoneId pulumi.StringPtrInput
}

func (ZoneAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneAssociationState)(nil)).Elem()
}

type zoneAssociationArgs struct {
	// The VPC to associate with the private hosted zone.
	VpcId string `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion *string `pulumi:"vpcRegion"`
	// The private hosted zone to associate.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ZoneAssociation resource.
type ZoneAssociationArgs struct {
	// The VPC to associate with the private hosted zone.
	VpcId pulumi.StringInput
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringPtrInput
	// The private hosted zone to associate.
	ZoneId pulumi.StringInput
}

func (ZoneAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneAssociationArgs)(nil)).Elem()
}
