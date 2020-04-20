// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 deployment.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
//
// > **Note:** Creating a deployment for an API requires at least one `apigatewayv2.Route` resource associated with that API.
type Deployment struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Whether the deployment was automatically released.
	AutoDeployed pulumi.BoolOutput `pulumi:"autoDeployed"`
	// The description for the deployment resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil {
		args = &DeploymentArgs{}
	}
	var resource Deployment
	err := ctx.RegisterResource("aws:apigatewayv2/deployment:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("aws:apigatewayv2/deployment:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
	// The API identifier.
	ApiId *string `pulumi:"apiId"`
	// Whether the deployment was automatically released.
	AutoDeployed *bool `pulumi:"autoDeployed"`
	// The description for the deployment resource.
	Description *string `pulumi:"description"`
}

type DeploymentState struct {
	// The API identifier.
	ApiId pulumi.StringPtrInput
	// Whether the deployment was automatically released.
	AutoDeployed pulumi.BoolPtrInput
	// The description for the deployment resource.
	Description pulumi.StringPtrInput
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The description for the deployment resource.
	Description *string `pulumi:"description"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// The description for the deployment resource.
	Description pulumi.StringPtrInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}
