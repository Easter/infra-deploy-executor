package provider

type Provider struct {
	User_name   string
	Password    string
	Tenant_name string
	Auth_url    string
	Region      string
}

func (provider *Provider) Getconfiginfo(vmconfig map[string]map[string]string) Provider {
	name := vmconfig["provider"]["user_name"]
	password, _ := vmconfig["provider"]["password"]
	tenant_name, _ := vmconfig["provider"]["tenant_name"]
	auth_url, _ := vmconfig["provider"]["auth_url"]
	region, _ := vmconfig["provider"]["region"]

	return Provider{name, password, tenant_name, auth_url, region}
}
