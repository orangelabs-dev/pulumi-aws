// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a S3 bucket resource.
type Bucket struct {
	pulumi.CustomResourceState

	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus pulumi.StringOutput `pulumi:"accelerationStatus"`
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName pulumi.StringOutput `pulumi:"bucketDomainName"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrOutput `pulumi:"bucketPrefix"`
	// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
	BucketRegionalDomainName pulumi.StringOutput `pulumi:"bucketRegionalDomainName"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants BucketGrantArrayOutput `pulumi:"grants"`
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings BucketLoggingArrayOutput `pulumi:"loggings"`
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration BucketObjectLockConfigurationPtrOutput `pulumi:"objectLockConfiguration"`
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document.
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region pulumi.StringOutput `pulumi:"region"`
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration BucketReplicationConfigurationPtrOutput `pulumi:"replicationConfiguration"`
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer pulumi.StringOutput `pulumi:"requestPayer"`
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration BucketServerSideEncryptionConfigurationPtrOutput `pulumi:"serverSideEncryptionConfiguration"`
	// A mapping of tags to assign to the bucket.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningOutput `pulumi:"versioning"`
	// A website object (documented below).
	Website BucketWebsitePtrOutput `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain pulumi.StringOutput `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringOutput `pulumi:"websiteEndpoint"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}
	var resource Bucket
	err := ctx.RegisterResource("aws:s3/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("aws:s3/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus *string `pulumi:"accelerationStatus"`
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl *string `pulumi:"acl"`
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn *string `pulumi:"arn"`
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket *string `pulumi:"bucket"`
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName *string `pulumi:"bucketDomainName"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
	BucketRegionalDomainName *string `pulumi:"bucketRegionalDomainName"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants []BucketGrant `pulumi:"grants"`
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId *string `pulumi:"hostedZoneId"`
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings []BucketLogging `pulumi:"loggings"`
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration *BucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document.
	Policy *string `pulumi:"policy"`
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region *string `pulumi:"region"`
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration *BucketReplicationConfiguration `pulumi:"replicationConfiguration"`
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer *string `pulumi:"requestPayer"`
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration *BucketServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// A mapping of tags to assign to the bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning *BucketVersioning `pulumi:"versioning"`
	// A website object (documented below).
	Website *BucketWebsite `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

type BucketState struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus pulumi.StringPtrInput
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl pulumi.StringPtrInput
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn pulumi.StringPtrInput
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket pulumi.StringPtrInput
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName pulumi.StringPtrInput
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrInput
	// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
	BucketRegionalDomainName pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayInput
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants BucketGrantArrayInput
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId pulumi.StringPtrInput
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings BucketLoggingArrayInput
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration BucketObjectLockConfigurationPtrInput
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document.
	Policy pulumi.StringPtrInput
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region pulumi.StringPtrInput
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration BucketReplicationConfigurationPtrInput
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer pulumi.StringPtrInput
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration BucketServerSideEncryptionConfigurationPtrInput
	// A mapping of tags to assign to the bucket.
	Tags pulumi.MapInput
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningPtrInput
	// A website object (documented below).
	Website BucketWebsitePtrInput
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus *string `pulumi:"accelerationStatus"`
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl interface{} `pulumi:"acl"`
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn *string `pulumi:"arn"`
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket *string `pulumi:"bucket"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants []BucketGrant `pulumi:"grants"`
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId *string `pulumi:"hostedZoneId"`
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings []BucketLogging `pulumi:"loggings"`
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration *BucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document.
	Policy interface{} `pulumi:"policy"`
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region *string `pulumi:"region"`
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration *BucketReplicationConfiguration `pulumi:"replicationConfiguration"`
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer *string `pulumi:"requestPayer"`
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration *BucketServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// A mapping of tags to assign to the bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning *BucketVersioning `pulumi:"versioning"`
	// A website object (documented below).
	Website *BucketWebsite `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus pulumi.StringPtrInput
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl pulumi.Input
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn pulumi.StringPtrInput
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket pulumi.StringPtrInput
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayInput
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants BucketGrantArrayInput
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId pulumi.StringPtrInput
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings BucketLoggingArrayInput
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration BucketObjectLockConfigurationPtrInput
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document.
	Policy pulumi.Input
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region pulumi.StringPtrInput
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration BucketReplicationConfigurationPtrInput
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer pulumi.StringPtrInput
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration BucketServerSideEncryptionConfigurationPtrInput
	// A mapping of tags to assign to the bucket.
	Tags pulumi.MapInput
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningPtrInput
	// A website object (documented below).
	Website BucketWebsitePtrInput
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}
