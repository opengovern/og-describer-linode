package linode

import (
	"context"
	"fmt"
	"github.com/linode/linodego"
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"strings"
)

func tableLinodeInstance(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_instance",
		Description: "Instances in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListInstance,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetInstance,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this Instance."},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "The Instance’s label is for display purposes only."},
			// Other columns
			{
				Name:        "alerts",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Alerts"),
				Description: "Alerts are triggered if CPU, IO, etc exceed these limits."},
			{
				Name:        "backups",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Backups"),
				Description: "Information about this Linode’s backups status."},
			{
				Name:        "created",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Created"),
				Description: "When this Instance was created."},
			{
				Name:        "hypervisor",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Hypervisor"),
				Description: "The virtualization software powering this Linode, e.g. kvm."},
			{
				Name:        "image",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image"),
				Description: "An Image ID to deploy the Disk from."},
			{
				Name:        "instance_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Type"),
				Description: "This is the Linode Type that this Linode was deployed with."},
			{
				Name:        "ipv4",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IPv4"),
				Description: "Array of this Linode’s IPv4 Addresses."},
			{
				Name:        "ipv6",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.IPv6"),
				Description: "This Linode’s IPv6 SLAAC address."},
			{
				Name:        "region",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region"),
				Description: "Region where the instance resides."},
			{
				Name:        "specs",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Specs"),
				Description: "Information about the resources available to this Linode, e.g. disk space."},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
				Description: "The current status of the instance: creating, active, resizing, contact_support."},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(transform.StringArrayToMap),
				Description: "Tags applied to this instance as a map."},
			{
				Name:        "tags_src",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
				Description: "List of Tags applied to this instance."},
			{
				Name:        "updated",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Updated"),
				Description: "When this Instance was last updated."},
			{
				Name:        "watchdog_enabled",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.WatchdogEnabled"),
				Description: "The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly."},
		}),
	}
}

func listInstance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_instance.listInstance", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}
	keyQuals := d.EqualsQuals
	if keyQuals["filter"] != nil {
		opts.Filter = keyQuals["filter"].GetStringValue()
	} else {
		filterParts := []string{}
		/*
			// Image filters seem to return "Error: [400] [X-Filter] Could not apply filter"
			if keyQuals["image"] != nil {
				filterParts = append(filterParts, fmt.Sprintf(`"image":"%s"`, keyQuals["image"].GetStringValue()))
			}
		*/
		if keyQuals["label"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"label":"%s"`, keyQuals["label"].GetStringValue()))
		}
		if keyQuals["region"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"region":"%s"`, keyQuals["region"].GetStringValue()))
		}
		// TODO: ipv4
		// TODO: tags
		if len(filterParts) > 0 {
			opts.Filter = fmt.Sprintf("{%s}", strings.Join(filterParts, ","))
		}
	}

	items, err := conn.ListInstances(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_instance.listInstance", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getInstance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_instance.getInstance", "connection_error", err)
		return nil, err
	}
	id := int(d.EqualsQuals["id"].GetInt64Value())
	item, err := conn.GetInstance(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_instance.getInstance", "query_error", err)
		return nil, err
	}
	return item, err
}
