resource "openstack_networking_network_v2" "external_network" {
  name          = "test_external_network"
  admin_state_up = "true"
  external      = "true"
  shared = "true"
  segments {
    physical_network = "test"
    network_type = "test"
  }
}

resource "openstack_networking_subnet_v2" "external_subnet" {
  network_id   = "${openstack_networking_network_v2.external_network.id}"
  cidr         = "10.41.100.0/24"
  name         = "external-subnet"
  ip_version   = 4
  enable_dhcp  = false
  allocation_pool {
    start = "10.41.100.249"
    end   = "10.41.100.251"
  }
  gateway_ip = "10.41.100.254"
  dns_nameservers = ["8.8.8.8"]
}
