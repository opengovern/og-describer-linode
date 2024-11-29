package linode

import (
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"golang.org/x/net/context"
)

func tableLinodeStackScript(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_stack_script",
		Description: "Instances in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStackScript,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetStackScript,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name: "id", Type: proto.ColumnType_INT,
				Description: "The unique ID of this Instance.",
				Transform:   transform.FromField("Description.ID"),
			},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Description: "The Instance’s label is for display purposes only.",
				Transform:   transform.FromField("Description.Label"),
			},
			{
				Name:      "username",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Username"),
			},
			{
				Name:      "description",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Description"),
			},
			{
				Name:      "ordinal",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.Ordinal"),
			},
			{
				Name:      "logo_url",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.LogoURL"),
			},
			{
				Name:      "images",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Images"),
			},
			{
				Name:      "deployments_total",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.DeploymentsTotal"),
			},
			{
				Name:      "deployments_active",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.DeploymentsActive"),
			},
			{
				Name:      "is_public",
				Type:      proto.ColumnType_BOOL,
				Transform: transform.FromField("Description.IsPublic"),
			},
			{
				Name:      "mine",
				Type:      proto.ColumnType_BOOL,
				Transform: transform.FromField("Description.Mine"),
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
			{
				Name:      "rev_note",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.RevNote"),
			},
			{
				Name:      "script",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Script"),
			},
			{
				Name:      "user_defined_fields",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.UserDefinedFields"),
			},
			{
				Name:      "user_gravatar_id",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.UserGravatarID"),
			},
		}),
	}
}
