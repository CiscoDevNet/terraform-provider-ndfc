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
  value       = keys(var.child_fabrics)
}

output "switches_by_child" {
  description = "Switches grouped by child fabric name"
  value       = local.switches_by_child
}
