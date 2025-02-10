package maps

import (
	describer "github.com/opengovern/og-describer-linode/discovery/describers"
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
		ListDescriber:   provider.DescribeListByLinode(describer.ListAccounts),
		GetDescriber:    nil,
	},

	"Linode/Database": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Database",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListDatabases),
		GetDescriber:    nil,
	},

	"Linode/Domain": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListDomains),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetDomain),
	},

	"Linode/Firewall": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Firewall",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListFirewalls),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetFirewall),
	},

	"Linode/Image": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Image",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListImages),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetImage),
	},

	"Linode/Kubernetes/Cluster": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Kubernetes/Cluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListKubernetesClusters),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetKubernetesCluster),
	},

	"Linode/Event": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Event",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListEvents),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetEvent),
	},

	"Linode/Instance": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Instance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListLinodeInstances),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetLinodeInstance),
	},

	"Linode/Longview/Client": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Longview/Client",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListLongViewClients),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetLongViewClient),
	},

	"Linode/NodeBalancer": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/NodeBalancer",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListNodeBalancers),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetNodeBalancer),
	},

	"Linode/ObjectStorage": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/ObjectStorage",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListObjectStorages),
		GetDescriber:    nil,
	},

	"Linode/StackScript": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/StackScript",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListStackScripts),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetStackScript),
	},

	"Linode/Vpc": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Vpc",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListVPCs),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetVPC),
	},

	"Linode/Volume": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Volume",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListVolumes),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetVolume),
	},

	"Linode/IPAddress": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/IPAddress",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListIPAddresses),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetIPAddress),
	},

	"Linode/NodeBalancer/Config": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/NodeBalancer/Config",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListNodeBalancerConfigs),
		GetDescriber:    nil,
	},

	"Linode/NodeBalancer/Node": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/NodeBalancer/Node",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListNodeBalancerNodes),
		GetDescriber:    nil,
	},

	"Linode/Kubernetes/Node": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Kubernetes/Node",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListNodes),
		GetDescriber:    provider.DescribeSingleByLinode(describer.GetNode),
	},

	"Linode/Kubernetes/NodePool": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Linode/Kubernetes/NodePool",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByLinode(describer.ListNodePools),
		GetDescriber:    nil,
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

	"Linode/NodeBalancer/Config": {
		Name:            "Linode/NodeBalancer/Config",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/NodeBalancer/Node": {
		Name:            "Linode/NodeBalancer/Node",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Kubernetes/Node": {
		Name:            "Linode/Kubernetes/Node",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Linode/Kubernetes/NodePool": {
		Name:            "Linode/Kubernetes/NodePool",
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
	"Linode/NodeBalancer/Config",
	"Linode/NodeBalancer/Node",
	"Linode/Kubernetes/Node",
	"Linode/Kubernetes/NodePool",
}
