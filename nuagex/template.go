package nuagex

import (
	"encoding/json"
	"log"
)

// Template represents NuageX Lab template
type Template struct {
	ID   string `yaml:"_id" json:"_id"`
	Name string
	Tags []string
}

// GetTemplates retrives Lab JSON object
func GetTemplates(u *User, id string) ([]*Template, error) {
	URL := buildURL("/templates?limit=100")
	b, _, err := SendHTTPRequest("GET", URL, u.Token, nil)
	// fmt.Printf("%s", b)
	if err != nil {
		log.Fatal(err)
	}
	var t []*Template
	json.Unmarshal(b, &t)
	return t, nil
}
