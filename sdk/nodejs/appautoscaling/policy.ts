// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Application AutoScaling Policy resource.
 * 
 * ## Example Usage
 * 
 * ### DynamoDB Table Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const dynamodbTableReadTarget = new aws.appautoscaling.Target("dynamodb_table_read_target", {
 *     maxCapacity: 100,
 *     minCapacity: 5,
 *     resourceId: "table/tableName",
 *     roleArn: aws_iam_role_DynamoDBAutoscaleRole.arn.apply(arn => arn),
 *     scalableDimension: "dynamodb:table:ReadCapacityUnits",
 *     serviceNamespace: "dynamodb",
 * });
 * const dynamodbTableReadPolicy = new aws.appautoscaling.Policy("dynamodb_table_read_policy", {
 *     name: dynamodbTableReadTarget.resourceId.apply(resourceId => `DynamoDBReadCapacityUtilization:${resourceId}`),
 *     policyType: "TargetTrackingScaling",
 *     resourceId: dynamodbTableReadTarget.resourceId,
 *     scalableDimension: dynamodbTableReadTarget.scalableDimension,
 *     serviceNamespace: dynamodbTableReadTarget.serviceNamespace,
 *     targetTrackingScalingPolicyConfiguration: {
 *         predefinedMetricSpecification: {
 *             predefinedMetricType: "DynamoDBReadCapacityUtilization",
 *         },
 *         targetValue: 70,
 *     },
 * });
 * ```
 * 
 * ### ECS Service Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ecsTarget = new aws.appautoscaling.Target("ecs_target", {
 *     maxCapacity: 4,
 *     minCapacity: 1,
 *     resourceId: "service/clusterName/serviceName",
 *     roleArn: var_ecs_iam_role,
 *     scalableDimension: "ecs:service:DesiredCount",
 *     serviceNamespace: "ecs",
 * });
 * const ecsPolicy = new aws.appautoscaling.Policy("ecs_policy", {
 *     name: "scale-down",
 *     policyType: "StepScaling",
 *     resourceId: "service/clusterName/serviceName",
 *     scalableDimension: "ecs:service:DesiredCount",
 *     serviceNamespace: "ecs",
 *     stepScalingPolicyConfigurations: [{
 *         adjustmentType: "ChangeInCapacity",
 *         cooldown: 60,
 *         metricAggregationType: "Maximum",
 *         stepAdjustments: [{
 *             metricIntervalUpperBound: "0",
 *             scalingAdjustment: -1,
 *         }],
 *     }],
 * }, {dependsOn: [ecsTarget]});
 * ```
 * 
 * ### Preserve desired count when updating an autoscaled ECS Service
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ecsService = new aws.ecs.Service("ecs_service", {
 *     cluster: "clusterName",
 *     desiredCount: 2,
 *     name: "serviceName",
 *     taskDefinition: "taskDefinitionFamily:1",
 * });
 * ```
 * 
 * ### Aurora Read Replica Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const replicasTarget = new aws.appautoscaling.Target("replicas", {
 *     maxCapacity: 15,
 *     minCapacity: 1,
 *     resourceId: aws_rds_cluster_example.id.apply(id => `cluster:${id}`),
 *     scalableDimension: "rds:cluster:ReadReplicaCount",
 *     serviceNamespace: "rds",
 * });
 * const replicasPolicy = new aws.appautoscaling.Policy("replicas", {
 *     name: "cpu-auto-scaling",
 *     policyType: "TargetTrackingScaling",
 *     resourceId: replicasTarget.resourceId,
 *     scalableDimension: replicasTarget.scalableDimension,
 *     serviceNamespace: replicasTarget.serviceNamespace,
 *     targetTrackingScalingPolicyConfiguration: {
 *         predefinedMetricSpecification: {
 *             predefinedMetricType: "RDSReaderAverageCPUUtilization",
 *         },
 *         scaleInCooldown: 300,
 *         scaleOutCooldown: 300,
 *         targetValue: 75,
 *     },
 * });
 * ```
 * 
 * ## Nested fields
 * 
 * ### `target_tracking_scaling_policy_configuration`
 * 
 * * `target_value` - (Required) The target value for the metric.
 * * `disable_scale_in` - (Optional) Indicates whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is `false`.
 * * `scale_in_cooldown` - (Optional) The amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
 * * `scale_out_cooldown` - (Optional) The amount of time, in seconds, after a scale out activity completes before another scale out activity can start.
 * * `customized_metric_specification` - (Optional) Reserved for future use. See supported fields below.
 * * `predefined_metric_specification` - (Optional) A predefined metric. See supported fields below.
 * 
 * ### `customized_metric_specification`
 * 
 * * `dimensions` - (Optional) The dimensions of the metric.
 * * `metric_name` - (Required) The name of the metric.
 * * `namespace` - (Required) The namespace of the metric.
 * * `statistic` - (Required) The statistic of the metric.
 * * `unit` - (Optional) The unit of the metric.
 * 
 * ### `predefined_metric_specification`
 * 
 * * `predefined_metric_type` - (Required) The metric type.
 * * `resource_label` - (Optional) Reserved for future use.
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /**
     * The scaling policy's adjustment type.
     */
    public readonly adjustmentType: pulumi.Output<string | undefined>;
    public readonly alarms: pulumi.Output<string[] | undefined>;
    /**
     * The ARN assigned by AWS to the scaling policy.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    public readonly cooldown: pulumi.Output<number | undefined>;
    public readonly metricAggregationType: pulumi.Output<string | undefined>;
    public readonly minAdjustmentMagnitude: pulumi.Output<number | undefined>;
    /**
     * The name of the policy.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
     */
    public readonly policyType: pulumi.Output<string | undefined>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly resourceId: pulumi.Output<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly scalableDimension: pulumi.Output<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly serviceNamespace: pulumi.Output<string>;
    public readonly stepAdjustments: pulumi.Output<{ metricIntervalLowerBound?: string, metricIntervalUpperBound?: string, scalingAdjustment: number }[] | undefined>;
    /**
     * Step scaling policy configuration, requires `policy_type = "StepScaling"` (default). See supported fields below.
     */
    public readonly stepScalingPolicyConfigurations: pulumi.Output<{ adjustmentType?: string, cooldown?: number, metricAggregationType?: string, minAdjustmentMagnitude?: number, stepAdjustments?: { metricIntervalLowerBound?: string, metricIntervalUpperBound?: string, scalingAdjustment: number }[] }[] | undefined>;
    /**
     * A target tracking policy, requires `policy_type = "TargetTrackingScaling"`. See supported fields below.
     */
    public readonly targetTrackingScalingPolicyConfiguration: pulumi.Output<{ customizedMetricSpecification?: { dimensions?: { name: string, value: string }[], metricName: string, namespace: string, statistic: string, unit?: string }, disableScaleIn?: boolean, predefinedMetricSpecification?: { predefinedMetricType: string, resourceLabel?: string }, scaleInCooldown?: number, scaleOutCooldown?: number, targetValue: number } | undefined>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: PolicyState = argsOrState as PolicyState | undefined;
            inputs["adjustmentType"] = state ? state.adjustmentType : undefined;
            inputs["alarms"] = state ? state.alarms : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["cooldown"] = state ? state.cooldown : undefined;
            inputs["metricAggregationType"] = state ? state.metricAggregationType : undefined;
            inputs["minAdjustmentMagnitude"] = state ? state.minAdjustmentMagnitude : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyType"] = state ? state.policyType : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["scalableDimension"] = state ? state.scalableDimension : undefined;
            inputs["serviceNamespace"] = state ? state.serviceNamespace : undefined;
            inputs["stepAdjustments"] = state ? state.stepAdjustments : undefined;
            inputs["stepScalingPolicyConfigurations"] = state ? state.stepScalingPolicyConfigurations : undefined;
            inputs["targetTrackingScalingPolicyConfiguration"] = state ? state.targetTrackingScalingPolicyConfiguration : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.scalableDimension === undefined) {
                throw new Error("Missing required property 'scalableDimension'");
            }
            if (!args || args.serviceNamespace === undefined) {
                throw new Error("Missing required property 'serviceNamespace'");
            }
            inputs["adjustmentType"] = args ? args.adjustmentType : undefined;
            inputs["alarms"] = args ? args.alarms : undefined;
            inputs["cooldown"] = args ? args.cooldown : undefined;
            inputs["metricAggregationType"] = args ? args.metricAggregationType : undefined;
            inputs["minAdjustmentMagnitude"] = args ? args.minAdjustmentMagnitude : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["scalableDimension"] = args ? args.scalableDimension : undefined;
            inputs["serviceNamespace"] = args ? args.serviceNamespace : undefined;
            inputs["stepAdjustments"] = args ? args.stepAdjustments : undefined;
            inputs["stepScalingPolicyConfigurations"] = args ? args.stepScalingPolicyConfigurations : undefined;
            inputs["targetTrackingScalingPolicyConfiguration"] = args ? args.targetTrackingScalingPolicyConfiguration : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:appautoscaling/policy:Policy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * The scaling policy's adjustment type.
     */
    readonly adjustmentType?: pulumi.Input<string>;
    readonly alarms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN assigned by AWS to the scaling policy.
     */
    readonly arn?: pulumi.Input<string>;
    readonly cooldown?: pulumi.Input<number>;
    readonly metricAggregationType?: pulumi.Input<string>;
    readonly minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * The name of the policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
     */
    readonly policyType?: pulumi.Input<string>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly scalableDimension?: pulumi.Input<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly serviceNamespace?: pulumi.Input<string>;
    readonly stepAdjustments?: pulumi.Input<pulumi.Input<{ metricIntervalLowerBound?: pulumi.Input<string>, metricIntervalUpperBound?: pulumi.Input<string>, scalingAdjustment: pulumi.Input<number> }>[]>;
    /**
     * Step scaling policy configuration, requires `policy_type = "StepScaling"` (default). See supported fields below.
     */
    readonly stepScalingPolicyConfigurations?: pulumi.Input<pulumi.Input<{ adjustmentType?: pulumi.Input<string>, cooldown?: pulumi.Input<number>, metricAggregationType?: pulumi.Input<string>, minAdjustmentMagnitude?: pulumi.Input<number>, stepAdjustments?: pulumi.Input<pulumi.Input<{ metricIntervalLowerBound?: pulumi.Input<string>, metricIntervalUpperBound?: pulumi.Input<string>, scalingAdjustment: pulumi.Input<number> }>[]> }>[]>;
    /**
     * A target tracking policy, requires `policy_type = "TargetTrackingScaling"`. See supported fields below.
     */
    readonly targetTrackingScalingPolicyConfiguration?: pulumi.Input<{ customizedMetricSpecification?: pulumi.Input<{ dimensions?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, metricName: pulumi.Input<string>, namespace: pulumi.Input<string>, statistic: pulumi.Input<string>, unit?: pulumi.Input<string> }>, disableScaleIn?: pulumi.Input<boolean>, predefinedMetricSpecification?: pulumi.Input<{ predefinedMetricType: pulumi.Input<string>, resourceLabel?: pulumi.Input<string> }>, scaleInCooldown?: pulumi.Input<number>, scaleOutCooldown?: pulumi.Input<number>, targetValue: pulumi.Input<number> }>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * The scaling policy's adjustment type.
     */
    readonly adjustmentType?: pulumi.Input<string>;
    readonly alarms?: pulumi.Input<pulumi.Input<string>[]>;
    readonly cooldown?: pulumi.Input<number>;
    readonly metricAggregationType?: pulumi.Input<string>;
    readonly minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * The name of the policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
     */
    readonly policyType?: pulumi.Input<string>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly scalableDimension: pulumi.Input<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly serviceNamespace: pulumi.Input<string>;
    readonly stepAdjustments?: pulumi.Input<pulumi.Input<{ metricIntervalLowerBound?: pulumi.Input<string>, metricIntervalUpperBound?: pulumi.Input<string>, scalingAdjustment: pulumi.Input<number> }>[]>;
    /**
     * Step scaling policy configuration, requires `policy_type = "StepScaling"` (default). See supported fields below.
     */
    readonly stepScalingPolicyConfigurations?: pulumi.Input<pulumi.Input<{ adjustmentType?: pulumi.Input<string>, cooldown?: pulumi.Input<number>, metricAggregationType?: pulumi.Input<string>, minAdjustmentMagnitude?: pulumi.Input<number>, stepAdjustments?: pulumi.Input<pulumi.Input<{ metricIntervalLowerBound?: pulumi.Input<string>, metricIntervalUpperBound?: pulumi.Input<string>, scalingAdjustment: pulumi.Input<number> }>[]> }>[]>;
    /**
     * A target tracking policy, requires `policy_type = "TargetTrackingScaling"`. See supported fields below.
     */
    readonly targetTrackingScalingPolicyConfiguration?: pulumi.Input<{ customizedMetricSpecification?: pulumi.Input<{ dimensions?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, metricName: pulumi.Input<string>, namespace: pulumi.Input<string>, statistic: pulumi.Input<string>, unit?: pulumi.Input<string> }>, disableScaleIn?: pulumi.Input<boolean>, predefinedMetricSpecification?: pulumi.Input<{ predefinedMetricType: pulumi.Input<string>, resourceLabel?: pulumi.Input<string> }>, scaleInCooldown?: pulumi.Input<number>, scaleOutCooldown?: pulumi.Input<number>, targetValue: pulumi.Input<number> }>;
}
