package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeVPC(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_vpc",
		Description: "Images in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListVPC,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetVPC,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this VPC."},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "A short description of the VPC."},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Description"),
				Description: "A description of the VPC."},
			{
				Name:      "region",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Region"),
			},
			{
				Name:      "subnets",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Subnets"),
			},
			{
				Name:      "created",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Created"),
			},
			{
				Name:      "updated",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Updated"),
			},
		}),
	}
}
