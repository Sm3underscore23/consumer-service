package config

//убрать
import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

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

func (cfg *config) DbConfigLoad() string {
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
	return cfg.KafkaConfig
}
