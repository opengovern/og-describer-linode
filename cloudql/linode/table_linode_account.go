package linode

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-linode/discovery/pkg/es"
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
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Email"),
				Description: "The email address of the person associated with this Account."},
			// Other columns
			{
				Name:        "address_1",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address1"),
				Description: "First line of this Account’s billing address."},
			{
				Name:        "address_2",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address2"),
				Description: "Second line of this Account’s billing address."},
			{
				Name:        "balance",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Balance"),
				Description: "This Account’s balance, in US dollars."},
			{
				Name:        "balance_uninvoiced",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BalanceUninvoiced"),
				Description: "This Account’s current estimated invoice in US dollars. This is not your final invoice balance. Transfer charges are not included in the estimate."},
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
				Name:        "credit_card",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CreditCard"),
				Description: "Credit Card information associated with this Account."},
			{
				Name:        "first_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirstName"),
				Description: "The first name of the person associated with this Account."},
			{
				Name:        "last_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LastName"),
				Description: "The last name of the person associated with this Account."},
			{
				Name:        "account",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account"),
				Description: "An external unique identifier for this account.",
			},
			{
				Name:        "phone",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Phone"),
				Description: "The phone number associated with this Account."},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.State"),
				Description: "The state for this Account’s billing address."},
			{
				Name:        "tax_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaxID"),
				Description: "The tax identification number associated with this Account, for tax calculations in some countries. If you do not live in a country that collects tax, this should be null."},
			{
				Name:        "zip",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Zip"),
				Description: "The zip code of this Account’s billing address."},
		}),
	}
}
