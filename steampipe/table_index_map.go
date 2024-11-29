package steampipe

import (
	"repo-url/pkg/sdk/es"
)

var Map = map[string]string{
  "Linode/Account": "linode_account",
  "Linode/Database": "linode_database",
  "Linode/Domain": "linode_domain",
  "Linode/Event": "linode_event",
  "Linode/Firewall": "linode_firewall",
  "Linode/Image": "linode_image",
  "Linode/Kubernetes/Cluster": "linode_kubernetes_cluster",
  "Linode/Instance": "linode_instance",
  "Linode/Longview/Client": "linode_longview_client",
  "Linode/NodeBalancer": "linode_node_balancer",
  "Linode/ObjectStorage": "linode_object_storage",
  "Linode/SlackScripts": "linode_slack_script",
}

var DescriptionMap = map[string]interface{}{
  "Linode/Account": opengovernance.Account{},
  "Linode/Database": opengovernance.Database{},
  "Linode/Domain": opengovernance.Domain{},
  "Linode/Event": opengovernance.Event{},
  "Linode/Firewall": opengovernance.Firewall{},
  "Linode/Image": opengovernance.Image{},
  "Linode/Kubernetes/Cluster": opengovernance.KubernetesCluster{},
  "Linode/Instance": opengovernance.Instance{},
  "Linode/Longview/Client": opengovernance.LongViewClient{},
  "Linode/NodeBalancer": opengovernance.NodeBalancer{},
  "Linode/ObjectStorage": opengovernance.ObjectStorage{},
  "Linode/SlackScripts": opengovernance.StackScript{},
}

var ReverseMap = map[string]string{
  "linode_account": "Linode/Account",
  "linode_database": "Linode/Database",
  "linode_domain": "Linode/Domain",
  "linode_event": "Linode/Event",
  "linode_firewall": "Linode/Firewall",
  "linode_image": "Linode/Image",
  "linode_kubernetes_cluster": "Linode/Kubernetes/Cluster",
  "linode_instance": "Linode/Instance",
  "linode_longview_client": "Linode/Longview/Client",
  "linode_node_balancer": "Linode/NodeBalancer",
  "linode_object_storage": "Linode/ObjectStorage",
  "linode_slack_script": "Linode/SlackScripts",
}
