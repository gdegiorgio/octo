package pkg

type Package struct {
	Name         string     `yaml:"name"`
	Version      string     `yaml:"version"`
	Description  string     `yaml:"description"`
	Homepage     string     `yaml:"homepage,omitempty"`
	License      string     `yaml:"license,omitempty"`
	Dependencies []string   `yaml:"dependencies,omitempty"`
	Platforms    []Platform `yaml:"platforms"`
}

type Platform struct {
	OS       string `yaml:"os"`
	Arch     string `yaml:"arch"`
	URL      string `yaml:"url"`
	Checksum string `yaml:"checksum,omitempty"`
	Bin      string `yaml:"bin"`
	Extract  struct {
		Format          string `yaml:"format"`
		StripComponents int    `yaml:"strip_components,omitempty"`
	} `yaml:"extract"`
	Hooks struct {
		PreInstall  string `yaml:"pre_install,omitempty"`
		PostInstall string `yaml:"post_install,omitempty"`
	} `yaml:"hooks,omitempty"`
}
