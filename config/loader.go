package config

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Tz       string
}

type Config struct {
	Server   Server
	Database Database
}

func InitializedLoader() *Config {

	return &Config{
		Server: Server{
			Host: "127.0.0.1",
			Port: "3000",
		},
		Database: Database{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			Name:     "crud",
			Tz:       "Asia/Jakarta",
		},
	}
}
