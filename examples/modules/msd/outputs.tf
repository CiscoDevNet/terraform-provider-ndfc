output "msd_fabric_name" {
  description = "Name of the MSD parent fabric"
  value       = ndfc_fabric_vxlan_msd.parent.fabric_name
}

output "msd_fabric_id" {
  description = "Terraform ID of the MSD parent fabric"
  value       = ndfc_fabric_vxlan_msd.parent.id
}

output "child_fabric_names" {
  description = "List of child fabric names"
  value       = keys(ndfc_fabric_vxlan_evpn.child)
}

output "child_fabric_ids" {
  description = "Map of child fabric name to Terraform ID"
  value       = { for k, v in ndfc_fabric_vxlan_evpn.child : k => v.id }
}

output "switches_by_child" {
  description = "Switches grouped by child fabric name"
  value       = local.switches_by_child
}
