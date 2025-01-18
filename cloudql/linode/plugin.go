package linode

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-linode",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"linode_account":            tableLinodeAccount(ctx),
			"linode_domain":             tableLinodeDomain(ctx),
			"linode_firewall":           tableLinodeFirewall(ctx),
			"linode_image":              tableLinodeImage(ctx),
			"linode_instance":           tableLinodeInstance(ctx),
			"linode_kubernetes_cluster": tableLinodeKubernetesCluster(ctx),
			"linode_node_balancer":      tableLinodeNodeBalancer(ctx),
			"linode_volume":             tableLinodeVolume(ctx),
			"linode_stack_script":       tableLinodeStackScript(ctx),
			"linode_database":           tableLinodeDatabase(ctx),
			"linode_longview_client":    tableLinodeLongviewClient(ctx),
			"linode_object_storage":     tableLinodeObjectStorageBucket(ctx),
			"linode_vpc":                tableLinodeVPC(ctx),
			"linode_ip_address":         tableLinodeIPAddress(ctx),
			"linode_event":              tableLinodeEvent(ctx),
		},
	}
	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "platform_integration_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"platform_integration_id", "platform_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"platform_integration_id", "platform_resource_id"})...)
			}
		}
	}
	return p
}
