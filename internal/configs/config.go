package configs

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
	"strings"
)

var Config config

type config struct {
	Host     string
	Port     string
	GRPCPort string
	Db       DbConfig
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func init() {
	if err := Load(); err != nil {
		panic("error in loading config. err: " + err.Error())
	}
}

// Load config from environment variables using viper package
func Load() error {
	viper.SetConfigName("development")

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("couldn't load file"))
	}
	pathArr := strings.Split(file, "/")

	dir := strings.Join(pathArr[:len(pathArr)-1], "/")
	dir = path.Join(dir, "../../configs")

	viper.AddConfigPath(dir)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return nil
}
