package ndfc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func (c NDFC) SaveConfiguration(ctx context.Context, diags *diag.Diagnostics, fabricName string) {
	_, err := c.apiClient.Post(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/config-save", fabricName), "")
	if err != nil {
		diags.AddError("Save failed", err.Error())
		res, _ := c.apiClient.Get(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/errors", fabricName))
		// TODO determine which error should be returned
		diags.AddError("Fabric Errors", res.String())
	}
}

func (c NDFC) DeployConfiguration(ctx context.Context, diags *diag.Diagnostics, fabricName string, serialNumbers []string) {
	_, err := c.apiClient.Post(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/config-deploy/%s", fabricName, strings.Join(serialNumbers, ",")), "")
	if err != nil {
		diags.AddError("Deploy failed", err.Error())
	}
}
