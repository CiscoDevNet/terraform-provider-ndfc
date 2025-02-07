package ndfc

import (
	"context"
	"testing"

	"terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/stretchr/testify/require"
)

func TestCheckNwAttachmentsAction(t *testing.T) {
	tests := []struct {
		name                string
		plan                *resource_networks.NDFCNetworksValue
		state               *resource_networks.NDFCNetworksValue
		vpcPairMap          map[string]string
		nwModified          bool
		globalDeploy        bool
		expectedFlagsGlobal uint16
		expectedFlagsAttach map[string]uint16
	}{
		{
			name: "New_Attachments_NwDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				NetworkName:       "test_net",
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				NetworkName:       "test_net",
				Attachments:       map[string]resource_network_attachments.NDFCAttachmentsValue{},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        false,
			expectedFlagsGlobal: (NewEntry | Deploy),
			expectedFlagsAttach: map[string]uint16{
				"serial1": (NewEntry | Deploy),
				"serial2": (NewEntry | Deploy),
				"serial3": (NewEntry | Deploy),
			},
		},
		{
			name: "New_Attachments_GlobalDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				NetworkName:       "test_net",
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				NetworkName:       "test_net",
				Attachments:       map[string]resource_network_attachments.NDFCAttachmentsValue{},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        true,
			expectedFlagsGlobal: (NewEntry | Deploy),
			expectedFlagsAttach: map[string]uint16{
				"serial1": (NewEntry | Deploy),
				"serial2": (NewEntry | Deploy),
				"serial3": (NewEntry | Deploy),
			},
		},
		{
			name: "New_Attachments_AttachDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				NetworkName:       "test_net",
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: true},
					"serial2": {DeployThisAttachment: true},
					"serial3": {DeployThisAttachment: true},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				NetworkName:       "test_net",
				Attachments:       map[string]resource_network_attachments.NDFCAttachmentsValue{},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        false,
			expectedFlagsGlobal: (NewEntry | Deploy),
			expectedFlagsAttach: map[string]uint16{
				"serial1": (NewEntry | Deploy),
				"serial2": (NewEntry | Deploy),
				"serial3": (NewEntry | Deploy),
			},
		},
		{
			name: "New_Attachments_AttachDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				NetworkName:       "test_net",
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: true},
					"serial2": {DeployThisAttachment: true},
					"serial3": {DeployThisAttachment: true},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				NetworkName:       "test_net",
				Attachments:       map[string]resource_network_attachments.NDFCAttachmentsValue{},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        false,
			expectedFlagsGlobal: (NewEntry | Deploy),
			expectedFlagsAttach: map[string]uint16{
				"serial1": (NewEntry | Deploy),
				"serial2": (NewEntry | Deploy),
				"serial3": (NewEntry | Deploy),
			},
		},
		{
			name: "New_Attachments_GlobalDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				NetworkName:       "test_net",
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				NetworkName:       "test_net",
				Attachments:       map[string]resource_network_attachments.NDFCAttachmentsValue{},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        true,
			expectedFlagsGlobal: (NewEntry | DeployAll | Deploy),
			expectedFlagsAttach: map[string]uint16{
				"serial1": (NewEntry | Deploy),
				"serial2": (NewEntry | Deploy),
				"serial3": (NewEntry | Deploy),
			},
		},
		{
			name: "New_Attachments_NwDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				NetworkName:       "test_net",
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				NetworkName:       "test_net",
				Attachments:       map[string]resource_network_attachments.NDFCAttachmentsValue{},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        false,
			expectedFlagsGlobal: (NewEntry | DeployAll | Deploy),
			expectedFlagsAttach: map[string]uint16{
				"serial1": (NewEntry | Deploy),
				"serial2": (NewEntry | Deploy),
				"serial3": (NewEntry | Deploy),
			},
		},
		// Modify attachments
		{
			name: "Modify_attachments_GlobalDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false, FreeformConfig: "new config"},
					"serial2": {DeployThisAttachment: false, FreeformConfig: "new config"},
					"serial3": {DeployThisAttachment: false, FreeformConfig: "new config"},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: false},
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        true,
			expectedFlagsGlobal: Update | Deploy,
			expectedFlagsAttach: map[string]uint16{
				"serial0": NoChange,
				"serial1": Update | Deploy,
				"serial2": Update | Deploy,
				"serial3": Update | Deploy,
			},
		},
		{
			name: "Modify_attachments_GlobalDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false, FreeformConfig: "new config"},
					"serial2": {DeployThisAttachment: false, FreeformConfig: "new config"},
					"serial3": {DeployThisAttachment: false, FreeformConfig: "new config"},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: false},
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        true,
			expectedFlagsGlobal: Update | Deploy | DeployAll,
			expectedFlagsAttach: map[string]uint16{
				"serial0": NoChange,
				"serial1": Update | Deploy,
				"serial2": Update | Deploy,
				"serial3": Update | Deploy,
			},
		},
		{
			name: "Modify_attachments_NwDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: false},
					"serial1": {DeployThisAttachment: false, FreeformConfig: "new config"},
					"serial2": {DeployThisAttachment: false, FreeformConfig: "new config"},
					"serial3": {DeployThisAttachment: false, FreeformConfig: "new config"},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: false},
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        false,
			expectedFlagsGlobal: Update | Deploy,
			expectedFlagsAttach: map[string]uint16{
				"serial0": NoChange,
				"serial1": Update | Deploy,
				"serial2": Update | Deploy,
				"serial3": Update | Deploy,
			},
		},
		{
			name: "Modify_attachments_NwDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: false},
					"serial1": {DeployThisAttachment: false, FreeformConfig: "new config"},
					"serial2": {DeployThisAttachment: false, FreeformConfig: "new config"},
					"serial3": {DeployThisAttachment: false, FreeformConfig: "new config"},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: false},
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        false,
			expectedFlagsGlobal: Update | Deploy | DeployAll,
			expectedFlagsAttach: map[string]uint16{
				"serial0": Deploy,
				"serial1": Update | Deploy,
				"serial2": Update | Deploy,
				"serial3": Update | Deploy,
			},
		},
		{
			name: "Modify_attachments_AttachDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: true},
					"serial1": {DeployThisAttachment: true, FreeformConfig: "new config"},
					"serial2": {DeployThisAttachment: true, FreeformConfig: "new config"},
					"serial3": {DeployThisAttachment: true, FreeformConfig: "new config"},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: true},
					"serial1": {DeployThisAttachment: true},
					"serial2": {DeployThisAttachment: true},
					"serial3": {DeployThisAttachment: true},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        false,
			expectedFlagsGlobal: Update | Deploy,
			expectedFlagsAttach: map[string]uint16{
				"serial0": Deploy,
				"serial1": Update | Deploy,
				"serial2": Update | Deploy,
				"serial3": Update | Deploy,
			},
		},
		{
			name: "Modify_attachments_AttachDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: true},
					"serial1": {DeployThisAttachment: true, FreeformConfig: "new config"},
					"serial2": {DeployThisAttachment: true, FreeformConfig: "new config"},
					"serial3": {DeployThisAttachment: true, FreeformConfig: "new config"},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial0": {DeployThisAttachment: true},
					"serial1": {DeployThisAttachment: true},
					"serial2": {DeployThisAttachment: true},
					"serial3": {DeployThisAttachment: true},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        false,
			expectedFlagsGlobal: Update | Deploy,
			expectedFlagsAttach: map[string]uint16{
				"serial0": NoChange,
				"serial1": Update | Deploy,
				"serial2": Update | Deploy,
				"serial3": Update | Deploy,
			},
		},
		{
			name: "Unchanged_Attachment_AttachDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: true},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: true},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        false,
			expectedFlagsGlobal: NoChange,
			expectedFlagsAttach: map[string]uint16{
				"serial1": NoChange,
			},
		},
		{
			name: "Unchanged_Attachment_AttachDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: true},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: true},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        false,
			expectedFlagsGlobal: Deploy,
			expectedFlagsAttach: map[string]uint16{
				"serial1": Deploy,
				"serial2": NoChange,
				"serial3": NoChange,
			},
		},
		{
			name: "Unchanged_Attachment_GlobalDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        true,
			expectedFlagsGlobal: NoChange,
			expectedFlagsAttach: map[string]uint16{
				"serial1": NoChange,
				"serial2": NoChange,
				"serial3": NoChange,
			},
		},
		{
			name: "Unchanged_Attachment_GlobalDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        true,
			expectedFlagsGlobal: DeployAll | Deploy,
			expectedFlagsAttach: map[string]uint16{
				"serial1": Deploy,
				"serial2": Deploy,
				"serial3": Deploy,
			},
		},
		{
			name: "Unchanged_Attachment_NwDeploy_NwUnModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        false,
			expectedFlagsGlobal: NoChange,
			expectedFlagsAttach: map[string]uint16{
				"serial1": NoChange,
				"serial2": NoChange,
				"serial3": NoChange,
			},
		},
		{
			name: "Unchanged_Attachment_NwDeploy_NwModified",
			plan: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
					"serial3": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          true,
			globalDeploy:        false,
			expectedFlagsGlobal: DeployAll | Deploy,
			expectedFlagsAttach: map[string]uint16{
				"serial1": Deploy,
				"serial2": Deploy,
				"serial3": Deploy,
			},
		},
		{
			name: "NwDeployModified",
			plan: &resource_networks.NDFCNetworksValue{
				NetworkName:       "test_net",
				DeployAttachments: true,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        false,
			expectedFlagsGlobal: DeployAll,
			expectedFlagsAttach: map[string]uint16{
				"serial1": NoChange,
			},
		},
		{
			// No change here
			name: "GlobalDeployModified",
			plan: &resource_networks.NDFCNetworksValue{
				NetworkName:       "test_net",
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        true,
			expectedFlagsGlobal: NoChange,
			expectedFlagsAttach: map[string]uint16{
				"serial1": NoChange,
			},
		},
		{
			name: "AttachmentDeployModified",
			plan: &resource_networks.NDFCNetworksValue{
				NetworkName:       "test_net",
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: true},
					"serial2": {DeployThisAttachment: false},
				},
			},
			state: &resource_networks.NDFCNetworksValue{
				DeployAttachments: false,
				Attachments: map[string]resource_network_attachments.NDFCAttachmentsValue{
					"serial1": {DeployThisAttachment: false},
					"serial2": {DeployThisAttachment: false},
				},
			},
			vpcPairMap:          map[string]string{},
			nwModified:          false,
			globalDeploy:        false,
			expectedFlagsGlobal: Deploy,
			expectedFlagsAttach: map[string]uint16{
				"serial1": Deploy,
				"serial2": NoChange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NDFC{}
			actionFlag := c.checkNwAttachmentsAction(context.Background(), tt.plan, tt.state, tt.vpcPairMap, tt.nwModified, tt.globalDeploy)
			require.Equal(t, tt.expectedFlagsGlobal, actionFlag)
			for serial, attach := range tt.plan.Attachments {
				t.Logf("Checking Serial: %s %v %v", serial, attach.UpdateAction, tt.expectedFlagsAttach[serial])
				require.Equal(t, tt.expectedFlagsAttach[serial], attach.UpdateAction)
			}
		})
	}
}
