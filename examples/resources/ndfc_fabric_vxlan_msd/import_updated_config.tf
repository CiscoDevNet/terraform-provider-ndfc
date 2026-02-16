# =============================================================================
# Step 3: Updated configuration with proper dependencies after import
# =============================================================================

# MSD Fabric - update with actual configuration
resource "ndfc_fabric_vxlan_msd" "msd_fabric" {
  fabric_name   = "msd_fabric"
  deploy        = false
  child_fabrics = ["child_fabric1", "child_fabric2"]
}

# Parent VRF resource - add dependency on MSD fabric
resource "ndfc_vrfs" "msd_parent_vrfs" {
  depends_on  = [ndfc_fabric_vxlan_msd.msd_fabric]
  fabric_name = ndfc_fabric_vxlan_msd.msd_fabric.fabric_name
  vrfs = {
    "vrf_01" = {
      vlan_id = 234
      attach_list = {
        "9PJFRPZZ26Z" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric1"
        }
      }
    }
  }
}

# Child fabric 1 VRFs - add dependency on parent VRF resource
resource "ndfc_vrfs" "child_fabric1_vrfs" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs]
  fabric_name = "child_fabric1"
  vrfs = {
    "vrf_01" = {
      vlan_id = 234
      attach_list = {
        "9PJFRPZZ26Z" = { deploy_this_attachment = true }
      }
    }
  }
}

# Child fabric 2 VRFs - add dependency on parent VRF resource
resource "ndfc_vrfs" "child_fabric2_vrfs" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs]
  fabric_name = "child_fabric2"
  vrfs = {
    "vrf_01" = {
      vlan_id = 234
      attach_list = {
        "92Z298RRNPO" = { deploy_this_attachment = true }
      }
    }
  }
}

# Parent Network resource - add dependency on ALL VRF resources
resource "ndfc_networks" "msd_parent_networks" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs, ndfc_vrfs.child_fabric1_vrfs, ndfc_vrfs.child_fabric2_vrfs]
  fabric_name = ndfc_fabric_vxlan_msd.msd_fabric.fabric_name
  networks = {
    "nw_01" = {
      vrf_name   = "vrf_01"
      network_id = 31001
      attach_list = {
        "9PJFRPZZ26Z" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric1"
        }
      }
    }
  }
}

# Child fabric 1 Networks - add dependency on parent Network resource
resource "ndfc_networks" "child_fabric1_networks" {
  depends_on  = [ndfc_networks.msd_parent_networks]
  fabric_name = "child_fabric1"
  networks = {
    "nw_01" = {
      vrf_name   = "vrf_01"
      network_id = 31001
      attach_list = {
        "9PJFRPZZ26Z" = { deploy_this_attachment = true }
      }
    }
  }
}

# Child fabric 2 Networks - add dependency on parent Network resource
resource "ndfc_networks" "child_fabric2_networks" {
  depends_on  = [ndfc_networks.msd_parent_networks]
  fabric_name = "child_fabric2"
  networks = {
    "nw_01" = {
      vrf_name   = "vrf_01"
      network_id = 31001
      attach_list = {
        "92Z298RRNPO" = { deploy_this_attachment = true }
      }
    }
  }
}
