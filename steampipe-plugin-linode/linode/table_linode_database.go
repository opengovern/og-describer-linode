package linode

import (
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"golang.org/x/net/context"
)

func tableLinodeDatabase(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_database",
		Description: "Instances in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDatabase,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDatabase,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this Instance."},
			{
				Name:      "status",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Status"),
			},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "The Instance’s label is for display purposes only.",
			},
			{
				Name:      "hosts",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Hosts"),
			},
			{
				Name:      "region",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Region"),
			},
			{
				Name:      "type",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Type"),
			},
			{
				Name:      "engine",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Engine"),
			},
			{
				Name:      "version",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Version"),
			},
			{
				Name:      "cluster_size",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.ClusterSize"),
			},
			{
				Name:      "replication_type",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.ReplicationType"),
			},
			{
				Name:      "ssl_connection",
				Type:      proto.ColumnType_BOOL,
				Transform: transform.FromField("Description.SSLConnection"),
			},
			{
				Name:      "encrypted",
				Type:      proto.ColumnType_BOOL,
				Transform: transform.FromField("Description.Encrypted"),
			},
			{
				Name:      "allow_list",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.AllowList"),
			},
			{
				Name:      "instance_uri",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.InstanceURI"),
			},
			{
				Name:      "created",
				Type:      proto.ColumnType_TIMESTAMP,
				Transform: transform.FromField("Description.Created"),
			},
			{
				Name:      "updated",
				Type:      proto.ColumnType_TIMESTAMP,
				Transform: transform.FromField("Description.Updated"),
			},
		}),
	}
}
