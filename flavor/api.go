package flavor

import (
	"infra-deploy-executor/provider"
	"strconv"
)

type Flavor struct {
	provider.Provider
	// flavor的结构体
	Name      string
	Ram       int
	Vcpus     int
	Disk      int
	Swap      int
	Is_public bool
}

func (flavor *Flavor) Getconfiginfo(provider provider.Provider, vmconfig map[string]map[string]string) Flavor {
	name := vmconfig["instance_flavor"]["name"]
	ram, _ := strconv.Atoi(vmconfig["instance_flavor"]["ram"])

	vcpus, _ := strconv.Atoi(vmconfig["instance_flavor"]["vcpus"])
	disk, _ := strconv.Atoi(vmconfig["instance_flavor"]["disk"])
	swap, _ := strconv.Atoi(vmconfig["instance_flavor"]["swap"])
	is_public, _ := strconv.ParseBool(vmconfig["instance_flavor"]["is_pudlic"])

	return Flavor{provider, name, ram, vcpus, disk, swap, is_public}
}

func (flavor *Flavor) Init(tf_path string) bool {
	return true
}

func (flavor *Flavor) Apply(tf_path string) bool {
	return true
}

func (flavor *Flavor) Destroy(tf_path string) bool {
	return true
}
