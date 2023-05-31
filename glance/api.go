package glance

import (
	"infra-deploy-executor/provider"
)

type Glance struct {
	provider.Provider
	Name             string
	Disk_format      string
	Container_format string
	Visibility       string
	Local_file_path  string
	// glance的结构体
	// openstack glance结构体
}

func (glance *Glance) Getconfiginfo(provider provider.Provider, vmconfig map[string]map[string]string) Glance {
	name := vmconfig["image"]["name"]
	disk_format, _ := vmconfig["image"]["name"]
	container_format, _ := vmconfig["image"]["container_format"]
	visibility, _ := vmconfig["image"]["visibility"]
	local_file_path, _ := vmconfig["image"]["local_file_path"]

	return Glance{provider, name, disk_format, container_format, visibility, local_file_path}
}
func (flavor *Glance) Init(tf_path string) bool {
	return true
}

func (flavor *Glance) Apply(tf_path string) bool {
	return true
}

func (flavor *Glance) Destroy(tf_path string) bool {
	return true
}
