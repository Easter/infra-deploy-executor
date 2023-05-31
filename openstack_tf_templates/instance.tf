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

#resource "openstack_images_image_v2" "example_image2" {
#  name = "example_image2"
#  disk_format = "qcow2"
#  container_format = "bare"
#  visibility = "public"
#  local_file_path = "/root/cirros-0.5.1-x86_64-disk.img"
#}
#
#resource "openstack_networking_network_v2" "weiyuan_terraform_net" {
#  name = "weiyuan_terraform_net"
#  admin_state_up = "true"
#}
#
#resource "openstack_networking_subnet_v2" "weiyuan_terraform_subnet" {
#  name = "weiyuan_terraform_subnet"
#  network_id = "${openstack_networking_network_v2.weiyuan_terraform_net.id}"
#  # network_id = "457188b9-00fc-4366-9011-cb52e8173d3c"
#  cidr = "192.168.1.0/24"
#  ip_version = 4
#}
#
#resource "openstack_compute_flavor_v2" "test-terraform-flavor-mini" {
#  name = "test-terraform-flavor-mini"
#  ram = "512"
#  vcpus = 1
#  disk = "1"
#  is_public = true
#}

data "openstack_images_image_v2" image_recent {
  most_recent = true
  name = "{{.Image_name}}"
}

data "openstack_networking_network_v2" "net_recent" {
  name = "{{.Net_name}}"
}

data "openstack_compute_flavor_v2" "flavor_recent" {
  name = "{{.Flavor_name}}"
}

resource "openstack_compute_instance_v2" "test_instance2" {
  name = "{{.Name}}"
  image_id = "${data.openstack_images_image_v2.image_recent.id}"
  # flavor_id = "${openstack_compute_flavor_v2.test-terraform-flavor-mini.id}"
  flavor_id = "${data.openstack_compute_flavor_v2.flavor_recent.id}"

  network {
   #  name = "${openstack_networking_network_v2.weiyuan_terraform_net.name}"
    name = "${data.openstack_networking_network_v2.net_recent.name}"
  }
}