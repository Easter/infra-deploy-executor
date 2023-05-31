package neutron

type External_network struct {
	Name             string
	Admin_state_up   bool
	External         bool
	Shared           bool
	Physical_network string
	Network_type     string
}

func (external_network *External_network) Getconfiginfo() External_network {
	// var segments map[string]string = map[string]string{"Physical_network": "test", "Network_type": "test"}
	return External_network{Name: "test_external_network", Admin_state_up: true, External: true, Shared: true, Physical_network: "test", Network_type: "test"}
}

func (external_network *External_network) Init(tf_path string) bool {
	return true
}

func (external_network *External_network) Apply(tf_path string) bool {
	return true
}

func (external_network *External_network) Destroy(tf_path string) bool {
	return true
}
