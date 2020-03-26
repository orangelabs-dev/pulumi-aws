// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates and manages an AWS IoT certificate.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_certificate.html.markdown.
type Certificate struct {
	pulumi.CustomResourceState

	// Boolean flag to indicate if the certificate should be active
	Active pulumi.BoolOutput `pulumi:"active"`
	// The ARN of the created certificate.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The certificate data, in PEM format.
	CertificatePem pulumi.StringOutput `pulumi:"certificatePem"`
	// The certificate signing request. Review
	// [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
	// for more information on generating a certificate from a certificate signing request (CSR).
	// If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
	// for more information on generating keys and a certificate.
	Csr pulumi.StringPtrOutput `pulumi:"csr"`
	// When no CSR is provided, the private key.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// When no CSR is provided, the public key.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil || args.Active == nil {
		return nil, errors.New("missing required argument 'Active'")
	}
	if args == nil {
		args = &CertificateArgs{}
	}
	var resource Certificate
	err := ctx.RegisterResource("aws:iot/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("aws:iot/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Boolean flag to indicate if the certificate should be active
	Active *bool `pulumi:"active"`
	// The ARN of the created certificate.
	Arn *string `pulumi:"arn"`
	// The certificate data, in PEM format.
	CertificatePem *string `pulumi:"certificatePem"`
	// The certificate signing request. Review
	// [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
	// for more information on generating a certificate from a certificate signing request (CSR).
	// If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
	// for more information on generating keys and a certificate.
	Csr *string `pulumi:"csr"`
	// When no CSR is provided, the private key.
	PrivateKey *string `pulumi:"privateKey"`
	// When no CSR is provided, the public key.
	PublicKey *string `pulumi:"publicKey"`
}

type CertificateState struct {
	// Boolean flag to indicate if the certificate should be active
	Active pulumi.BoolPtrInput
	// The ARN of the created certificate.
	Arn pulumi.StringPtrInput
	// The certificate data, in PEM format.
	CertificatePem pulumi.StringPtrInput
	// The certificate signing request. Review
	// [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
	// for more information on generating a certificate from a certificate signing request (CSR).
	// If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
	// for more information on generating keys and a certificate.
	Csr pulumi.StringPtrInput
	// When no CSR is provided, the private key.
	PrivateKey pulumi.StringPtrInput
	// When no CSR is provided, the public key.
	PublicKey pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Boolean flag to indicate if the certificate should be active
	Active bool `pulumi:"active"`
	// The certificate signing request. Review
	// [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
	// for more information on generating a certificate from a certificate signing request (CSR).
	// If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
	// for more information on generating keys and a certificate.
	Csr *string `pulumi:"csr"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Boolean flag to indicate if the certificate should be active
	Active pulumi.BoolInput
	// The certificate signing request. Review
	// [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
	// for more information on generating a certificate from a certificate signing request (CSR).
	// If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
	// for more information on generating keys and a certificate.
	Csr pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}
