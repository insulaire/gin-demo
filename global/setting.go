package global

import (
	"gin-demo/internal/config"
	"gin-demo/pkg/ip"
	"log"
)

var (
	ServiceSetting *config.Service
	ConsulSetting  *config.Consul
)

func init() {
	if err := setupServiceSetting(); err != nil {
		log.Fatal(err)
	}
	if err := setupConsulSetting(); err != nil {
		log.Fatal(err)
	}
}

func setupServiceSetting() error {
	ServiceSetting = new(config.Service)
	str_ip, err := ip.GetLocalIP()
	if err != nil {
		return err
	}
	ServiceSetting.Host = str_ip
	ServiceSetting.Port = 9999
	ServiceSetting.Name = "Test"
	return nil
}

func setupConsulSetting() error {
	ConsulSetting = new(config.Consul)
	ConsulSetting.Address = "http://localhost:8500"
	return nil
}
