package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	DB       DBConfig
	Aria2    Aria2Config
	OpenList OpenListConfig
	Log      LogConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DBConfig struct {
	Path string
}

type Aria2Config struct {
	RPCURL string
	Secret string
}

type OpenListConfig struct {
	BaseURL string
	Token   string
}

type LogConfig struct {
	Level  string
	Format string
}

// 读取配置
func ReadConfig() Config {
	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// 从环境变量中读取配置
	return Config{
		App: AppConfig{
			Name: os.Getenv("APP_NAME"),
			Env:  os.Getenv("APP_ENV"),
			Port: os.Getenv("APP_PORT"),
		},
		DB: DBConfig{
			Path: os.Getenv("DB_PATH"),
		},
		Aria2: Aria2Config{
			RPCURL: os.Getenv("ARIA2_RPC_URL"),
			Secret: os.Getenv("ARIA2_SECRET"),
		},
		OpenList: OpenListConfig{
			BaseURL: os.Getenv("OPENLIST_BASE_URL"),
			Token:   os.Getenv("OPENLIST_TOKEN"),
		},
		Log: LogConfig{
			Level:  os.Getenv("LOG_LEVEL"),
			Format: os.Getenv("LOG_FORMAT"),
		},
	}
}