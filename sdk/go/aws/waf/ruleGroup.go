// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Rule Group Resource
type RuleGroup struct {
	pulumi.CustomResourceState

	// A list of activated rules, see below
	ActivatedRules RuleGroupActivatedRuleArrayOutput `pulumi:"activatedRules"`
	// The ARN of the WAF rule group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A friendly name for the metrics from the rule group
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// A friendly name of the rule group
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewRuleGroup(ctx *pulumi.Context,
	name string, args *RuleGroupArgs, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	if args == nil {
		args = &RuleGroupArgs{}
	}
	var resource RuleGroup
	err := ctx.RegisterResource("aws:waf/ruleGroup:RuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleGroup gets an existing RuleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleGroupState, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	var resource RuleGroup
	err := ctx.ReadResource("aws:waf/ruleGroup:RuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleGroup resources.
type ruleGroupState struct {
	// A list of activated rules, see below
	ActivatedRules []RuleGroupActivatedRule `pulumi:"activatedRules"`
	// The ARN of the WAF rule group.
	Arn *string `pulumi:"arn"`
	// A friendly name for the metrics from the rule group
	MetricName *string `pulumi:"metricName"`
	// A friendly name of the rule group
	Name *string `pulumi:"name"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

type RuleGroupState struct {
	// A list of activated rules, see below
	ActivatedRules RuleGroupActivatedRuleArrayInput
	// The ARN of the WAF rule group.
	Arn pulumi.StringPtrInput
	// A friendly name for the metrics from the rule group
	MetricName pulumi.StringPtrInput
	// A friendly name of the rule group
	Name pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (RuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupState)(nil)).Elem()
}

type ruleGroupArgs struct {
	// A list of activated rules, see below
	ActivatedRules []RuleGroupActivatedRule `pulumi:"activatedRules"`
	// A friendly name for the metrics from the rule group
	MetricName string `pulumi:"metricName"`
	// A friendly name of the rule group
	Name *string `pulumi:"name"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a RuleGroup resource.
type RuleGroupArgs struct {
	// A list of activated rules, see below
	ActivatedRules RuleGroupActivatedRuleArrayInput
	// A friendly name for the metrics from the rule group
	MetricName pulumi.StringInput
	// A friendly name of the rule group
	Name pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (RuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupArgs)(nil)).Elem()
}
