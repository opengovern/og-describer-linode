package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_account",
		Description: "Account information.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccount,
		},
		Columns: []*plugin.Column{
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Email"),
				Description: "The email address of the person associated with this Account."},
			{
				Name:        "city",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.City"),
				Description: "The city for this Account’s billing address."},
			{
				Name:        "company",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Company"),
				Description: "The company name associated with this Account."},
			{
				Name:        "country",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Country"),
				Description: "The two-letter country code of this Account’s billing address."},
			{
				Name:        "euuid",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Euuid"),
				Description: "An external unique identifier for this account.",
			},
		},
	}
}
