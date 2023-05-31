provider "openstack" {
user_name = "admin"
password = "qwer123456"
tenant_name = "admin"
auth_url = "http://10.41.1000.243:5000/v3"
region = "RegionOne"
}

resource "openstack_networking_network_v2" "weiyuan_net" {
  name = "openstack-net"
  admin_state_up = "true"
}

resource "openstack_networking_subnet_v2" "weiyuan_subnet" {
  name = "openstack-net-subnet"
  network_id = "${openstack_networking_network_v2.weiyuan_net.id}"
  cidr = "192.168.1.0/24"
  gateway_ip = "192.168.1.254"
  ip_version = "4"
  enable_dhcp = "false"
}