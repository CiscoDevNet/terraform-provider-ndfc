# =============================================================================
# Step 1: Create resource blocks for import
# =============================================================================

# MSD Fabric resource block
resource "ndfc_fabric_vxlan_msd" "msd_fabric" {
  fabric_name = "msd_fabric"
  deploy      = false
}

# Parent VRF resource block
resource "ndfc_vrfs" "msd_parent_vrfs" {
  fabric_name = "msd_fabric"
  vrfs        = {}
}

# Child fabric 1 VRF resource block
resource "ndfc_vrfs" "child_fabric1_vrfs" {
  fabric_name = "child_fabric1"
  vrfs        = {}
}

# Child fabric 2 VRF resource block
resource "ndfc_vrfs" "child_fabric2_vrfs" {
  fabric_name = "child_fabric2"
  vrfs        = {}
}

# Parent Network resource block
resource "ndfc_networks" "msd_parent_networks" {
  fabric_name = "msd_fabric"
  networks    = {}
}

# Child fabric 1 Network resource block
resource "ndfc_networks" "child_fabric1_networks" {
  fabric_name = "child_fabric1"
  networks    = {}
}

# Child fabric 2 Network resource block
resource "ndfc_networks" "child_fabric2_networks" {
  fabric_name = "child_fabric2"
  networks    = {}
}
