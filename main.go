package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"

	"./nux"
)

type labConf struct {
	Name     string    `yaml:"name" json:"name"`
	Reason   string    `yaml:"reason" json:"reason"`
	Expires  time.Time `yaml:"expires" json:"expires"`
	Template string    `yaml:"template" json:"template"`
	SSHKeys  []struct {
		Name string `yaml:"name" json:"name"`
		Key  string `yaml:"key" json:"key"`
	} `yaml:"sshKeys" json:"sshKeys"`
	Services []struct {
		Name        string   `yaml:"name" json:"name"`
		Type        string   `yaml:"type" json:"type"`
		Port        int      `yaml:"port" json:"port"`
		URLScheme   string   `yaml:"urlScheme,omitempty" json:"urlScheme,omitempty"`
		Protocols   []string `yaml:"protocols" json:"protocols"`
		Destination struct {
			Port    int    `yaml:"port" json:"port"`
			Address string `yaml:"address" json:"address"`
		} `yaml:"destination" json:"destination"`
	} `yaml:"services" json:"services"`
	Networks []struct {
		Name string      `yaml:"name" json:"name"`
		Cidr string      `yaml:"cidr" json:"cidr"`
		DNS  interface{} `yaml:"dns" json:"dns"`
		Dhcp bool        `yaml:"dhcp" json:"dhcp"`
	} `yaml:"networks" json:"networks"`
	Servers []struct {
		Name       string `yaml:"name" json:"name"`
		Image      string `yaml:"image" json:"image"`
		Flavor     string `yaml:"flavor" json:"flavor"`
		Interfaces []struct {
			Index   int    `yaml:"index" json:"index"`
			Network string `yaml:"network" json:"network"`
			Address string `yaml:"address" json:"address"`
		} `yaml:"interfaces" json:"interfaces"`
	} `yaml:"servers" json:"servers"`
}

func (c *labConf) getConf(fn string) *labConf {
	fmt.Printf("Loading lab configuration from '%s' file\n", fn)
	yamlFile, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Printf("Lab Configuration Load error   #%v ", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	usrCredF := flag.String("u", "user_creds.yml", "path to the credentials YAML file")
	labConfF := flag.String("l", "lab.yml", "path to the lab definition YAML file")
	flag.Parse()
	var user nux.User

	user.LoadCredentials(*usrCredF)

	token, err := nux.UserLogin(user)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf(token)

	// // flavors, err := nux.GetAllFlavors(token)
	// flavors, _ := nux.GetAllFlavors(token)
	// fmt.Println()
	// fmt.Println()
	// fmt.Printf("%#v", &flavors)
	// tj := []byte(`{"name":"nuxctltest","reason":"Nux'em all","expires":"2019-12-31T17:13:11.278Z","template":"5980ee745a38da00012d158d","sshKeys":[{"name":"cats","key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0jLLF2c7sSUHCwFJ1cpj0mTNRfemi6XMKxAf7H4gIzs/joL18W+wlSrnHnu801bDLc2RNg8dvOvXmzjzZrKKMInWMXrzb0zjljCPumYWlI/koWAIMzpENjeRjWrB22WFVOVDW4GkphzhchFWFPSF9xxU4i3MPHn3HYZiy6ieLOvknplQDivTXcLwRjvenK35gjLPxo4nFhQWyUdzPLLCb/NUx/CGjkh10qwVfw+AWw3x3boqxwBXitPjLJ+ocfZVohAWupKuvsTNzlh8m39imI5pi9qTFT1x9YtkYsTYCy2iz2HCirqy+8BC/CPCyHuEjs8ZT9OFHHQtjB/qgzfTp"}],"servers":[{"name":"vsc4","image":"nux_vsc_5.2.3","flavor":"m1.small","interfaces":[{"index":0,"network":"private","address":"10.0.0.11"}]}]}`)

	var c labConf
	c.getConf(*labConfF)

	j, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Sending request to create a lab...\n")
	nux.CreateLab(token, j)
	fmt.Printf("Lab has been successfully queued for creation!\n")
	// fmt.Println(string(j))
	// b, _ := json.Marshal(c)

	// fmt.Printf("%#v", &c)
	// fmt.Println()
	// fmt.Printf("%#v", string(b))
}
