package models

import (
  "os"
  "gopkg.in/yaml.v2"

  
)


// 使用 YAML 配置文件中的键值对进行初始化
type MySQLConfig struct {
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
}


// 加一层
type GlobalConfig struct{
   MySQLConf MySQLConfig `yaml:"mysql"`
}

// method
func (c *GlobalConfig) getConf() *GlobalConfig{
    yamlFile, err := os.ReadFile("config/application.yaml")
    if err != nil{
      panic(err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if(err != nil){
        panic(err)
    }
    return c
    
}

func InitDB()(err error){
    var c GlobalConfig
    conf := c.getConf()
    err = InitMySql(conf.MySQLConf)
    if(err != nil){
        panic(err)
    }
    return nil
    
}

func Close(){
    CloseMySQL()
}