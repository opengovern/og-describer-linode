package maps

import (
	"github.com/opengovern/og-describer-linode/discovery/pkg/es"
)

var ResourceTypesToTables = map[string]string{
  "Linode/Account": "linode_account",
  "Linode/Database": "linode_database",
  "Linode/Domain": "linode_domain",
  "Linode/Firewall": "linode_firewall",
  "Linode/Image": "linode_image",
  "Linode/Kubernetes/Cluster": "linode_kubernetes_cluster",
  "Linode/Event": "linode_event",
  "Linode/Instance": "linode_instance",
  "Linode/Longview/Client": "linode_longview_client",
  "Linode/NodeBalancer": "linode_node_balancer",
  "Linode/ObjectStorage": "linode_object_storage",
  "Linode/StackScript": "linode_stack_script",
  "Linode/Vpc": "linode_vpc",
  "Linode/Volume": "linode_volume",
  "Linode/IPAddress": "linode_ip_address",
}

var ResourceTypeToDescription = map[string]interface{}{
  "Linode/Account": opengovernance.Account{},
  "Linode/Database": opengovernance.Database{},
  "Linode/Domain": opengovernance.Domain{},
  "Linode/Firewall": opengovernance.Firewall{},
  "Linode/Image": opengovernance.Image{},
  "Linode/Kubernetes/Cluster": opengovernance.KubernetesCluster{},
  "Linode/Event": opengovernance.Event{},
  "Linode/Instance": opengovernance.Instance{},
  "Linode/Longview/Client": opengovernance.LongViewClient{},
  "Linode/NodeBalancer": opengovernance.NodeBalancer{},
  "Linode/ObjectStorage": opengovernance.ObjectStorageBucket{},
  "Linode/StackScript": opengovernance.StackScript{},
  "Linode/Vpc": opengovernance.VPC{},
  "Linode/Volume": opengovernance.Volume{},
  "Linode/IPAddress": opengovernance.IPAddress{},
}

var TablesToResourceTypes = map[string]string{
  "linode_account": "Linode/Account",
  "linode_database": "Linode/Database",
  "linode_domain": "Linode/Domain",
  "linode_firewall": "Linode/Firewall",
  "linode_image": "Linode/Image",
  "linode_kubernetes_cluster": "Linode/Kubernetes/Cluster",
  "linode_event": "Linode/Event",
  "linode_instance": "Linode/Instance",
  "linode_longview_client": "Linode/Longview/Client",
  "linode_node_balancer": "Linode/NodeBalancer",
  "linode_object_storage": "Linode/ObjectStorage",
  "linode_stack_script": "Linode/StackScript",
  "linode_vpc": "Linode/Vpc",
  "linode_volume": "Linode/Volume",
  "linode_ip_address": "Linode/IPAddress",
}
