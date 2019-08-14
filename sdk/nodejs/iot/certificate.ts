// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages an AWS IoT certificate.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_certificate.html.markdown.
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateState, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iot/certificate:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * Boolean flag to indicate if the certificate should be active
     */
    public readonly active!: pulumi.Output<boolean>;
    /**
     * The ARN of the created certificate.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The certificate data, in PEM format.
     */
    public /*out*/ readonly certificatePem!: pulumi.Output<string>;
    /**
     * The certificate signing request. Review
     * [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
     * for more information on generating a certificate from a certificate signing request (CSR).
     * If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
     * for more information on generating keys and a certificate.
     */
    public readonly csr!: pulumi.Output<string | undefined>;
    /**
     * When no CSR is provided, the private key.
     */
    public /*out*/ readonly privateKey!: pulumi.Output<string>;
    /**
     * When no CSR is provided, the public key.
     */
    public /*out*/ readonly publicKey!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CertificateState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificatePem"] = state ? state.certificatePem : undefined;
            inputs["csr"] = state ? state.csr : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["publicKey"] = state ? state.publicKey : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if (!args || args.active === undefined) {
                throw new Error("Missing required property 'active'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["csr"] = args ? args.csr : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificatePem"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
            inputs["publicKey"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * Boolean flag to indicate if the certificate should be active
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * The ARN of the created certificate.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The certificate data, in PEM format.
     */
    readonly certificatePem?: pulumi.Input<string>;
    /**
     * The certificate signing request. Review
     * [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
     * for more information on generating a certificate from a certificate signing request (CSR).
     * If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
     * for more information on generating keys and a certificate.
     */
    readonly csr?: pulumi.Input<string>;
    /**
     * When no CSR is provided, the private key.
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * When no CSR is provided, the public key.
     */
    readonly publicKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Boolean flag to indicate if the certificate should be active
     */
    readonly active: pulumi.Input<boolean>;
    /**
     * The certificate signing request. Review
     * [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
     * for more information on generating a certificate from a certificate signing request (CSR).
     * If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
     * for more information on generating keys and a certificate.
     */
    readonly csr?: pulumi.Input<string>;
}
