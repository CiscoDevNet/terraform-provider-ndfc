//Test 1: CreateAndDeploy 10 VRFS with one of them having 3 attachments and deploy 
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

locals {
  vrfs = ndfc_vrf_bulk.test_evpn_vxlan_deployment_vrf1.vrfs
  network_vrfs = flatten([
    for key, value in local.vrfs :
    [key]
  ])
}

resource "ndfc_vrf_bulk" "test_evpn_vxlan_deployment_vrf1" {
  fabric_name = "test_evpn_vxlan"
  vrfs = {
    "Murali_vrf_01" : {
      attach_list : {
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
    "Murali_vrf_02" : {
      attach_list : {
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
    "Murali_vrf_03" : {
      attach_list : {
        "9FE076D8EJL" : {
          deploy_this_attachment = false
          vlan                   = 3200
        },
        "9TQYTJSZ1VJ" : {
          deploy_this_attachment = false
          vlan                   = 3100
        },
        "9QBCTIN0FMY" : {
          deploy_this_attachment = false
        }
      }

    },
    "Murali_vrf_04" : {

    },
    "Murali_vrf_05" : {
      "9FE076D8EJL" : {
        deploy_this_attachment = false
      },
      "9TQYTJSZ1VJ" : {
        deploy_this_attachment = false
      },
      "9QBCTIN0FMY" : {
        deploy_this_attachment = false
      }

    },
    "Murali_vrf_06" : {

    },
    "Murali_vrf_07" : {

    },
    "Murali_vrf_08" : {

    },
    "Murali_vrf_09" : {
      vrf_id = 51109
    },
    "Murali_vrf_10" : {
      vrf_id = 51110
    }
  }
}

resource "ndfc_networks" "test_evpn_vxlan_deployment_nw1" {
  fabric_name = "test_evpn_vxlan"
  networks = {
    "Murali_nw_01" : {
      vrf_name              = local.network_vrfs[0]
      network_id            = 30100
      gateway_ipv4_address  = "192.168.1.1/24"
      gateway_ipv6_address  = "2001:db8:1::1/64"
      vlan_id               = 2
      vlan_name             = "vlan2"
      layer2_only           = false
      interface_description = "Nice interface description"
      mtu                   = 9100
      secondary_gateway_1   = "192.168.2.1/24"
      secondary_gateway_2   = "192.168.3.1/24"
      secondary_gateway_3   = "192.168.4.1/24"
      secondary_gateway_4   = "192.168.5.1/24"
      arp_suppression       = false
      ingress_replication   = false
      multicast_group       = "239.1.1.2"
      dhcp_relay_servers = [
        {
          address = "10.1.1.1",
          vrf     = "management"
        },
        {
          address = "20.1.1.1"
          vrf     = "default"
        },
        {
          address = "20.1.1.5"
          vrf     = "default"
        }
      ]
      dhcp_relay_loopback_id = 999
      //tag                    = 1234
      l3_gatway_border       = true
      igmp_version           = 3
      attachments = {
        "9FE076D8EJL" : {
          switch_ports           = ["Ethernet1/1", "Ethernet1/2"]
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          switch_ports           = ["Ethernet1/2", "Ethernet1/1"]
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          switch_ports           = ["Ethernet1/2","Ethernet1/1"]
          deploy_this_attachment = true
        }
      }
    },
     "Murali_nw_02" : {
      vrf_name              = local.network_vrfs[1]
      vlan_id = 3
       attachments = {
        "9FE076D8EJL" : {
          switch_ports           = ["Ethernet1/3", "Ethernet1/4", "Ethernet1/5"]
          deploy_this_attachment = true
        },
        "9TQYTJSZ1VJ" : {
          switch_ports           = ["Ethernet1/4", "Ethernet1/5", "Ethernet1/3"]
          deploy_this_attachment = true
        },
        "9QBCTIN0FMY" : {
          switch_ports           = ["Ethernet1/5", "Ethernet1/3", "Ethernet1/4"]
          deploy_this_attachment = true
        }
      }
     }
  }
}

data "ndfc_vrf_bulk" "test_evpn_vxlan_deployment_vrf1" {
  fabric_name = "test_evpn_vxlan"
}


data "ndfc_networks" "test_evpn_vxlan_deployment_nw1" {
  fabric_name = "test_evpn_vxlan"
}