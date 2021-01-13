package goft

import "gopkg.in/yaml.v2"

type ServerConfig struct {
	Port int32
	Name string
}
type SysConfig struct {
	Server *ServerConfig
	Config map[string]interface{}
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{Port: 8080, Name: "myweb"}}
}
func InitConfig() *SysConfig {
	config := NewSysConfig()
	if file := LoadConfigFile(); file != nil {
		err := yaml.Unmarshal(file, &config)
		Error(err)
	}
	return config
}
