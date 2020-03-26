// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package waf

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Rule Resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_rule.html.markdown.
type Rule struct {
	pulumi.CustomResourceState

	// The ARN of the WAF rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates RulePredicateArrayOutput `pulumi:"predicates"`
	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	if args == nil {
		args = &RuleArgs{}
	}
	var resource Rule
	err := ctx.RegisterResource("aws:waf/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("aws:waf/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The ARN of the WAF rule.
	Arn *string `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RulePredicate `pulumi:"predicates"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

type RuleState struct {
	// The ARN of the WAF rule.
	Arn pulumi.StringPtrInput
	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName pulumi.StringPtrInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RulePredicateArrayInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RulePredicate `pulumi:"predicates"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName pulumi.StringInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RulePredicateArrayInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}
