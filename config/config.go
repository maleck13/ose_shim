package config

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	
	"io/ioutil"
	"os"
)

//use an interface to limit access to the config object to read only
type Configuration interface {
	GetPProfEnabled() bool
	GetDockerUser()string
	GetDockerPass()string
	GetDockerRegistry()string
}

type config struct {
	PProfEnabled  bool `json:"pprof_enabled"`
	DockerUser string `json:"-"`
	DockerPass string `json:"-"`
	DockerRegistry string `json:"-"`
}



func (c *config) GetDockerUser()string{
	return os.Getenv("docker_user")
}

func (c *config) GetDockerPass()string{
	return os.Getenv("docker_pass")
}

func (c *config) GetDockerRegistry()string{
	return os.Getenv("docker_registry")
}

func (c *config) GetPProfEnabled() bool {
	return c.PProfEnabled
}


var Conf Configuration

func SetGlobalConfig(path string) {
	Conf = &config{}
	file, err := os.Open(path)
	if nil != err {
		logrus.Panic("failed to open config file ", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if nil != err {
		logrus.Panic("failed to read config file ", err)
		return
	}
	if err = json.Unmarshal(data, Conf); err != nil {
		logrus.Panic("failed to decode config file ", err)
		return
	}
}

