package models

type ServiceLink struct {
	Link      string      `json:"link"`
	Countries interface{} `json:"countries"`
}
