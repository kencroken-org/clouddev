// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("digitalocean:index/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// The name of certificate.
	Name string `pulumi:"name"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	Domains []string `pulumi:"domains"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	Name            string `pulumi:"name"`
	NotAfter        string `pulumi:"notAfter"`
	Sha1Fingerprint string `pulumi:"sha1Fingerprint"`
	State           string `pulumi:"state"`
	Type            string `pulumi:"type"`
}
