//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

import (
	"net"
	"strings"
	"time"
)

type Metadata struct{}

type CreditCard struct {
	LastFour string `json:"last_four"`
	Expiry   string `json:"expiry"`
}

// Promotion represents a Promotion object
type Promotion struct {
	// The amount available to spend per month.
	CreditMonthlyCap string `json:"credit_monthly_cap"`

	// The total amount of credit left for this promotion.
	CreditRemaining string `json:"credit_remaining"`

	// A detailed description of this promotion.
	Description string `json:"description"`

	// When this promotion's credits expire.
	ExpirationDate *TimeStamp `json:"-"`

	// The location of an image for this promotion.
	ImageURL string `json:"image_url"`

	// The service to which this promotion applies.
	ServiceType string `json:"service_type"`

	// Short details of this promotion.
	Summary string `json:"summary"`

	// The amount of credit left for this month for this promotion.
	ThisMonthCreditRemaining string `json:"this_month_credit_remaining"`
}

type Account struct {
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	Email             string      `json:"email"`
	Company           string      `json:"company"`
	Address1          string      `json:"address_1"`
	Address2          string      `json:"address_2"`
	Balance           float32     `json:"balance"`
	BalanceUninvoiced float32     `json:"balance_uninvoiced"`
	City              string      `json:"city"`
	State             string      `json:"state"`
	Zip               string      `json:"zip"`
	Country           string      `json:"country"`
	TaxID             string      `json:"tax_id"`
	Phone             string      `json:"phone"`
	CreditCard        *CreditCard `json:"credit_card"`
	EUUID             string      `json:"euuid"`
	BillingSource     string      `json:"billing_source"`
	Capabilities      []string    `json:"capabilities"`
	ActiveSince       *TimeStamp  `json:"active_since"`
	ActivePromotions  []Promotion `json:"active_promotions"`
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
	Created         *TimeStamp   `json:"created"`
	Updated         *TimeStamp   `json:"updated"`
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
	Created         *TimeStamp   `json:"created"`
	Message         string       `json:"message"`
	Duration        float64      `json:"duration"`
}

type TimeStamp struct {
	time.Time
}

func (ct *TimeStamp) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`) // Remove quotes
	t, err := time.Parse("2006-01-02T15:04:05", str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
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
	Data  []InstanceDescription `json:"data"`
	Page  int                   `json:"page"`
	Pages int                   `json:"pages"`
}

type InstanceDescription struct {
	ID              int                     `json:"id"`
	Created         *TimeStamp              `json:"created"`
	Updated         *TimeStamp              `json:"updated"`
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
	Created *TimeStamp      `json:"created"`
	Updated *TimeStamp      `json:"updated"`
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
	Updated      *TimeStamp    `json:"updated"`
	Created      *TimeStamp    `json:"created"`
	Expiry       *TimeStamp    `json:"expiry"`
	EOL          *TimeStamp    `json:"eol"`
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
	ID           int                    `json:"id"`
	Created      *TimeStamp             `json:"created"`
	Updated      *TimeStamp             `json:"updated"`
	Label        string                 `json:"label"`
	Region       string                 `json:"region"`
	Status       string                 `json:"status"`
	K8sVersion   string                 `json:"k8s_version"`
	Tags         []string               `json:"tags"`
	ControlPlane LKEClusterControlPlane `json:"control_plane"`
}

type LongViewClientListResponse struct {
	Data  []LongViewClientDescription `json:"data"`
	Page  int                         `json:"page"`
	Pages int                         `json:"pages"`
}

type LongViewClientDescription struct {
	ID          int        `json:"id"`
	APIKey      string     `json:"api_key"`
	Created     *TimeStamp `json:"created"`
	InstallCode string     `json:"install_code"`
	Label       string     `json:"label"`
	Updated     *TimeStamp `json:"updated"`
	Apps        struct {
		Apache any `json:"apache"`
		MySQL  any `json:"mysql"`
		NginX  any `json:"nginx"`
	} `json:"apps"`
}

type NodeBalancerTransfer struct {
	Total *float64 `json:"total"`
	Out   *float64 `json:"out"`
	In    *float64 `json:"in"`
}

type NodeBalancerListResponse struct {
	Data  []NodeBalancerDescription `json:"data"`
	Page  int                       `json:"page"`
	Pages int                       `json:"pages"`
}

type NodeBalancerDescription struct {
	ID                 int                  `json:"id"`
	Label              *string              `json:"label"`
	Region             string               `json:"region"`
	Hostname           *string              `json:"hostname"`
	IPv4               *string              `json:"ipv4"`
	IPv6               *string              `json:"ipv6"`
	ClientConnThrottle int                  `json:"client_conn_throttle"`
	Transfer           NodeBalancerTransfer `json:"transfer"`
	Tags               []string             `json:"tags"`
	Created            *TimeStamp           `json:"created"`
	Updated            *TimeStamp           `json:"updated"`
}

type ObjectStorageDescription struct {
	AmountUsed int `json:"used"`
}

type StackScriptUDF struct {
	Label   string `json:"label"`
	Name    string `json:"name"`
	Example string `json:"example"`
	OneOf   string `json:"oneOf,omitempty"`
	ManyOf  string `json:"manyOf,omitempty"`
	Default string `json:"default,omitempty"`
}

type StackScriptListResponse struct {
	Data  []StackScriptDescription `json:"data"`
	Page  int                      `json:"page"`
	Pages int                      `json:"pages"`
}

type StackScriptDescription struct {
	ID                int               `json:"id"`
	Username          string            `json:"username"`
	Label             string            `json:"label"`
	Description       string            `json:"description"`
	Ordinal           int               `json:"ordinal"`
	LogoURL           string            `json:"logo_url"`
	Images            []string          `json:"images"`
	DeploymentsTotal  int               `json:"deployments_total"`
	DeploymentsActive int               `json:"deployments_active"`
	IsPublic          bool              `json:"is_public"`
	Mine              bool              `json:"mine"`
	Created           *TimeStamp        `json:"created"`
	Updated           *TimeStamp        `json:"updated"`
	RevNote           string            `json:"rev_note"`
	Script            string            `json:"script"`
	UserDefinedFields *[]StackScriptUDF `json:"user_defined_fields"`
	UserGravatarID    string            `json:"user_gravatar_id"`
}

type VolumeListResponse struct {
	Data  []VolumeDescription `json:"data"`
	Page  int                 `json:"page"`
	Pages int                 `json:"pages"`
}

type VolumeDescription struct {
	ID             int        `json:"id"`
	Label          string     `json:"label"`
	Status         string     `json:"status"`
	Region         string     `json:"region"`
	Size           int        `json:"size"`
	LinodeID       *int       `json:"linode_id"`
	FilesystemPath string     `json:"filesystem_path"`
	Tags           []string   `json:"tags"`
	Created        *TimeStamp `json:"created"`
	Updated        *TimeStamp `json:"updated"`
	Encryption     string     `json:"encryption"`
}

type VPCSubnetLinodeInterface struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
}

type VPCSubnetLinode struct {
	ID         int                        `json:"id"`
	Interfaces []VPCSubnetLinodeInterface `json:"interfaces"`
}

type VPCSubnet struct {
	ID      int               `json:"id"`
	Label   string            `json:"label"`
	IPv4    string            `json:"ipv4"`
	Linodes []VPCSubnetLinode `json:"linodes"`
	Created *TimeStamp        `json:"created"`
	Updated *TimeStamp        `json:"updated"`
}

type VPCListResponse struct {
	Data  []VPCDescription `json:"data"`
	Page  int              `json:"page"`
	Pages int              `json:"pages"`
}

type VPCDescription struct {
	ID          int         `json:"id"`
	Label       string      `json:"label"`
	Description string      `json:"description"`
	Region      string      `json:"region"`
	Subnets     []VPCSubnet `json:"subnets"`
	Created     *TimeStamp  `json:"created"`
	Updated     *TimeStamp  `json:"updated"`
}

type InstanceIPNAT1To1 struct {
	Address  string `json:"address"`
	SubnetID int    `json:"subnet_id"`
	VPCID    int    `json:"vpc_id"`
}

type IPAddressListResponse struct {
	Data  []IPAddressDescription `json:"data"`
	Page  int                    `json:"page"`
	Pages int                    `json:"pages"`
}

type IPAddressDescription struct {
	Address    string             `json:"address"`
	Gateway    string             `json:"gateway"`
	SubnetMask string             `json:"subnet_mask"`
	Prefix     int                `json:"prefix"`
	Type       string             `json:"type"`
	Public     bool               `json:"public"`
	RDNS       string             `json:"rdns"`
	LinodeID   int                `json:"linode_id"`
	Region     string             `json:"region"`
	VPCNAT1To1 *InstanceIPNAT1To1 `json:"vpc_nat_1_1"`
	Reserved   bool               `json:"reserved"`
}
