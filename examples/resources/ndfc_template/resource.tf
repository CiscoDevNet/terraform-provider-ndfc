
resource "ndfc_template" "test_resource_template_1" {
  template_name       = "template1"
  description         = "This is a template"
  tags                = [key1, value1]
  supported_platforms = "All"
  template_type       = "CLI"
  template_content    = ""
  content_type        = "CLI"
  template_sub_type   = "CONFIG"
}