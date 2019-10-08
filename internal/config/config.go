package config

type Config struct {
	Mongo *Mongo
	Token *Token
}

type Mongo struct {
	Url string
	Database string
}

type Token struct {
	Key string
}

func InitDefaults() *Config {
	mongo := Mongo {
		Url: "mongodb://localhost:27017",
		Database: "gojournal",
	}

	token := Token {
		Key: "DEFAULTKEY",
	}

	return &Config {
		Mongo: &mongo,
		Token: &token,
	}
}
