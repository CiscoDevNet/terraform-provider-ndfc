package resource_vrf_attachments

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceVrfAttachments(t *testing.T) {
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
			name: "Test_resource_vrf_attachments",
			args: args{
				rscType:  "resource",
				rscName:  "vrf_attachments",
				dataFile: "/examples/ndfc_payloads/data_vrf_attachments.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"model_read", func(t *testing.T) {
			fileName := os.Getenv("GOPATH") + "/src/terraform-provider-ndfc" + tt.args.dataFile
			rsType := tt.args.rscType
			rscName := tt.args.rscName
			dataFromFile, err := os.ReadFile(fileName)
			if err != nil {
				t.Errorf("File read failure %v", err)
				return
			}
			modelData := NDFCVrfAttachmentsModel{}
			v := VrfAttachmentsModel{}

			err = json.Unmarshal(dataFromFile, &modelData)
			if err != nil {
				t.Errorf("Json Unmarshal failed %s_%s: %v", rsType, rscName, err)
			}
			if err := v.SetModelData(&modelData); err != nil {
				t.Errorf("SetModelData failed %s_%s: %v", rsType, rscName, err)
			}
			t.Logf("%s_%s Read and Set ok", rsType, rscName)

			var dataFromModel []byte

			modelDataRead := v.GetModelData()

			dataFromModel, err = json.Marshal(&modelDataRead)
			if err != nil {
				t.Errorf("Json marshal failed %s_%s: %v", rsType, rscName, err)
			}
			t.Logf("%s_%s Marshall ok", rsType, rscName)

			require.JSONEq(t, string(dataFromModel), string(dataFromFile))

		})
	}
}
