# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    address: pulumi.Output[str]
    """
    The hostname of the RDS instance. See also `endpoint` and `port`.
    """
    allocated_storage: pulumi.Output[int]
    """
    (Required unless a `snapshot_identifier` or
    `replicate_source_db` is provided) The allocated storage in gibibytes.
    """
    allow_major_version_upgrade: pulumi.Output[bool]
    """
    Indicates that major version
    upgrades are allowed. Changing this parameter does not result in an outage and
    the change is asynchronously applied as soon as possible.
    """
    apply_immediately: pulumi.Output[bool]
    """
    Specifies whether any database modifications
    are applied immediately, or during the next maintenance window. Default is
    `false`. See [Amazon RDS Documentation for more
    information.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)
    for more information.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the RDS instance.
    """
    auto_minor_version_upgrade: pulumi.Output[bool]
    """
    Indicates that minor engine upgrades
    will be applied automatically to the DB instance during the maintenance window.
    Defaults to true.
    """
    availability_zone: pulumi.Output[str]
    """
    The AZ for the RDS instance.
    """
    backup_retention_period: pulumi.Output[int]
    """
    The days to retain backups for. Must be
    between `0` and `35`. When creating a Read Replica the value must be greater than `0`. [See Read Replica][1].
    """
    backup_window: pulumi.Output[str]
    """
    The daily time range (in UTC) during which
    automated backups are created if they are enabled. Example: "09:46-10:16". Must
    not overlap with `maintenance_window`.
    """
    ca_cert_identifier: pulumi.Output[str]
    """
    Specifies the identifier of the CA certificate for the
    DB instance.
    """
    character_set_name: pulumi.Output[str]
    """
    The character set name to use for DB
    encoding in Oracle instances. This can't be changed. See [Oracle Character Sets
    Supported in Amazon
    RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html)
    for more information.
    """
    copy_tags_to_snapshot: pulumi.Output[bool]
    """
    Copy all Instance `tags` to snapshots. Default is `false`.
    """
    db_subnet_group_name: pulumi.Output[str]
    """
    Name of [DB subnet group](https://www.terraform.io/docs/providers/aws/r/db_subnet_group.html). DB instance will
    be created in the VPC associated with the DB subnet group. If unspecified, will
    be created in the `default` VPC, or in EC2 Classic, if available. When working
    with read replicas, it needs to be specified only if the source database
    specifies an instance in another AWS Region. See [DBSubnetGroupName in API
    action CreateDBInstanceReadReplica](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstanceReadReplica.html)
    for additional read replica contraints.
    """
    deletion_protection: pulumi.Output[bool]
    """
    If the DB instance should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
    """
    domain: pulumi.Output[str]
    """
    The ID of the Directory Service Active Directory domain to create the instance in.
    """
    domain_iam_role_name: pulumi.Output[str]
    """
    The name of the IAM role to be used when making API calls to the Directory Service.
    """
    enabled_cloudwatch_logs_exports: pulumi.Output[list]
    """
    List of log types to enable for exporting to CloudWatch logs. If omitted, no logs will be exported. Valid values (depending on `engine`): `alert`, `audit`, `error`, `general`, `listener`, `slowquery`, `trace`, `postgresql` (PostgreSQL), `upgrade` (PostgreSQL).
    """
    endpoint: pulumi.Output[str]
    """
    The connection endpoint in `address:port` format.
    """
    engine: pulumi.Output[str]
    """
    (Required unless a `snapshot_identifier` or `replicate_source_db`
    is provided) The database engine to use.  For supported values, see the Engine parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html).
    Note that for Amazon Aurora instances the engine must match the [DB cluster](https://www.terraform.io/docs/providers/aws/r/rds_cluster.html)'s engine'.
    For information on the difference between the available Aurora MySQL engines
    see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
    in the Amazon RDS User Guide.
    """
    engine_version: pulumi.Output[str]
    """
    The engine version to use. If `auto_minor_version_upgrade`
    is enabled, you can provide a prefix of the version such as `5.7` (for `5.7.10`) and
    this attribute will ignore differences in the patch version automatically (e.g. `5.7.17`).
    For supported values, see the EngineVersion parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html).
    Note that for Amazon Aurora instances the engine version must match the [DB cluster](https://www.terraform.io/docs/providers/aws/r/rds_cluster.html)'s engine version'.
    """
    final_snapshot_identifier: pulumi.Output[str]
    """
    The name of your final DB snapshot
    when this DB instance is deleted. If omitted, no final snapshot will be made.
    """
    hosted_zone_id: pulumi.Output[str]
    """
    The canonical hosted zone ID of the DB instance (to be used
    in a Route 53 Alias record).
    """
    iam_database_authentication_enabled: pulumi.Output[bool]
    """
    Specifies whether or
    mappings of AWS Identity and Access Management (IAM) accounts to database
    accounts is enabled.
    """
    identifier: pulumi.Output[str]
    """
    The name of the RDS instance,
    if omitted, Terraform will assign a random, unique identifier.
    """
    identifier_prefix: pulumi.Output[str]
    """
    Creates a unique
    identifier beginning with the specified prefix. Conflicts with `identifier`.
    """
    instance_class: pulumi.Output[str]
    """
    The instance type of the RDS instance.
    """
    iops: pulumi.Output[int]
    """
    The amount of provisioned IOPS. Setting this implies a
    storage_type of "io1".
    """
    kms_key_id: pulumi.Output[str]
    """
    The ARN for the KMS encryption key. If creating an
    encrypted replica, set this to the destination KMS ARN.
    """
    license_model: pulumi.Output[str]
    """
    (Optional, but required for some DB engines, i.e. Oracle
    SE1) License model information for this DB instance.
    """
    maintenance_window: pulumi.Output[str]
    """
    The window to perform maintenance in.
    Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00". See [RDS
    Maintenance Window
    docs](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow)
    for more information.
    """
    monitoring_interval: pulumi.Output[int]
    """
    The interval, in seconds, between points
    when Enhanced Monitoring metrics are collected for the DB instance. To disable
    collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid
    Values: 0, 1, 5, 10, 15, 30, 60.
    """
    monitoring_role_arn: pulumi.Output[str]
    """
    The ARN for the IAM role that permits RDS
    to send enhanced monitoring metrics to CloudWatch Logs. You can find more
    information on the [AWS
    Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
    what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
    """
    multi_az: pulumi.Output[bool]
    """
    Specifies if the RDS instance is multi-AZ
    """
    name: pulumi.Output[str]
    """
    The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance. Note that this does not apply for Oracle or SQL Server engines. See the [AWS documentation](http://docs.aws.amazon.com/cli/latest/reference/rds/create-db-instance.html) for more details on what applies for those engines.
    """
    option_group_name: pulumi.Output[str]
    """
    Name of the DB option group to associate.
    """
    parameter_group_name: pulumi.Output[str]
    """
    Name of the DB parameter group to
    associate.
    """
    password: pulumi.Output[str]
    """
    (Required unless a `snapshot_identifier` or `replicate_source_db`
    is provided) Password for the master DB user. Note that this may show up in
    logs, and it will be stored in the state file.
    """
    port: pulumi.Output[int]
    """
    The port on which the DB accepts connections.
    """
    publicly_accessible: pulumi.Output[bool]
    """
    Bool to control if instance is publicly
    accessible. Default is `false`.
    """
    replicas: pulumi.Output[list]
    replicate_source_db: pulumi.Output[str]
    """
    Specifies that this resource is a Replicate
    database, and to use this value as the source database. This correlates to the
    `identifier` of another Amazon RDS Database to replicate. Note that if you are
    creating a cross-region replica of an encrypted database you will also need to
    specify a `kms_key_id`. See [DB Instance Replication][1] and [Working with
    PostgreSQL and MySQL Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)
    for more information on using Replication.
    """
    resource_id: pulumi.Output[str]
    """
    The RDS Resource ID of this instance.
    """
    s3_import: pulumi.Output[dict]
    """
    Restore from a Percona Xtrabackup in S3.  See [Importing Data into an Amazon RDS MySQL DB Instance](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Importing.html)
    """
    security_group_names: pulumi.Output[list]
    """
    List of DB Security Groups to
    associate. Only used for [DB Instances on the _EC2-Classic_
    Platform](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html#USER_VPC.FindDefaultVPC).
    """
    skip_final_snapshot: pulumi.Output[bool]
    """
    Determines whether a final DB snapshot is
    created before the DB instance is deleted. If true is specified, no DBSnapshot
    is created. If false is specified, a DB snapshot is created before the DB
    instance is deleted, using the value from `final_snapshot_identifier`. Default
    is `false`.
    """
    snapshot_identifier: pulumi.Output[str]
    """
    Specifies whether or not to create this
    database from a snapshot. This correlates to the snapshot ID you'd find in the
    RDS console, e.g: rds:production-2015-06-26-06-05.
    """
    status: pulumi.Output[str]
    """
    The RDS instance status.
    """
    storage_encrypted: pulumi.Output[bool]
    """
    Specifies whether the DB instance is
    encrypted. Note that if you are creating a cross-region read replica this field
    is ignored and you should instead declare `kms_key_id` with a valid ARN. The
    default is `false` if not specified.
    """
    storage_type: pulumi.Output[str]
    """
    One of "standard" (magnetic), "gp2" (general
    purpose SSD), or "io1" (provisioned IOPS SSD). The default is "io1" if `iops` is
    specified, "standard" if not. Note that this behaviour is different from the AWS
    web console, where the default is "gp2".
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    timezone: pulumi.Output[str]
    """
    Time zone of the DB instance. `timezone` is currently
    only supported by Microsoft SQL Server. The `timezone` can only be set on
    creation. See [MSSQL User
    Guide](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone)
    for more information.
    """
    username: pulumi.Output[str]
    """
    (Required unless a `snapshot_identifier` or `replicate_source_db`
    is provided) Username for the master DB user.
    """
    vpc_security_group_ids: pulumi.Output[list]
    """
    List of VPC security groups to
    associate.
    """
    def __init__(__self__, __name__, __opts__=None, allocated_storage=None, allow_major_version_upgrade=None, apply_immediately=None, auto_minor_version_upgrade=None, availability_zone=None, backup_retention_period=None, backup_window=None, character_set_name=None, copy_tags_to_snapshot=None, db_subnet_group_name=None, deletion_protection=None, domain=None, domain_iam_role_name=None, enabled_cloudwatch_logs_exports=None, engine=None, engine_version=None, final_snapshot_identifier=None, iam_database_authentication_enabled=None, identifier=None, identifier_prefix=None, instance_class=None, iops=None, kms_key_id=None, license_model=None, maintenance_window=None, monitoring_interval=None, monitoring_role_arn=None, multi_az=None, name=None, option_group_name=None, parameter_group_name=None, password=None, port=None, publicly_accessible=None, replicate_source_db=None, s3_import=None, security_group_names=None, skip_final_snapshot=None, snapshot_identifier=None, storage_encrypted=None, storage_type=None, tags=None, timezone=None, username=None, vpc_security_group_ids=None):
        """
        Provides an RDS instance resource.  A DB instance is an isolated database
        environment in the cloud.  A DB instance can contain multiple user-created
        databases.
        
        Changes to a DB instance can occur when you manually change a parameter, such as
        `allocated_storage`, and are reflected in the next maintenance window. Because
        of this, Terraform may report a difference in its planning phase because a
        modification has not yet taken place. You can use the `apply_immediately` flag
        to instruct the service to apply the change immediately (see documentation
        below).
        
        When upgrading the major version of an engine, `allow_major_version_upgrade`
        must be set to `true`.
        
        > **Note:** using `apply_immediately` can result in a brief downtime as the
        server reboots. See the AWS Docs on [RDS Maintenance][2] for more information.
        
        > **Note:** All arguments including the username and password will be stored in
        the raw state as plain-text. [Read more about sensitive data in
        state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        ## RDS Instance Class Types
        
        Amazon RDS supports three types of instance classes: Standard, Memory Optimized,
        and Burstable Performance. For more information please read the AWS RDS documentation
        about [DB Instance Class Types](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[int] allocated_storage: (Required unless a `snapshot_identifier` or
               `replicate_source_db` is provided) The allocated storage in gibibytes.
        :param pulumi.Input[bool] allow_major_version_upgrade: Indicates that major version
               upgrades are allowed. Changing this parameter does not result in an outage and
               the change is asynchronously applied as soon as possible.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any database modifications
               are applied immediately, or during the next maintenance window. Default is
               `false`. See [Amazon RDS Documentation for more
               information.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)
               for more information.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Indicates that minor engine upgrades
               will be applied automatically to the DB instance during the maintenance window.
               Defaults to true.
        :param pulumi.Input[str] availability_zone: The AZ for the RDS instance.
        :param pulumi.Input[int] backup_retention_period: The days to retain backups for. Must be
               between `0` and `35`. When creating a Read Replica the value must be greater than `0`. [See Read Replica][1].
        :param pulumi.Input[str] backup_window: The daily time range (in UTC) during which
               automated backups are created if they are enabled. Example: "09:46-10:16". Must
               not overlap with `maintenance_window`.
        :param pulumi.Input[str] character_set_name: The character set name to use for DB
               encoding in Oracle instances. This can't be changed. See [Oracle Character Sets
               Supported in Amazon
               RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html)
               for more information.
        :param pulumi.Input[bool] copy_tags_to_snapshot: Copy all Instance `tags` to snapshots. Default is `false`.
        :param pulumi.Input[str] db_subnet_group_name: Name of [DB subnet group](https://www.terraform.io/docs/providers/aws/r/db_subnet_group.html). DB instance will
               be created in the VPC associated with the DB subnet group. If unspecified, will
               be created in the `default` VPC, or in EC2 Classic, if available. When working
               with read replicas, it needs to be specified only if the source database
               specifies an instance in another AWS Region. See [DBSubnetGroupName in API
               action CreateDBInstanceReadReplica](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstanceReadReplica.html)
               for additional read replica contraints.
        :param pulumi.Input[bool] deletion_protection: If the DB instance should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        :param pulumi.Input[str] domain: The ID of the Directory Service Active Directory domain to create the instance in.
        :param pulumi.Input[str] domain_iam_role_name: The name of the IAM role to be used when making API calls to the Directory Service.
        :param pulumi.Input[list] enabled_cloudwatch_logs_exports: List of log types to enable for exporting to CloudWatch logs. If omitted, no logs will be exported. Valid values (depending on `engine`): `alert`, `audit`, `error`, `general`, `listener`, `slowquery`, `trace`, `postgresql` (PostgreSQL), `upgrade` (PostgreSQL).
        :param pulumi.Input[str] engine: (Required unless a `snapshot_identifier` or `replicate_source_db`
               is provided) The database engine to use.  For supported values, see the Engine parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html).
               Note that for Amazon Aurora instances the engine must match the [DB cluster](https://www.terraform.io/docs/providers/aws/r/rds_cluster.html)'s engine'.
               For information on the difference between the available Aurora MySQL engines
               see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
               in the Amazon RDS User Guide.
        :param pulumi.Input[str] engine_version: The engine version to use. If `auto_minor_version_upgrade`
               is enabled, you can provide a prefix of the version such as `5.7` (for `5.7.10`) and
               this attribute will ignore differences in the patch version automatically (e.g. `5.7.17`).
               For supported values, see the EngineVersion parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html).
               Note that for Amazon Aurora instances the engine version must match the [DB cluster](https://www.terraform.io/docs/providers/aws/r/rds_cluster.html)'s engine version'.
        :param pulumi.Input[str] final_snapshot_identifier: The name of your final DB snapshot
               when this DB instance is deleted. If omitted, no final snapshot will be made.
        :param pulumi.Input[bool] iam_database_authentication_enabled: Specifies whether or
               mappings of AWS Identity and Access Management (IAM) accounts to database
               accounts is enabled.
        :param pulumi.Input[str] identifier: The name of the RDS instance,
               if omitted, Terraform will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique
               identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[str] instance_class: The instance type of the RDS instance.
        :param pulumi.Input[int] iops: The amount of provisioned IOPS. Setting this implies a
               storage_type of "io1".
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key. If creating an
               encrypted replica, set this to the destination KMS ARN.
        :param pulumi.Input[str] license_model: (Optional, but required for some DB engines, i.e. Oracle
               SE1) License model information for this DB instance.
        :param pulumi.Input[str] maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00". See [RDS
               Maintenance Window
               docs](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow)
               for more information.
        :param pulumi.Input[int] monitoring_interval: The interval, in seconds, between points
               when Enhanced Monitoring metrics are collected for the DB instance. To disable
               collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid
               Values: 0, 1, 5, 10, 15, 30, 60.
        :param pulumi.Input[str] monitoring_role_arn: The ARN for the IAM role that permits RDS
               to send enhanced monitoring metrics to CloudWatch Logs. You can find more
               information on the [AWS
               Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
               what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
        :param pulumi.Input[bool] multi_az: Specifies if the RDS instance is multi-AZ
        :param pulumi.Input[str] name: The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance. Note that this does not apply for Oracle or SQL Server engines. See the [AWS documentation](http://docs.aws.amazon.com/cli/latest/reference/rds/create-db-instance.html) for more details on what applies for those engines.
        :param pulumi.Input[str] option_group_name: Name of the DB option group to associate.
        :param pulumi.Input[str] parameter_group_name: Name of the DB parameter group to
               associate.
        :param pulumi.Input[str] password: (Required unless a `snapshot_identifier` or `replicate_source_db`
               is provided) Password for the master DB user. Note that this may show up in
               logs, and it will be stored in the state file.
        :param pulumi.Input[int] port: The port on which the DB accepts connections.
        :param pulumi.Input[bool] publicly_accessible: Bool to control if instance is publicly
               accessible. Default is `false`.
        :param pulumi.Input[str] replicate_source_db: Specifies that this resource is a Replicate
               database, and to use this value as the source database. This correlates to the
               `identifier` of another Amazon RDS Database to replicate. Note that if you are
               creating a cross-region replica of an encrypted database you will also need to
               specify a `kms_key_id`. See [DB Instance Replication][1] and [Working with
               PostgreSQL and MySQL Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)
               for more information on using Replication.
        :param pulumi.Input[dict] s3_import: Restore from a Percona Xtrabackup in S3.  See [Importing Data into an Amazon RDS MySQL DB Instance](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Importing.html)
        :param pulumi.Input[list] security_group_names: List of DB Security Groups to
               associate. Only used for [DB Instances on the _EC2-Classic_
               Platform](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html#USER_VPC.FindDefaultVPC).
        :param pulumi.Input[bool] skip_final_snapshot: Determines whether a final DB snapshot is
               created before the DB instance is deleted. If true is specified, no DBSnapshot
               is created. If false is specified, a DB snapshot is created before the DB
               instance is deleted, using the value from `final_snapshot_identifier`. Default
               is `false`.
        :param pulumi.Input[str] snapshot_identifier: Specifies whether or not to create this
               database from a snapshot. This correlates to the snapshot ID you'd find in the
               RDS console, e.g: rds:production-2015-06-26-06-05.
        :param pulumi.Input[bool] storage_encrypted: Specifies whether the DB instance is
               encrypted. Note that if you are creating a cross-region read replica this field
               is ignored and you should instead declare `kms_key_id` with a valid ARN. The
               default is `false` if not specified.
        :param pulumi.Input[str] storage_type: One of "standard" (magnetic), "gp2" (general
               purpose SSD), or "io1" (provisioned IOPS SSD). The default is "io1" if `iops` is
               specified, "standard" if not. Note that this behaviour is different from the AWS
               web console, where the default is "gp2".
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] timezone: Time zone of the DB instance. `timezone` is currently
               only supported by Microsoft SQL Server. The `timezone` can only be set on
               creation. See [MSSQL User
               Guide](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone)
               for more information.
        :param pulumi.Input[str] username: (Required unless a `snapshot_identifier` or `replicate_source_db`
               is provided) Username for the master DB user.
        :param pulumi.Input[list] vpc_security_group_ids: List of VPC security groups to
               associate.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allocated_storage'] = allocated_storage

        __props__['allow_major_version_upgrade'] = allow_major_version_upgrade

        __props__['apply_immediately'] = apply_immediately

        __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade

        __props__['availability_zone'] = availability_zone

        __props__['backup_retention_period'] = backup_retention_period

        __props__['backup_window'] = backup_window

        __props__['character_set_name'] = character_set_name

        __props__['copy_tags_to_snapshot'] = copy_tags_to_snapshot

        __props__['db_subnet_group_name'] = db_subnet_group_name

        __props__['deletion_protection'] = deletion_protection

        __props__['domain'] = domain

        __props__['domain_iam_role_name'] = domain_iam_role_name

        __props__['enabled_cloudwatch_logs_exports'] = enabled_cloudwatch_logs_exports

        __props__['engine'] = engine

        __props__['engine_version'] = engine_version

        __props__['final_snapshot_identifier'] = final_snapshot_identifier

        __props__['iam_database_authentication_enabled'] = iam_database_authentication_enabled

        __props__['identifier'] = identifier

        __props__['identifier_prefix'] = identifier_prefix

        if not instance_class:
            raise TypeError('Missing required property instance_class')
        __props__['instance_class'] = instance_class

        __props__['iops'] = iops

        __props__['kms_key_id'] = kms_key_id

        __props__['license_model'] = license_model

        __props__['maintenance_window'] = maintenance_window

        __props__['monitoring_interval'] = monitoring_interval

        __props__['monitoring_role_arn'] = monitoring_role_arn

        __props__['multi_az'] = multi_az

        __props__['name'] = name

        __props__['option_group_name'] = option_group_name

        __props__['parameter_group_name'] = parameter_group_name

        __props__['password'] = password

        __props__['port'] = port

        __props__['publicly_accessible'] = publicly_accessible

        __props__['replicate_source_db'] = replicate_source_db

        __props__['s3_import'] = s3_import

        __props__['security_group_names'] = security_group_names

        __props__['skip_final_snapshot'] = skip_final_snapshot

        __props__['snapshot_identifier'] = snapshot_identifier

        __props__['storage_encrypted'] = storage_encrypted

        __props__['storage_type'] = storage_type

        __props__['tags'] = tags

        __props__['timezone'] = timezone

        __props__['username'] = username

        __props__['vpc_security_group_ids'] = vpc_security_group_ids

        __props__['address'] = None
        __props__['arn'] = None
        __props__['ca_cert_identifier'] = None
        __props__['endpoint'] = None
        __props__['hosted_zone_id'] = None
        __props__['replicas'] = None
        __props__['resource_id'] = None
        __props__['status'] = None

        super(Instance, __self__).__init__(
            'aws:rds/instance:Instance',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

