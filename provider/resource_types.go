package provider
import (
	"repo-url/provider/describer"
	"repo-url/provider/configs"
	model "github.com/opengovern/og-describer-integrationType/pkg/sdk/models"
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
		ListDescriber:        DescribeListByLinode(describer.GetAccount),
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
		GetDescriber:         nil,
	},

	"Linode/Event": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/Event",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListEvents),
		GetDescriber:         nil,
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
		GetDescriber:         nil,
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
		ListDescriber:        ,
		GetDescriber:         DescribeSingleByLinode(describer.GetObjectStorage),
	},

	"Linode/SlackScripts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Linode/SlackScripts",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByLinode(describer.ListStackScripts),
		GetDescriber:         DescribeSingleByLinode(describer.GetStackScript),
	},
}
