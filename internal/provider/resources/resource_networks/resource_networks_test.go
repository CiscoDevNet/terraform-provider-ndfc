package resource_networks

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceNetworks(t *testing.T) {
	type args struct {
		rscType  string
		rscName  string
		dataFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_resource_networks",
			args: args{
				rscType:  "resource",
				rscName:  "networks",
				dataFile: "/examples/ndfc_payloads/data_networks.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"model_read", func(t *testing.T) {
			fileName := os.Getenv("GOPATH") + "/src/terraform-provider-ndfc" + tt.args.dataFile
			RsType := tt.args.rscType
			rscName := tt.args.rscName
			dataFromFile, err := os.ReadFile(fileName)
			if err != nil {
				t.Errorf("File read failure %v", err)
				return
			}
			modelData := NDFCNetworksModel{}
			v := NetworksModel{}
			vP := NDFCNetworksPayload{}

			err = json.Unmarshal(dataFromFile, &vP.Networks)

			if err != nil {
				t.Errorf("Json Unmarshal failed %s_%s: %v", RsType, rscName, err)
			}
			modelData.FillNetworksFromPayload(&vP)
			for _, network := range vP.Networks {
				t.Logf("Network: %s", network.NetworkName)
				t.Logf("Network: %s", network.NetworkType)
				for _, subnet := range network.NetworkTemplateConfig.DhcpRelayServers {
					t.Logf("Address: %s Vrf %s", subnet.Address, subnet.Vrf)
				}
			}
			t.Logf("%v", modelData)

			if err := v.SetModelData(&modelData); err != nil {
				t.Errorf("SetModelData failed %s_%s: %v", RsType, rscName, err)
			}
			t.Logf("%s_%s Read and Set ok", RsType, rscName)

			var dataFromModel []byte

			modelDataRead := v.GetModelData()

			payloadData := modelDataRead.FillNetworksPayloadFromModel()

			dataFromModel, err = json.Marshal(&payloadData.Networks)
			if err != nil {
				t.Errorf("Json marshal failed %s_%s: %v", RsType, rscName, err)
				t.Fatalf("Json marshal failed %s_%s: %v", RsType, rscName, err)
			}
			t.Logf("%s_%s Marshall ok\n%v", RsType, rscName, string(dataFromModel))
			require.JSONEq(t, string(dataFromModel), string(dataFromFile))

		})
	}
}
