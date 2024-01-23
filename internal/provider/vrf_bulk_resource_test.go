package provider

import (
	"context"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Test_vrfBulkResource_Create(t *testing.T) {
	type fields struct {
		client *ndfc.NDFC
	}
	type args struct {
		ctx  context.Context
		req  resource.CreateRequest
		resp *resource.CreateResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &vrfBulkResource{
				client: tt.fields.client,
			}
			r.Create(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}
