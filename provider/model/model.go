//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

import (
	"github.com/linode/linodego"
	"net"
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

type InstanceAlert struct {
	CPU           int `json:"cpu"`
	IO            int `json:"io"`
	NetworkIn     int `json:"network_in"`
	NetworkOut    int `json:"network_out"`
	TransferQuota int `json:"transfer_quota"`
}

type InstanceBackup struct {
	Available bool `json:"available,omitempty"`
	Enabled   bool `json:"enabled,omitempty"`
	Schedule  struct {
		Day    string `json:"day,omitempty"`
		Window string `json:"window,omitempty"`
	} `json:"schedule,omitempty"`
}

type InstanceSpec struct {
	Disk     int `json:"disk"`
	Memory   int `json:"memory"`
	VCPUs    int `json:"vcpus"`
	Transfer int `json:"transfer"`
	GPUs     int `json:"gpus"`
}

type InstancePlacementGroup struct {
	ID                   int    `json:"id"`
	Label                string `json:"label"`
	PlacementGroupType   string `json:"placement_group_type"`
	PlacementGroupPolicy string `json:"placement_group_policy"`
}

type LinodeListResponse struct {
	Data  []LinodeDescription `json:"data"`
	Page  int                 `json:"page"`
	Pages int                 `json:"pages"`
}

type LinodeDescription struct {
	ID              int                     `json:"id"`
	Created         *time.Time              `json:"-"`
	Updated         *time.Time              `json:"-"`
	Region          string                  `json:"region"`
	Alerts          *InstanceAlert          `json:"alerts"`
	Backups         *InstanceBackup         `json:"backups"`
	Image           string                  `json:"image"`
	Group           string                  `json:"group"`
	IPv4            []*net.IP               `json:"ipv4"`
	IPv6            string                  `json:"ipv6"`
	Label           string                  `json:"label"`
	Type            string                  `json:"type"`
	Status          string                  `json:"status"`
	HasUserData     bool                    `json:"has_user_data"`
	Hypervisor      string                  `json:"hypervisor"`
	HostUUID        string                  `json:"host_uuid"`
	Specs           *InstanceSpec           `json:"specs"`
	WatchdogEnabled bool                    `json:"watchdog_enabled"`
	Tags            []string                `json:"tags"`
	PlacementGroup  *InstancePlacementGroup `json:"placement_group"`
	DiskEncryption  string                  `json:"disk_encryption"`
	LKEClusterID    int                     `json:"lke_cluster_id"`
	Capabilities    []string                `json:"capabilities"`
}

type NetworkAddresses struct {
	IPv4 *[]string `json:"ipv4,omitempty"`
	IPv6 *[]string `json:"ipv6,omitempty"`
}

type FirewallRule struct {
	Action      string           `json:"action"`
	Label       string           `json:"label"`
	Description string           `json:"description,omitempty"`
	Ports       string           `json:"ports,omitempty"`
	Protocol    string           `json:"protocol"`
	Addresses   NetworkAddresses `json:"addresses"`
}

type FirewallRuleSet struct {
	Inbound        []FirewallRule `json:"inbound"`
	InboundPolicy  string         `json:"inbound_policy"`
	Outbound       []FirewallRule `json:"outbound"`
	OutboundPolicy string         `json:"outbound_policy"`
}

type FirewallListResponse struct {
	Data  []FirewallDescription `json:"data"`
	Page  int                   `json:"page"`
	Pages int                   `json:"pages"`
}

type FirewallDescription struct {
	ID      int             `json:"id"`
	Label   string          `json:"label"`
	Status  string          `json:"status"`
	Tags    []string        `json:"tags,omitempty"`
	Rules   FirewallRuleSet `json:"rules"`
	Created *time.Time      `json:"created"`
	Updated *time.Time      `json:"updated"`
}

type ImageRegion struct {
	Region string `json:"region"`
	Status string `json:"status"`
}

type ImageListResponse struct {
	Data  []ImageDescription `json:"data"`
	Page  int                `json:"page"`
	Pages int                `json:"pages"`
}

type ImageDescription struct {
	ID           string        `json:"id"`
	CreatedBy    string        `json:"created_by"`
	Capabilities []string      `json:"capabilities"`
	Label        string        `json:"label"`
	Description  string        `json:"description"`
	Type         string        `json:"type"`
	Vendor       string        `json:"vendor"`
	Status       string        `json:"status"`
	Size         int           `json:"size"`
	TotalSize    int           `json:"total_size"`
	IsPublic     bool          `json:"is_public"`
	Deprecated   bool          `json:"deprecated"`
	Regions      []ImageRegion `json:"regions"`
	Tags         []string      `json:"tags"`
	Updated      *time.Time    `json:"updated"`
	Created      *time.Time    `json:"created"`
	Expiry       *time.Time    `json:"expiry"`
	EOL          *time.Time    `json:"eol"`
}

type LKEClusterControlPlane struct {
	HighAvailability bool `json:"high_availability"`
}

type KubernetesClusterListResponse struct {
	Data  []KubernetesClusterDescription `json:"data"`
	Page  int                            `json:"page"`
	Pages int                            `json:"pages"`
}

type KubernetesClusterDescription struct {
	linodego.LKECluster
	ID           int                    `json:"id"`
	Created      *time.Time             `json:"-"`
	Updated      *time.Time             `json:"-"`
	Label        string                 `json:"label"`
	Region       string                 `json:"region"`
	Status       string                 `json:"status"`
	K8sVersion   string                 `json:"k8s_version"`
	Tags         []string               `json:"tags"`
	ControlPlane LKEClusterControlPlane `json:"control_plane"`
}
