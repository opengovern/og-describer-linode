package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeEventClient(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_event",
		Description: "",
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
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this Instance."},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
				Description: ""},
			// Other columns
			{
				Name:        "action",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Action"),
				Description: ""},
			{
				Name:        "percent_complete",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PercentComplete"),
				Description: ""},
			{
				Name:        "rate",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rate"),
				Description: ""},
			{
				Name:        "read",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Read"),
				Description: ""},
			{
				Name:        "seen",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Seen"),
				Description: ""},
			{
				Name:        "time_remaining",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TimeRemaining"),
				Description: ""},
			{
				Name:        "username",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Username"),
				Description: ""},
			{
				Name:        "entity",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Entity"),
				Description: ""},
			{
				Name:        "secondary_entity",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecondaryEntity"),
				Description: ""},
			{
				Name:        "created",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Created"),
				Description: ""},
			{
				Name:        "message",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Message"),
				Description: ""},
			{
				Name:        "duration",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Duration"),
				Description: ""},
		}),
	}
}
