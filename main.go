package main

import (
	"context"
	"fmt"
	"log"

	"github.com/newstack-cloud/bluelink/libs/plugin-framework/plugin"
	"github.com/newstack-cloud/bluelink/libs/plugin-framework/pluginservicev1"
	"github.com/newstack-cloud/bluelink/libs/plugin-framework/sdk/pluginutils"
	"github.com/newstack-cloud/bluelink/libs/plugin-framework/sdk/providerv1"
	"github.com/newstack-cloud/celerity-provider-test/providertest"
)

func main() {
	serviceClient, closeService, err := pluginservicev1.NewEnvServiceClient()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer closeService()

	hostInfoContainer := pluginutils.NewHostInfoContainer()
	providerServer := providerv1.NewProviderPlugin(
		providertest.NewProvider(),
		hostInfoContainer,
		serviceClient,
	)
	config := plugin.ServePluginConfiguration{
		ID: "newstack-cloud/test",
		PluginMetadata: &pluginservicev1.PluginMetadata{
			PluginVersion: "1.0.0",
			DisplayName:   "AWS",
			FormattedDescription: "AWS provider for the Deploy Engine including `resources`, `data sources`," +
				" `links` and `custom variable types` for interacting with AWs services.",
			RepositoryUrl: "https://github.com/newstack-cloud/bluelink-provider-aws",
			Author:        "Two Hundred",
		},
		ProtocolVersion: "1.0",
	}

	fmt.Println("Starting provider plugin server...")
	close, err := plugin.ServeProviderV1(
		context.Background(),
		providerServer,
		serviceClient,
		hostInfoContainer,
		config,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	pluginutils.WaitForShutdown(close)
}
