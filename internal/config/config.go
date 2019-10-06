package config

type Config struct {
	Mongo *Mongo
}

type Mongo struct {
	Url string
	Database string
}

func InitDefaults() *Config {
	mongo := Mongo {
		Url: "mongodb://localhost:27017",
		Database: "gojournal",
	}

	return &Config {
		Mongo: &mongo,
	}
}
