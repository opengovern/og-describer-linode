package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeClusterNodePool(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_kubernetes_node_pool",
		Description: "Node pools in the cluster.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNodePool,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetNodePool,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique identifier of the node pool.",
			},
			{
				Name:        "autoscaler",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Autoscaler"),
				Description: "The autoscaler configuration for this node pool.",
			},
			{
				Name:        "count",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Count"),
				Description: "The number of nodes in the pool.",
			},
			{
				Name:        "disk_encryption",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskEncryption"),
				Description: "The disk encryption status of the node pool.",
			},
			{
				Name:        "disks",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Disks"),
				Description: "List of disks attached to the nodes in this pool.",
			},
			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Labels"),
				Description: "A map of labels assigned to the node pool.",
			},
			{
				Name:        "nodes",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodes"),
				Description: "List of nodes in this pool.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
				Description: "Tags associated with this node pool.",
			},
			{
				Name:        "taints",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Taints"),
				Description: "Taints applied to the nodes in this pool.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Type"),
				Description: "The type of node in the pool.",
			},
			{
				Name:        "account",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account"),
				Description: "The account associated with the node pool.",
			},
		}),
	}
}
