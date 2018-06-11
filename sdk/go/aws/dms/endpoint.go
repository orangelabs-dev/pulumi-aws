// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.
// 
// ~> **Note:** All arguments including the password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](/docs/state/sensitive-data.html).
type Endpoint struct {
	s *pulumi.ResourceState
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOpt) (*Endpoint, error) {
	if args == nil || args.EndpointId == nil {
		return nil, errors.New("missing required argument 'EndpointId'")
	}
	if args == nil || args.EndpointType == nil {
		return nil, errors.New("missing required argument 'EndpointType'")
	}
	if args == nil || args.EngineName == nil {
		return nil, errors.New("missing required argument 'EngineName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["certificateArn"] = nil
		inputs["databaseName"] = nil
		inputs["endpointId"] = nil
		inputs["endpointType"] = nil
		inputs["engineName"] = nil
		inputs["extraConnectionAttributes"] = nil
		inputs["kmsKeyArn"] = nil
		inputs["mongodbSettings"] = nil
		inputs["password"] = nil
		inputs["port"] = nil
		inputs["s3Settings"] = nil
		inputs["serverName"] = nil
		inputs["serviceAccessRole"] = nil
		inputs["sslMode"] = nil
		inputs["tags"] = nil
		inputs["username"] = nil
	} else {
		inputs["certificateArn"] = args.CertificateArn
		inputs["databaseName"] = args.DatabaseName
		inputs["endpointId"] = args.EndpointId
		inputs["endpointType"] = args.EndpointType
		inputs["engineName"] = args.EngineName
		inputs["extraConnectionAttributes"] = args.ExtraConnectionAttributes
		inputs["kmsKeyArn"] = args.KmsKeyArn
		inputs["mongodbSettings"] = args.MongodbSettings
		inputs["password"] = args.Password
		inputs["port"] = args.Port
		inputs["s3Settings"] = args.S3Settings
		inputs["serverName"] = args.ServerName
		inputs["serviceAccessRole"] = args.ServiceAccessRole
		inputs["sslMode"] = args.SslMode
		inputs["tags"] = args.Tags
		inputs["username"] = args.Username
	}
	inputs["endpointArn"] = nil
	s, err := ctx.RegisterResource("aws:dms/endpoint:Endpoint", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Endpoint{s: s}, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EndpointState, opts ...pulumi.ResourceOpt) (*Endpoint, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["certificateArn"] = state.CertificateArn
		inputs["databaseName"] = state.DatabaseName
		inputs["endpointArn"] = state.EndpointArn
		inputs["endpointId"] = state.EndpointId
		inputs["endpointType"] = state.EndpointType
		inputs["engineName"] = state.EngineName
		inputs["extraConnectionAttributes"] = state.ExtraConnectionAttributes
		inputs["kmsKeyArn"] = state.KmsKeyArn
		inputs["mongodbSettings"] = state.MongodbSettings
		inputs["password"] = state.Password
		inputs["port"] = state.Port
		inputs["s3Settings"] = state.S3Settings
		inputs["serverName"] = state.ServerName
		inputs["serviceAccessRole"] = state.ServiceAccessRole
		inputs["sslMode"] = state.SslMode
		inputs["tags"] = state.Tags
		inputs["username"] = state.Username
	}
	s, err := ctx.ReadResource("aws:dms/endpoint:Endpoint", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Endpoint{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Endpoint) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Endpoint) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The Amazon Resource Name (ARN) for the certificate.
func (r *Endpoint) CertificateArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["certificateArn"])
}

// The name of the endpoint database.
func (r *Endpoint) DatabaseName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["databaseName"])
}

// The Amazon Resource Name (ARN) for the endpoint.
func (r *Endpoint) EndpointArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpointArn"])
}

// The database endpoint identifier.
func (r *Endpoint) EndpointId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpointId"])
}

// The type of endpoint. Can be one of `source | target`.
func (r *Endpoint) EndpointType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpointType"])
}

// The type of engine for the endpoint. Can be one of `mysql | oracle | postgres | mariadb | aurora | redshift | sybase | sqlserver | dynamodb | mongodb | s3 | azuredb`.
func (r *Endpoint) EngineName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["engineName"])
}

// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
func (r *Endpoint) ExtraConnectionAttributes() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["extraConnectionAttributes"])
}

// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
func (r *Endpoint) KmsKeyArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["kmsKeyArn"])
}

// Settings for the source MongoDB endpoint. Available settings are `auth_type` (default: `PASSWORD`), `auth_mechanism` (default: `DEFAULT`), `nesting_level` (default: `NONE`), `extract_doc_id` (default: `false`), `docs_to_investigate` (default: `1000`) and `auth_source` (default: `admin`). For more details, see [Using MongoDB as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html).
func (r *Endpoint) MongodbSettings() *pulumi.Output {
	return r.s.State["mongodbSettings"]
}

// The password to be used to login to the endpoint database.
func (r *Endpoint) Password() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["password"])
}

// The port used by the endpoint database.
func (r *Endpoint) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// Settings for the target S3 endpoint. Available settings are `service_access_role_arn`, `external_table_definition`, `csv_row_delimiter` (default: `\\n`), `csv_delimiter` (default: `,`), `bucket_folder`, `bucket_name` and `compression_type` (default: `NONE`). For more details, see [Using Amazon S3 as a Target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html).
func (r *Endpoint) S3Settings() *pulumi.Output {
	return r.s.State["s3Settings"]
}

// The host name of the server.
func (r *Endpoint) ServerName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serverName"])
}

// The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
func (r *Endpoint) ServiceAccessRole() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceAccessRole"])
}

// The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
func (r *Endpoint) SslMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sslMode"])
}

// A mapping of tags to assign to the resource.
func (r *Endpoint) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The user name to be used to login to the endpoint database.
func (r *Endpoint) Username() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["username"])
}

// Input properties used for looking up and filtering Endpoint resources.
type EndpointState struct {
	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn interface{}
	// The name of the endpoint database.
	DatabaseName interface{}
	// The Amazon Resource Name (ARN) for the endpoint.
	EndpointArn interface{}
	// The database endpoint identifier.
	EndpointId interface{}
	// The type of endpoint. Can be one of `source | target`.
	EndpointType interface{}
	// The type of engine for the endpoint. Can be one of `mysql | oracle | postgres | mariadb | aurora | redshift | sybase | sqlserver | dynamodb | mongodb | s3 | azuredb`.
	EngineName interface{}
	// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
	ExtraConnectionAttributes interface{}
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn interface{}
	// Settings for the source MongoDB endpoint. Available settings are `auth_type` (default: `PASSWORD`), `auth_mechanism` (default: `DEFAULT`), `nesting_level` (default: `NONE`), `extract_doc_id` (default: `false`), `docs_to_investigate` (default: `1000`) and `auth_source` (default: `admin`). For more details, see [Using MongoDB as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html).
	MongodbSettings interface{}
	// The password to be used to login to the endpoint database.
	Password interface{}
	// The port used by the endpoint database.
	Port interface{}
	// Settings for the target S3 endpoint. Available settings are `service_access_role_arn`, `external_table_definition`, `csv_row_delimiter` (default: `\\n`), `csv_delimiter` (default: `,`), `bucket_folder`, `bucket_name` and `compression_type` (default: `NONE`). For more details, see [Using Amazon S3 as a Target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html).
	S3Settings interface{}
	// The host name of the server.
	ServerName interface{}
	// The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
	ServiceAccessRole interface{}
	// The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
	SslMode interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The user name to be used to login to the endpoint database.
	Username interface{}
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn interface{}
	// The name of the endpoint database.
	DatabaseName interface{}
	// The database endpoint identifier.
	EndpointId interface{}
	// The type of endpoint. Can be one of `source | target`.
	EndpointType interface{}
	// The type of engine for the endpoint. Can be one of `mysql | oracle | postgres | mariadb | aurora | redshift | sybase | sqlserver | dynamodb | mongodb | s3 | azuredb`.
	EngineName interface{}
	// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
	ExtraConnectionAttributes interface{}
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn interface{}
	// Settings for the source MongoDB endpoint. Available settings are `auth_type` (default: `PASSWORD`), `auth_mechanism` (default: `DEFAULT`), `nesting_level` (default: `NONE`), `extract_doc_id` (default: `false`), `docs_to_investigate` (default: `1000`) and `auth_source` (default: `admin`). For more details, see [Using MongoDB as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html).
	MongodbSettings interface{}
	// The password to be used to login to the endpoint database.
	Password interface{}
	// The port used by the endpoint database.
	Port interface{}
	// Settings for the target S3 endpoint. Available settings are `service_access_role_arn`, `external_table_definition`, `csv_row_delimiter` (default: `\\n`), `csv_delimiter` (default: `,`), `bucket_folder`, `bucket_name` and `compression_type` (default: `NONE`). For more details, see [Using Amazon S3 as a Target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html).
	S3Settings interface{}
	// The host name of the server.
	ServerName interface{}
	// The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
	ServiceAccessRole interface{}
	// The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
	SslMode interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The user name to be used to login to the endpoint database.
	Username interface{}
}