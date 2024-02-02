package ndfc

import (
	"reflect"
	"testing"
)

func TestVrfAttachmentsSplitID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []string
	}{
		{name: "test1", args: args{id: "testf/v1{1,2,3}/v2{1,2,3}"}, want: "testf", want1: []string{"v1", "v2"}},
		{name: "test1", args: args{id: "testc/v5{1,2,3}/v6{1,2,3}"}, want: "testc", want1: []string{"v5", "v6"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := VrfAttachmentsSplitID(tt.args.id)
			if got != tt.want {
				t.Errorf("VrfAttachmentsSplitID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("VrfAttachmentsSplitID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
