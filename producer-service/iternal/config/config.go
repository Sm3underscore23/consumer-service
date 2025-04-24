package config

import (
	"log"
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

type kafkaConfig struct {
	BrokerList []string `yaml:"broker_list"`
	TopicName  string   `yaml:"topic"`
}

type config struct {
	Host        string      `yaml:"host"`
	Port        string      `yaml:"port"`
	KafkaConfig kafkaConfig `yaml:"kafka"`
}

func InitMainConfig(configPath string) (config, error) {
	var mainConfig config

	_, err := os.Stat(configPath)
	if err != nil {
		return config{}, err
	}

	rowConfig, err := os.ReadFile(configPath)
	if err != nil {
		return config{}, err
	}

	err = yaml.Unmarshal(rowConfig, &mainConfig)
	if err != nil {
		return config{}, err
	}

	return mainConfig, nil
}

func (cfg *config) ServerAdressLoad() string {

	log.Printf(
		"server config succeccfuly load\n%29sAdres: %s:%s",
		cfg.Host,
		"",
		cfg.Port,
	)

	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *config) KafkaConfigLoad() ([]string, string) {

	log.Printf(
		"kafka broker list: %v\n%20swrite to topic: %s\n",
		cfg.KafkaConfig.BrokerList,
		"",
		cfg.KafkaConfig.TopicName,
	)

	return cfg.KafkaConfig.BrokerList, cfg.KafkaConfig.TopicName
}
