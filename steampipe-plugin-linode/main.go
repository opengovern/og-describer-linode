package main

import (
	"github.com/opengovern/og-describer-linode/steampipe-plugin-linode/linode"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: linode.Plugin})
}
