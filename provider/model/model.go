//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

import (
	"time"
)

type Metadata struct{}

type CreditCard struct {
	LastFour string `json:"last_four"`
	Expiry   string `json:"expiry"`
}

type Account struct {
	Email             string      `json:"email"`
	Address1          string      `json:"address_1"`
	Address2          string      `json:"address_2"`
	Balance           float32     `json:"balance"`
	BalanceUninvoiced float32     `json:"balance_uninvoiced"`
	City              string      `json:"city"`
	Company           string      `json:"company"`
	Country           string      `json:"country"`
	CreditCard        *CreditCard `json:"credit_card"`
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	Euuid             string      `json:"euuid"`
	Phone             string      `json:"phone"`
	State             string      `json:"state"`
	TaxID             string      `json:"tax_id"`
	Zip               string      `json:"zip"`
}

type AccountDescription struct {
	Email             string      `json:"email"`
	Address1          string      `json:"address_1"`
	Address2          string      `json:"address_2"`
	Balance           string      `json:"balance"`
	BalanceUninvoiced string      `json:"balance_uninvoiced"`
	City              string      `json:"city"`
	Company           string      `json:"company"`
	Country           string      `json:"country"`
	CreditCard        *CreditCard `json:"credit_card"`
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	Euuid             string      `json:"euuid"`
	Phone             string      `json:"phone"`
	State             string      `json:"state"`
	TaxID             string      `json:"tax_id"`
	Zip               string      `json:"zip"`
}

type DatabaseListResponse struct {
	Data  []DatabaseDescription `json:"data"`
	Page  int                   `json:"page"`
	Pages int                   `json:"pages"`
}

type DatabaseHost struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary,omitempty"`
}

type DatabaseDescription struct {
	ID              int          `json:"id"`
	Status          string       `json:"status"`
	Label           string       `json:"label"`
	Hosts           DatabaseHost `json:"hosts"`
	Region          string       `json:"region"`
	Type            string       `json:"type"`
	Engine          string       `json:"engine"`
	Version         string       `json:"version"`
	ClusterSize     int          `json:"cluster_size"`
	ReplicationType string       `json:"replication_type"`
	SSLConnection   bool         `json:"ssl_connection"`
	Encrypted       bool         `json:"encrypted"`
	AllowList       []string     `json:"allow_list"`
	InstanceURI     string       `json:"instance_uri"`
	Created         *time.Time   `json:"created"`
	Updated         *time.Time   `json:"updated"`
}

type DomainListResponse struct {
	Data  []DomainDescription `json:"data"`
	Page  int                 `json:"page"`
	Pages int                 `json:"pages"`
}

type DomainDescription struct {
	ID          int      `json:"id"`
	Domain      string   `json:"domain"`
	Type        string   `json:"type"`
	Group       string   `json:"group"`
	Status      string   `json:"status"`
	Description string   `json:"description"`
	SOAEmail    string   `json:"soa_email"`
	RetrySec    int      `json:"retry_sec"`
	MasterIPs   []string `json:"master_ips"`
	AXfrIPs     []string `json:"axfr_ips"`
	Tags        []string `json:"tags"`
	ExpireSec   int      `json:"expire_sec"`
	RefreshSec  int      `json:"refresh_sec"`
	TTLSec      int      `json:"ttl_sec"`
}

type EventEntity struct {
	ID     any    `json:"id"`
	Label  string `json:"label"`
	Type   string `json:"type"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type EventListResponse struct {
	Data  []EventDescription `json:"data"`
	Page  int                `json:"page"`
	Pages int                `json:"pages"`
}

type EventDescription struct {
	ID              int          `json:"id"`
	Status          string       `json:"status"`
	Action          string       `json:"action"`
	PercentComplete int          `json:"percent_complete"`
	Rate            *string      `json:"rate"`
	Read            bool         `json:"read"`
	Seen            bool         `json:"seen"`
	TimeRemaining   *int         `json:"time_remaining"`
	Username        string       `json:"username"`
	Entity          *EventEntity `json:"entity"`
	SecondaryEntity *EventEntity `json:"secondary_entity"`
	Created         *time.Time   `json:"created"`
	Message         string       `json:"message"`
	Duration        float64      `json:"duration"`
}
