//Test UpdateAttachmentNoDeploy add attachments without deploy to one of the VRF {

terraform {
  required_providers {
    ndfc = {
      source = "registry.terraform.io/cisco/ndfc"
    }
  }
}

provider "ndfc" {
  username = "admin"
  password = "admin!@#"
  #  password = "ins3965!"
  #  host      = "https://10.195.225.193"
  host     = "https://10.78.210.161"
  insecure = true
}

resource "ndfc_vrf_bulk" "test_evpn_vxlan_deployments1" {
  fabric_name = "test_evpn_vxlan"
  vrfs = {
    "Murali_vrf_01" : {
      //  vrf_name                       = "Murali_vrf_03"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
          vlan                   = 3000
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
          vlan                   = 3001
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }

      }
    },
    "Murali_vrf_02" : {
      // vrf_name                       = "Murali_vrf_01"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
          vlan                   = 3500
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
          vlan                   = 3500
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
          vlan                   = 3500
        }
      }

    },
    "Murali_vrf_03" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
    },
    "Murali_vrf_04" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
    },
    "Murali_vrf_05" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
    },
    "Murali_vrf_06" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
    },
    "Murali_vrf_07" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
    },
    "Murali_vrf_08" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
    },
    "Murali_vrf_09" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
    },
    "Murali_vrf_10" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"

    }
  }

}