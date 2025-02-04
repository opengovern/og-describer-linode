package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeClusterNode(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_kubernetes_node",
		Description: "Nodes in the cluster.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListClusterNode,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetClusterNode,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique identifier of the node.",
			},
			{
				Name:        "instance_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceID"),
				Description: "The instance ID of the node.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
				Description: "The current status of the node.",
			},
			{
				Name:        "account",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account"),
				Description: "The account associated with the node.",
			},
		}),
	}
}
