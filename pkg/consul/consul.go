package consul

import (
	"fmt"
	"gin-demo/global"
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

func Registra() {
	conf := consulapi.DefaultConfig()
	conf.Address = global.ConsulSetting.Address
	client, err := consulapi.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = global.ServiceSetting.Name
	registration.Name = global.ServiceSetting.Name
	registration.Port = global.ServiceSetting.Port
	registration.Tags = nil
	registration.Address = global.ServiceSetting.Host
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP: fmt.Sprintf("HTTP://%s:%d%s", registration.Address, registration.Port, "/ping"),
		//HTTP:                           "https://www.baidu.com/",
		Timeout:                        "1s",
		Interval:                       "1s",
		DeregisterCriticalServiceAfter: "10s",
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}

}
