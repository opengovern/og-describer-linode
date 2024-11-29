package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeIPAddress(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_ip_address",
		Description: "Images in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIPAddress,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("address"),
			Hydrate:    opengovernance.GetIPAddress,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:      "address",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Address"),
			},
			{
				Name:      "gateway",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Gateway"),
			},
			{
				Name:      "subnet_mask",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.SubnetMask"),
			},
			{
				Name:      "prefix",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.Prefix"),
			},
			{
				Name:      "type",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Type"),
			},
			{
				Name:      "public",
				Type:      proto.ColumnType_BOOL,
				Transform: transform.FromField("Description.Public"),
			},
			{
				Name:      "rdns",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.RDNS"),
			},
			{
				Name:      "linode_id",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.LinodeID"),
			},
			{
				Name:      "region",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Region"),
			},
			{
				Name:      "vpc_nat_1_1",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.VPCNAT1To1"),
			},
			{
				Name:      "reserved",
				Type:      proto.ColumnType_BOOL,
				Transform: transform.FromField("Description.Reserved"),
			},
		}),
	}
}
