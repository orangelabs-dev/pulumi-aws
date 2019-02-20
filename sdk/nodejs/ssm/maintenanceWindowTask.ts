// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SSM Maintenance Window Task resource
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const instance = new aws.ec2.Instance("instance", {
 *     ami: "ami-4fccb37f",
 *     instanceType: "m1.small",
 * });
 * const window = new aws.ssm.MaintenanceWindow("window", {
 *     cutoff: 1,
 *     duration: 3,
 *     name: "maintenance-window-%s",
 *     schedule: "cron(0 16 ? * TUE *)",
 * });
 * const task = new aws.ssm.MaintenanceWindowTask("task", {
 *     description: "This is a maintenance window task",
 *     maxConcurrency: "2",
 *     maxErrors: "1",
 *     name: "maintenance-window-task",
 *     priority: 1,
 *     serviceRoleArn: "arn:aws:iam::187416307283:role/service-role/AWS_Events_Invoke_Run_Command_112316643",
 *     targets: [{
 *         key: "InstanceIds",
 *         values: [instance.id],
 *     }],
 *     taskArn: "AWS-RunShellScript",
 *     taskParameters: [{
 *         name: "commands",
 *         values: ["pwd"],
 *     }],
 *     taskType: "RUN_COMMAND",
 *     windowId: window.id,
 * });
 * ```
 */
export class MaintenanceWindowTask extends pulumi.CustomResource {
    /**
     * Get an existing MaintenanceWindowTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MaintenanceWindowTaskState, opts?: pulumi.CustomResourceOptions): MaintenanceWindowTask {
        return new MaintenanceWindowTask(name, <any>state, { ...opts, id: id });
    }

    /**
     * The description of the maintenance window task.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * A structure containing information about an Amazon S3 bucket to write instance-level logs to. Documented below.
     */
    public readonly loggingInfo: pulumi.Output<{ s3BucketName: string, s3BucketPrefix?: string, s3Region: string } | undefined>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    public readonly maxConcurrency: pulumi.Output<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    public readonly maxErrors: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    public readonly priority: pulumi.Output<number | undefined>;
    /**
     * The role that should be assumed when executing the task.
     */
    public readonly serviceRoleArn: pulumi.Output<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    public readonly targets: pulumi.Output<{ key: string, values: string[] }[]>;
    /**
     * The ARN of the task to execute.
     */
    public readonly taskArn: pulumi.Output<string>;
    /**
     * A structure containing information about parameters required by the particular `task_arn`. Documented below.
     */
    public readonly taskParameters: pulumi.Output<{ name: string, values: string[] }[] | undefined>;
    /**
     * The type of task being registered. The only allowed value is `RUN_COMMAND`.
     */
    public readonly taskType: pulumi.Output<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    public readonly windowId: pulumi.Output<string>;

    /**
     * Create a MaintenanceWindowTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MaintenanceWindowTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MaintenanceWindowTaskArgs | MaintenanceWindowTaskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: MaintenanceWindowTaskState = argsOrState as MaintenanceWindowTaskState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["loggingInfo"] = state ? state.loggingInfo : undefined;
            inputs["maxConcurrency"] = state ? state.maxConcurrency : undefined;
            inputs["maxErrors"] = state ? state.maxErrors : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["serviceRoleArn"] = state ? state.serviceRoleArn : undefined;
            inputs["targets"] = state ? state.targets : undefined;
            inputs["taskArn"] = state ? state.taskArn : undefined;
            inputs["taskParameters"] = state ? state.taskParameters : undefined;
            inputs["taskType"] = state ? state.taskType : undefined;
            inputs["windowId"] = state ? state.windowId : undefined;
        } else {
            const args = argsOrState as MaintenanceWindowTaskArgs | undefined;
            if (!args || args.maxConcurrency === undefined) {
                throw new Error("Missing required property 'maxConcurrency'");
            }
            if (!args || args.maxErrors === undefined) {
                throw new Error("Missing required property 'maxErrors'");
            }
            if (!args || args.serviceRoleArn === undefined) {
                throw new Error("Missing required property 'serviceRoleArn'");
            }
            if (!args || args.targets === undefined) {
                throw new Error("Missing required property 'targets'");
            }
            if (!args || args.taskArn === undefined) {
                throw new Error("Missing required property 'taskArn'");
            }
            if (!args || args.taskType === undefined) {
                throw new Error("Missing required property 'taskType'");
            }
            if (!args || args.windowId === undefined) {
                throw new Error("Missing required property 'windowId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["loggingInfo"] = args ? args.loggingInfo : undefined;
            inputs["maxConcurrency"] = args ? args.maxConcurrency : undefined;
            inputs["maxErrors"] = args ? args.maxErrors : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["taskArn"] = args ? args.taskArn : undefined;
            inputs["taskParameters"] = args ? args.taskParameters : undefined;
            inputs["taskType"] = args ? args.taskType : undefined;
            inputs["windowId"] = args ? args.windowId : undefined;
        }
        super("aws:ssm/maintenanceWindowTask:MaintenanceWindowTask", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MaintenanceWindowTask resources.
 */
export interface MaintenanceWindowTaskState {
    /**
     * The description of the maintenance window task.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A structure containing information about an Amazon S3 bucket to write instance-level logs to. Documented below.
     */
    readonly loggingInfo?: pulumi.Input<{ s3BucketName: pulumi.Input<string>, s3BucketPrefix?: pulumi.Input<string>, s3Region: pulumi.Input<string> }>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    readonly maxConcurrency?: pulumi.Input<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    readonly maxErrors?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The role that should be assumed when executing the task.
     */
    readonly serviceRoleArn?: pulumi.Input<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    readonly targets?: pulumi.Input<pulumi.Input<{ key: pulumi.Input<string>, values: pulumi.Input<pulumi.Input<string>[]> }>[]>;
    /**
     * The ARN of the task to execute.
     */
    readonly taskArn?: pulumi.Input<string>;
    /**
     * A structure containing information about parameters required by the particular `task_arn`. Documented below.
     */
    readonly taskParameters?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, values: pulumi.Input<pulumi.Input<string>[]> }>[]>;
    /**
     * The type of task being registered. The only allowed value is `RUN_COMMAND`.
     */
    readonly taskType?: pulumi.Input<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    readonly windowId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MaintenanceWindowTask resource.
 */
export interface MaintenanceWindowTaskArgs {
    /**
     * The description of the maintenance window task.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A structure containing information about an Amazon S3 bucket to write instance-level logs to. Documented below.
     */
    readonly loggingInfo?: pulumi.Input<{ s3BucketName: pulumi.Input<string>, s3BucketPrefix?: pulumi.Input<string>, s3Region: pulumi.Input<string> }>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    readonly maxConcurrency: pulumi.Input<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    readonly maxErrors: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The role that should be assumed when executing the task.
     */
    readonly serviceRoleArn: pulumi.Input<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    readonly targets: pulumi.Input<pulumi.Input<{ key: pulumi.Input<string>, values: pulumi.Input<pulumi.Input<string>[]> }>[]>;
    /**
     * The ARN of the task to execute.
     */
    readonly taskArn: pulumi.Input<string>;
    /**
     * A structure containing information about parameters required by the particular `task_arn`. Documented below.
     */
    readonly taskParameters?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, values: pulumi.Input<pulumi.Input<string>[]> }>[]>;
    /**
     * The type of task being registered. The only allowed value is `RUN_COMMAND`.
     */
    readonly taskType: pulumi.Input<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    readonly windowId: pulumi.Input<string>;
}
