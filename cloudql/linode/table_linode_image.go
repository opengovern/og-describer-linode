package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_image",
		Description: "Images in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListImage,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetImage,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this Image."},
			{
				Name:        "account",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account"),
				Description: "An external unique identifier for this account.",
			},
			{Name: "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "A short description of the Image."},
			// Other columns
			{Name: "created",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Created"),
				Description: "When this Image was created."},
			{Name: "created_by",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CreatedBy"),
				Description: "The name of the User who created this Image, or 'linode' for official Images."},
			{Name: "capabilities",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Capabilities"),
				Description: ""},
			{Name: "deprecated",
				Transform:   transform.FromField("Description.Deprecated"),
				Type:        proto.ColumnType_BOOL,
				Description: "Whether or not this Image is deprecated. Will only be true for deprecated public Images."},
			{Name: "description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Description"),
				Description: "A detailed description of this Image."},
			{Name: "type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Type"),
				Description: ""},
			{Name: "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
				Description: ""},
			{Name: "expiry",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Expiry"),
				Description: "Only Images created automatically (from a deleted Linode; type=automatic) will expire."},
			{
				Name:        "is_public",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.IsPublic"),
				Description: "True if the Image is public."},
			{
				Name:        "size",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Size"),
				Description: "The minimum size this Image needs to deploy. Size is in MB."},
			{
				Name:        "vendor",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vendor"),
				Description: "The upstream distribution vendor. None for private Images."},
			{Name: "total_size",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TotalSize"),
				Description: ""},
			{Name: "regions",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Regions"),
				Description: ""},
			{Name: "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
				Description: ""},
			{Name: "updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Updated"),
				Description: ""},
			{Name: "eol",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EOL"),
				Description: ""},
		}),
	}
}
