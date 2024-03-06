package resource_vrf_attachments

import (
	"terraform-provider-ndfc/internal/provider/types"
	"testing"
)

func TestNDFCInstanceValuesValue_MarshalJSON(t *testing.T) {
	loopback := new(types.Int64Custom)
	*loopback = types.Int64Custom(1001)

	type fields struct {
		LoopbackId   *types.Int64Custom
		LoopbackIpv4 string
		LoopbackIpv6 string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				LoopbackId:   loopback,
				LoopbackIpv4: "10.1.1.1",
				LoopbackIpv6: "2001:db8::68",
			},
			want: `"{\"loopbackId\":\"1001\",\"loopbackIpv4\":\"10.1.1.1\",\"loopbackIpv6\":\"2001:db8::68\"}"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NDFCInstanceValuesValue{
				LoopbackId:   tt.fields.LoopbackId,
				LoopbackIpv4: tt.fields.LoopbackIpv4,
				LoopbackIpv6: tt.fields.LoopbackIpv6,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("NDFCInstanceValuesValue.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %s\nwant %s", got, tt.want)
			if string(got) != tt.want {
				t.Errorf("NDFCInstanceValuesValue.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
