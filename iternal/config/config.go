package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MainConfigLoader interface {
	DbConfigLoad() string
	KafkaConfigLoad() kafkaConfig
}

type dbConfig struct {
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbSSL      string `yaml:"db_sslmode"`
}

type kafkaConfig struct {
	ConsumerGroup string   `yaml:"kafka_consumer_group"`
	BrokerList    []string `yaml:"broker_list"`
	TopicName     string   `yaml:"topic"`
	Assignor      string   `yaml:"assignor"`
}

type config struct {
	KafkaConfig kafkaConfig `yaml:"kafka"`
	DbConfig    dbConfig    `yaml:"db"`
}

func NewMainConfig(configPath string) (MainConfigLoader, error) {
	var mainConfig config

	_, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	rowConfig, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(rowConfig, &mainConfig)
	if err != nil {
		return nil, err
	}

	return &config{
		KafkaConfig: mainConfig.KafkaConfig,
		DbConfig:    mainConfig.DbConfig,
	}, nil
}

func (cfg *config) DbConfigLoad() string {

	log.Printf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.DbConfig.DbHost,
		cfg.DbConfig.DbPort,
		cfg.DbConfig.DbName,
		cfg.DbConfig.DbUser,
		cfg.DbConfig.DbPassword,
		cfg.DbConfig.DbSSL,
	)

	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.DbConfig.DbHost,
		cfg.DbConfig.DbPort,
		cfg.DbConfig.DbName,
		cfg.DbConfig.DbUser,
		cfg.DbConfig.DbPassword,
		cfg.DbConfig.DbSSL,
	)
}

func (cfg *config) KafkaConfigLoad() kafkaConfig {

	log.Printf(
		"cafka consumer grouop: %s\n%20skafka broker list: %v\n%20swrite to topic: %s\n",
		cfg.KafkaConfig.ConsumerGroup,
		"",
		cfg.KafkaConfig.BrokerList,
		"",
		cfg.KafkaConfig.TopicName,
	)

	return cfg.KafkaConfig
}
