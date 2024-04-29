package client

type ConfigClientList struct {
	Google Google `yaml:"google"`
}

type Google struct {
	ProjectID string `yaml:"projectId"`
}
