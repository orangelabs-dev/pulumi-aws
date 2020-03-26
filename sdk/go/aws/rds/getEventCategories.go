// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rds

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func GetEventCategories(ctx *pulumi.Context, args *GetEventCategoriesArgs, opts ...pulumi.InvokeOption) (*GetEventCategoriesResult, error) {
	var rv GetEventCategoriesResult
	err := ctx.Invoke("aws:rds/getEventCategories:getEventCategories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEventCategories.
type GetEventCategoriesArgs struct {
	// The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot.
	SourceType *string `pulumi:"sourceType"`
}

// A collection of values returned by getEventCategories.
type GetEventCategoriesResult struct {
	// A list of the event categories.
	EventCategories []string `pulumi:"eventCategories"`
	// id is the provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	SourceType *string `pulumi:"sourceType"`
}
