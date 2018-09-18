package nuagex

// Network represents NuageX network definition
type Network struct {
	Name string      `yaml:"name" json:"name"`
	Cidr string      `yaml:"cidr" json:"cidr"`
	DNS  interface{} `yaml:"dns" json:"dns"`
	Dhcp bool        `yaml:"dhcp" json:"dhcp"`
}
