
terraform {
  required_providers {
    ndfc = {
      source = "registry.terraform.io/cisco/ndfc"
    }
  }
}

provider "ndfc" {
  username = "test"
  password = "test"
  host     = "https://10.78.210.161"
  insecure = true
}

resource "ndfc_vrfs" "vrf_test_1" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_1" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_2" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_2" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_3" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_3" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_4" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_4" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_5" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_5" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_6" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_6" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_7" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_7" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_8" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_8" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_9" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_9" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_10" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_10" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_11" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_11" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_12" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_12" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_13" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_13" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_14" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_14" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_15" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_15" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_16" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_16" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_17" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_17" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_18" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_18" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_19" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_19" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}

resource "ndfc_vrfs" "vrf_test_20" {

  fabric_name            = "test_evpn_vxlan"
  deploy_all_attachments = false
  vrfs = {
    "vrf_acc_20" : {
      deploy_attachments = false
      attach_list = {
        "9FE076D8EJL" : {
          serial_number          = "9FE076D8EJL"
          deploy_this_attachment = true
        }

        "9TQYTJSZ1VJ" : {
          serial_number          = "9TQYTJSZ1VJ"
          deploy_this_attachment = true
        }

      }
    }

  }

}
