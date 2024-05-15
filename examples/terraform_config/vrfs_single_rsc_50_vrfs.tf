
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

resource "ndfc_vrf_bulk" "vrf_test" {

    fabric_name = "test_evpn_vxlan"
    deploy_all_attachments = false
    vrfs = {
        "vrf_acc_1" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_10" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_11" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_12" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_13" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_14" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_15" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_16" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_17" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_18" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_19" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_2" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_20" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_21" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_22" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_23" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_24" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_25" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_26" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_27" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_28" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_29" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_3" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_30" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_31" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_32" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_33" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_34" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_35" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_36" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_37" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_38" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_39" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_4" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_40" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_41" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_42" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_43" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_44" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_45" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_46" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_47" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_48" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_49" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_5" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_50" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_6" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_7" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_8" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
        "vrf_acc_9" : {
            deploy_attachments = false
            attach_list = {
                "9FE076D8EJL" : {
                    serial_number = "9FE076D8EJL"
                    deploy_this_attachment = true
                }
                
                "9QBCTIN0FMY" : {
                    serial_number = "9QBCTIN0FMY"
                    deploy_this_attachment = true
                }
                
                "9TQYTJSZ1VJ" : {
                    serial_number = "9TQYTJSZ1VJ"
                    deploy_this_attachment = true
                }
                
            }
        }
    
    }
  
}
