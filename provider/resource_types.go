package provider
import (
	"github.com/opengovern/og-describer-linode/provider/describer"
	"github.com/opengovern/og-describer-linode/provider/configs"
	model "github.com/opengovern/og-describer-linode/pkg/sdk/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"Linode/Account": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Account",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListAccounts),
		GetDescriber:         nil,
	},

	"Linode/Database": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Database",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListDatabases),
		GetDescriber:         nil,
	},

	"Linode/Domain": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListDomains),
		GetDescriber:         DescribeSingleByLinode(describer.GetDomain),
	},

	"Linode/Firewall": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Firewall",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListFirewalls),
		GetDescriber:         DescribeSingleByLinode(describer.GetFirewall),
	},

	"Linode/Image": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Image",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListImages),
		GetDescriber:         DescribeSingleByLinode(describer.GetImage),
	},

	"Linode/Kubernetes/Cluster": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Kubernetes/Cluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListKubernetesClusters),
		GetDescriber:         DescribeSingleByLinode(describer.GetKubernetesCluster),
	},

	"Linode/Instance": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Instance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListLinodeInstances),
		GetDescriber:         DescribeSingleByLinode(describer.GetLinodeInstance),
	},

	"Linode/Longview/Client": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Longview/Client",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListLongViewClients),
		GetDescriber:         DescribeSingleByLinode(describer.GetLongViewClient),
	},

	"Linode/NodeBalancer": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/NodeBalancer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListNodeBalancers),
		GetDescriber:         DescribeSingleByLinode(describer.GetNodeBalancer),
	},

	"Linode/ObjectStorage": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/ObjectStorage",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListObjectStorages),
		GetDescriber:         nil,
	},

	"Linode/StackScript": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/StackScript",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListStackScripts),
		GetDescriber:         DescribeSingleByLinode(describer.GetStackScript),
	},

	"Linode/Vpc": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Vpc",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListVPCs),
		GetDescriber:         DescribeSingleByLinode(describer.GetVPC),
	},

	"Linode/Volume": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Volume",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListVolumes),
		GetDescriber:         DescribeSingleByLinode(describer.GetVolume),
	},

	"Linode/IPAddress": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/IPAddress",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListIPAddresses),
		GetDescriber:         DescribeSingleByLinode(describer.GetIPAddress),
	},
}
