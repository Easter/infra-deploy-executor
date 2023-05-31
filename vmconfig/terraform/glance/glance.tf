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
user_name = "admin"
password = "qwer123456"
tenant_name = "admin"
auth_url = "http://10.41.1000.243:5000/v3"
region = "RegionOne"
}
# 创建资源
resource "openstack_images_image_v2" "example_image2" {
 name = "k8s-host-image"
 disk_format = "k8s-host-image"
 container_format = "bare"
 visibility = "public"
 local_file_path = "/root/kylin-sp3.qcow2"
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