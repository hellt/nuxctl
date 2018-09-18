package nuagex

// Server represents NuageX server definition
type Server struct {
	Name       string      `yaml:"name" json:"name"`
	Image      string      `yaml:"image" json:"image"`
	Flavor     string      `yaml:"flavor" json:"flavor"`
	Interfaces []Interface `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`
}

// Interface represents NuageX server's interface definition
type Interface struct {
	Index   int    `yaml:"index" json:"index"`
	Network string `yaml:"network" json:"network"`
	Address string `yaml:"address" json:"address"`
}
