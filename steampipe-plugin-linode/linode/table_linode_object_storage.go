package linode

import (
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"golang.org/x/net/context"
)

func tableLinodeObjectStorageBucket(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_object_storage",
		Description: "Instances in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListObjectStorageBucket,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"cluster", "label"}),
			Hydrate:    opengovernance.GetObjectStorageBucket,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "The name of this bucket."},
			// Other columns
			{Name: "cluster",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster"),
				Description: "The ID of the Object Storage Cluster this bucket is in."},
			{Name: "region",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Region")},
			{Name: "created",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Created"),
				Description: "When this bucket was created."},
			{Name: "hostname",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Hostname"),
				Description: "The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public."},
			{Name: "objects",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.Objects")},
			{Name: "size",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.Size")},
		}),
	}
}
