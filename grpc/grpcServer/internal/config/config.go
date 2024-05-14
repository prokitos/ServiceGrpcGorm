package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// yaml: чтобы переменная могла иметь любое название, но из yaml бралось именно определенное
// env-default: если не указано, то по стандарту будет local
// env-required: true приложение крашнется или будет ошибка, если не указать этот параметр
// go get github.com/ilyakaznacheev/cleanenv

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GrpcConfig    `yaml:"grpc"`
}

type GrpcConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// слово Must использовать нужно редко. Используется когда не надо возвращать ошибку если она была
// например тут при старте не загрузились пути к бд, значит нет смысла запускать сервер и мы его крашим
func MustLoad() *Config {

	path := FetchConfigPath()

	// если ничего не получили
	if path == "" {
		panic("empty path to config")
	}

	// если путь получили, но ничего нет
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("nothing from this path")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read")
	}

	return &cfg
}

// получение пути к файлу через флаг который указываем при запуске приложения --config=./path
func FetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
