package maps

import (
	"github.com/opengovern/og-describer-linode/discovery/describers"
	model "github.com/opengovern/og-describer-linode/discovery/pkg/models"
	"github.com/opengovern/og-describer-linode/discovery/provider"
	"github.com/opengovern/og-describer-linode/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypes = map[string]model.ResourceType{

	"Linode/Account": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Account",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListAccounts),
		GetDescriber:    nil,
	},

	"Linode/Database": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Database",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListDatabases),
		GetDescriber:    nil,
	},

	"Linode/Domain": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListDomains),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetDomain),
	},

	"Linode/Firewall": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Firewall",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListFirewalls),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetFirewall),
	},

	"Linode/Image": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Image",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListImages),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetImage),
	},

	"Linode/Kubernetes/Cluster": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Kubernetes/Cluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListKubernetesClusters),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetKubernetesCluster),
	},

	"Linode/Event": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Event",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListEvents),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetEvent),
	},

	"Linode/Instance": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Instance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListLinodeInstances),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetLinodeInstance),
	},

	"Linode/Longview/Client": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Longview/Client",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListLongViewClients),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetLongViewClient),
	},

	"Linode/NodeBalancer": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/NodeBalancer",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListNodeBalancers),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetNodeBalancer),
	},

	"Linode/ObjectStorage": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/ObjectStorage",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListObjectStorages),
		GetDescriber:    nil,
	},

	"Linode/StackScript": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/StackScript",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListStackScripts),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetStackScript),
	},

	"Linode/Vpc": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Vpc",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListVPCs),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetVPC),
	},

	"Linode/Volume": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Volume",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListVolumes),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetVolume),
	},

	"Linode/IPAddress": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/IPAddress",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describers.ListIPAddresses),
		GetDescriber:    provider.DescribeSingleByLinode(describers.GetIPAddress),
	},
}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Linode/Account": {
		Name:            "Linode/Account",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Database": {
		Name:            "Linode/Database",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Domain": {
		Name:            "Linode/Domain",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Firewall": {
		Name:            "Linode/Firewall",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Image": {
		Name:            "Linode/Image",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Kubernetes/Cluster": {
		Name:            "Linode/Kubernetes/Cluster",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Event": {
		Name:            "Linode/Event",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Instance": {
		Name:            "Linode/Instance",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Longview/Client": {
		Name:            "Linode/Longview/Client",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/NodeBalancer": {
		Name:            "Linode/NodeBalancer",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/ObjectStorage": {
		Name:            "Linode/ObjectStorage",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/StackScript": {
		Name:            "Linode/StackScript",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Vpc": {
		Name:            "Linode/Vpc",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Volume": {
		Name:            "Linode/Volume",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/IPAddress": {
		Name:            "Linode/IPAddress",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},
}

var ResourceTypesList = []string{
	"Linode/Account",
	"Linode/Database",
	"Linode/Domain",
	"Linode/Firewall",
	"Linode/Image",
	"Linode/Kubernetes/Cluster",
	"Linode/Event",
	"Linode/Instance",
	"Linode/Longview/Client",
	"Linode/NodeBalancer",
	"Linode/ObjectStorage",
	"Linode/StackScript",
	"Linode/Vpc",
	"Linode/Volume",
	"Linode/IPAddress",
}
