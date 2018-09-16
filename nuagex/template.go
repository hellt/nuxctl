package nuagex

import (
	"encoding/json"
	"log"
	"net/http"
)

// TemplateShort represents NuageX Template with just the fields needed to list-tempalates command
type TemplateShort struct {
	ID   string `yaml:"_id" json:"_id"`
	Name string
	Tags []string
}

// Template defines a NuageX environment
type Template struct {
	ID string `yaml:"_id,omitempty" json:"_id"`
	// `template` field is not present in the template structure
	// its injected here to assign ID value to this field, making the resulting
	// YAML to be suitable for lab definition
	Template string    `yaml:"template,omitempty"`
	Name     string    `yaml:"name" json:"name"`
	Services []Service `yaml:"services" json:"services"`
	Networks []Network `yaml:"networks" json:"networks"`
	Servers  []Server  `yaml:"servers" json:"servers"`
}

// GetTemplates retrives Lab JSON object
func GetTemplates(u *User, id string) ([]*TemplateShort, error) {
	URL := buildURL("/templates?limit=100")
	b, _, err := SendHTTPRequest("GET", URL, u.Token, nil)
	if err != nil {
		log.Fatal(err)
	}
	var t []*TemplateShort
	json.Unmarshal(b, &t)
	return t, nil
}

// GetTemplate retrives a single Template
func GetTemplate(u *User, id string) (*Template, *http.Response, error) {
	// URL := buildURL(fmt.Sprintf("/templates/%v?expand=true", id))
	// b, r, err := SendHTTPRequest("GET", URL, u.Token, nil)
	// local testing
	b := []byte(`{"_id":"5822cc4493ce9900018dc664","name":"Nuage Networks 5.2.2 - Ocata ML2 VSS","parameters":[],"__v":42,"allowedGroups":["employee"],"allowedOrganizations":["58263cc147057600012dc4ba","58d3ead1f0f5f100015eca20"],"draft":false,"updated":"2016-11-09T07:12:04.254Z","created":"2016-11-09T07:12:04.254Z","numDeployment":0,"services":[{"name":"vsd","type":"public","port":8443,"urlScheme":"https","_id":"5822cc4493ce9900018dc670","protocols":["tcp"],"destination":{"port":18443,"address":"10.0.0.2"}},{"name":"Openstack-Horizon","type":"public","port":80,"urlScheme":"http","_id":"5b89bfae52d68000015467e0","credentials":{"username":"admin","password":"96f81e07a8e8467d"},"protocols":["tcp"],"destination":{"address":"10.0.0.10","port":8888}},{"name":"Openstack-Console","type":"public","port":6080,"urlScheme":"","_id":"5b89bfae52d68000015467df","protocols":["tcp"],"destination":{"address":"10.0.0.10","port":6080}},{"name":"elasticsearch-api","type":"public","port":6200,"urlScheme":"","_id":"5b89c02752d68000015467f1","protocols":["tcp"],"destination":{"address":"10.0.0.5","port":6200}}],"networks":[{"name":"private","cidr":"10.0.0.0/24","_id":"5822cc4493ce9900018dc66d","dns":["10.0.0.1"],"dhcp":true}],"servers":[{"name":"vsd01","image":"vsd01_NUX_5.2.2_OCATA_VSS","flavor":"ceph.vsd","_id":"5822cc4493ce9900018dc66b","interfaces":[{"network":"private","address":"10.0.0.2","index":0,"_id":"5822cc4493ce9900018dc66c"}]},{"name":"vsc01","image":"vsc01_NUX_5.2.2_OCATA_VSS","flavor":"ceph.vsc","_id":"5822cc4493ce9900018dc669","interfaces":[{"network":"private","address":"10.0.0.3","index":0,"_id":"5822cc4493ce9900018dc66a"}]},{"name":"elastic","image":"elastic_NUX_5.2.2_OCATA_VSS","flavor":"ceph.vsd","_id":"5b89bfae52d68000015467e7","interfaces":[{"index":0,"network":"private","address":"10.0.0.5","_id":"5b89bfae52d68000015467e8"}]},{"name":"os-controller","image":"os-controller_NUX_5.2.2_OCATA_VSS","flavor":"ceph.m1.xlarge","_id":"5b89bfae52d68000015467e5","interfaces":[{"index":0,"network":"private","address":"10.0.0.10","_id":"5b89bfae52d68000015467e6"}]},{"name":"os-compute01","image":"os-compute01_NUX_5.2.2_OCATA_VSS","flavor":"ceph.m1.xlarge","_id":"5b89bfae52d68000015467e3","interfaces":[{"index":0,"network":"private","address":"10.0.0.11","_id":"5b89bfae52d68000015467e4"}]},{"name":"os-compute02","image":"os-compute02_NUX_5.2.2_OCATA_VSS","flavor":"ceph.m1.xlarge","_id":"5b89bfae52d68000015467e1","interfaces":[{"index":0,"network":"private","address":"10.0.0.12","_id":"5b89bfae52d68000015467e2"}]}],"tags":["vcs","vss","ocata","ml2","openstack"],"metadatas":[{"key":"vsp_release","value":"5.2.2","_id":"5b9c0f6bbc12f10001fb3537"},{"key":"bypass_vsd_init","value":"false","_id":"5b9c0f6bbc12f10001fb3536"},{"key":"public","value":"false","_id":"5b9c0f6bbc12f10001fb3535"},{"key":"ssh_access","value":"true","_id":"5b9c0f6bbc12f10001fb3534"},{"key":"remote_vrs","value":"true","_id":"5b9c0f6bbc12f10001fb3533"},{"key":"external_vrs","value":"true","_id":"5b9c0f6bbc12f10001fb3532"},{"key":"docker_examples","value":"true","_id":"5b9c0f6bbc12f10001fb3531"},{"key":"kvm_examples","value":"true","_id":"5b9c0f6bbc12f10001fb3530"}]}`)
	r := &http.Response{} // delete after local testing
	var err error         // delete after local testing
	if err != nil {
		log.Fatal(err)
	}
	var t *Template
	json.Unmarshal(b, &t)
	return t, r, nil
}
