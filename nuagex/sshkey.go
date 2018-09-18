package nuagex

// SSHKey represents NuageX SSH key definition
type SSHKey struct {
	Name string `yaml:"name" json:"name"`
	Key  string `yaml:"key" json:"key"`
}
