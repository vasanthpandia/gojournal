package config

import (
	"time"
)

type Config struct {
	MongoConfig *MongoConfig
	Token *Token
}

type MongoConfig struct {
	Url string
	Database string
}

type Token struct {
	Key []byte
	Validity time.Duration
}

type ServerConfig struct {
	DBConnection *MongoConnection
	Token	*Token
}

func InitDefaults() *Config {
	mConfig := MongoConfig {
		Url: "mongodb://localhost:27017",
		Database: "gojournal",
	}

	dur, _ := time.ParseDuration("24h")

	token := Token {
		Key: []byte("DEFAULTKEY"),
		Validity: dur,
	}

	return &Config {
		MongoConfig: &mConfig,
		Token: &token,
	}
}

func getConfigFor(env string) *Config {
	//TODO - Fetch config from appropriate toml file
	return &Config{}
}

func GetServerConfig(env string) *ServerConfig {
	srvConfig := &ServerConfig {}

	var config *Config

	if env == "development" {
		config = InitDefaults()
	} else {
		config = getConfigFor(env)
	}

	connection, err := GetMongoConnection(config.MongoConfig)
	if err != nil {
		panic(err)
	}

	srvConfig.DBConnection = connection
	srvConfig.Token = config.Token

	return srvConfig
}
