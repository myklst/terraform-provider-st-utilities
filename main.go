package main

import (
	"context"
	"terraform-provider-st-utilities/utilities"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name test

func main() {
	providerserver.Serve(context.Background(), utilities.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/myklst/st-utilities",
	})
}
