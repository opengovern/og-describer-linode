package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeNodeBalancerConfig(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_node_balancer_config",
		Description: "Configurations for NodeBalancers assigned to this Linode and readable by the requesting User.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNodeBalancerConfig,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetNodeBalancerConfig,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique ID of this NodeBalancer configuration.",
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
				Name:        "algorithm",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Algorithm"),
				Description: "The balancing algorithm used by this configuration.",
			},
			{
				Name:        "check",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Check"),
				Description: "The type of health check used for nodes.",
			},
			{
				Name:        "check_attempts",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CheckAttempts"),
				Description: "The number of failed health check attempts before marking a node down.",
			},
			{
				Name:        "check_body",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckBody"),
				Description: "The expected body content for HTTP checks.",
			},
			{
				Name:        "check_interval",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CheckInterval"),
				Description: "The interval (in seconds) between health checks.",
			},
			{
				Name:        "check_passive",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CheckPassive"),
				Description: "Whether passive checks are enabled.",
			},
			{
				Name:        "check_path",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckPath"),
				Description: "The path used for HTTP health checks.",
			},
			{
				Name:        "check_timeout",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CheckTimeout"),
				Description: "The timeout (in seconds) for health checks.",
			},
			{
				Name:        "cipher_suite",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CipherSuite"),
				Description: "The cipher suite used for SSL termination.",
			},
			{
				Name:        "port",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Port"),
				Description: "The port this configuration listens on.",
			},
			{
				Name:        "protocol",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Protocol"),
				Description: "The protocol used by this configuration (e.g., HTTP, HTTPS, TCP).",
			},
			{
				Name:        "proxy_protocol",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProxyProtocol"),
				Description: "The proxy protocol setting.",
			},
			{
				Name:        "ssl_cert",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSLCert"),
				Description: "The SSL certificate used for this configuration (if applicable).",
			},
			{
				Name:        "ssl_commonname",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSLCommonName"),
				Description: "The common name (CN) from the SSL certificate.",
			},
			{
				Name:        "ssl_fingerprint",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSLFingerprint"),
				Description: "The SSL certificate fingerprint.",
			},
			{
				Name:        "ssl_key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSLKey"),
				Description: "The SSL private key used for this configuration (if applicable).",
			},
			{
				Name:        "stickiness",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stickiness"),
				Description: "The session stickiness setting.",
			},
			{
				Name:        "nodes",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodes"),
				Description: "A list of nodes associated with this configuration.",
			},
			{
				Name:        "nodes_status",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NodesStatus"),
				Description: "The status of nodes within this configuration.",
			},
		}),
	}
}
