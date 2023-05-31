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
resource "openstack_compute_flavor_v2" "terraform-flavor" {
  name = "{{.Name}}"
  ram = "{{.Ram}}"
  vcpus = "{{.Vcpus}}"
  disk = "{{.Disk}}"
  is_public = "{{.Is_public}}"
}
