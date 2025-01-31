//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package provider

import (
	"net"
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
	ExpirationDate string `json:"-"`

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
	FirstName         string
	LastName          string
	Email             string
	Company           string
	Address1          string
	Address2          string
	Balance           float32
	BalanceUninvoiced float32
	City              string
	State             string
	Zip               string
	Country           string
	TaxID             string
	Phone             string
	CreditCard        CreditCard `json:"credit_card"`
	EUUID             string
	BillingSource     string
	Capabilities      []string `json:"capabilities"`
	ActiveSince       string
	ActivePromotions  []Promotion `json:"active_promotions"`
}

type AccountDescription struct {
	Email   string
	City    string
	Company string
	Country string
	Euuid   string
}

type DatabaseListResponse struct {
	Data  []DatabaseSingleResponse `json:"data"`
	Page  int                      `json:"page"`
	Pages int                      `json:"pages"`
}

type DatabaseHost struct {
	Primary   string
	Secondary string
}
type DatabaseHostSingle struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary,omitempty"`
}
type DatabaseSingleResponse struct {
	ID              int                `json:"id"`
	Status          string             `json:"status"`
	Label           string             `json:"label"`
	Hosts           DatabaseHostSingle `json:"hosts"`
	Region          string             `json:"region"`
	Type            string             `json:"type"`
	Engine          string             `json:"engine"`
	Version         string             `json:"version"`
	ClusterSize     int                `json:"cluster_size"`
	ReplicationType string             `json:"replication_type"`
	SSLConnection   bool               `json:"ssl_connection"`
	Encrypted       bool               `json:"encrypted"`
	AllowList       []string           `json:"allow_list"`
	InstanceURI     string             `json:"instance_uri"`
	Created         string             `json:"created"`
	Updated         string             `json:"updated"`
}

type DatabaseDescription struct {
	ID              int
	Status          string
	Label           string
	Hosts           DatabaseHost
	Region          string
	Type            string
	Engine          string
	Version         string
	ClusterSize     int
	ReplicationType string
	SSLConnection   bool
	Encrypted       bool
	AllowList       []string
	InstanceURI     string
	Created         string
	Updated         string
}

type DomainListResponse struct {
	Data  []DomainRecord `json:"data"`
	Page  int            `json:"page"`
	Pages int            `json:"pages"`
}

type DomainRecord struct {
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

type DomainDescription struct {
	ID          int
	Domain      string
	Type        string
	Group       string
	Status      string
	Description string
	SOAEmail    string
	RetrySec    int
	MasterIPs   []string
	AXfrIPs     []string
	Tags        []string
	ExpireSec   int
	RefreshSec  int
	TTLSec      int
}

type EventEntity struct {
	ID     any
	Label  string
	Type   string
	Status string
	URL    string
}

type EventListResponse struct {
	Data  []EventResp `json:"data"`
	Page  int         `json:"page"`
	Pages int         `json:"pages"`
}
type EventResp struct {
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
	Created         string       `json:"created"`
	Message         string       `json:"message"`
	Duration        float64      `json:"duration"`
}

type EventDescription struct {
	ID              int
	Status          string
	Action          string
	PercentComplete int
	Rate            *string
	Read            bool
	Seen            bool
	TimeRemaining   *int
	Username        string
	Entity          *EventEntity
	SecondaryEntity *EventEntity
	Created         string
	Message         string
	Duration        float64
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
	Data  []LinodeSingleResponse `json:"data"`
	Page  int                    `json:"page"`
	Pages int                    `json:"pages"`
}
type LinodeSingleResponse struct {
	ID              int                     `json:"id"`
	Created         string                  `json:"created"`
	Updated         string                  `json:"updated"`
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

type InstanceDescription struct {
	ID              int
	Created         string
	Updated         string
	Region          string
	Alerts          *InstanceAlert
	Backups         *InstanceBackup
	Image           string
	Group           string
	IPv4            []*net.IP
	IPv6            string
	Label           string
	Type            string
	Status          string
	HasUserData     bool
	Hypervisor      string
	HostUUID        string
	Specs           *InstanceSpec
	WatchdogEnabled bool
	Tags            []string
	PlacementGroup  *InstancePlacementGroup
	DiskEncryption  string
	LKEClusterID    int
	Capabilities    []string
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
	ID      int
	Label   string
	Status  string
	Tags    []string
	Rules   FirewallRuleSet
	Created string
	Updated string
}

type ImageRegion struct {
	Region string `json:"region"`
	Status string `json:"status"`
}

type ImageListResponse struct {
	Data  []ImageResponseSingle `json:"data"`
	Page  int                   `json:"page"`
	Pages int                   `json:"pages"`
}

type ImageResponseSingle struct {
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
	Updated      string        `json:"updated"`
	Created      string        `json:"created"`
	Expiry       string        `json:"expiry"`
	EOL          string        `json:"eol"`
}

type ImageDescription struct {
	ID           string
	CreatedBy    string
	Capabilities []string
	Label        string
	Description  string
	Type         string
	Vendor       string
	Status       string
	Size         int
	TotalSize    int
	IsPublic     bool
	Deprecated   bool
	Regions      []ImageRegion
	Tags         []string
	Updated      string
	Created      string
	Expiry       string
	EOL          string
}

type LKEClusterControlPlane struct {
	HighAvailability bool
}
type LKEClusterControlPlaneResp struct {
	HighAvailability bool `json:"high_availability"`
}

type KubernetesClusterListResponse struct {
	Data  []KubernetesClusterResp `json:"data"`
	Page  int                     `json:"page"`
	Pages int                     `json:"pages"`
}

type KubernetesClusterResp struct {
	ID           int                        `json:"id"`
	Created      string                     `json:"created"`
	Updated      string                     `json:"updated"`
	Label        string                     `json:"label"`
	Region       string                     `json:"region"`
	Status       string                     `json:"status"`
	K8sVersion   string                     `json:"k8s_version"`
	Tags         []string                   `json:"tags"`
	ControlPlane LKEClusterControlPlaneResp `json:"control_plane"`
}

type KubernetesClusterDescription struct {
	ID           int
	Created      string
	Updated      string
	Label        string
	Region       string
	Status       string
	K8sVersion   string
	Tags         []string
	ControlPlane LKEClusterControlPlane
}

type LongViewClientListResponse struct {
	Data  []LongViewClientDescription `json:"data"`
	Page  int                         `json:"page"`
	Pages int                         `json:"pages"`
}

type LongViewClientDescription struct {
	ID          int    `json:"id"`
	APIKey      string `json:"api_key"`
	Created     string `json:"created"`
	InstallCode string `json:"install_code"`
	Label       string `json:"label"`
	Updated     string `json:"updated"`
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
	Data  []NodeBalancerResp `json:"data"`
	Page  int                `json:"page"`
	Pages int                `json:"pages"`
}
type NodeBalancerResp struct {
	ID                 int                  `json:"id"`
	Label              *string              `json:"label"`
	Region             string               `json:"region"`
	Hostname           *string              `json:"hostname"`
	IPv4               *string              `json:"ipv4"`
	IPv6               *string              `json:"ipv6"`
	ClientConnThrottle int                  `json:"client_conn_throttle"`
	Transfer           NodeBalancerTransfer `json:"transfer"`
	Tags               []string             `json:"tags"`
	Created            string               `json:"created"`
	Updated            string               `json:"updated"`
}

type NodeBalancerDescription struct {
	ID                 int
	Label              *string
	Region             string
	Hostname           *string
	IPv4               *string
	IPv6               *string
	ClientConnThrottle int
	Transfer           NodeBalancerTransfer
	Tags               []string
	Created            string
	Updated            string
}

type ObjectStorageBucketListResponse struct {
	Data  []ObjectStorageBucketDescription `json:"data"`
	Page  int                              `json:"page"`
	Pages int                              `json:"pages"`
}

// ObjectStorageBucketDescription represents a ObjectStorage object
type ObjectStorageBucketDescription struct {
	Label    string
	Cluster  string
	Region   string
	Created  string
	Hostname string
	Objects  int
	Size     int
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
	Data  []StackScriptResp `json:"data"`
	Page  int               `json:"page"`
	Pages int               `json:"pages"`
}
type StackScriptResp struct {
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
	Created           string            `json:"created"`
	Updated           string            `json:"updated"`
	RevNote           string            `json:"rev_note"`
	Script            string            `json:"script"`
	UserDefinedFields *[]StackScriptUDF `json:"user_defined_fields"`
	UserGravatarID    string            `json:"user_gravatar_id"`
}

type StackScriptDescription struct {
	ID                int
	Username          string
	Label             string
	Description       string
	Ordinal           int
	LogoURL           string
	Images            []string
	DeploymentsTotal  int
	DeploymentsActive int
	IsPublic          bool
	Mine              bool
	Created           string
	Updated           string
	RevNote           string
	Script            string
	UserDefinedFields *[]StackScriptUDF
	UserGravatarID    string
}

type VolumeListResponse struct {
	Data  []VolumeSingleResponse `json:"data"`
	Page  int                    `json:"page"`
	Pages int                    `json:"pages"`
}

type VolumeSingleResponse struct {
	ID             int      `json:"id"`
	Label          string   `json:"label"`
	Status         string   `json:"status"`
	Region         string   `json:"region"`
	Size           int      `json:"size"`
	LinodeID       *int     `json:"linode_id"`
	FilesystemPath string   `json:"filesystem_path"`
	Tags           []string `json:"tags"`
	Created        string   `json:"created"`
	Updated        string   `json:"updated"`
	Encryption     string   `json:"encryption"`
}

type VolumeDescription struct {
	ID             int
	Label          string
	Status         string
	Region         string
	Size           int
	LinodeID       *int
	FilesystemPath string
	Tags           []string
	Created        string
	Updated        string
	Encryption     string
}

type VPCSubnetLinodeInterface struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
}

type VPCSubnetLinode struct {
	ID         int
	Interfaces []VPCSubnetLinodeInterface
}

type VPCSubnet struct {
	ID      int
	Label   string
	IPv4    string
	Linodes []VPCSubnetLinode
	Created string
	Updated string
}

type VPCListResponse struct {
	Data  []VPCDescription `json:"data"`
	Page  int              `json:"page"`
	Pages int              `json:"pages"`
}

type VPCDescription struct {
	ID          int
	Label       string
	Description string
	Region      string
	Subnets     []VPCSubnet
	Created     string
	Updated     string
}

type InstanceIPNAT1To1 struct {
	Address  string `json:"address"`
	SubnetID int    `json:"subnet_id"`
	VPCID    int    `json:"vpc_id"`
}

type IPAddressListResponse struct {
	Data  []IPAddressResp `json:"data"`
	Page  int             `json:"page"`
	Pages int             `json:"pages"`
}
type IPAddressResp struct {
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

type IPAddressDescription struct {
	Address    string
	Gateway    string
	SubnetMask string
	Prefix     int
	Type       string
	Public     bool
	RDNS       string
	LinodeID   int
	Region     string
	VPCNAT1To1 *InstanceIPNAT1To1
	Reserved   bool
}

type NodeJSON struct {
	Address        string `json:"address"`
	ConfigID       int    `json:"config_id"`
	ID             int    `json:"id"`
	Label          string `json:"label"`
	Mode           string `json:"mode"`
	NodeBalancerID int    `json:"nodebalancer_id"`
	Status         string `json:"status"`
	Weight         int    `json:"weight"`
}

type Node struct {
	Address        string
	ConfigID       int
	ID             int
	Label          string
	Mode           string
	NodeBalancerID int
	Status         string
	Weight         int
}

type NodesStatusJSON struct {
	Down int `json:"down"`
	Up   int `json:"up"`
}

type NodesStatus struct {
	Down int
	Up   int
}

type NodeBalancerConfigListResponse struct {
	Data  []NodeBalancerConfigJSON `json:"data"`
	Page  int                      `json:"page"`
	Pages int                      `json:"pages"`
}

type NodeBalancerConfigJSON struct {
	Algorithm      string          `json:"algorithm"`
	Check          string          `json:"check"`
	CheckAttempts  int             `json:"check_attempts"`
	CheckBody      string          `json:"check_body"`
	CheckInterval  int             `json:"check_interval"`
	CheckPassive   bool            `json:"check_passive"`
	CheckPath      string          `json:"check_path"`
	CheckTimeout   int             `json:"check_timeout"`
	CipherSuite    string          `json:"cipher_suite"`
	ID             int             `json:"id"`
	NodeBalancerID int             `json:"nodebalancer_id"`
	Nodes          []NodeJSON      `json:"nodes"`
	NodesStatus    NodesStatusJSON `json:"nodes_status"`
	Port           int             `json:"port"`
	Protocol       string          `json:"protocol"`
	ProxyProtocol  string          `json:"proxy_protocol"`
	SSLCert        *string         `json:"ssl_cert"`
	SSLCommonName  string          `json:"ssl_commonname"`
	SSLFingerprint string          `json:"ssl_fingerprint"`
	SSLKey         *string         `json:"ssl_key"`
	Stickiness     string          `json:"stickiness"`
}

type NodeBalancerConfigDescription struct {
	Algorithm      string
	Check          string
	CheckAttempts  int
	CheckBody      string
	CheckInterval  int
	CheckPassive   bool
	CheckPath      string
	CheckTimeout   int
	CipherSuite    string
	ID             int
	NodeBalancerID int
	Nodes          []Node
	NodesStatus    NodesStatus
	Port           int
	Protocol       string
	ProxyProtocol  string
	SSLCert        *string
	SSLCommonName  string
	SSLFingerprint string
	SSLKey         *string
	Stickiness     string
}

type NodeBalancerNodeListResponse struct {
	Data  []NodeRespJSON `json:"data"`
	Page  int            `json:"page"`
	Pages int            `json:"pages"`
}

type NodeRespJSON struct {
	Address        string `json:"address"`
	ConfigID       int    `json:"config_id"`
	ID             int    `json:"id"`
	Label          string `json:"label"`
	Mode           string `json:"mode"`
	NodeBalancerID int    `json:"nodebalancer_id"`
	Status         string `json:"status"`
	Weight         int    `json:"weight"`
}

type NodeDescription struct {
	Address        string
	ConfigID       int
	ID             int
	Label          string
	Mode           string
	NodeBalancerID int
	Status         string
	Weight         int
}
