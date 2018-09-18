package nuagex

// Service represents NuageX service (port forwarding rule on Jumpbox)
type Service struct {
	Name        string   `yaml:"name" json:"name"`
	Type        string   `yaml:"type" json:"type"`
	Port        int      `yaml:"port" json:"port"`
	URLScheme   string   `yaml:"urlScheme,omitempty" json:"urlScheme,omitempty"`
	Protocols   []string `yaml:"protocols" json:"protocols"`
	Destination struct {
		Port    int    `yaml:"port" json:"port"`
		Address string `yaml:"address" json:"address"`
	} `yaml:"destination" json:"destination"`
}
