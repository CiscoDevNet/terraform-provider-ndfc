//Test UpdateVRFNoDeploy add 10 more VRFs 

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

        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false

        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        }
      }

    },
    "Murali_vrf_03" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true

        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true

        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_04" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true

        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true

        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_05" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true

        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true

        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_06" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_07" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_08" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_09" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_10" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_11" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_12" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_13" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_14" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_15" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_16" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_17" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_18" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_19" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_20" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_21" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_22" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_23" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_24" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_25" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        }
      }
    },
    "Murali_vrf_26" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_27" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_28" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_29" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_30" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_31" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_32" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_33" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_34" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_35" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_36" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_37" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_38" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_39" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_40" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_41" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_42" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_43" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_44" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_45" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_46" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_47" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_48" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_49" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_50" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = true
        },
      }
    },
    "Murali_vrf_51" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_52" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_53" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_54" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_55" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_56" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_57" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_58" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_59" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    }
    "Murali_vrf_59" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_60" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_61" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_62" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_63" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_64" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_65" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_66" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_67" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_68" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_69" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    },
    "Murali_vrf_70" : {
      // vrf_name                       = "Murali_vrf_02"
      vrf_template           = "Default_VRF_Universal"
      vrf_extension_template = "Default_VRF_Extension_Universal"
      attach_list = {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        },
      }
    }
  }
}
