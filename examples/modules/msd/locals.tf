locals {
  # Group switches by their child fabric
  switches_by_child = {
    for child_name in keys(var.child_fabrics) : child_name => [
      for serial, fabric in var.switch_fabric_map : serial if fabric == child_name
    ]
  }

  # Parent-level VRF attachments: ALL switches, with fabric field
  parent_vrf_attach = {
    for serial, child_name in var.switch_fabric_map : serial => {
      deploy_this_attachment = true
      fabric                 = child_name
    }
  }

  # Parent-level network attachments: ALL switches, with fabric field
  parent_net_attach = {
    for serial, child_name in var.switch_fabric_map : serial => {
      deploy_this_attachment = true
      fabric                 = child_name
    }
  }

  # All serial numbers (for deploy)
  all_serial_numbers = keys(var.switch_fabric_map)
}
