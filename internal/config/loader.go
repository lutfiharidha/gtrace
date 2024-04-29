package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/lutfiharidha/google-trace/internal/config/client"
	"github.com/lutfiharidha/google-trace/internal/config/logging"
	"github.com/lutfiharidha/google-trace/internal/config/server"
	"github.com/lutfiharidha/google-trace/pkg/shared/util"
	"github.com/spf13/viper"
)

type config struct {
	APP struct {
		ENV string `mapstructure:"env"`
	} `mapstructure:"app"`
	Server server.ServerList
	Logger logging.LoggerConfig
	Client client.ConfigClientList
	// Database db.DatabaseList
	// Message  message.MessageList
}

var cfg config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(dir + "/internal/config/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("server.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load server config: %v", err))
	}

	viper.AddConfigPath(dir + "/internal/config/logging")
	viper.SetConfigType("yaml")
	viper.SetConfigName("logger.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load server config: %v", err))
	}

	viper.AddConfigPath(dir + "/internal/config/client")
	viper.SetConfigType("yaml")
	viper.SetConfigName("client.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load client config: %v", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	viper.Unmarshal(&cfg)

	fmt.Println("=============================")
	fmt.Println(util.Stringify(cfg))
	fmt.Println("=============================")

}

func GetConfig() *config {
	return &cfg
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
