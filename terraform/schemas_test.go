package terraform

import (
	"github.com/hashicorp/terraform/configs/configschema"
)

func simpleTestSchemas() *Schemas {
	provider := simpleMockProvider()
	provisioner := simpleMockProvisioner()
	return &Schemas{
		providers: map[string]*ProviderSchema{
			"test": provider.GetSchemaReturn,
		},
		provisioners: map[string]*configschema.Block{
			"test": provisioner.GetSchemaResponse.Provisioner,
		},
	}
}