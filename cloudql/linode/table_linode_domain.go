package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_domain",
		Description: "Domains in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDomain,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDomain,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this Domain."},
			{
				Name:        "domain",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain"),
				Description: "The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain."},
			// Other columns
			{
				Name:        "axfr_ips",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AXfrIPs"),
				Description: "The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it."},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Description"),
				Description: "A description for this Domain. This is for display purposes only."},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Type"),
				Description: "If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave)."},
			{
				Name:        "expire_sec",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ExpireSec"),
				Description: "The amount of time in seconds that may pass before this Domain is no longer authoritative."},
			{
				Name:        "master_ips",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MasterIPs"),
				Description: "The IP addresses representing the master DNS for this Domain."},
			{
				Name:        "refresh_sec",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RefreshSec"),
				Description: "The amount of time in seconds before this Domain should be refreshed."},
			{
				Name:        "retry_sec",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RetrySec"),
				Description: "The interval, in seconds, at which a failed refresh should be retried."},
			{
				Name:        "soa_email",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SOAEmail"),
				Description: "Start of Authority email address. This is required for master Domains."},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
				Description: "Used to control whether this Domain is currently being rendered: disabled, active, edit_mode, has_errors."},
			{
				Name:        "group",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group"),
				Description: ""},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(transform.StringArrayToMap),
				Description: "Tags applied to this domain as a map."},
			{
				Name:        "ttl_sec",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TTLSec").NullIfZero(),
				Description: "Time to Live - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers."},
		}),
	}
}
