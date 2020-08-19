// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SpacesBucketObject struct {
	pulumi.CustomResourceState

	// The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl pulumi.StringPtrOutput `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 pulumi.StringPtrOutput `pulumi:"contentBase64"`
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition pulumi.StringPtrOutput `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding pulumi.StringPtrOutput `pulumi:"contentEncoding"`
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage pulumi.StringPtrOutput `pulumi:"contentLanguage"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// the ETag generated for the object (an MD5 sum of the object content). The hash is an MD5 digest of the
	// object data. For objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5
	// digest. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The name of the object once it is in the bucket.
	Key pulumi.StringOutput `pulumi:"key"`
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The region where the bucket resides (Defaults to `nyc3`)
	Region pulumi.StringOutput `pulumi:"region"`
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect pulumi.StringPtrOutput `pulumi:"websiteRedirect"`
}

// NewSpacesBucketObject registers a new resource with the given unique name, arguments, and options.
func NewSpacesBucketObject(ctx *pulumi.Context,
	name string, args *SpacesBucketObjectArgs, opts ...pulumi.ResourceOption) (*SpacesBucketObject, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil {
		args = &SpacesBucketObjectArgs{}
	}
	var resource SpacesBucketObject
	err := ctx.RegisterResource("digitalocean:index/spacesBucketObject:SpacesBucketObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpacesBucketObject gets an existing SpacesBucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpacesBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpacesBucketObjectState, opts ...pulumi.ResourceOption) (*SpacesBucketObject, error) {
	var resource SpacesBucketObject
	err := ctx.ReadResource("digitalocean:index/spacesBucketObject:SpacesBucketObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpacesBucketObject resources.
type spacesBucketObjectState struct {
	// The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)
	Acl *string `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket *string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl *string `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `pulumi:"content"`
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 *string `pulumi:"contentBase64"`
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `pulumi:"contentType"`
	// the ETag generated for the object (an MD5 sum of the object content). The hash is an MD5 digest of the
	// object data. For objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5
	// digest. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).
	Etag *string `pulumi:"etag"`
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The name of the object once it is in the bucket.
	Key *string `pulumi:"key"`
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]string `pulumi:"metadata"`
	// The region where the bucket resides (Defaults to `nyc3`)
	Region *string `pulumi:"region"`
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source *string `pulumi:"source"`
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId *string `pulumi:"versionId"`
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect *string `pulumi:"websiteRedirect"`
}

type SpacesBucketObjectState struct {
	// The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)
	Acl pulumi.StringPtrInput
	// The name of the bucket to put the file in.
	Bucket pulumi.StringPtrInput
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl pulumi.StringPtrInput
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrInput
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 pulumi.StringPtrInput
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition pulumi.StringPtrInput
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding pulumi.StringPtrInput
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringPtrInput
	// the ETag generated for the object (an MD5 sum of the object content). The hash is an MD5 digest of the
	// object data. For objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5
	// digest. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).
	Etag pulumi.StringPtrInput
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy pulumi.BoolPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringPtrInput
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata pulumi.StringMapInput
	// The region where the bucket resides (Defaults to `nyc3`)
	Region pulumi.StringPtrInput
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source pulumi.StringPtrInput
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId pulumi.StringPtrInput
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect pulumi.StringPtrInput
}

func (SpacesBucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*spacesBucketObjectState)(nil)).Elem()
}

type spacesBucketObjectArgs struct {
	// The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)
	Acl *string `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl *string `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `pulumi:"content"`
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 *string `pulumi:"contentBase64"`
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `pulumi:"contentType"`
	// the ETag generated for the object (an MD5 sum of the object content). The hash is an MD5 digest of the
	// object data. For objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5
	// digest. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).
	Etag *string `pulumi:"etag"`
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The name of the object once it is in the bucket.
	Key string `pulumi:"key"`
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]string `pulumi:"metadata"`
	// The region where the bucket resides (Defaults to `nyc3`)
	Region string `pulumi:"region"`
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source *string `pulumi:"source"`
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect *string `pulumi:"websiteRedirect"`
}

// The set of arguments for constructing a SpacesBucketObject resource.
type SpacesBucketObjectArgs struct {
	// The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)
	Acl pulumi.StringPtrInput
	// The name of the bucket to put the file in.
	Bucket pulumi.StringInput
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl pulumi.StringPtrInput
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrInput
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 pulumi.StringPtrInput
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition pulumi.StringPtrInput
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding pulumi.StringPtrInput
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringPtrInput
	// the ETag generated for the object (an MD5 sum of the object content). The hash is an MD5 digest of the
	// object data. For objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5
	// digest. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).
	Etag pulumi.StringPtrInput
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy pulumi.BoolPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringInput
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata pulumi.StringMapInput
	// The region where the bucket resides (Defaults to `nyc3`)
	Region pulumi.StringInput
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source pulumi.StringPtrInput
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect pulumi.StringPtrInput
}

func (SpacesBucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spacesBucketObjectArgs)(nil)).Elem()
}