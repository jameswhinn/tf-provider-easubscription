package main

import (
	// "github.com/jameswhinn/tf-provider-easubscription"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return subscriptioncreator.Provider()
		},
	})
}
