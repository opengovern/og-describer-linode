package linode

import (
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"golang.org/x/net/context"
)

func tableLinodeLongviewClient(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_longview_client",
		Description: "Instances in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLongViewClient,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetLongViewClient,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "account",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account"),
				Description: "An external unique identifier for this account.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this Instance."},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "The Instance’s label is for display purposes only."},
			// Other columns
			{
				Name:        "api_key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIKey"),
				Description: "Alerts are triggered if CPU, IO, etc exceed these limits."},
			{
				Name:        "created",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Created"),
				Description: "When this Instance was created."},
			{
				Name:        "updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Updated"),
				Description: "When this Instance was created."},
			{
				Name:        "install_code",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstallCode"),
				Description: "Alerts are triggered if CPU, IO, etc exceed these limits."},
			{
				Name:        "apps",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Apps"),
				Description: "Alerts are triggered if CPU, IO, etc exceed these limits."},
		}),
	}
}
