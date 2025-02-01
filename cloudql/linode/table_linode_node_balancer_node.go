package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeNodeBalancerNode(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_node_balancer_node",
		Description: "Nodes assigned to a NodeBalancer and readable by the requesting User.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNode,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetNode,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this NodeBalancer Node.",
			},
			{
				Name:        "account",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account"),
				Description: "An external unique identifier for this account.",
			},
			{
				Name:        "nodebalancer_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.NodeBalancerID"),
				Description: "The ID of the associated NodeBalancer.",
			},
			{
				Name:        "config_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ConfigID"),
				Description: "The ID of the associated NodeBalancer configuration.",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address"),
				Description: "The address and port of the backend node.",
			},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "The label for the backend node.",
			},
			{
				Name:        "mode",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Mode"),
				Description: "The mode of the node (e.g., accept, reject, drain).",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
				Description: "The status of the node (e.g., UP, DOWN).",
			},
			{
				Name:        "weight",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Weight"),
				Description: "The weight assigned to the node for load balancing.",
			},
		}),
	}
}
