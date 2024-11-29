package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "linode"                                    // example: aws, azure
	IntegrationName      = integration.Type("LINODE_ACCOUNT")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-linode" // example: github.com/opengovern/og-describer-aws
)
