// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an EKS Cluster.
 * 
 * ## Example Usage
 * 
 * ### Example IAM Role for EKS Cluster
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.iam.Role("example", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *       "Principal": {
 *         "Service": "eks.amazonaws.com"
 *       },
 *       "Action": "sts:AssumeRole"
 *     }
 *   ]
 * }
 * `,
 * });
 * const example_AmazonEKSClusterPolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy", {
 *     policyArn: "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
 *     role: example.name,
 * });
 * const example_AmazonEKSServicePolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSServicePolicy", {
 *     policyArn: "arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
 *     role: example.name,
 * });
 * ```
 * 
 * ### Enabling Control Plane Logging
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const clusterName = config.get("clusterName") || "example";
 * 
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("example", {
 *     retentionInDays: 7,
 * });
 * const exampleCluster = new aws.eks.Cluster("example", {
 *     enabledClusterLogTypes: [
 *         "api",
 *         "audit",
 *     ],
 * }, {dependsOn: [exampleLogGroup]});
 * ```
 * 
 * ### Enabling IAM Roles for Service Accounts
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleCluster = new aws.eks.Cluster("example", {});
 * const exampleOpenIdConnectProvider = new aws.iam.OpenIdConnectProvider("example", {
 *     clientIdLists: ["sts.amazonaws.com"],
 *     thumbprintLists: [],
 *     url: exampleCluster.identities[0].oidcs[0].issuer,
 * });
 * const current = aws.getCallerIdentity();
 * const exampleAssumeRolePolicy = pulumi.all([exampleOpenIdConnectProvider.url, exampleOpenIdConnectProvider.arn]).apply(([url, arn]) => aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRoleWithWebIdentity"],
 *         conditions: [{
 *             test: "StringEquals",
 *             values: ["system:serviceaccount:kube-system:aws-node"],
 *             variable: `${url.replace("https://", "")}:sub`,
 *         }],
 *         effect: "Allow",
 *         principals: [{
 *             identifiers: [arn],
 *             type: "Federated",
 *         }],
 *     }],
 * }));
 * const exampleRole = new aws.iam.Role("example", {
 *     assumeRolePolicy: exampleAssumeRolePolicy.json,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/eks_cluster.html.markdown.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:eks/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Nested attribute containing `certificate-authority-data` for your cluster.
     */
    public /*out*/ readonly certificateAuthority!: pulumi.Output<outputs.eks.ClusterCertificateAuthority>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
     */
    public readonly enabledClusterLogTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     */
    public readonly encryptionConfig!: pulumi.Output<outputs.eks.ClusterEncryptionConfig | undefined>;
    /**
     * The endpoint for your Kubernetes API server.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
     */
    public /*out*/ readonly identities!: pulumi.Output<outputs.eks.ClusterIdentity[]>;
    /**
     * Name of the cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The platform version for the cluster.
     */
    public /*out*/ readonly platformVersion!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the [`aws.iam.RolePolicy` resource](https://www.terraform.io/docs/providers/aws/r/iam_role_policy.html) or [`aws.iam.RolePolicyAttachment` resource](https://www.terraform.io/docs/providers/aws/r/iam_role_policy_attachment.html), otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`. 
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Key-value mapping of resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
     */
    public readonly vpcConfig!: pulumi.Output<outputs.eks.ClusterVpcConfig>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificateAuthority"] = state ? state.certificateAuthority : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["enabledClusterLogTypes"] = state ? state.enabledClusterLogTypes : undefined;
            inputs["encryptionConfig"] = state ? state.encryptionConfig : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["identities"] = state ? state.identities : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platformVersion"] = state ? state.platformVersion : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            if (!args || args.vpcConfig === undefined) {
                throw new Error("Missing required property 'vpcConfig'");
            }
            inputs["enabledClusterLogTypes"] = args ? args.enabledClusterLogTypes : undefined;
            inputs["encryptionConfig"] = args ? args.encryptionConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificateAuthority"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["identities"] = undefined /*out*/;
            inputs["platformVersion"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Nested attribute containing `certificate-authority-data` for your cluster.
     */
    readonly certificateAuthority?: pulumi.Input<inputs.eks.ClusterCertificateAuthority>;
    readonly createdAt?: pulumi.Input<string>;
    /**
     * A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
     */
    readonly enabledClusterLogTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     */
    readonly encryptionConfig?: pulumi.Input<inputs.eks.ClusterEncryptionConfig>;
    /**
     * The endpoint for your Kubernetes API server.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
     */
    readonly identities?: pulumi.Input<pulumi.Input<inputs.eks.ClusterIdentity>[]>;
    /**
     * Name of the cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The platform version for the cluster.
     */
    readonly platformVersion?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the [`aws.iam.RolePolicy` resource](https://www.terraform.io/docs/providers/aws/r/iam_role_policy.html) or [`aws.iam.RolePolicyAttachment` resource](https://www.terraform.io/docs/providers/aws/r/iam_role_policy_attachment.html), otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`. 
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
     */
    readonly vpcConfig?: pulumi.Input<inputs.eks.ClusterVpcConfig>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
     */
    readonly enabledClusterLogTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     */
    readonly encryptionConfig?: pulumi.Input<inputs.eks.ClusterEncryptionConfig>;
    /**
     * Name of the cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the [`aws.iam.RolePolicy` resource](https://www.terraform.io/docs/providers/aws/r/iam_role_policy.html) or [`aws.iam.RolePolicyAttachment` resource](https://www.terraform.io/docs/providers/aws/r/iam_role_policy_attachment.html), otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
     */
    readonly vpcConfig: pulumi.Input<inputs.eks.ClusterVpcConfig>;
}
