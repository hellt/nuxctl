package nuagex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// Lab defines a NuageX environment
type Lab struct {
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

// LabResponse : NuageX Lab response JSON object mapping
type LabResponse struct {
	ID       string `json:"_id"`
	Name     string
	Password string
	Status   string
}

// Conf loads nuagex lab configuration from a YAML file
func (c *Lab) Conf(fn string) *Lab {
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

// CreateLab : Create a Lab in NuageX
func CreateLab(u *User, reqb []byte) (LabResponse, *http.Response, error) {
	URL := buildURL("/labs")
	b, r, err := SendHTTPRequest("POST", URL, u.Token, reqb)
	if err != nil {
		return LabResponse{}, r, err
	}
	var lr LabResponse
	json.Unmarshal(b, &lr)
	if r.StatusCode != 200 {
		var eresp ErrorResponse
		json.Unmarshal(b, &eresp)
		log.Fatalf("Failed to create lab. Reason: %s", eresp.Message)
	}
	return lr, r, nil
}

// DumpLab retrives Lab JSON object
func DumpLab(u *User, id string) (Lab, *http.Response, error) {
	URL := buildURL(fmt.Sprintf("/labs/%v?expand=true", id))
	b, r, err := SendHTTPRequest("GET", URL, u.Token, nil)
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != 200 {
		var eresp ErrorResponse
		json.Unmarshal(b, &eresp)
		log.Fatalf("Failed to dump the lab. Reason: %s", eresp.Message)
	}
	var result Lab
	json.Unmarshal(b, &result)
	return result, r, nil
}
