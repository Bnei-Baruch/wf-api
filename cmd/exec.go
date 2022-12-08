package cmd

import (
	"fmt"
	"github.com/Bnei-Baruch/udb/models"
	"strings"

	"github.com/spf13/viper"
)

var cfgFile string

func Exec() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Could not read config, using: ", viper.ConfigFileUsed(), err.Error())
		return
	}

	models.InitDB()
	InitMQTT()
	InitHTTP()
}
