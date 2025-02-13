package goft

import (
	"gopkg.in/yaml.v2"
)

type UserConfig map[interface{}]interface{}
//递归读取用户配置文件
func GetConfigValue(m UserConfig,prefix []string,index int) interface{}  {
	key:=prefix[index]
	if v,ok:=m[key];ok{
		if index==len(prefix)-1{ //到了最后一个
			return v
		}else{
			index=index+1
			if mv,ok:=v.(UserConfig);ok{ //值必须是UserConfig类型
				return GetConfigValue(mv,prefix,index)
			}else{
				return  nil
			}
		}
	}
	return  nil
}
type ServerConfig struct {
	Port int32
	Name string
	Html string
}

//system config
type SysConfig struct {
	Server *ServerConfig
	Config UserConfig
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
func(this *SysConfig) Name() string{
	return "SysConfig"
}
