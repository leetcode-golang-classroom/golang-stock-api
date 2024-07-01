package config

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-stock-api/internal/util"
	"github.com/spf13/viper"
)

type Config struct {
	Port  int    `mapstructure:"PORT"`
	DbURL string `mapstructure:"POSTGRES_URL"`
}

var AppConfig *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	util.FailOnError(v.BindEnv("PORT", "POSTGRES_URL"), "failed to bind PORT, POSTGRES_URL")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("Load from environment variable")
	}
	err = v.Unmarshal(&AppConfig)
	if err != nil {
		util.FailOnError(err, "Failed to read enivronment")
	}
}
