package main

import (
	"github.com/arctiqtim/terraform-provider-jsondecode/decode"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: decode.Provider,
	})
}
