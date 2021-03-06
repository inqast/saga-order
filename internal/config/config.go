package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const configFilePath = "./config.yml"

var cfg *Config

func GetConfigFromFile() *Config {
	cfg := GetConfig()
	if !cfg.loaded {
		if err := ReadConfig(configFilePath); err != nil {
			panic("Could not load configuration file")
		}
	}

	return cfg
}

type GrpcConfig struct {
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	MaxConnectionIdle int    `yaml:"maxConnectionIdle"`
	Timeout           int    `yaml:"timeout"`
	MaxConnectionAge  int    `yaml:"maxConnectionAge"`
}

func (t GrpcConfig) Address() string {
	return fmt.Sprintf("%s:%s", t.Host, t.Port)
}

type DatabaseConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Pass           string `yaml:"password"`
	Name           string `yaml:"name"`
	Migrations     string `yaml:"migrations"`
	SSLmode        string `yaml:"sslmode"`
	Driver         string `yaml:"driver"`
	ConnectRetries int    `yaml:"connectRetries"`
}

func (dbconf *DatabaseConfig) GetConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbconf.Host, dbconf.Port, dbconf.User, dbconf.Pass, dbconf.Name, dbconf.SSLmode)
}

func (dbconf *DatabaseConfig) GetMigrationsPath() string {
	return dbconf.Migrations
}

type Topics struct {
	ResetOrder  string `yaml:"resetOrder"`
	NewOrder    string `yaml:"newOrder"`
	NewReserves string `yaml:"newReserves"`
}

type ConsumerGroups struct {
	ResetOrder  string `yaml:"resetOrder"`
	NewOrder    string `yaml:"newOrder"`
	NewReserves string `yaml:"newReserves"`
}

type KafkaConfig struct {
	Brokers        string         `yaml:"brokers"`
	Topics         Topics         `yaml:"topics"`
	ConsumerGroups ConsumerGroups `yaml:"consumerGroups"`
}

type ServiceConfig struct {
	Grpc     GrpcConfig     `yaml:"grpc"`
	Database DatabaseConfig `yaml:"database"`
}

type Config struct {
	loaded  bool
	Cart    ServiceConfig `yaml:"cart"`
	Order   ServiceConfig `yaml:"order"`
	Reserve ServiceConfig `yaml:"reserve"`
	Kafka   KafkaConfig   `yaml:"kafka"`
}

func ReadConfig(filePath string) error {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	if cfg == nil {
		cfg = &Config{
			loaded: false,
		}
	}

	return cfg
}
