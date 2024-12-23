package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_volume",
		Description: "Volumes in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListVolume,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetVolume,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this Volume."},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "The Volume’s label is for display purposes only."},
			{
				Name:        "size",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Size"),
				Description: "The Volume’s size, in GiB."},
			// Other columns
			{
				Name:        "created",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Created"),
				Description: "When this Volume was created."},
			{
				Name:        "filesystem_path",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FilesystemPath"),
				Description: "The full filesystem path for the Volume based on the Volume’s label."},
			{
				Name:        "linode_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LinodeID"),
				Description: "If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here."},
			{
				Name:        "region",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region"),
				Description: "Region where the volume resides."},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
				Description: "The current status of the volume: creating, active, resizing, contact_support."},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(transform.StringArrayToMap),
				Description: "Tags applied to this volume as a map."},
			{
				Name:        "encryption",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Encryption"),
				Description: ""},
			{
				Name:        "updated",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Updated"),
				Description: "When this Volume was last updated."},
		}),
	}
}
