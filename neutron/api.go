package neutron

import (
	"infra-deploy-executor/provider"
	"strconv"
)

type Neutron struct {
	// neutron的结构体
	Name           string
	Admin_state_up bool
	// Project          string
	// Provider_type    string
	// shared           bool
	// External_network bool
	// External_subnet  bool
}

type SubNet struct {
	provider.Provider
	Neutron
	Subnet_name string
	Cidr        string
	Ip_version  int
	Gateway_ip  string
	Enable_dhcp bool
	// Dns_nameservers []string
}

func (subnet *SubNet) Getconfiginfo(provider provider.Provider, vmconfig map[string]map[string]string) SubNet {
	name := vmconfig["network"]["name"]
	admin_state_up, _ := strconv.ParseBool(vmconfig["network"]["admin_state_up"])
	net := Neutron{name, admin_state_up}
	subnet_name := vmconfig["subnet"]["subnet_name"]
	cidr, _ := vmconfig["subnet"]["cidr"]
	ip_version, _ := strconv.Atoi(vmconfig["subnet"]["ip_version"])
	gateway_ip, _ := vmconfig["subnet"]["gateway_ip"]
	enable_dhcp, _ := strconv.ParseBool(vmconfig["subnet"]["enable_dhcp"])
	return SubNet{provider, net, subnet_name, cidr, ip_version, gateway_ip, enable_dhcp}
	//return Glance{provider, name, disk_format, container_format, visibility, local_file_path}
}

func (flavor *Neutron) Init(tf_path string) bool {
	return true
}

func (flavor *Neutron) Apply(tf_path string) bool {
	return true
}

func (flavor *Neutron) Destroy(tf_path string) bool {
	return true
}

func (flavor *SubNet) Init(tf_path string) bool {
	return true
}

func (flavor *SubNet) Apply(tf_path string) bool {
	return true
}

func (flavor *SubNet) Destroy(tf_path string) bool {
	return true
}
