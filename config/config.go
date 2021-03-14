package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Listen_Addr     string `yaml:"listen_addr"`
	TCP_Enabled     bool   `yaml:"tcp_enabled"`
	UDP_Enabled     bool   `yaml:"udp_enabled"`
	Upstream_Server string `yaml:"upstream_server"`
	Upstream_Port   string `yaml:"upstream_port"`
}

var AppConfig Config

func LoadConfig() {

	viper.SetDefault("LISTEN_ADDR", ":53")
	viper.SetDefault("UDP_ENABLED", true)
	viper.SetDefault("UPSTREAM_SERVER", "1.1.1.1")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error parsing configuration file: %s", err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Fatal error unmarshal configuration: %s \n", err)
	}
}
