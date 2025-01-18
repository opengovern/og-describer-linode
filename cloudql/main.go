package main

import (
	"github.com/opengovern/og-describer-linode/cloudql/linode"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: linode.Plugin})
}
