package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	MongoConfig *MongoConfig
	Token       *Token
}

type MongoConfig struct {
	Url      string
	Database string
}

type Token struct {
	Key      []byte
	Validity time.Duration
}

type ServerConfig struct {
	DBConnection *MongoConnection
	Token        *Token
}

func InitDefaults() *Config {
	mConfig := MongoConfig{
		Url:      "mongodb://localhost:27017",
		Database: "gojournal",
	}

	dur, _ := time.ParseDuration("24h")

	token := Token{
		Key:      []byte("DEFAULTKEY"),
		Validity: dur,
	}

	return &Config{
		MongoConfig: &mConfig,
		Token:       &token,
	}
}

func getConfigFor(env string) *Config {
	mConfig := MongoConfig{
		Url:      os.Getenv("MONGODB_URI"),
		Database: os.Getenv("DB_NAME"),
	}

	dur, _ := time.ParseDuration(os.Getenv("TOKEN_EXPIRY"))

	token := Token{
		Key:      []byte(os.Getenv("JWT_KEY")),
		Validity: dur,
	}

	return &Config{
		MongoConfig: &mConfig,
		Token:       &token,
	}
}

func GetServerConfig() *ServerConfig {
	srvConfig := &ServerConfig{}

	env := os.Getenv("env")

	fmt.Println(env)

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

func logConfig(config *Config) {
	fmt.Println("MongoUrl : ", config.MongoConfig.Url)
	fmt.Println("Database : ", config.MongoConfig.Database)
	fmt.Println("Token_Expiry : ", config.Token.Validity)
	fmt.Println("Token_Key : ", config.Token.Key)
}
