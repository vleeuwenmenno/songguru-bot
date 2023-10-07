package models

type Service struct {
	Name  string   `yaml:"name"`
	Icon  string   `yaml:"icon"`
	Color int      `yaml:"color"`
	Urls  []string `yaml:"urls"`
}
