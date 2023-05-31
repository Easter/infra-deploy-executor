terraform {
required_version = ">= 0.14.0"
  required_providers {
    openstack = {
      source  = "hashicorp/openstack"
      version = ">= 1.48.0"
    }
  }
}

provider "openstack" {
user_name = "{{.User_name}}"
password = "{{.Password}}"
tenant_name = "{{.Tenant_name}}"
auth_url = "{{.Auth_url}}"
region = "{{.Region}}"
}
# 创建资源
resource "openstack_images_image_v2" "example_image2" {
 name = "{{.Name}}"
 disk_format = "{{.Disk_format}}"
 container_format = "{{.Container_format}}"
 visibility = "{{.Visibility}}"
 local_file_path = "{{.Local_file_path}}"
}
# 查询资源
#data "openstack_images_image_v2" "images" {
#  most_recent = true
#  name = "example_image2"
#}

# output "image_names" {
# value = [for image in data.openstack_images_image_v2.images : image.name]
#   value = data.openstack_images_image_v2.images.id
# }