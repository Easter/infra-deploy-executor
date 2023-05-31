provider "openstack" {
user_name = "{{.User_name}}"
password = "{{.Password}}"
tenant_name = "{{.Tenant_name}}"
auth_url = "{{.Auth_url}}"
region = "{{.Region}}"
}

resource "openstack_networking_network_v2" "weiyuan_net" {
  name = "{{.Name}}"
  admin_state_up = "{{.Admin_state_up}}"
}

resource "openstack_networking_subnet_v2" "weiyuan_subnet" {
  name = "{{.Subnet_name}}"
  network_id = "${openstack_networking_network_v2.weiyuan_net.id}"
  cidr = "{{.Cidr}}"
  gateway_ip = "{{.Gateway_ip}}"
  ip_version = "{{.Ip_version}}"
  enable_dhcp = "{{.Enable_dhcp}}"
}