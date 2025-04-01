package main

import "fmt"

func main() {
	// plugin framework isn't ready yet, so we'll just print a message.
	fmt.Println("Plugin framework is not ready yet. This is a sample plugin for testing the registry.")
	// serviceClient, closeService, err := pluginservicev1.NewEnvServiceClient()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// defer closeService()

	// hostInfoContainer := pluginutils.NewHostInfoContainer()
	// providerServer := providerv1.NewProviderPlugin(
	// 	aws.NewProvider(),
	// 	hostInfoContainer,
	// 	serviceClient,
	// )
	// config := plugin.ServeProviderConfiguration{
	// 	ID: "two-hundred/test",
	// 	PluginMetadata: &pluginservicev1.PluginMetadata{
	// 		PluginVersion: "1.0.0",
	// 		DisplayName:   "AWS",
	// 		FormattedDescription: "AWS provider for the Deploy Engine including `resources`, `data sources`," +
	// 			" `links` and `custom variable types` for interacting with AWs services.",
	// 		RepositoryUrl: "https://github.com/two-hundred/celerity-provider-aws",
	// 		Author:        "Two Hundred",
	// 	},
	// 	ProtocolVersion: "1.0.0",
	// }

	// close, err := plugin.ServeProviderV1(
	// 	context.Background(),
	// 	providerServer,
	// 	serviceClient,
	// 	hostInfoContainer,
	// 	config,
	// )
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// pluginutils.WaitForShutdown(close)
}
