// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDroplet(ctx *pulumi.Context, args *LookupDropletArgs, opts ...pulumi.InvokeOption) (*LookupDropletResult, error) {
	var rv LookupDropletResult
	err := ctx.Invoke("digitalocean:index/getDroplet:getDroplet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDroplet.
type LookupDropletArgs struct {
	// The ID of the Droplet
	Id *int `pulumi:"id"`
	// The name of the Droplet.
	Name *string `pulumi:"name"`
	// A tag applied to the Droplet.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by getDroplet.
type LookupDropletResult struct {
	// Whether backups are enabled.
	Backups   bool   `pulumi:"backups"`
	CreatedAt string `pulumi:"createdAt"`
	// The size of the Droplets disk in GB.
	Disk int  `pulumi:"disk"`
	Id   *int `pulumi:"id"`
	// The Droplet image ID or slug.
	Image string `pulumi:"image"`
	// The Droplets public IPv4 address
	Ipv4Address string `pulumi:"ipv4Address"`
	// The Droplets private IPv4 address
	Ipv4AddressPrivate string `pulumi:"ipv4AddressPrivate"`
	// Whether IPv6 is enabled.
	Ipv6 bool `pulumi:"ipv6"`
	// The Droplets public IPv6 address
	Ipv6Address string `pulumi:"ipv6Address"`
	// The Droplets private IPv6 address
	Ipv6AddressPrivate string `pulumi:"ipv6AddressPrivate"`
	// Whether the Droplet is locked.
	Locked bool `pulumi:"locked"`
	// The amount of the Droplets memory in MB.
	Memory int `pulumi:"memory"`
	// Whether monitoring agent is installed.
	Monitoring bool    `pulumi:"monitoring"`
	Name       *string `pulumi:"name"`
	// Droplet hourly price.
	PriceHourly float64 `pulumi:"priceHourly"`
	// Droplet monthly price.
	PriceMonthly float64 `pulumi:"priceMonthly"`
	// Whether private networks are enabled.
	PrivateNetworking bool `pulumi:"privateNetworking"`
	// The region the Droplet is running in.
	Region string `pulumi:"region"`
	// The unique slug that indentifies the type of Droplet.
	Size string `pulumi:"size"`
	// The status of the Droplet.
	Status string  `pulumi:"status"`
	Tag    *string `pulumi:"tag"`
	// A list of the tags associated to the Droplet.
	Tags []string `pulumi:"tags"`
	// The uniform resource name of the Droplet
	Urn string `pulumi:"urn"`
	// The number of the Droplets virtual CPUs.
	Vcpus int `pulumi:"vcpus"`
	// List of the IDs of each volumes attached to the Droplet.
	VolumeIds []string `pulumi:"volumeIds"`
	// The ID of the VPC where the Droplet is located.
	VpcUuid string `pulumi:"vpcUuid"`
}
