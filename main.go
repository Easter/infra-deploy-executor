package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"infra-deploy-executor/flavor"
	"infra-deploy-executor/glance"
	"infra-deploy-executor/instance"
	"infra-deploy-executor/neutron"
	"infra-deploy-executor/provider"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Terraformif interface {
	// Getconfiginfo(provider.Provider, map[string]map[string]string)
	Init(string) bool
	Apply(string) bool
	Destroy(string) bool
}

func read_config(config_json_path string) map[string]map[string]string {
	content, err := ioutil.ReadFile(config_json_path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var vmconfig map[string]map[string]string

	err = json.Unmarshal(content, &vmconfig)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return vmconfig
}

func Generate_tf(server Terraformif, tf_templates string, serverdir string) (string, error) {
	tf_template, err := os.ReadFile(tf_templates)
	if err != nil {
		return "", err
	}
	flavor_tf_template := string(tf_template)
	tmpl, err := template.New("tfTemplate").Parse(flavor_tf_template)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, server)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return "", err
	}
	tf_file := buf.String()

	// return true
	// 将填充后的内容写入文件
	os.MkdirAll("./vmconfig/terraform/"+serverdir, os.ModePerm)
	file, err := os.Create("./vmconfig/terraform/" + serverdir + "/" + serverdir + ".tf")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return "", err
	}

	defer file.Close()
	_, err = file.WriteString(tf_file)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return "", err
	}
	fmt.Println("tf content written to file successfully.")
	tf_file_path := "./vmconfig/terraform/" + serverdir + "/" + serverdir + ".tf"
	fmt.Println(tf_file_path)
	return tf_file_path, nil

}

func main() {
	vmconfig_file := "./vmconfig/vmconfig.example.json"
	vmconfig := read_config(vmconfig_file)

	provider := provider.Provider{}
	provider = provider.Getconfiginfo(vmconfig)

	flavor_obj := &flavor.Flavor{}
	flavor := flavor_obj.Getconfiginfo(provider, vmconfig)
	flavor_tf_template := "./openstack_tf_templates/flavor.tf"
	tf_file_path, err := Generate_tf(&flavor, flavor_tf_template, "flavor")
	if err != nil {
		return
	} else {
		flavor.Init(tf_file_path)
		flavor.Apply(tf_file_path)
	}

	glance_obj := &glance.Glance{}
	glance := glance_obj.Getconfiginfo(provider, vmconfig)
	glance_tf_template := "./openstack_tf_templates/glance.tf"
	tf_file_path, err = Generate_tf(&glance, glance_tf_template, "glance")
	if err != nil {
		return
	} else {
		glance.Init(tf_file_path)
		glance.Apply(tf_file_path)
	}

	subnet_obj := &neutron.SubNet{}
	subnet := subnet_obj.Getconfiginfo(provider, vmconfig)
	neutron_tf_template := "./openstack_tf_templates/neutron.tf"
	tf_file_path, err = Generate_tf(&subnet, neutron_tf_template, "subnet")
	if err != nil {
		return
	} else {
		subnet.Init(tf_file_path)
		subnet.Apply(tf_file_path)
	}

	instance_obj := &instance.Instance{}
	instance := instance_obj.Getconfiginfo(provider, vmconfig)
	instance_tf_template := "./openstack_tf_templates/instance.tf"
	tf_file_path, err = Generate_tf(&instance, instance_tf_template, "instance")
	if err != nil {
		return
	} else {
		instance.Init(tf_file_path)
		instance.Apply(tf_file_path)
	}

	external_network_obj := &neutron.External_network{}
	external_network := external_network_obj.Getconfiginfo()
	external_network_tf_template := "./openstack_tf_templates/external_network.tf"
	tf_file_path, err = Generate_tf(&external_network, external_network_tf_template, "external_network")
	if err != nil {
		return
	} else {
		instance.Init(tf_file_path)
		instance.Apply(tf_file_path)
	}

	fmt.Println(provider)

}
