package google

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google/provider"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

type Config transport_tpg.Config

func ConfigureBasePaths(c *Config) {
	transport_tpg.ConfigureBasePaths((*transport_tpg.Config)(c))
}

func Provider() *schema.Provider {
	return provider.Provider()
}

func (c *Config) LoadAndValidate(ctx context.Context) error {
	return (*transport_tpg.Config)(c).LoadAndValidate(ctx)
}
