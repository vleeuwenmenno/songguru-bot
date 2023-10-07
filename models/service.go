package models

type Service struct {
	Name  string   `yaml:"name"`
	Icon  string   `yaml:"icon"`
	Color string   `yaml:"color"`
	Urls  []string `yaml:"urls"`
}
