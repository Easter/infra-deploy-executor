package instance

import (
	"infra-deploy-executor/provider"
	"strconv"
)

type Instance struct {
	provider.Provider
	Name        string
	Num         int
	Image_name  string
	Net_name    string
	Flavor_name string
}

func (instance *Instance) Getconfiginfo(provider provider.Provider, vmconfig map[string]map[string]string) Instance {
	name := vmconfig["instance"]["name"]
	num, _ := strconv.Atoi(vmconfig["instance"]["num"])
	image_name := vmconfig["image"]["name"]
	net_name := vmconfig["network"]["name"]
	flavor_name, _ := vmconfig["instance_flavor"]["name"]

	return Instance{provider, name, num, image_name, net_name, flavor_name}
}

func (flavor *Instance) Init(tf_path string) bool {
	return true
}

func (flavor *Instance) Apply(tf_path string) bool {
	return true
}

func (flavor *Instance) Destroy(tf_path string) bool {
	return true
}
