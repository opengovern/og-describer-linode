package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLinodeFirewall(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_firewall",
		Description: "Firewalls in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFirewall,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetFirewall,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this Firewall."},
			{
				Name:        "account",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account"),
				Description: "An external unique identifier for this account.",
			},
			{
				Name:        "created",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Created"),
				Description: "The date and time this firewall was created."},
			{
				Name:        "updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Updated"),
				Description: "The date and time this firewall was updated."},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Label"),
				Description: "The firewall’s label is for display purposes only."},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
				Description: "The status of the firewall. Possible values are 'enabled', 'disabled', or 'deleted'."},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
				Description: "Tags applied to this firewall."},
			{
				Name:        "rules",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rules"),
				Description: "The rules associated with the firewall."},
		}),
	}
}
