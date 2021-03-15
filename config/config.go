package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Listen_Addr      string   `yaml:"listen_addr"`
	TCP_Enabled      bool     `yaml:"tcp_enabled"`
	UDP_Enabled      bool     `yaml:"udp_enabled"`
	Upstream_Servers []string `yaml:"upstream_servers"`
}

var AppConfig *Config

func LoadConfig() {

	viper.SetEnvPrefix("PXY")

	viper.SetDefault("listen_addr", ":53")
	viper.SetDefault("udp_enabled", true)
	viper.SetDefault("upstream_servers", []string{"1.1.1.1", "8.8.8.8", "8.8.4.4"})
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Cannot read configuration file: %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Fatal error unmarshal configuration: %s \n", err)
	}
}
