// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides information about an Elastic File System Mount Target (EFS).
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/efs_mount_target.html.markdown.
 */
export function getMountTarget(args: GetMountTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetMountTargetResult> & GetMountTargetResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetMountTargetResult> = pulumi.runtime.invoke("aws:efs/getMountTarget:getMountTarget", {
        "mountTargetId": args.mountTargetId,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getMountTarget.
 */
export interface GetMountTargetArgs {
    /**
     * ID of the mount target that you want to have described
     */
    readonly mountTargetId: string;
}

/**
 * A collection of values returned by getMountTarget.
 */
export interface GetMountTargetResult {
    /**
     * The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
     */
    readonly dnsName: string;
    /**
     * Amazon Resource Name of the file system for which the mount target is intended.
     */
    readonly fileSystemArn: string;
    /**
     * ID of the file system for which the mount target is intended.
     */
    readonly fileSystemId: string;
    /**
     * Address at which the file system may be mounted via the mount target.
     */
    readonly ipAddress: string;
    readonly mountTargetId: string;
    /**
     * The ID of the network interface that Amazon EFS created when it created the mount target.
     */
    readonly networkInterfaceId: string;
    /**
     * List of VPC security group IDs attached to the mount target.
     */
    readonly securityGroups: string[];
    /**
     * ID of the mount target's subnet.
     */
    readonly subnetId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
