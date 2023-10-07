package models

type Services struct {
	StreamingServices map[string]Service `yaml:"streaming_services"`
}
