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
resource "openstack_compute_flavor_v2" "terraform-flavor" {
  name = "k8s-host-flavor"
  ram = "8192"
  vcpus = "4"
  disk = "128"
  is_public = "true"
}
