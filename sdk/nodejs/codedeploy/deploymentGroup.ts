// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CodeDeploy Deployment Group for a CodeDeploy Application
 * 
 * > **NOTE on blue/green deployments:** When using `green_fleet_provisioning_option` with the `COPY_AUTO_SCALING_GROUP` action, CodeDeploy will create a new ASG with a different name. This ASG is _not_ managed by terraform and will conflict with existing configuration and state. You may want to use a different approach to managing deployments that involve multiple ASG, such as `DISCOVER_EXISTING` with separate blue and green ASG.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleApplication = new aws.codedeploy.Application("example", {
 *     name: "example-app",
 * });
 * const exampleRole = new aws.iam.Role("example", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Sid": "",
 *       "Effect": "Allow",
 *       "Principal": {
 *         "Service": "codedeploy.amazonaws.com"
 *       },
 *       "Action": "sts:AssumeRole"
 *     }
 *   ]
 * }
 * `,
 *     name: "example-role",
 * });
 * const exampleTopic = new aws.sns.Topic("example", {
 *     name: "example-topic",
 * });
 * const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("example", {
 *     alarmConfiguration: {
 *         alarms: ["my-alarm-name"],
 *         enabled: true,
 *     },
 *     appName: exampleApplication.name,
 *     autoRollbackConfiguration: {
 *         enabled: true,
 *         events: ["DEPLOYMENT_FAILURE"],
 *     },
 *     deploymentGroupName: "example-group",
 *     ec2TagSets: [{
 *         ec2TagFilters: [
 *             {
 *                 key: "filterkey1",
 *                 type: "KEY_AND_VALUE",
 *                 value: "filtervalue",
 *             },
 *             {
 *                 key: "filterkey2",
 *                 type: "KEY_AND_VALUE",
 *                 value: "filtervalue",
 *             },
 *         ],
 *     }],
 *     serviceRoleArn: exampleRole.arn,
 *     triggerConfigurations: [{
 *         triggerEvents: ["DeploymentFailure"],
 *         triggerName: "example-trigger",
 *         triggerTargetArn: exampleTopic.arn,
 *     }],
 * });
 * const aWSCodeDeployRole = new aws.iam.RolePolicyAttachment("AWSCodeDeployRole", {
 *     policyArn: "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole",
 *     role: exampleRole.name,
 * });
 * ```
 * 
 * ### Blue Green Deployments with ECS
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleApplication = new aws.codedeploy.Application("example", {
 *     computePlatform: "ECS",
 *     name: "example",
 * });
 * const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("example", {
 *     appName: exampleApplication.name,
 *     autoRollbackConfiguration: {
 *         enabled: true,
 *         events: ["DEPLOYMENT_FAILURE"],
 *     },
 *     blueGreenDeploymentConfig: {
 *         deploymentReadyOption: {
 *             actionOnTimeout: "CONTINUE_DEPLOYMENT",
 *         },
 *         terminateBlueInstancesOnDeploymentSuccess: {
 *             action: "TERMINATE",
 *             terminationWaitTimeInMinutes: 5,
 *         },
 *     },
 *     deploymentConfigName: "CodeDeployDefault.ECSAllAtOnce",
 *     deploymentGroupName: "example",
 *     deploymentStyle: {
 *         deploymentOption: "WITH_TRAFFIC_CONTROL",
 *         deploymentType: "BLUE_GREEN",
 *     },
 *     ecsService: {
 *         clusterName: aws_ecs_cluster_example.name,
 *         serviceName: aws_ecs_service_example.name,
 *     },
 *     loadBalancerInfo: {
 *         targetGroupPairInfo: {
 *             prodTrafficRoute: {
 *                 listenerArns: [aws_lb_listener_example.arn],
 *             },
 *             targetGroups: [
 *                 {
 *                     name: aws_lb_target_group_blue.name,
 *                 },
 *                 {
 *                     name: aws_lb_target_group_green.name,
 *                 },
 *             ],
 *         },
 *     },
 *     serviceRoleArn: aws_iam_role_example.arn,
 * });
 * ```
 * 
 * ### Blue Green Deployments with Servers and Classic ELB
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleApplication = new aws.codedeploy.Application("example", {
 *     name: "example-app",
 * });
 * const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("example", {
 *     appName: exampleApplication.name,
 *     blueGreenDeploymentConfig: {
 *         deploymentReadyOption: {
 *             actionOnTimeout: "STOP_DEPLOYMENT",
 *             waitTimeInMinutes: 60,
 *         },
 *         greenFleetProvisioningOption: {
 *             action: "DISCOVER_EXISTING",
 *         },
 *         terminateBlueInstancesOnDeploymentSuccess: {
 *             action: "KEEP_ALIVE",
 *         },
 *     },
 *     deploymentGroupName: "example-group",
 *     deploymentStyle: {
 *         deploymentOption: "WITH_TRAFFIC_CONTROL",
 *         deploymentType: "BLUE_GREEN",
 *     },
 *     loadBalancerInfo: {
 *         elbInfos: [{
 *             name: aws_elb_example.name,
 *         }],
 *     },
 *     serviceRoleArn: aws_iam_role_example.arn,
 * });
 * ```
 */
export class DeploymentGroup extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentGroupState, opts?: pulumi.CustomResourceOptions): DeploymentGroup {
        return new DeploymentGroup(name, <any>state, { ...opts, id: id });
    }

    /**
     * Configuration block of alarms associated with the deployment group (documented below).
     */
    public readonly alarmConfiguration: pulumi.Output<{ alarms?: string[], enabled?: boolean, ignorePollAlarmFailure?: boolean } | undefined>;
    /**
     * The name of the application.
     */
    public readonly appName: pulumi.Output<string>;
    /**
     * Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
     */
    public readonly autoRollbackConfiguration: pulumi.Output<{ enabled?: boolean, events?: string[] } | undefined>;
    /**
     * Autoscaling groups associated with the deployment group.
     */
    public readonly autoscalingGroups: pulumi.Output<string[] | undefined>;
    /**
     * Configuration block of the blue/green deployment options for a deployment group (documented below).
     */
    public readonly blueGreenDeploymentConfig: pulumi.Output<{ deploymentReadyOption?: { actionOnTimeout?: string, waitTimeInMinutes?: number }, greenFleetProvisioningOption: { action?: string }, terminateBlueInstancesOnDeploymentSuccess?: { action?: string, terminationWaitTimeInMinutes?: number } }>;
    /**
     * The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
     */
    public readonly deploymentConfigName: pulumi.Output<string | undefined>;
    /**
     * The name of the deployment group.
     */
    public readonly deploymentGroupName: pulumi.Output<string>;
    /**
     * Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
     */
    public readonly deploymentStyle: pulumi.Output<{ deploymentOption?: string, deploymentType?: string }>;
    /**
     * Tag filters associated with the deployment group. See the AWS docs for details.
     */
    public readonly ec2TagFilters: pulumi.Output<{ key?: string, type?: string, value?: string }[] | undefined>;
    /**
     * Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
     */
    public readonly ec2TagSets: pulumi.Output<{ ec2TagFilters?: { key?: string, type?: string, value?: string }[] }[] | undefined>;
    /**
     * Configuration block(s) of the ECS services for a deployment group (documented below).
     */
    public readonly ecsService: pulumi.Output<{ clusterName: string, serviceName: string } | undefined>;
    /**
     * Single configuration block of the load balancer to use in a blue/green deployment (documented below).
     */
    public readonly loadBalancerInfo: pulumi.Output<{ elbInfos?: { name?: string }[], targetGroupInfos?: { name?: string }[], targetGroupPairInfo?: { prodTrafficRoute: { listenerArns: string[] }, targetGroups: { name: string }[], testTrafficRoute?: { listenerArns: string[] } } }>;
    /**
     * On premise tag filters associated with the group. See the AWS docs for details.
     */
    public readonly onPremisesInstanceTagFilters: pulumi.Output<{ key?: string, type?: string, value?: string }[] | undefined>;
    /**
     * The service role ARN that allows deployments.
     */
    public readonly serviceRoleArn: pulumi.Output<string>;
    /**
     * Configuration block(s) of the triggers for the deployment group (documented below).
     */
    public readonly triggerConfigurations: pulumi.Output<{ triggerEvents: string[], triggerName: string, triggerTargetArn: string }[] | undefined>;

    /**
     * Create a DeploymentGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentGroupArgs | DeploymentGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DeploymentGroupState = argsOrState as DeploymentGroupState | undefined;
            inputs["alarmConfiguration"] = state ? state.alarmConfiguration : undefined;
            inputs["appName"] = state ? state.appName : undefined;
            inputs["autoRollbackConfiguration"] = state ? state.autoRollbackConfiguration : undefined;
            inputs["autoscalingGroups"] = state ? state.autoscalingGroups : undefined;
            inputs["blueGreenDeploymentConfig"] = state ? state.blueGreenDeploymentConfig : undefined;
            inputs["deploymentConfigName"] = state ? state.deploymentConfigName : undefined;
            inputs["deploymentGroupName"] = state ? state.deploymentGroupName : undefined;
            inputs["deploymentStyle"] = state ? state.deploymentStyle : undefined;
            inputs["ec2TagFilters"] = state ? state.ec2TagFilters : undefined;
            inputs["ec2TagSets"] = state ? state.ec2TagSets : undefined;
            inputs["ecsService"] = state ? state.ecsService : undefined;
            inputs["loadBalancerInfo"] = state ? state.loadBalancerInfo : undefined;
            inputs["onPremisesInstanceTagFilters"] = state ? state.onPremisesInstanceTagFilters : undefined;
            inputs["serviceRoleArn"] = state ? state.serviceRoleArn : undefined;
            inputs["triggerConfigurations"] = state ? state.triggerConfigurations : undefined;
        } else {
            const args = argsOrState as DeploymentGroupArgs | undefined;
            if (!args || args.appName === undefined) {
                throw new Error("Missing required property 'appName'");
            }
            if (!args || args.deploymentGroupName === undefined) {
                throw new Error("Missing required property 'deploymentGroupName'");
            }
            if (!args || args.serviceRoleArn === undefined) {
                throw new Error("Missing required property 'serviceRoleArn'");
            }
            inputs["alarmConfiguration"] = args ? args.alarmConfiguration : undefined;
            inputs["appName"] = args ? args.appName : undefined;
            inputs["autoRollbackConfiguration"] = args ? args.autoRollbackConfiguration : undefined;
            inputs["autoscalingGroups"] = args ? args.autoscalingGroups : undefined;
            inputs["blueGreenDeploymentConfig"] = args ? args.blueGreenDeploymentConfig : undefined;
            inputs["deploymentConfigName"] = args ? args.deploymentConfigName : undefined;
            inputs["deploymentGroupName"] = args ? args.deploymentGroupName : undefined;
            inputs["deploymentStyle"] = args ? args.deploymentStyle : undefined;
            inputs["ec2TagFilters"] = args ? args.ec2TagFilters : undefined;
            inputs["ec2TagSets"] = args ? args.ec2TagSets : undefined;
            inputs["ecsService"] = args ? args.ecsService : undefined;
            inputs["loadBalancerInfo"] = args ? args.loadBalancerInfo : undefined;
            inputs["onPremisesInstanceTagFilters"] = args ? args.onPremisesInstanceTagFilters : undefined;
            inputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            inputs["triggerConfigurations"] = args ? args.triggerConfigurations : undefined;
        }
        super("aws:codedeploy/deploymentGroup:DeploymentGroup", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeploymentGroup resources.
 */
export interface DeploymentGroupState {
    /**
     * Configuration block of alarms associated with the deployment group (documented below).
     */
    readonly alarmConfiguration?: pulumi.Input<{ alarms?: pulumi.Input<pulumi.Input<string>[]>, enabled?: pulumi.Input<boolean>, ignorePollAlarmFailure?: pulumi.Input<boolean> }>;
    /**
     * The name of the application.
     */
    readonly appName?: pulumi.Input<string>;
    /**
     * Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
     */
    readonly autoRollbackConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, events?: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * Autoscaling groups associated with the deployment group.
     */
    readonly autoscalingGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block of the blue/green deployment options for a deployment group (documented below).
     */
    readonly blueGreenDeploymentConfig?: pulumi.Input<{ deploymentReadyOption?: pulumi.Input<{ actionOnTimeout?: pulumi.Input<string>, waitTimeInMinutes?: pulumi.Input<number> }>, greenFleetProvisioningOption?: pulumi.Input<{ action?: pulumi.Input<string> }>, terminateBlueInstancesOnDeploymentSuccess?: pulumi.Input<{ action?: pulumi.Input<string>, terminationWaitTimeInMinutes?: pulumi.Input<number> }> }>;
    /**
     * The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
     */
    readonly deploymentConfigName?: pulumi.Input<string>;
    /**
     * The name of the deployment group.
     */
    readonly deploymentGroupName?: pulumi.Input<string>;
    /**
     * Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
     */
    readonly deploymentStyle?: pulumi.Input<{ deploymentOption?: pulumi.Input<string>, deploymentType?: pulumi.Input<string> }>;
    /**
     * Tag filters associated with the deployment group. See the AWS docs for details.
     */
    readonly ec2TagFilters?: pulumi.Input<pulumi.Input<{ key?: pulumi.Input<string>, type?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]>;
    /**
     * Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
     */
    readonly ec2TagSets?: pulumi.Input<pulumi.Input<{ ec2TagFilters?: pulumi.Input<pulumi.Input<{ key?: pulumi.Input<string>, type?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]> }>[]>;
    /**
     * Configuration block(s) of the ECS services for a deployment group (documented below).
     */
    readonly ecsService?: pulumi.Input<{ clusterName: pulumi.Input<string>, serviceName: pulumi.Input<string> }>;
    /**
     * Single configuration block of the load balancer to use in a blue/green deployment (documented below).
     */
    readonly loadBalancerInfo?: pulumi.Input<{ elbInfos?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string> }>[]>, targetGroupInfos?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string> }>[]>, targetGroupPairInfo?: pulumi.Input<{ prodTrafficRoute: pulumi.Input<{ listenerArns: pulumi.Input<pulumi.Input<string>[]> }>, targetGroups: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string> }>[]>, testTrafficRoute?: pulumi.Input<{ listenerArns: pulumi.Input<pulumi.Input<string>[]> }> }> }>;
    /**
     * On premise tag filters associated with the group. See the AWS docs for details.
     */
    readonly onPremisesInstanceTagFilters?: pulumi.Input<pulumi.Input<{ key?: pulumi.Input<string>, type?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]>;
    /**
     * The service role ARN that allows deployments.
     */
    readonly serviceRoleArn?: pulumi.Input<string>;
    /**
     * Configuration block(s) of the triggers for the deployment group (documented below).
     */
    readonly triggerConfigurations?: pulumi.Input<pulumi.Input<{ triggerEvents: pulumi.Input<pulumi.Input<string>[]>, triggerName: pulumi.Input<string>, triggerTargetArn: pulumi.Input<string> }>[]>;
}

/**
 * The set of arguments for constructing a DeploymentGroup resource.
 */
export interface DeploymentGroupArgs {
    /**
     * Configuration block of alarms associated with the deployment group (documented below).
     */
    readonly alarmConfiguration?: pulumi.Input<{ alarms?: pulumi.Input<pulumi.Input<string>[]>, enabled?: pulumi.Input<boolean>, ignorePollAlarmFailure?: pulumi.Input<boolean> }>;
    /**
     * The name of the application.
     */
    readonly appName: pulumi.Input<string>;
    /**
     * Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
     */
    readonly autoRollbackConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, events?: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * Autoscaling groups associated with the deployment group.
     */
    readonly autoscalingGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block of the blue/green deployment options for a deployment group (documented below).
     */
    readonly blueGreenDeploymentConfig?: pulumi.Input<{ deploymentReadyOption?: pulumi.Input<{ actionOnTimeout?: pulumi.Input<string>, waitTimeInMinutes?: pulumi.Input<number> }>, greenFleetProvisioningOption?: pulumi.Input<{ action?: pulumi.Input<string> }>, terminateBlueInstancesOnDeploymentSuccess?: pulumi.Input<{ action?: pulumi.Input<string>, terminationWaitTimeInMinutes?: pulumi.Input<number> }> }>;
    /**
     * The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
     */
    readonly deploymentConfigName?: pulumi.Input<string>;
    /**
     * The name of the deployment group.
     */
    readonly deploymentGroupName: pulumi.Input<string>;
    /**
     * Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
     */
    readonly deploymentStyle?: pulumi.Input<{ deploymentOption?: pulumi.Input<string>, deploymentType?: pulumi.Input<string> }>;
    /**
     * Tag filters associated with the deployment group. See the AWS docs for details.
     */
    readonly ec2TagFilters?: pulumi.Input<pulumi.Input<{ key?: pulumi.Input<string>, type?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]>;
    /**
     * Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
     */
    readonly ec2TagSets?: pulumi.Input<pulumi.Input<{ ec2TagFilters?: pulumi.Input<pulumi.Input<{ key?: pulumi.Input<string>, type?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]> }>[]>;
    /**
     * Configuration block(s) of the ECS services for a deployment group (documented below).
     */
    readonly ecsService?: pulumi.Input<{ clusterName: pulumi.Input<string>, serviceName: pulumi.Input<string> }>;
    /**
     * Single configuration block of the load balancer to use in a blue/green deployment (documented below).
     */
    readonly loadBalancerInfo?: pulumi.Input<{ elbInfos?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string> }>[]>, targetGroupInfos?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string> }>[]>, targetGroupPairInfo?: pulumi.Input<{ prodTrafficRoute: pulumi.Input<{ listenerArns: pulumi.Input<pulumi.Input<string>[]> }>, targetGroups: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string> }>[]>, testTrafficRoute?: pulumi.Input<{ listenerArns: pulumi.Input<pulumi.Input<string>[]> }> }> }>;
    /**
     * On premise tag filters associated with the group. See the AWS docs for details.
     */
    readonly onPremisesInstanceTagFilters?: pulumi.Input<pulumi.Input<{ key?: pulumi.Input<string>, type?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]>;
    /**
     * The service role ARN that allows deployments.
     */
    readonly serviceRoleArn: pulumi.Input<string>;
    /**
     * Configuration block(s) of the triggers for the deployment group (documented below).
     */
    readonly triggerConfigurations?: pulumi.Input<pulumi.Input<{ triggerEvents: pulumi.Input<pulumi.Input<string>[]>, triggerName: pulumi.Input<string>, triggerTargetArn: pulumi.Input<string> }>[]>;
}
