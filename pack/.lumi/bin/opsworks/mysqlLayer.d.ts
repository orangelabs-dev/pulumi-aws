import * as lumi from "@lumi/lumi";
export declare class MysqlLayer extends lumi.NamedResource implements MysqlLayerArgs {
    readonly autoAssignElasticIps?: boolean;
    readonly autoAssignPublicIps?: boolean;
    readonly autoHealing?: boolean;
    readonly customConfigureRecipes?: string[];
    readonly customDeployRecipes?: string[];
    readonly customInstanceProfileArn?: string;
    readonly customJson?: string;
    readonly customSecurityGroupIds?: string[];
    readonly customSetupRecipes?: string[];
    readonly customShutdownRecipes?: string[];
    readonly customUndeployRecipes?: string[];
    readonly drainElbOnShutdown?: boolean;
    readonly ebsVolume?: {
        iops?: number;
        mountPoint: string;
        numberOfDisks: number;
        raidLevel?: string;
        size: number;
        type?: string;
    }[];
    readonly elasticLoadBalancer?: string;
    readonly layerId: string;
    readonly installUpdatesOnBoot?: boolean;
    readonly instanceShutdownTimeout?: number;
    readonly mysqlLayerName?: string;
    readonly rootPassword?: string;
    readonly rootPasswordOnAllInstances?: boolean;
    readonly stackId: string;
    readonly systemPackages?: string[];
    readonly useEbsOptimizedInstances?: boolean;
    constructor(name: string, args: MysqlLayerArgs);
}
export interface MysqlLayerArgs {
    readonly autoAssignElasticIps?: boolean;
    readonly autoAssignPublicIps?: boolean;
    readonly autoHealing?: boolean;
    readonly customConfigureRecipes?: string[];
    readonly customDeployRecipes?: string[];
    readonly customInstanceProfileArn?: string;
    readonly customJson?: string;
    readonly customSecurityGroupIds?: string[];
    readonly customSetupRecipes?: string[];
    readonly customShutdownRecipes?: string[];
    readonly customUndeployRecipes?: string[];
    readonly drainElbOnShutdown?: boolean;
    readonly ebsVolume?: {
        iops?: number;
        mountPoint: string;
        numberOfDisks: number;
        raidLevel?: string;
        size: number;
        type?: string;
    }[];
    readonly elasticLoadBalancer?: string;
    readonly installUpdatesOnBoot?: boolean;
    readonly instanceShutdownTimeout?: number;
    readonly mysqlLayerName?: string;
    readonly rootPassword?: string;
    readonly rootPasswordOnAllInstances?: boolean;
    readonly stackId: string;
    readonly systemPackages?: string[];
    readonly useEbsOptimizedInstances?: boolean;
}
