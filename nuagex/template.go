package nuagex

import (
	"encoding/json"
	"fmt"
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
	URL := buildURL(fmt.Sprintf("/templates/%v?expand=true", id))
	b, r, err := SendHTTPRequest("GET", URL, u.Token, nil)
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != 200 {
		var eresp ErrorResponse
		json.Unmarshal(b, &eresp)
		log.Fatalf("Failed to dump the lab. Reason: %s", eresp.Message)
	}
	var t *Template
	json.Unmarshal(b, &t)
	return t, r, nil
}
